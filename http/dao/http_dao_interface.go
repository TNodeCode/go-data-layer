package httpdao

import (
	"io"
	"net/http"
)

type IHttpDao interface {
	Get(urlString string) (*http.Response, error)
	GetWithHeaders(urlString string, headers map[string]string) (*http.Response, error)
	GetAsInterface(urlString string, entity interface{}) error
	Post(urlString string, body io.Reader) (*http.Response, error)
	PostWithHeaders(urlString string, headers map[string]string, body io.Reader) (*http.Response, error)
	Put(urlString string, body io.Reader) (*http.Response, error)
	PutWithHeaders(urlString string, headers map[string]string, body io.Reader) (*http.Response, error)
	Delete(urlString string) (*http.Response, error)
	DeleteWithHeaders(urlString string, headers map[string]string) (*http.Response, error)
}
