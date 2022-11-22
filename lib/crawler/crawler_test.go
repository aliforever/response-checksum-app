package crawler_test

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"response-checksum-app/lib/crawler"
	"response-checksum-app/shared/httpclient"
	httpclientMock "response-checksum-app/shared/mocks/httpclient"
	"testing"
)

type loggerMock struct {
	data []string
	errs []string
}

func (l *loggerMock) Println(a ...any) {
	l.data = append(l.data, fmt.Sprint(a...))
}

func (l *loggerMock) Errorf(format string, a ...any) {
	l.errs = append(l.errs, fmt.Sprintf(format, a...))
}

func Test_crawler_Crawl(t *testing.T) {
	type fields struct {
		client *httpclient.Client
	}

	type args struct {
		numberOfWorkers int
		addresses       []string
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantedResult map[string]string
		wantedErrs   map[string]string
	}{
		{
			name: "Successful",
			fields: fields{
				client: httpclient.New(httpclientMock.New().SetResponse(&http.Response{
					Body: io.NopCloser(bytes.NewReader([]byte("Hello"))),
				})),
			},
			args: args{
				numberOfWorkers: 1,
				addresses:       []string{"google.com"},
			},
			// wantedResult: []string{fmt.Sprintf("%s - %s", cases[0].Address, cases[0].Md5Checksum)},
			wantedResult: map[string]string{
				"google.com": "google.com - 8b1a9953c4611296a827abf8c47804d7",
			},
		},
		{
			name: "Failed",
			fields: fields{
				client: httpclient.New(httpclientMock.New().SetError(errors.New("host_not_found"))),
			},
			args: args{
				numberOfWorkers: 1,
				addresses:       []string{"google.com"},
			},
			// wantedResult: []string{fmt.Sprintf("%s - %s", cases[0].Address, cases[0].Md5Checksum)},
			wantedResult: map[string]string{},
			wantedErrs: map[string]string{
				"google.com": "Error: google.com - host_not_found",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &loggerMock{}

			c := crawler.New(tt.fields.client, logger)

			c.Crawl(tt.args.numberOfWorkers, tt.args.addresses...)

			for _, result := range tt.wantedResult {
				found := false
				for _, printedResult := range logger.data {
					if result == printedResult {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("couldn't find %s in output", result)
				}
			}

			for _, err := range tt.wantedErrs {
				found := false
				for _, s := range logger.errs {
					if err == s {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("couldn't find %s in output", err)
				}
			}
		})
	}
}
