package converter

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

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

func ConvertMarkdownToHTMLDirectory(inputDirectory, outputDirectory string) error {
    files, err := os.ReadDir(inputDirectory)
    if err != nil {
        return err
    }

	if len(files) == 0	{
		log.Fatal("No files found in the directory")
		return nil
	}

    var wg sync.WaitGroup

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        wg.Add(1)

        go func(file os.DirEntry) {
            defer wg.Done()
            
            inputFile := filepath.Join(inputDirectory, file.Name())
            outputFile := filepath.Join(outputDirectory, strings.ReplaceAll(file.Name(), ".md", "") + ".html")

            if err := ConvertMarkdownToHTML(inputFile, outputFile); err != nil {
                log.Printf("Error converting %s: %v", inputFile, err)
            }
        }(file)
    }

    wg.Wait() 
    return nil
}
