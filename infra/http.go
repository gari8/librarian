package infra

import (
	"io"
	"net/http"
)

type Http struct{}

func NewHttp() *Http {
	return &Http{}
}

func (h Http) Get(url string, queries map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if queries != nil {
		for k, v := range queries {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
