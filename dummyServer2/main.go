
 package main

 import (
 
	 "log"
	 "net/http"
	 "github.com/gorilla/handlers"
	 "os"
	 sw "server.com/gin/go"
 
 )
 
 func main() {
 
	 log.Printf("Server started")
	 router := sw.NewRouter()
	 server := &http.Server{
 
		 Addr:         "0.0.0.0:8080",
		 Handler:      handlers.LoggingHandler(os.Stdout, router),
	 }

	 log.Fatal(server.ListenAndServe())
 
 }