package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string // This will store the hashed password
}

var (
	users     = make(map[string]string) // Map to store username:hashedPassword
	userMutex = sync.Mutex{}            // To prevent concurrent access issues
)

// Save users to a file
func saveUsers() {
	file, err := os.Create("users.txt")
	if err != nil {
		log.Println("Error creating users file:", err)
		return
	}
	defer file.Close()

	for username, hashedPassword := range users {
		file.WriteString(username + ":" + hashedPassword + "\n")
	}
}

// Load users from a file
func loadUsers() {
	file, err := os.Open("users.txt")
	if err != nil {
		log.Println("Error loading users file:", err)
		return
	}
	defer file.Close()

	var username, hashedPassword string
	for {
		_, err := fmt.Fscanf(file, "%s:%s\n", &username, &hashedPassword)
		if err != nil {
			break
		}
		users[username] = hashedPassword
	}
}

func main() {
	loadUsers()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("login.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		userMutex.Lock()
		hashedPassword, exists := users[username]
		userMutex.Unlock()

		if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Login successful!"))
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error creating account", http.StatusInternalServerError)
			return
		}

		userMutex.Lock()
		users[username] = string(hashedPassword)
		userMutex.Unlock()

		saveUsers()

		w.Write([]byte("Registration successful!"))
	})

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
