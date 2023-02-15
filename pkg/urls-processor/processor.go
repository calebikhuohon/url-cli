package urls_processor

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"golang.org/x/exp/slices"
)

type Pair struct {
	Url      string `json:"url"`
	BodySize int64  `json:"body_size"`
}

func VisitUrls(ctx context.Context, urls []string) []Pair {
	var (
		pairs []Pair
		wg    sync.WaitGroup
	)

	for _, u := range urls {
		link, err := url.Parse(u)
		if err != nil {
			log.Println("error: ", err)
			continue
		}

		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			bodySize, err := doReq(ctx, http.MethodGet, url)
			if err != nil {
				log.Printf("failed to get body size of url - %s: error - %v", url, err)
				return
			}

			pairs = append(pairs, Pair{
				Url:      url,
				BodySize: bodySize,
			})
		}(link.String())
	}
	wg.Wait()

	slices.SortFunc(pairs, func(a, b Pair) bool {
		return a.BodySize > b.BodySize
	})

	return pairs
}

func doReq(ctx context.Context, method string, path string) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, method, path, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		var data interface{}
		if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("request failed with status: %d and err: %v", res.StatusCode, data)
	}

	if res.Header.Get("Content-Length") != "" {
		bodySize, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			return 0, err
		}

		return bodySize, nil
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return int64(len(body)), nil
}
