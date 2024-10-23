package main

import (
	"github.com/joho/godotenv"
	"github.com/luiz-eduardo-gs/estrategia-concursos-scraper/internal/services"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}

func main() {
	loadEnv()

	srcFolder := os.Getenv("RESOURCES_FOLDER")

	cli := httpclient.NewClient()
	svc := services.NewService(cli)
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	courses, err := svc.ListCourses()
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses.Data.Courses {
		cInfo, err := svc.GetCourseByID(course.Id)
		if err != nil {
			log.Fatal(err)
		}

		for i, class := range cInfo.Data.Classes {
			filename := filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name) + ".pdf"

			if _, err := os.Stat(filename); err == nil {
				filename = filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name+"("+strconv.Itoa(i)+")") + ".pdf"
			}

			go svc.SavePDF(class.Pdf, filename)
		}

		time.Sleep(60 * time.Second)
	}
}
