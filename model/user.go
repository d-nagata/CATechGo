package model

type (
	UserCreateRequest struct {
		Name string `json:"name"`
	}
	UserCreateResponse struct {
		Token string `json:"token"`
	}
	LoginRequest struct {
		Name string `json:"name"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}
	UserGetResponse struct {
		Name string `json:"name"`
	}
	UserUpdateRequest struct {
		Name string `json:"name"`
	}
)
