package userapis

type CreateUserReq struct {
	ID         string
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	Role       string `json:"role" validate:"required,eq=Administrator|eq=Driver|eq=Provider"`
}

type LoginUserReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Methods related to the User struct, such as password hashing
