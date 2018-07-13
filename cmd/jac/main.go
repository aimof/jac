package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"log"
	"os"

	"github.com/aimof/jac"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	runes := make([]rune, 0, 16384)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		runes = append(runes, r)
	}

	u, err := jac.Encode(runes)
	if err != nil {
		log.Fatalln(err)
	}

	for _, n := range u {
		err = binary.Write(os.Stdout, binary.BigEndian, n)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
