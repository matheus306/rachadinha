package service

import (
	config "rachadinha/internal/configuration"
	"rachadinha/internal/entity"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gorm.io/gorm"
)

func RecuperarGrupos(grupos *[]entity.Grupo) {
	db := config.Conectar()
	db.Preload("GrupoUsuario", func(tx *gorm.DB) *gorm.DB {
		return tx.Preload("Usuario")
	}).Find(&grupos)

}

func SalvarGrupo(grupo *entity.Grupo) {
	db := config.Conectar()
	db.Save(grupo)
}

func RecuperarUsuarioDoMes(grupo *entity.Grupo) (entity.Usuario, int) {

	var usuarioDoMes entity.Usuario
	var proximoSequencial int

	for _, grupoUsuario := range grupo.GrupoUsuario {
		if grupoUsuario.SequencialGrupo == grupo.SquencialAtual {
			usuarioDoMes = grupoUsuario.Usuario
			proximoSequencial = grupoUsuario.SequencialGrupo + 1
		}
	}

	if (entity.Usuario{}) == usuarioDoMes {
		sort.Slice(grupo.GrupoUsuario, func(i, j int) bool {
			return grupo.GrupoUsuario[i].SequencialGrupo < grupo.GrupoUsuario[j].SequencialGrupo
		})
		usuarioDoMes = grupo.GrupoUsuario[0].Usuario
		proximoSequencial = grupo.GrupoUsuario[0].SequencialGrupo + 1
	}

	return usuarioDoMes, proximoSequencial
}

func EnviarEmail(usuario *entity.Usuario) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})

	if err != nil {
		panic(err)
	}

	svc := ses.New(sess)

	from := "matheus306@gmail.com"

	template := "GrupoNetflix"

	to := usuario.Email

	firstName := usuario.Nome

	data := "{ \"firstName\":\"" + firstName + "\"}"

	input := &ses.SendTemplatedEmailInput{
		Source:   &from,
		Template: &template,
		Destination: &ses.Destination{
			ToAddresses: []*string{&to},
		},
		TemplateData: &data,
	}

	svc.SendTemplatedEmail(input)
}
