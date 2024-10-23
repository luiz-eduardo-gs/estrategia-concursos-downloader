package main

import (
	"github.com/joho/godotenv"
	"github.com/luiz-eduardo-gs/estrategia-concursos-scraper/internal/services"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-scraper/pkg/clients/http"
	"log"
	"os"
	"path/filepath"
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
	listCourses := services.NewListCoursesService(cli)
	getCourse := services.NewGetCourseService(cli)
	savePdf := services.NewSavePdfService(cli)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	courses, err := listCourses.Execute()
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses.Data.Courses {
		cInfo, err := getCourse.Execute(course.Id)
		if err != nil {
			log.Fatal(err)
		}

		for _, class := range cInfo.Data.Classes {
			filename := filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name+"_"+time.Now().String()) + ".pdf"
			go savePdf.Execute(class.Pdf, filename)
		}

		time.Sleep(60 * time.Second)
	}
}
