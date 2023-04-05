package entities

type Email struct {
	ID        string
	Email     string
	Create_at string
}

type SendEmail struct {
	Sender   string   `form:"sender"`
	Receiver []string `form:"receiver"`
	CC       []string `form:"cc"`
	Subject  string   `form:"subject"`
	Content  string   `form:"content"`
	Password string
}


type SendEmailRegister struct {
	Email string
}