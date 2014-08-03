package storage

import (
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type rPack struct {
	key []byte
	ch chan<- []byte
}

type wPack struct {
	key []byte
	data []byte
	ch chan<- error
}

var (
	badgeRChan chan rPack
	badgeWChan chan wPack

	assertionRChan chan rPack
	assertionWChan chan wPack

	issuerRChan chan rPack
	issuerWChan chan wPack
)

func dbReadkey() {
	select {
	case p := <- badgeRChan :
	    data, err := DB["badge"].Get(p.key, nil)
	    if err != nil {
	        close(p.ch)
	    } else {
	        p.ch <- data
	    }
	case p := <- assertionRChan :
	    data, err := DB["assertion"].Get(p.key, nil)
	    if err != nil {
	        close(p.ch)
	    } else {
	        p.ch <- data
	    }
	case p := <- issuerRChan :
	    data, err := DB["issuer"].Get(p.key, nil)
	    if err != nil {
	        close(p.ch)
	    } else {
	        p.ch <- data
	    }
	}
}

func dbWrite() {
	select {
	case p := <- badgeWChan :
	    p.ch <- DB["badge"].Put(p.key, p.data, &opt.WriteOptions{})
	case p := <- assertionWChan :
	    p.ch <- DB["assertion"].Put(p.key, p.data, &opt.WriteOptions{})
	case p := <- issuerWChan :
	    p.ch <- DB["issuer"].Put(p.key, p.data, &opt.WriteOptions{})
	}
}