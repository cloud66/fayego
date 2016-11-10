package main

import (
	"fmt"

	"github.com/cloud66/fayego/fayeserver"
)

func main() {
	fmt.Println("Starting faye server on port 3002")
	fayeserver.Start(":3002")
}
