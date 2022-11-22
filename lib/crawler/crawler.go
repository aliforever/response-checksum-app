package crawler

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"response-checksum-app/shared/httpclient"
	"response-checksum-app/shared/stdlogger"
	"sync"
)

// Crawler is the package responsible to fetch given addresses
type Crawler struct {
	client *httpclient.Client

	logger Logger
}

// New is the factory method for Crawler
func New(client *httpclient.Client, logger Logger) *Crawler {
	if logger == nil {
		logger = stdlogger.StdLogger{}
	}

	return &Crawler{client: client, logger: logger}
}

// Crawl crawls given addresses with given numberOfWorkers and logs the result to passed logger
func (c *Crawler) Crawl(numberOfWorkers int, addresses ...string) {
	var addressListener = make(chan string)

	var resultListener = make(chan result)
	var errorListener = make(chan failure)

	wg := &sync.WaitGroup{}
	wg.Add(len(addresses))

	for i := 0; i < numberOfWorkers; i++ {
		go c.worker(addressListener, resultListener, errorListener)
	}

	for _, address := range addresses {
		addressListener <- address
	}

	close(addressListener)

	go c.logResults(resultListener, wg)
	go c.logErrors(errorListener, wg)

	wg.Wait()

	close(resultListener)
	close(errorListener)
}

func (c *Crawler) logResults(resultListener <-chan result, wg *sync.WaitGroup) {
	for result := range resultListener {
		c.logger.Println(fmt.Sprintf("%s - %s", result.Address, result.Checksum))
		wg.Done()
	}
}

func (c *Crawler) logErrors(errListener <-chan failure, wg *sync.WaitGroup) {
	for err := range errListener {
		c.logger.Errorf("Error: %s - %s", err.Address, err.Error)
		wg.Done()
	}
}

func (c *Crawler) worker(addresses <-chan string, results chan<- result, errs chan<- failure) {
	for address := range addresses {
		resp, err := c.client.Get(c.client.ParseUrl(address))
		if err != nil {
			errs <- failure{
				Address: address,
				Error:   err,
			}
			continue
		}

		md5Sum := md5.Sum(resp)

		results <- result{
			Address:  address,
			Checksum: hex.EncodeToString(md5Sum[:]),
		}
	}
}
