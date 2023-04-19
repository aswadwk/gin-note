package dto

type ActivityCreateDTO struct {
	Email string `json:"email"`
	Title string `json:"title"`
}

type ActivityUpdateByIDDTO struct {
	Email string `json:"email"`
	Title string `json:"title"`
}
