package httpdao

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HttpDao struct {
	Url            string
	DefaultHeaders map[string]string
	AuthHeader     string
}

func (dao HttpDao) SetAuthHeader(authHeader string) {
	dao.DefaultHeaders["Authorization"] = "Bearer " + authHeader
}

func (dao HttpDao) NewRequest(method string, urlString string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, dao.Url+urlString, body)

	if headers == nil && dao.DefaultHeaders != nil {
		for key, value := range dao.DefaultHeaders {
			req.Header.Add(key, value)
		}
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	return req, err
}

func (dao HttpDao) NewGetRequest(urlString string, headers map[string]string) (*http.Request, error) {
	req, err := dao.NewRequest("GET", urlString, headers, nil)
	return req, err
}

func (dao HttpDao) NewFormDataPostRequest(urlString string, headers map[string]string, data url.Values) (*http.Request, error) {
	if headers == nil {
		headers = make(map[string]string)
	}

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Content-Length"] = strconv.Itoa(len(data.Encode()))

	req, err := dao.NewRequest("POST", urlString, headers, strings.NewReader(data.Encode()))
	return req, err
}

func (dao HttpDao) NewPostRequest(urlString string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := dao.NewRequest("POST", urlString, headers, body)
	return req, err
}

func (dao HttpDao) NewPutRequest(urlString string, headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := dao.NewRequest("PUT", urlString, headers, body)
	return req, err
}

func (dao HttpDao) NewDeleteRequest(urlString string, headers map[string]string) (*http.Request, error) {
	req, err := dao.NewRequest("DELETE", urlString, headers, nil)
	return req, err
}

func (dao HttpDao) Get(urlString string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}

	if headers == nil {
		headers = dao.DefaultHeaders
	}

	req, err := dao.NewGetRequest(urlString, headers)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (dao HttpDao) Post(urlString string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	if headers == nil {
		headers = dao.DefaultHeaders
	}

	req, err := dao.NewPostRequest(urlString, headers, body)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (dao HttpDao) Put(urlString string, headers map[string]string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}

	if headers == nil {
		headers = dao.DefaultHeaders
	}

	req, err := dao.NewPutRequest(urlString, headers, body)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (dao HttpDao) Delete(urlString string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}

	if headers == nil {
		headers = dao.DefaultHeaders
	}

	req, err := dao.NewDeleteRequest(urlString, headers)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}

func (dao HttpDao) PostFormData(urlString string, headers map[string]string, data url.Values) (*http.Response, error) {
	client := &http.Client{}
	req, err := dao.NewFormDataPostRequest(urlString, headers, data)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, err
}
