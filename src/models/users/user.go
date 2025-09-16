package user

type UserRegister struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Pass  string `json:"pass" xml:"pass" form:"pass"`
	Email string `json:"email" xml:"email" form:"email"`
}

type User struct {
	Id        int    `json:"id" xml:"id" form:"id"`
	Name      string `json:"name" xml:"name" form:"name"`
	Email     string `json:"email" xml:"email" form:"email"`
	Pass      string `json:"pass" xml:"pass" form:"pass"`
	Create_at string `json:"create_at" xml:"create_at" form:"create_at"`
}

type UserLogin struct {
	Email string `json:"email" xml:"email" form:"email"`
	Pass  string `json:"pass" xml:"pass" form:"pass"`
}

type UserResponse struct {
	Id        int    `json:"id" xml:"id" form:"id"`
	Name      string `json:"name" xml:"name" form:"name"`
	Email     string `json:"email" xml:"email" form:"email"`
	Create_at string `json:"create_at" xml:"create_at" form:"create_at"`
}
