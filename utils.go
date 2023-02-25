package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var ids map[string]string
var sldb *sql.DB
var try string

// RandStringBytes generates random string of given length
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// sqllite3 initialization for trying out as key value store.
func initializesqlite3() {
	db, err := sql.Open("sqlite3", "/tmp/sqlitetest1")
	if err != nil {
		log.Fatal(err)
	}
	sldb = db
	createTable()
}

// nativemap initialization for trying out as key value store.
func initializemap() {
	ids = make(map[string]string)
	for i := 0; i < 1000000; i++ {
		id := uuid.New()
		ids[id.String()] = RandStringBytes(90)
		if i < 10 {
			fmt.Println(id)
		}
	}
}

// This method is used for sqlite kv store creates a table key and value and marks key as index
func createTable() {
	_, err := sldb.Exec("CREATE TABLE IF NOT EXISTS idtable (idkey TEXT NOT NULL PRIMARY KEY, value TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	sldb.Exec("CREATE INDEX name_index ON idtable (idkey)")
}

// This method is used for sqlite kv store to set a key and value
func setString(key string, value string) {
	// fmt.Println(key)
	query := "INSERT INTO idtable (idkey, value) VALUES('" + key + "', '" + value + "')"
	_, err := sldb.Exec(query, key, value)
	if err != nil {
		panic(err)
	}
}

// This method is used for sqlite kv store to get a value for a given key
func getkeysqlite(key string) string {
	var value string
	query := "SELECT value FROM idtable WHERE idkey = '" + key + "'"
	err := sldb.QueryRow(query, key).Scan(&value)

	if err != nil {
		if err == sql.ErrNoRows {
			return "not-found"
		} else {
			return "error found"
		}
	}
	return value
}
