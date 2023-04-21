package dto

type TodoCreateDTO struct {
	ActivityGroupID uint   `json:"activity_group_id"`
	Title string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool `json:"is_active"`
}

type TodoUpdateByIDDTO struct {
	Title string `json:"title"`
	Priority string `json:"priority"`
	IsActive bool `json:"is_active"`
}

type TodoDeleteByIDDTO struct {
	ID uint `json:"id"`
}

type TodoSearchDTO struct {
	ActivityGroupID string `param:"activity_group_id"`
}