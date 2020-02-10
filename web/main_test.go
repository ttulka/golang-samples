package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestHome(t *testing.T) {
  rr := httptest.NewRecorder()
  r, err := http.NewRequest(http.MethodGet, "/", nil)
  if err != nil {
    t.Fatal(err)
  }
  
  home(rr, r)
  
  rs := rr.Result()  
  if rs.StatusCode != http.StatusOK {
    t.Errorf("Expected %d to be %d", rs.StatusCode, http.StatusOK)
  }
  defer rs.Body.Close()
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil {
    t.Fatal(err)
  }
  if !strings.Contains(string(body), "<h1>Hello!</h1>") {
    t.Errorf("Expected body to contain %q", "<h1>Hello!</h1>")
  }
}

func TestHomeName(t *testing.T) {
  rr := httptest.NewRecorder()
  r, err := http.NewRequest(http.MethodGet, "/?name=Tomas", nil)
  if err != nil {
    t.Fatal(err)
  }
  
  home(rr, r)
  
  rs := rr.Result()  
  if rs.StatusCode != http.StatusOK {
    t.Errorf("Expected %d to be %d", rs.StatusCode, http.StatusOK)
  }
  defer rs.Body.Close()
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil {
    t.Fatal(err)
  }
  if !strings.Contains(string(body), "<h1>Hello, Tomas!</h1>") {
    t.Errorf("Expected body to contain %q", "<h1>Hello!</h1>")
  }
}

func TestPage(t *testing.T) {
  rr := httptest.NewRecorder()
  r, err := http.NewRequest(http.MethodGet, "/page/", nil)
  if err != nil {
    t.Fatal(err)
  }
  
  page(rr, r)
  
  rs := rr.Result()  
  if rs.StatusCode != http.StatusOK {
    t.Errorf("Expected %d to be %d", rs.StatusCode, http.StatusOK)
  }
  defer rs.Body.Close()
  body, err := ioutil.ReadAll(rs.Body)
  if err != nil {
    t.Fatal(err)
  }
  if !strings.Contains(string(body), "Page") {
    t.Errorf("Expected body to contain %q", "Page")
  }
}

func TestPageOnlyGetAllowed(t *testing.T) {
  rr := httptest.NewRecorder()
  r, err := http.NewRequest(http.MethodPost, "/page/", nil)
  if err != nil {
    t.Fatal(err)
  }
  
  page(rr, r)
  
  rs := rr.Result()  
  if rs.StatusCode != http.StatusMethodNotAllowed {
    t.Errorf("Expected %d to be %d", rs.StatusCode, http.StatusMethodNotAllowed)
  }
  defer rs.Body.Close()
}