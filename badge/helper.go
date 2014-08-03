package badge

import "github.com/sitcon-tw/open-badge/badge/server"
import "github.com/sitcon-tw/open-badge/badge/storage"

func Run() {
	storage.Init()
	server.Run()
	storage.Close()
}