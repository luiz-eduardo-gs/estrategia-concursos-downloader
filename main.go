package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	CoursesUrl = "https://api.estrategiaconcursos.com.br/api/aluno/pacote/223673"
	CourseUrl  = "https://api.estrategiaconcursos.com.br/api/aluno/curso/"
)

type CoursesResponse struct {
	Data struct {
		Id      int         `json:"id"`
		Name    string      `json:"nome"`
		LdiUrl  interface{} `json:"ldi_url"`
		Courses []struct {
			Id           int    `json:"id"`
			Name         string `json:"nome"`
			CourseTypeID int    `json:"tipo_curso_id"`
		} `json:"cursos"`
	} `json:"data"`
}

type CourseInfo struct {
	Data struct {
		Id           int    `json:"id"`
		Name         string `json:"nome"`
		DataInicio   string `json:"data_inicio"`
		DataRetirada string `json:"data_retirada"`
		Tipo         string `json:"tipo"`
		Professores  []struct {
			Id     int    `json:"id"`
			Nome   string `json:"nome"`
			Imagem string `json:"imagem"`
		} `json:"professores"`
		Classes []struct {
			Id                       int         `json:"id"`
			Name                     string      `json:"nome"`
			Conteudo                 string      `json:"conteudo"`
			IsDisponivel             bool        `json:"is_disponivel"`
			Visualizada              bool        `json:"visualizada"`
			DataPublicacao           string      `json:"data_publicacao"`
			Pdf                      string      `json:"pdf"`
			FuncionalidadePdfGrifado interface{} `json:"funcionalidade_pdf_grifado"`
			PdfGrifado               interface{} `json:"pdf_grifado"`
			PdfSimplificado          *string     `json:"pdf_simplificado"`
			IsAlunoFinalizado        bool        `json:"is_aluno_finalizado"`
			ConteudosPendentes       int         `json:"conteudos_pendentes"`
			Videos                   []struct {
				Id          int     `json:"id"`
				Titulo      string  `json:"titulo"`
				Visualizado bool    `json:"visualizado"`
				Resumo      *string `json:"resumo"`
				Slide       *string `json:"slide"`
				MapaMental  *string `json:"mapa_mental"`
				Posicao     int     `json:"posicao"`
				Resolucao   string  `json:"resolucao"`
				Resolucoes  struct {
					P  string `json:"720p"`
					P1 string `json:"480p"`
					P2 string `json:"360p"`
				} `json:"resolucoes"`
				Anotacoes interface{} `json:"anotacoes"`
				Audio     string      `json:"audio"`
				Thumbnail *string     `json:"thumbnail"`
			} `json:"videos"`
			TecConcursos    *string     `json:"tec_concursos"`
			Livestream      interface{} `json:"livestream"`
			LivestreamLink  interface{} `json:"livestream_link"`
			LivestreamData  interface{} `json:"livestream_data"`
			LivestreamSenha interface{} `json:"livestream_senha"`
			ConferenciaHora interface{} `json:"conferencia_hora"`
			ConferenciaLink interface{} `json:"conferencia_link"`
			Baixado         bool        `json:"baixado"`
		} `json:"aulas"`
		Nota                     int           `json:"nota"`
		PermiteForum             bool          `json:"permite_forum"`
		Discursivas              interface{}   `json:"discursivas"`
		AulasBaixadas            []int         `json:"aulas_baixadas"`
		AulasBaixadasHoje        []interface{} `json:"aulas_baixadas_hoje"`
		DownloadsRestantes       int           `json:"downloads_restantes"`
		PesquisaHabilitada       bool          `json:"pesquisa_habilitada"`
		Icone                    interface{}   `json:"icone"`
		ClienteLike              interface{}   `json:"cliente_like"`
		Likes                    int           `json:"likes"`
		Dislikes                 int           `json:"dislikes"`
		RaioX                    interface{}   `json:"raio_x"`
		MapaDaLei                interface{}   `json:"mapa_da_lei"`
		FuncionalidadeForum      bool          `json:"funcionalidade_forum"`
		FuncionalidadeMapaMental bool          `json:"funcionalidade_mapa_mental"`
		FuncionalidadePdfGrifado bool          `json:"funcionalidade_pdf_grifado"`
		FuncionalidadeResumo     bool          `json:"funcionalidade_resumo"`
		Modalidade               string        `json:"modalidade"`
		TotalAulas               int           `json:"total_aulas"`
		TotalAulasVisualizadas   int           `json:"total_aulas_visualizadas"`
		CertificadoPodeEmitir    bool          `json:"certificado_pode_emitir"`
		CertificadoJaEmitido     bool          `json:"certificado_ja_emitido"`
		CertificadoLink          string        `json:"certificado_link"`
		LdiUrl                   interface{}   `json:"ldi_url"`
	} `json:"data"`
}

func main() {
	var courses *CoursesResponse

	loadEnv()

	srcFolder := os.Getenv("RESOURCES_FOLDER")

	cli := http.Client{}
	req, err := http.NewRequest(http.MethodGet, CoursesUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&courses)
	if err != nil {
		log.Fatalf("Error trying to unmarshal json: %s", err.Error())
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, course := range courses.Data.Courses {
		req, err := http.NewRequest(http.MethodGet, CourseUrl+strconv.Itoa(course.Id), nil)
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+os.Getenv("TOKEN"))
		res, err := cli.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		cInfo := &CourseInfo{}
		err = json.NewDecoder(res.Body).Decode(&cInfo)
		if err != nil {
			log.Fatalf("Error trying to unmarshal json: %s", err.Error())
		}

		for i, class := range cInfo.Data.Classes {
			go func() {
				req, err := http.NewRequest(http.MethodGet, class.Pdf, nil)
				if err != nil {
					log.Fatal(err)
				}

				res, err := cli.Do(req)
				if err != nil {
					log.Fatal(err)
				}

				filename := filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name) + ".pdf"

				if _, err := os.Stat(filename); err == nil {
					filename = filepath.Join(cwd, srcFolder, cInfo.Data.Name, class.Name+"("+strconv.Itoa(i)+")") + ".pdf"
				}

				log.Printf("Downloading file: %s", filename)
				err = downloadFile(filename, res.Body)
				if err != nil {
					log.Fatalf("Error trying to save pdf: %s", err.Error())
				}

				res.Body.Close()
				time.Sleep(10 * time.Second)
			}()
		}

		time.Sleep(60 * time.Second)

		res.Body.Close()
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}

func downloadFile(path string, r io.ReadCloser) error {
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

	_, err = io.Copy(out, r)
	if err != nil {
		return err
	}

	return nil
}
