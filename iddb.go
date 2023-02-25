package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/akrylysov/pogreb"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

// The best choice is usnig pogreb as key value store. This is a simple kv store that is optimized for fast reads.
// More info can be found here https://artem.krylysov.com/blog/2018/03/24/pogreb-key-value-store/

type iddb interface {
	Write(string, string)
	Read(string) string
}

type pogorebdb struct {
	pgdb    *pogreb.DB
	idcache *cache.Cache
}

// NewPogoDB initlializes db with an LRU cache for 500 items.
func NewPogoDB(test bool) *pogorebdb {
	db, err := pogreb.Open("/tmp/pogreb.test", nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// test data initializes.
	if test {
		for i := 0; i < 1000000; i++ {
			id := uuid.New()
			err := db.Put([]byte(id.String()), []byte(RandStringBytes(90)))
			if err != nil {
				log.Fatal(err)
			}
			if i < 10 {
				fmt.Println(id)
			}
		}
	}
	return &pogorebdb{
		pgdb:    db,
		idcache: cache.NewFrom(0, 0, make(map[string]cache.Item, 500)),
	}
}

func (db *pogorebdb) Write(key, value string) {
	err := db.pgdb.Put([]byte(key), []byte(value))
	if err != nil {
		log.Fatal(err)
	}
}

// Read gets the key from cache and fetches from db for cache miss and add the value to the cache before returning.
func (db *pogorebdb) Read(key string) string {
	value, ok := db.idcache.Get(key)
	if ok {
		return value.(string)
	}
	val, err := db.pgdb.Get([]byte(key))
	if err != nil {
		return "not-found"
	}
	kvalue := string(val)
	db.idcache.Set(key, kvalue, 0)
	return kvalue
}

func intializeSampleData() {
	readFile, err := os.Open("example.data")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		keyandvalues := strings.SplitN(fileScanner.Text(), " ", 2)
		iddbkey.Write(keyandvalues[0], keyandvalues[1])
	}
	readFile.Close()
}
