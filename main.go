package main

import (
	"github.com/mishozz/library-cli/cli"
)

const (
	HOST           = "http://localhost:"
	PORT           = "8080"
	LIBRARY_API_V1 = "/library/api/v1/"
)

func main() {
	cli.Execute()
}
