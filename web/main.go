package main

import (
  "log"
  "net/http"
  "html/template"
  "crypto/tls"
  "time"
)

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
  
  err = ts.Execute(w, msg)
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

func routes() http.Handler {
  mux := http.NewServeMux()

  mux.HandleFunc("/", home)
  mux.HandleFunc("/page/", page)

  mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
  
  return mux
}

func main() {
  tlsConfig := &tls.Config{
    PreferServerCipherSuites: true,
    CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
  }
  server := &http.Server{
    Addr: ":4000",
    Handler: recoverPanic(logRequest(secureHeaders(routes()))),
    TLSConfig: tlsConfig,
    IdleTimeout: time.Minute,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }
  
  log.Println("Starting server on :4000")
  
  err := server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
  log.Fatal(err)
}