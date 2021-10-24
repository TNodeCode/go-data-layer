package httpdao

import (
	"io/ioutil"
	"strings"
	"testing"
)

type PostEntity struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

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

func TestGetAsMap(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	var data map[string]interface{}
	err := dao.GetAsInterface("/posts/1", &data)

	t.Log(data)

	if err != nil {
		t.Fatal(err)
	}

	if data["id"] != 1.0 {
		t.Fatalf("Id should be 1, but is %d", data["id"])
	}

	if data["userId"] != 1.0 {
		t.Fatalf("UserId should be 1, but is %d", data["userId"])
	}

	if data["title"] != "sunt aut facere repellat provident occaecati excepturi optio reprehenderit" {
		t.Fatalf("Title should be 'sunt aut facere repellat provident occaecati excepturi optio reprehenderit', but is '%s'", data["title"])
	}

	if data["body"] == "" {
		t.Fatalf("Body should not be empty")
	}
}

func TestGetAsStruct(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	var postEntity PostEntity
	err := dao.GetAsInterface("/posts/1", &postEntity)

	t.Log(postEntity)

	if err != nil {
		t.Fatal(err)
	}

	if postEntity.Id != 1 {
		t.Fatalf("Id should be 1, but is %d", postEntity.Id)
	}

	if postEntity.UserId != 1 {
		t.Fatalf("UserId should be 1, but is %d", postEntity.UserId)
	}

	if postEntity.Title != "sunt aut facere repellat provident occaecati excepturi optio reprehenderit" {
		t.Fatalf("Title should be 'sunt aut facere repellat provident occaecati excepturi optio reprehenderit', but is '%s'", postEntity.Title)
	}

	if postEntity.Body == "" {
		t.Fatalf("Body should not be empty")
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

func TestPostAsStruct(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	var postEntity PostEntity
	err := dao.PostAsInterface("/posts", strings.NewReader("{\"id\": 101}"), &postEntity)

	t.Log(postEntity)

	if err != nil {
		t.Fatal(err)
	}

	if postEntity.Id != 101 {
		t.Fatalf("Id should be 101, but is %d", postEntity.Id)
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

func TestPutAsStruct(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	var postEntity PostEntity
	err := dao.PutAsInterface("/posts/100", strings.NewReader("{\"id\": 100}"), &postEntity)

	t.Log(postEntity)

	if err != nil {
		t.Fatal(err)
	}

	if postEntity.Id != 100 {
		t.Fatalf("Id should be 100, but is %d", postEntity.Id)
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

func TestDeleteAsStruct(t *testing.T) {
	var dao IHttpDao

	// Create API Dao for Json Placeholder API
	defaultHeaders := make(map[string]string)

	dao = NewHttpDao(
		"jsonplaceholder",
		"https://jsonplaceholder.typicode.com",
		defaultHeaders,
	)

	var postEntity PostEntity
	err := dao.DeleteAsInterface("/posts/100", &postEntity)

	t.Log(postEntity)

	if err != nil {
		t.Fatal(err)
	}

	if postEntity.Id != 0 {
		t.Fatalf("Id should be 0, but is %d", postEntity.Id)
	}
}
