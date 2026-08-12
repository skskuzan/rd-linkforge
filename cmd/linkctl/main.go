// Command linkctl is the Linkforge operator CLI.
package main

import "fmt"

// version is overwritten at build time with -ldflags "-X main.version=...".
var version = "dev"

func main() {
	fmt.Printf("linkctl %s\n", version)
}
