# Charles-to-Go
Simple GUI to convert Charles headers to golang's default http client (net/http)

## Usage
1. Compile code to a binary, `go build -ldflags -H=windowsgui` is reccomended in order to not show the CLI
2. After opening the exe, copy the headers from the Charles request data and simply click the `Convert!` button
3. BAM, your clipboard is now formatted as `req.Header.Set("a", "b")` for easy pasting into your project
