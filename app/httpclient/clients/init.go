package clients

import (
	"net/http"
)

type HttpClient interface {
}

type httpClient struct {
	HttpClient *http.Client
}

func InitHttpClient(conn *http.Client) HttpClient {
	return &httpClient{
		HttpClient: conn,
	}
}
