package services

import (
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"log"
	"os"
	"path/filepath"
)

type SavePdf struct {
	adapter httpclient.Client
}

func NewSavePdfService(adapter httpclient.Client) *SavePdf {
	return &SavePdf{
		adapter: adapter,
	}
}

func (s *SavePdf) Execute(url string, out string) {
	bytes, _, err := s.adapter.Get(url)

	log.Printf("Downloading file: %s", out)

	err = saveFile(out, bytes)
	if err != nil {
		log.Fatalf("Error trying to save pdf: %s", err.Error())
	}
}

func saveFile(path string, b []byte) error {
	dir, _ := filepath.Split(path)

	err := os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
