package scheamas

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nome      string
	Telefone  string
	Email     string
	Documento string
	Endereco  string
}

type UsuarioResponse struct {
	Id        uint   `json:"id"`
	Nome      string `json:"nome"`
	Telefone  string `json:"telefone"`
	Email     string `json:"email"`
	Documento string `json:"documento"`
	Endereco  string `json:"endereco"`
}

func ToUsuarioResponse(usuario Usuario) UsuarioResponse {
	return UsuarioResponse{
		Id:        usuario.ID,
		Nome:      usuario.Nome,
		Telefone:  usuario.Telefone,
		Email:     usuario.Email,
		Documento: usuario.Documento,
		Endereco:  usuario.Endereco,
	}
}
