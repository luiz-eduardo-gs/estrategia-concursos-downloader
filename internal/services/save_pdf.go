package services

import (
	"encoding/json"
	"github.com/luiz-eduardo-gs/estrategia-concursos-scraper/internal/dtos"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Service struct {
	adapter httpclient.Client
}

func NewService(adapter httpclient.Client) *Service {
	return &Service{
		adapter: adapter,
	}
}

func (s *Service) ListCourses() (*dtos.CoursesResponse, error) {
	courses := &dtos.CoursesResponse{}

	bytes, _, err := s.adapter.Get(os.Getenv("COURSES_URL"))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, courses)
	if err != nil {
		return nil, err
	}

	return courses, nil
}

func (s *Service) GetCourseByID(id int) (*dtos.Course, error) {
	course := &dtos.Course{}

	bytes, _, err := s.adapter.Get(os.Getenv("COURSE_URL") + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, course)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (s *Service) SavePDF(url string, out string) {
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
