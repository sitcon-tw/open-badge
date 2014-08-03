package badge

import "github.com/sitcon-tw/open-badge/badge/storage"

func Run() {
	storage.Init()
	serverRun()
	storage.Close()
}