package main

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/aimof/jac"
)

func main() {
	b := make([]byte, 0, 16384)
	for {
		err := binary.Read(os.Stdin, binary.BigEndian, b)
		if err != nil {
			log.Fatalln(err)
		}
	}

	u := jac.Decode(b)
	if err != nil {
		log.Fatalln(err)
	}

	f.Write(os.Stdout)
}
