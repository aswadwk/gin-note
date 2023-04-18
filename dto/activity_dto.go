package dto

type ActivityCreateDTO struct {
	Email string `json:"email" binding:"required"`
	Title string `json:"title" form:"body" binding:"required=title cannot be null"`
}

type ActivityUpdateByIDDTO struct {
	Email string `json:"email"`
	Title string `json:"title" binding:"required"`
}
