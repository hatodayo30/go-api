package dto

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserGetResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	HighScore int32  `json:"highScore"`
	Coin      int32  `json:"coin"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}
