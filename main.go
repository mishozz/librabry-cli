package main

import (
	"github.com/mishozz/library-cli/cli"
)

type UserDetails struct {
	email string
}

func main() {
	cli.Execute()
}
