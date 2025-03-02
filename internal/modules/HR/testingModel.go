package hr

// `json:"" form:""`

type TUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
