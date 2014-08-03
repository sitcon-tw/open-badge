package storage

import (
	"os"
	"log"
	"github.com/syndtr/goleveldb/leveldb"
)

var DB map[string]*leveldb.DB
var Index map[string]*leveldb.DB

func mkdir() {
	if _, err := os.Stat("./storage"); os.IsNotExist(err) {
		if err := os.Mkdir("./storage", 0755); err != nil {
			log.Fatal("[BadgeStorage]", err)
			os.Exit(0)
		}
	}
	if _, err := os.Stat("./storage/index"); os.IsNotExist(err) {
		if err := os.Mkdir("./storage/index", 0755); err != nil {
			log.Fatal("[BadgeStorage]", err)
			os.Exit(0)
		}
	}
}

func loadDB() {
	DB = make(map[string]*leveldb.DB)
	Index = make(map[string]*leveldb.DB)
	for _, v := range([]string{"badge", "assertion", "issuer"}) {
		if db, err := leveldb.OpenFile("./storage/" + v + ".ldb", nil); err == nil {
			DB[v] = db
		} else {
			log.Fatal("[BadgeStorage]", err)
			os.Exit(0)
		}
	}
	for _, v := range([]string{}) {
		if db, err := leveldb.OpenFile("./storage/index/" + v + ".ldb", nil); err == nil {
			Index[v] = db
		} else {
			log.Fatal("[BadgeStorage]", err)
			os.Exit(0)
		}
	}
}

func Init() {
	mkdir()
	loadDB()
	go dbReadkey()
	go dbWrite()
}

func Close() {
	for k := range(DB) {
		DB[k].Close()
	}
	for k := range(Index) {
		Index[k].Close()
	}
}