package storage

func ReadKey(t string, key []byte, chdata chan<- []byte, cherr chan<- error) {
	p := rPack{
		key: key,
		chdata: chdata,
		cherr: cherr,
	}
	switch t {
	case "badge":
		badgeRChan <- p
	case "assertion":
		assertionRChan <- p
	case "issuer":
		issuerRChan <- p
	}
}

func Write(t string, key []byte, data []byte, ch chan<- error) {
	p := wPack{
		key: key,
		data: data,
		ch: ch,
	}
	switch t {
	case "badge":
		badgeWChan <- p
	case "assertion":
		assertionWChan <- p
	case "issuer":
		issuerWChan <- p
	}
}