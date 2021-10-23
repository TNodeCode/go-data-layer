package httpdao

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetRequest(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	res, err := dao.Get("/posts/1")

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(body))

	expectedResponseBody := `{` + "\n" +
		`  "userId": 1,` + "\n" +
		`  "id": 1,` + "\n" +
		`  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",` + "\n" +
		`  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"` + "\n" +
		`}`

	t.Logf(expectedResponseBody)

	if string(body) != expectedResponseBody {
		t.Fatalf("Expected GET reponse body does not match actual response body")
	}
}

func TestPostRequest(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	requestBody := `{` + "\n" +
		`  "userId": 1,` + "\n" +
		`  "id": 101,` + "\n" +
		`  "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",` + "\n" +
		`  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"` + "\n" +
		`}`

	res, err := dao.Post("/posts", strings.NewReader(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(body))

	expectedResponseBody := `{` + "\n" +
		`  "id": 101` + "\n" +
		`}`

	t.Logf(expectedResponseBody)

	if string(body) != expectedResponseBody {
		t.Fatalf("Expected POST reponse body does not match actual response body")
	}
}

func TestPutRequest(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	requestBody := `{` + "\n" +
		`  "userId": 1,` + "\n" +
		`  "id": 1,` + "\n" +
		`  "title": "xsunt aut facere repellat provident occaecati excepturi optio reprehenderit",` + "\n" +
		`  "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto"` + "\n" +
		`}`

	res, err := dao.Put("/posts/1", strings.NewReader(requestBody))

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(body))

	expectedResponseBody := `{` + "\n" +
		`  "id": 1` + "\n" +
		`}`

	t.Logf(expectedResponseBody)

	if string(body) != expectedResponseBody {
		t.Fatalf("Expected PUT reponse body does not match actual response body")
	}
}

func TestDeleteRequest(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	res, err := dao.Delete("/posts/1")

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf(string(body))

	expectedResponseBody := "{}"

	t.Logf(expectedResponseBody)

	if string(body) != expectedResponseBody {
		t.Fatalf("Expected DELETE reponse body does not match actual response body")
	}
}
