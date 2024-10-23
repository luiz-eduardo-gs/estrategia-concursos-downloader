package dtos

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
