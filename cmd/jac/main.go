package main

import (
	"bufio"
	"bytes"
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

	buf := bytes.NewBuffer(make([]byte, 0, 3*len(u)))
	for _, n := range u {
		err = binary.Write(buf, binary.BigEndian, n)
		if err != nil {
			log.Fatalln(err)
		}
	}

	f, err := os.Create("kokoro-utf8.txt.jac")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	f.Write(buf.Bytes())
}
