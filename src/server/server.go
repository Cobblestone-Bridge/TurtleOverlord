package main

import (
	"fmt"
	"io"
	"net/http"
)

func loginServerHandler(wrt http.ResponseWriter, req *http.Request) {
	io.WriteString(wrt, "Turtle Server");
	var turtleID string
	turtleID = req.FormValue("id")
	fmt.Printf(" [Login] Turtle has joined with ID: " + turtleID + "\n")
	//DB Check for turtle and if it is connected + create object
}

func jobServerHandler(wrt http.ResponseWriter, req *http.Request) {
	io.WriteString(wrt, "Turtle Server")
	fmt.Printf(" [Job] New Turtle has joined with ID: " + req.FormValue("id") + "\n")
	//DB Check for turtle and if it is connected + create object
}

func statusServerHandler(wrt http.ResponseWriter, req *http.Request) {
	io.WriteString(wrt, "Turtle Server")
	fmt.Printf(" [Status] New Turtle has joined with ID: " + req.FormValue("id") + "\n")
	//DB Check for turtle and if it is connected + create object
}

func webServerHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Web Server - Coming Soon")
	//Display current stats on server and turtles + resources
}

func startServer(port string){
	fmt.Printf(" [Server] Port " + port + " is in use!\n")
	fmt.Printf(" [Server] Please check your webbrowser!\n")

	//Assigning handlers
	http.HandleFunc("/turtle/job/", jobServerHandler)
	http.HandleFunc("/turtle/login/", loginServerHandler)
	http.HandleFunc("/turtle/status/", statusServerHandler)
	http.HandleFunc("/", webServerHandler)
	http.ListenAndServe(port, nil)
}

func main(){
	//Start up logs
	fmt.Printf("\n Turtle Overlord (v0.1)\n\n")
	fmt.Printf(" [Server] Server is starting!\n")

	//Config should be read just before
	//Starting the server
	startServer(":8000")
}	