package dtos

type Course struct {
	Data struct {
		Id           int    `json:"id"`
		Name         string `json:"nome"`
		StartDate    string `json:"data_inicio"`
		DataRetirada string `json:"data_retirada"`
		Type         string `json:"tipo"`
		Professors   []struct {
			Id    int    `json:"id"`
			Name  string `json:"nome"`
			Image string `json:"imagem"`
		} `json:"professores"`
		Classes []struct {
			Id              int         `json:"id"`
			Name            string      `json:"nome"`
			Content         string      `json:"conteudo"`
			IsAvailable     bool        `json:"is_disponivel"`
			Viewed          bool        `json:"visualizada"`
			PublicationDate string      `json:"data_publicacao"`
			Pdf             string      `json:"pdf"`
			MarkedPdfFunc   interface{} `json:"funcionalidade_pdf_grifado"`
			MarkedPdf       interface{} `json:"pdf_grifado"`
			SimplifiedPdf   *string     `json:"pdf_simplificado"`
			IsStudentDone   bool        `json:"is_aluno_finalizado"`
			PendingContent  int         `json:"conteudos_pendentes"`
			Videos          []struct {
				Id          int     `json:"id"`
				Title       string  `json:"titulo"`
				Watched     bool    `json:"visualizado"`
				Resumo      *string `json:"resumo"`
				Slide       *string `json:"slide"`
				MentalMap   *string `json:"mapa_mental"`
				Position    int     `json:"posicao"`
				Resolution  string  `json:"resolucao"`
				Resolutions struct {
					P  string `json:"720p"`
					P1 string `json:"480p"`
					P2 string `json:"360p"`
				} `json:"resolucoes"`
				Annotations interface{} `json:"anotacoes"`
				Audio       string      `json:"audio"`
				Thumbnail   *string     `json:"thumbnail"`
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
