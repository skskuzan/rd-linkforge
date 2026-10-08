package main

import (
	"fmt"
	"os"

	"github.com/victorkovalyov-teletec/linkforge/internal/link"
	"github.com/victorkovalyov-teletec/linkforge/internal/store"
)

func main() {
	s := store.New()
	var nextID uint64 = 1

	for _, target := range os.Args[1:] {
		l, err := link.New(nextID, target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", target, err)
			continue
		}

		err = store.Add(s, l)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding to store: %v\n", err)
			continue
		}
		nextID++
	}

	links := store.All(s)
	for _, l := range links {
		fmt.Printf("%s  %s\n", l.Code, l.Target)
	}

	if len(links) > 0 {
		fmt.Printf("\nTotal: %d\n", store.Count(s))
	}
}
