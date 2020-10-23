package dbi

import (
	"os"
)

// URI struct represents parts of connection string
type URI struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// CreateURIFromEnv create URI from environment variables
//
// This should be called after loading environment variables; otherwise, it will create an empty URI
func CreateURIFromEnv() *URI {
	user := os.Getenv("PGUSER")
	pw := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	database := os.Getenv("PGDATABASE")
	return &URI{Username: user, Password: pw, Host: host, Port: port, Database: database}
}

// String gives a valid connection string
func (uri *URI) String() string {
	var userpw string
	var hostport string
	var withDB string

	// No host
	if uri.Host == "" {
		return "postgresql://"
	}

	// username and password
	if uri.Username != "" && uri.Password != "" {
		userpw = uri.Username + ":" + uri.Password + "@"
	} else if uri.Username != "" {
		userpw = uri.Username + "@"
	}

	// host and port
	if uri.Port == "" {
		hostport = uri.Host
	} else {
		hostport = uri.Host + ":" + uri.Port
	}

	// database name
	if uri.Database != "" {
		withDB = "/" + uri.Database
	}

	return "postgresql://" + userpw + hostport + withDB
}
