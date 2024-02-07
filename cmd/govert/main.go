package main

import (
	"log"
	"os"

	"github.com/abdealijaroli/govert/pkg/converter"
	"github.com/spf13/cobra"
)

var directoryFlag bool

var rootCmd = &cobra.Command{
	Use:   "govert",
	Short: "A CLI for converting Markdown to HTML",
	Long:  `A Command Line Interface application for converting Markdown to HTML`,
	Run: func(cmd *cobra.Command, args []string) {
		var inputFile, outputFile string

		if len(args) < 1 {
			log.Fatalf("Not enough arguments: expected input and output file paths")
		} else if len(args) == 1 {
			inputFile = args[0]
			outputFile = "output.html"
		} else {
			inputFile = args[0]
			outputFile = args[1]
		}

		if directoryFlag {
			var inputDirectory, outputDirectory string

			if len(args) == 1 {
				inputDirectory = args[0]
				outputDirectory = "outputDir"

				if _, err := os.Stat(outputDirectory); os.IsNotExist(err) {
					if err := os.Mkdir(outputDirectory, 0755); err != nil {
						log.Fatalf("Error creating output directory: %v", err)
					}
				}
			} else {
				inputDirectory = args[0]
				outputDirectory = args[1]
			}

			err := converter.ConvertMarkdownToHTMLDirectory(inputDirectory, outputDirectory)
			if err != nil {
				log.Fatalf("Error converting Markdown to HTML: %v", err)
			}
			log.Printf("Directory %s converted to %s\n", inputDirectory, outputDirectory)
			return
		}

		err := converter.ConvertMarkdownToHTML(inputFile, outputFile)
		if err != nil {
			log.Fatalf("Error converting Markdown to HTML: %v", err)
		}

		log.Printf("File %s converted to %s\n", inputFile, outputFile)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&directoryFlag, "directory", "d", false, "Convert a directory of Markdown files to HTML")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func main() {
	Execute()
}
