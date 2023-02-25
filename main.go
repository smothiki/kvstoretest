package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"

	"database/sql"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var iddbkey iddb

// getkey is a factory method to get the value from a choice of kv store based on try value for a given key.
func getkey(key string) string {
	if try == "" {
		value, ok := ids[key]
		if ok {
			return value
		}
	}
	if try == "1" {
		return iddbkey.Read(key)
	}
	if try == "2" {
		return getkeysqlite(key)
	}
	return "not-found"
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", getkey(c.URLParams["key"]))
}

// A factory method to initialize a choice of key value implementation based on try value.
func initialize() {
	if try == "" {
		initializemap()
	}
	if try == "1" {
		iddbkey = NewPogoDB(false)
		intializeSampleData()
	}
	if try == "2" {
		// initializesqlite3()
		db, err := sql.Open("sqlite3", "/tmp/sqlitetest1")
		if err != nil {
			log.Fatal(err)
		}
		sldb = db
	}
}

func main() {
	// Try is a global variable to test various key value store scenarios
	// if try is empty a native hash map of string , string will be used.
	// try=1 uses pogreb with cache is used.
	try = "1"
	// initialize initializes the key value store with example data.
	initialize()
	goji.Get("/key/:key", hello)
	goji.Serve()
}
