package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

func main() {

	logfile, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file: ", err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)
	component := hello("Dave")

	http.Handle("/", templ.Handler(component))

	log.Println("App running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Could not start http server: ", err)
	}
}
