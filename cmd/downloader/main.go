package main

import (
	"github.com/joho/godotenv"
	"github.com/luiz-eduardo-gs/estrategia-concursos-downloader/internal/services"
	httpclient "github.com/luiz-eduardo-gs/estrategia-concursos-downloader/pkg/clients/http"
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
	listCoursesSvc := services.NewListCoursesService(cli)
	getCourseSvc := services.NewGetCourseService(cli)
	savePdfSvc := services.NewSavePdfService(cli)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	courses, err := listCoursesSvc.Execute()
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses.Data.Courses {
		cInfo, err := getCourseSvc.Execute(course.Id)
		if err != nil {
			log.Fatal(err)
		}

		for _, class := range cInfo.Data.Classes {
			filename := filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name+"_"+strconv.FormatInt(time.Now().UnixNano(), 10)) + ".pdf"
			go savePdfSvc.Execute(class.Pdf, filename)
		}

		time.Sleep(10 * time.Second)
	}
}
