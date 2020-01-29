package main

import (
  "log"
  "net/http"
  "fmt"
)

// Define a home handler function which writes a byte slice as the response body.
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }  
  name := r.URL.Query().Get("name")
  if name != "" {
    fmt.Fprintf(w, "Hello, %v!", name)
  } else {
    w.Write([]byte("Hello!"))
  }
}

func page(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    w.Header().Set("Allow", http.MethodGet)
    http.Error(w, "Method Not Allowed", 405)
    return
  }
  
  w.Write([]byte("Page"))
}

func main() {
  // Use the http.NewServeMux() function to initialize a new servemux, then
  // register the home function as the handler for the "/" URL pattern.
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/page/", page)
  
  // Use the http.ListenAndServe() function to start a new web server. We pass in
  // two parameters: the TCP network address to listen on (in this case ":4000")
  // and the servemux we just created. If http.ListenAndServe() returns an error
  // we use the log.Fatal() function to log the error message and exit. Note
  // that any error returned by http.ListenAndServe() is always non-nil.
  log.Println("Starting server on :4000")
  log.Fatal(http.ListenAndServe(":4000", mux))
}