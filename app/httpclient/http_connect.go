package httpclient

import (
	"net/http"

	"gitlab.com/chaihanij/evat/app/env"
)

func NewHttpClient() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	return &http.Client{
		Timeout:   env.HttpClientTimeOut,
		Transport: t,
	}
}
