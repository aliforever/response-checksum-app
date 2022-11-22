package main

import (
	"flag"
	"net/http"
	"response-checksum-app/lib/crawler"
	"response-checksum-app/shared/httpclient"
	"response-checksum-app/shared/stdlogger"
)

func main() {
	var numberOfWorkers int

	flag.IntVar(&numberOfWorkers, "parallel", 10, "-parallel 10")
	flag.Parse()

	websites := flag.Args()

	client := httpclient.New(http.DefaultClient)

	cr := crawler.New(client, stdlogger.StdLogger{})

	cr.Crawl(numberOfWorkers, websites...)
}
