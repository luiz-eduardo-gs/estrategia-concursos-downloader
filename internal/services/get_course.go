package services

import (
	"encoding/json"
	"github.com/luiz-eduardo-gs/estrategia-concursos-scraper/internal/dtos"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"os"
	"strconv"
)

type GetCourse struct {
	adapter httpclient.Client
}

func NewGetCourseService(adapter httpclient.Client) *GetCourse {
	return &GetCourse{
		adapter: adapter,
	}
}

func (s *GetCourse) Execute(id int) (*dtos.Course, error) {
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
