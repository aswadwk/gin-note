package dto

type TodoCreateDTO struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type TodoUpdateDTO struct {
	Title string `json:"title"`
}

type TodoUpdateByIDDTO struct {
	Title string `json:"title"`
}