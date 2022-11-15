package parameters

type CreateComment struct {
	Message string `json:"message" validate:"required"`
	PhotoID int    `json:"photo_id" validate:"required"`
}

type UpdateComment struct {
	Message string `json:"message"`
}
