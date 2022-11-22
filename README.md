# Response Checksum App
This is a simple application to send a get request to given urls and return the md5checksum of their response body

## Run:
- Cd into `cmd/crawler`
- Run `go run main.go` following the urls you want

#### Example:
`go run main.go google.com yahoo.com`

#### Note:
You can use `-parallel` following with the number of workers you want, to enable concurrent requests

For example to run with 3 workers:

`go run main.go -parallel 3 google.com yahoo.com live.com`

The default number of workers is set to 10 if the flag is not specified


### #TODO:
- Parsing flags can be separated to its own package
- Unit tests can be increased to cover more cases