# DBI

Postgresql Database Interface on top of pgx

[![PkgGoDev](https://pkg.go.dev/badge/github.com/junekimdev/dbi)](https://pkg.go.dev/github.com/junekimdev/dbi)
[![Go Report Card](https://goreportcard.com/badge/github.com/junekimdev/dbi)](https://goreportcard.com/report/github.com/junekimdev/dbi)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/junekimdev/dbi)
![GitHub](https://img.shields.io/github/license/junekimdev/dbi)

---

## Getting Started

### Prerequisite

- You need an active running instance of Postgresql DB somewhere you can access
- (optional) Create `.env` file in your root directory and add below variables
  - PGUSER
  - PGPASSWORD
  - PGHOST
  - PGDATABASE
  - PGPORT

### Installing

go get it (pun intended :smile_cat:)

```shell
go get github.com/junekimdev/dbi
```

## Usage

```golang
package main

import (
  "log"

  "github.com/junekimdev/dbi"
  "github.com/joho/godotenv"
)

func init(){
  // Load environment variables
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Failed to load .env file: %v", err)
  }
  // Connect to DB with env
  dbi.Connect(dbi.CreateURIFromEnv().String())
}

// User struct
type User struct {
  id     int
  name   string
  gender string
}

func main() {
  // Prep your sql
  sql := "SELECT id, name, gender FROM test_user WHERE id=($1)"

  // Query DB
  result := Query(sql, 7)

  // Destination of Scan
  var user User

  // Scan the result
  // You Need to pass "scan function" that has "Scan" method of pgx.Rows
  if err := Scan(result, func() { result.Scan(&user.id, &user.name, &user.gender) }); err != nil {
    panic(err)
  }

  log.Printf("%#+v", user)
}
```
