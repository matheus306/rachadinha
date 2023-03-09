package entity

type Grupo struct {
	ID             int
	Nome           string
	SquencialAtual int `gorm:"column:sequencial_atual"`
	GrupoUsuario   []*GrupoUsuario
}

func (u *Grupo) TableName() string {
	return "prd.tb_grupo"
}
