// hello project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	client := createClient()
	defer client.Close()

	stringOperation(client)
	listOperation(client)
	setOperation(client)
	hashOperation(client)

	connectPool(client)
}
