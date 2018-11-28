package main

import (
	"fmt"
	"github.com/rexitorg/sample_go_pkg/domain/user"
)

func main() {
	u := user.NewUser("konchan")
	fmt.Printf("print user name: %#v\n", u)
}
