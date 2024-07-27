package user_api

type CreateUserReq struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" validate:"required,eq=Administrator|eq=Driver|eq=Provider"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Methods related to the User struct, such as password hashing
