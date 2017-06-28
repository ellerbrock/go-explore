package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ellerbrock/go-check"
)

func main() {
	content, err := ioutil.ReadFile("x.txt")
	check.Error(err)

	// we only arrive here if no error ...
	fmt.Printf("file content:\n\n%v", content)
}
