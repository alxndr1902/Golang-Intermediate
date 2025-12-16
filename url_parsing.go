package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://example.com:8080/path?query=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=John&age=30"

	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println("Query Params:", queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	//building url
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "30")
	baseURL.RawQuery = query.Encode()
	fmt.Println("Built URL:", baseURL.String())

	values := url.Values{}
	values.Add("name", "Jane")
	values.Add("age", "30")
	values.Add("city", "london")
	encodedQuery := values.Encode()
	fmt.Println("Encoded Query:", encodedQuery)

	baseURL1 := "http://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery

	fmt.Println("Full URL:", fullURL)
}
