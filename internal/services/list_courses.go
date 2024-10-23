package services

import (
	"encoding/json"
	"github.com/luiz-eduardo-gs/estrategia-concursos-scraper/internal/dtos"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"os"
)

type ListCourses struct {
	adapter httpclient.Client
}

func NewListCoursesService(adapter httpclient.Client) *ListCourses {
	return &ListCourses{
		adapter: adapter,
	}
}

func (s *ListCourses) Execute() (*dtos.CoursesResponse, error) {
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
