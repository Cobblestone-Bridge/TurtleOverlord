package main

import (
	//"fmt"
	"net/http"

	"TurtleOverlord/controllers"

	"github.com/gorilla/mux"
)


func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", controllers.Index)
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../src/TurtleOverlord/static/"))))
		r.HandleFunc("/bootstrap", controllers.Bootstrap)
		r.HandleFunc("/job", controllers.Job)

		http.Handle("/", r)

    http.ListenAndServe(":8080", nil)
}
