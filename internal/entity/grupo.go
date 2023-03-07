package entity

type Grupo struct {
	ID      int
	Nome    string
	Usuario []*Usuario `gorm:"many2many:prd.tb_grupo_usuario;"`
}

func (u *Grupo) TableName() string {
	return "prd.tb_grupo"
}
