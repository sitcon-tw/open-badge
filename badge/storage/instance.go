package storage

func Read(t string, key []byte, ch chan<- []byte) {
	p := rPack{
		key: key,
		ch: ch,
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