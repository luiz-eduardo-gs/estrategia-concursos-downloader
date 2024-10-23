package http

import (
	"io"
	"net/http"
	"os"
)

type Client interface {
	Get(url string) ([]byte, int, error)
}

type Adapter struct {
	cli http.Client
}

func NewClient() *Adapter {
	return &Adapter{
		cli: http.Client{},
	}
}

func (a *Adapter) Get(url string) (data []byte, status int, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))

	res, err := a.cli.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, http.StatusOK, nil
}
