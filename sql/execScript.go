package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "strings"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  args := os.Args[1:]
  if len(args) < 6 {
    fmt.Println("Usage: <host> <port> <user> <pass> <dbname> <path/to/script.sql>")
    os.Exit(2)
  }  
  host := args[0]
  port := args[1]
  user := args[2]
  pass := args[3]
  dbname := args[4]
  script := args[5]
  
  content, err := ioutil.ReadFile(script)
  if err != nil {
    fmt.Println("Cannot read the file:", err)
    os.Exit(1)
  }
  query := string(content)
  
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, pass, host, port, dbname))
  if err != nil {
    fmt.Println("Cannot connect to the database:", err)
    os.Exit(1)
  }  
  
  for _, q := range strings.Split(query, ";") {
    if q != "" {
      if _, err := db.Exec(q); err != nil {
        fmt.Println("Cannot execute the sql:", q, err)
        os.Exit(1)
      }
    }
  }  
  
  fmt.Println("Sql script sucessfully executed.")
}