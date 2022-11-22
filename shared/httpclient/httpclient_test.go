package httpclient_test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"response-checksum-app/shared/httpclient"
	httpClientMock "response-checksum-app/shared/mocks/httpclient"
	"testing"
)

func Test_client_ParseUrl(t *testing.T) {
	type fields struct {
		httpClient *http.Client
	}
	type args struct {
		address string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Case 1",
			fields: fields{
				httpClient: http.DefaultClient,
			},
			args: args{
				address: "google.com",
			},
			want: "http://google.com",
		},
		{
			name: "Case 2",
			fields: fields{
				httpClient: http.DefaultClient,
			},
			args: args{
				address: "www.google.com",
			},
			want: "http://www.google.com",
		},
		{
			name: "Case 3",
			fields: fields{
				httpClient: http.DefaultClient,
			},
			args: args{
				address: "http://google.com",
			},
			want: "http://google.com",
		},
		{
			name: "Case 4",
			fields: fields{
				httpClient: http.DefaultClient,
			},
			args: args{
				address: "https://google.com",
			},
			want: "https://google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := httpclient.New(tt.fields.httpClient)

			if got := cl.ParseUrl(tt.args.address); got != tt.want {
				t.Errorf("ParseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Get(t *testing.T) {
	type fields struct {
		httpClient httpclient.HttpClient
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr error
	}{
		{
			name: "Successful response",
			fields: fields{
				httpClient: httpClientMock.New().SetResponse(&http.Response{
					Body: io.NopCloser(bytes.NewReader([]byte("<h1>Hi This is Google!</h1>"))),
				}),
			},
			args: args{
				address: "http://google.com",
			},
			want:    []byte("<h1>Hi This is Google!</h1>"),
			wantErr: nil,
		},
		{
			name: "Failed with Host not found",
			fields: fields{
				httpClient: httpClientMock.New().SetError(errors.New("host_not_found")),
			},
			args: args{
				address: "http://google.com",
			},
			want:    nil,
			wantErr: errors.New("host_not_found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := httpclient.New(tt.fields.httpClient)

			got, err := c.Get(tt.args.address)
			if (err != nil) && (tt.wantErr.Error() != err.Error()) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
