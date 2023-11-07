package pokecache

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (JsonResponse, error) {
	// Sets base URL if it doesn't have one passed in
	urlToUse := base + "/location-area?offset=0&limit=20"
	fmt.Println(urlToUse)
	if pageURL != nil {
		urlToUse = *pageURL
	}
	fmt.Printf("URL In Use: %v\n", urlToUse)

	// Check Cache
	fmt.Println("Checking Cache")
	if value, ok := c.Cache.Get(urlToUse); ok {
		resp := JsonResponse{}
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return JsonResponse{}, err
		}
		fmt.Println("Found Resposne")
		return resp, nil
	}

	// Call API if not in cache
	fmt.Println("Calling API")
	r := JsonResponse{}
	res, err := http.Get(urlToUse)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return JsonResponse{}, err
	}

	c.Cache.Add(urlToUse, body)
	return r, nil
}
