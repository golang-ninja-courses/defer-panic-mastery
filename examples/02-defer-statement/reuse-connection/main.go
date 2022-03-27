package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {
	const n = 5

	ct := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("connection was reused: %t", info.Reused)
		},
	}
	ctx := httptrace.WithClientTrace(context.Background(), ct)

	for i := 0; i < n; i++ {
		_, err := GetPage(ctx, "http://example.com/")
		if err != nil {
			log.Printf("%d: error: %v", i, err)
		}
	}
}

func GetPage(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build GET request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, res.Body)
		return nil, fmt.Errorf("not ok: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %v", err)
	}

	return body, nil
}
