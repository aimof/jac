package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aimof/jac"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	bytes := make([]byte, 0, 16384)
	for {
		b, err := buf.ReadByte()
		if err == io.EOF {
			break
		} else if err != io.EOF && err != nil {
			log.Fatalln(err)
		}
		bytes = append(bytes, b)
	}

	u := jac.Decode(bytes)

	for _, n := range u {
		fmt.Printf("%c", n)
	}
}
