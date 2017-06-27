package main

import (
	"bytes"
	"os"

	"github.com/dimiro1/banner"
)

func main() {
	isEnabled := true
	isColorEnabled := true
	banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBufferString("My Custom Banner"))
}
