package main

import (
	"encoding/base64"
	"flag"
	"io"
	"os"
)

var (
	flagURL = flag.Bool("u", false, "decode URL encoding")
	flagRaw = flag.Bool("r", false, "decode raw encoding - no padding")
)

func main() {
	flag.Parse()
	enc := base64.StdEncoding
	if *flagURL {
		enc = base64.URLEncoding
	}
	if *flagRaw {
		enc = enc.WithPadding(base64.NoPadding)
	}
	io.Copy(os.Stdout, base64.NewDecoder(enc, os.Stdin))
}
