// Command linkforged is the Linkforge service binary.
package main

import "fmt"

// version is overwritten at build time with -ldflags "-X main.version=...".
var version = "dev"

func main() {
	fmt.Printf("linkforged %s\n", version)
}
