package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"templ-tutorial/utils"

	"github.com/a-h/templ"
)

func Start() {
	logfile, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not open log file: ", err)
	}
	defer logfile.Close()

	log.SetOutput(logfile)

	name := utils.RandomString(10)
	fmt.Println(name)

	pageComponent := Page()

	http.Handle("/", templ.Handler(pageComponent))

	log.Println("App running on http://localhost:8080")
	// print to console
	fmt.Println("App running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Could not start http server: ", err)
	}
}
