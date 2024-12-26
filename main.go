package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Credentials struct {
	Username string
	Password string
}

func main() {
	http.HandleFunc("/", serveLandingPage) // Serve the landing page
	http.HandleFunc("/capture", captureCredentials)

	// Serve static files (images, CSS, JS, etc.) from the 'public' directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Serve the HTML landing page
func serveLandingPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/index.html"))
	tmpl.Execute(w, nil)
}

// Handle form submission
func captureCredentials(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("Captured Credentials:")
		fmt.Printf("Username: %s\nPassword: %s\n", username, password)

		// Respond to user
		fmt.Fprintf(w, `
			<h1>Phishing Awareness Simulation</h1>
			<p>Thank you for participating in this simulation. Below are the credentials you entered:</p>
			<p><strong>Username:</strong> %s</p>
			<p><strong>Password:</strong> %s</p>
			<p><em>This was a controlled exercise. Do not use your real credentials in such simulations!</em></p>
		`, username, password)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
