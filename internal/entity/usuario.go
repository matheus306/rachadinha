package entity

type Usuario struct {
	ID    int
	Nome  string
	Email string
}

func (u *Usuario) TableName() string {
	return "prd.tb_usuario"
}
