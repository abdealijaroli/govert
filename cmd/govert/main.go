package main

import (
	"fmt"
	"log"

	"github.com/abdealijaroli/govert/pkg/converter"
)

func main() {
	err := converter.ConvertMarkdownToHTML("test/input.md", "test/output.html")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File %s converted to %s\n", "test/input.md", "test/output.html")
}
