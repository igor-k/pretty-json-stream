package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	open  = "<PLAIN TEXT START>"
	close = "<PLAIN TEXT END>"
)

func init() {
	flag.StringVar(&open, "open", open, "")
	flag.StringVar(&close, "close", close, "")

	flag.Parse()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var buf bytes.Buffer

	for scanner.Scan() {
		buf.Reset()

		if err := json.Indent(&buf, scanner.Bytes(), "", "\t"); err != nil {
			fmt.Printf("%s\n%s\n%s\n", open, scanner.Text(), close)
			continue
		}

		fmt.Println(buf.String())
	}
}
