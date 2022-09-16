package http

type GetUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Id string `json:"id" validate:"required"`
}

type UserClient interface {
	GetUser(GetUserRequest) (*GetUserResponse, error)
}
