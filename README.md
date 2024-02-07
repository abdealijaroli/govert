# Govert

Govert is a Command Line Interface (CLI) application for converting Markdown to HTML with Concurrent Batch Conversion and Live Preview.

## Installation

To use Govert, follow these steps:

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/abdealijaroli/govert.git
    ```

2. Navigate to the project directory:

    ```bash
    cd govert
    ```

3. Build the binary:

    ```bash
    go build -o bin/govert ./cmd/govert
    ```

## Usage

To convert a Markdown file to HTML using Govert, run the binary with the input and output file paths as arguments:

```bash
./bin/govert input.md output.html
```

If you don't specify an output path, the output will be written to "output.html" by default.


```bash
./bin/govert input.md
```

To convert a directory of Markdown files to HTML using Govert, run the binary with the input and output directory paths as arguments:


```bash 
./bin/govert -d path/to/input-directory
```

You don't need to specify an output directory, the output will be written to "outputDir" by default.


## Contributing

Contributions are welcome! To contribute to Govert, follow these steps:

1. Fork this repository.

2. Create a new branch:

    ```bash
    git checkout -b feature/branch-name
    ```

3. Make your changes and commit them:

    ```bash
    git commit -m 'Add some feature'
    ```

4. Push to the main branch:

    ```bash
    git push origin feature/branch-name
    ```
5. Create a new Pull Request.

## License

This project is licensed under the MIT License.