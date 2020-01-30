package main

import (
  "log"
  "net/http"
  "html/template"
)

// Define a home handler function which writes a byte slice as the response body.
func home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  ts, err := template.ParseFiles("./html/home.tmpl", "./html/layout.tmpl")
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
    return
  }
  
  msg := "Hello"
  name := r.URL.Query().Get("name")
  if name != "" {
    msg += ", " + name + "!"
  } else {
    msg += "!"
  }
  
  err = ts.Execute(w, nil)
  if err != nil {
    log.Println(err.Error())
    http.Error(w, "Internal Server Error", 500)
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
  mux := http.NewServeMux()

  mux.HandleFunc("/", home)
  mux.HandleFunc("/page/", page)

  mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  
  server := &http.Server{
    Addr: ":4000",
    Handler: mux,
  }
  
  log.Println("Starting server on :4000")
  
  err := server.ListenAndServe(":4000", mux)
  log.Fatal(err)
}