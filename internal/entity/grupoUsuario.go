package entity

type GrupoUsuario struct {
	ID              int
	SequencialGrupo int
	GrupoId         int
	UsuarioId       int
	Usuario         Usuario
}

func (u *GrupoUsuario) TableName() string {
	return "prd.tb_grupo_usuario"
}
