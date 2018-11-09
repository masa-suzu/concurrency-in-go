package chapter4

import (
	"fmt"
	"net/http"
)

func GetFromURLs() {
	type Response struct {
		Error error
		Value *http.Response
	}

	get := func(done <-chan interface{}, urls ...string) <-chan Response {
		response := make(chan Response)
		go func() {
			defer close(response)

			for _, url := range urls {
				v, e := http.Get(url)
				res := Response{Error: e, Value: v}

				select {
				case <-done:
					return
				case response <- res:
				}
			}
		}()
		return response
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://211212.dsd"}

	for response := range get(done, urls...) {
		if response.Error != nil {
			fmt.Printf("error: %v\n", response.Error)
			continue
		}
		fmt.Printf("Response: %v\n", response.Value.Status)
	}
}
