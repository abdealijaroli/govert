package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abdealijaroli/govert/pkg/converter"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var directoryFlag bool
var liveFlag bool

var rootCmd = &cobra.Command{
	Use:   "govert",
	Short: "A CLI for converting Markdown to HTML",
	Long:  `A Command Line Interface application for converting Markdown to HTML`,
	Run: func(cmd *cobra.Command, args []string) {
		if liveFlag {
			go startServer()
			go watchFiles(args[0])
			select {}
		} else {
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
		}
	},
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir("test")))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func watchFiles(inputFile string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(event.Op) {
					log.Println("Modified file:", event.Name)
					converter.ConvertMarkdownToHTML(event.Name, "test/output.html")
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func init() {
	rootCmd.Flags().BoolVarP(&directoryFlag, "directory", "d", false, "Convert a directory of Markdown files to HTML")
	rootCmd.Flags().BoolVarP(&liveFlag, "live", "l", false, "Start a live server and watch for changes")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
}

func main() {
	Execute()
}
