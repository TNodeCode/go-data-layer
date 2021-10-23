# Go Data Layer

Package that abstracts data sources and provides useful methods to access them.

#### Installation

```bash
$ go get github.com/tnodecode/go-data-layer
```

## HTTP Dao

Send GET Request to HTTP Endpoint

```go
package main

import (
	"fmt"
	"io/ioutil"

	httpdao "github.com/tnodecode/go-data-layer/http/dao"
)

func main() {
    // Create Data Access Object for HTTP endpoint
	headers := make(map[string]string)
	dao := httpdao.NewHttpDao("dao1", "https://jsonplaceholder.typicode.com", headers)

    // Send GET Request
	res, err := dao.Get("/posts/1", headers)

    // Check for errors
	if err != nil {
		fmt.Println(err)
	}

    // Get response body
	body, err := ioutil.ReadAll(res.Body)

    // Print response body
	fmt.Println(string(body))
}
```
