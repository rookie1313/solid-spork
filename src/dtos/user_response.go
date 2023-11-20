package dtos

type UserResponse struct {
	ID    uint   `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
