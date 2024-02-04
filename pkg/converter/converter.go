package converter

import (
	"log"
	"os"

	"github.com/russross/blackfriday/v2"
)

func ConvertMarkdownToHTML(inputFile, outputFile string) error {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	output := blackfriday.Run(input)

	err = os.WriteFile(outputFile, output, 0666)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
