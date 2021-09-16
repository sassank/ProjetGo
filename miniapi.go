package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//routes
func main() {
	http.HandleFunc("/", Heure)
	http.HandleFunc("/add", Ajout)
	http.HandleFunc("/entries", Resultat)

	fmt.Println("Serveur lance sur le port 4567")
	http.ListenAndServe(":4567", nil)
}

//afficher l'heure actuelle
func Heure(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		notAllowed(w, req)
	} else {
		fmt.Fprintf(w, time.Now().Format("15:04"))
	}
}

func notAllowed(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, req.Method+" is not allowed.")
}

//ajouter des donnees
func Ajout(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	if req.Method != "POST" {
		notAllowed(w, req)
	}
}

//afficher les donnees
func Resultat(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		notAllowed(w, req)
	} else {
		entries := listEntries()

		for _, rawEntry := range entries {
			entry := strings.Split(rawEntry, ":")

			fmt.Fprintf(w, entry[1]+"\n")
		}
	}
}

func addEntry(author, message string) {
	f, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(author + ":" + message + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}

func listEntries() []string {
	raw, err := os.ReadFile("results.txt")

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(raw), "\n")

	return data
}
