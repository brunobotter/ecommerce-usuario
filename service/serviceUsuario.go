package service

import (
	"github.com/brunobotter/ecommerce-usuario/config"
	"github.com/brunobotter/ecommerce-usuario/scheamas"
	"github.com/brunobotter/ecommerce-usuario/vo"
	"gorm.io/gorm"
)

// db é a instância do banco de dados global
var db *gorm.DB

func InitializeService(database *gorm.DB) {
	db = database
}

func CreateUsuario(request vo.CreateUsuarioRequest) (scheamas.Usuario, error) {
	logger := config.GetLogger("CreateUsuario")
	user := scheamas.Usuario{
		Nome:      request.Nome,
		Telefone:  request.Telefone,
		Email:     request.Email,
		Documento: request.Documento,
		Endereco:  request.Endereco,
	}

	if err := db.Create(&user).Error; err != nil {
		return scheamas.Usuario{}, err
	}
	logger.Debugf("usuario: %v", user)
	return user, nil
}

func DeleteUsuario(id string) (scheamas.Usuario, error) {
	usuario := scheamas.Usuario{}
	if err := db.First(&usuario, id).Error; err != nil {
		return scheamas.Usuario{}, err
	}
	if err := db.Delete(&usuario).Error; err != nil {
		return scheamas.Usuario{}, err
	}
	return usuario, nil
}

func ListUsuarios() ([]scheamas.Usuario, error) {
	usuarios := []scheamas.Usuario{}
	if err := db.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func ShowUsuario(id string) (scheamas.UsuarioResponse, error) {
	logger := config.GetLogger("ShowUsuarios")
	usuario := scheamas.Usuario{}
	if err := db.First(&usuario, id).Error; err != nil {
		return scheamas.UsuarioResponse{}, err
	}
	responseUser := scheamas.ToUsuarioResponse(usuario)
	logger.Debugf("usuario: %v", responseUser)
	return responseUser, nil
}

func UpdateUsuario(id string, request vo.UpdateUsuarioRequest) (scheamas.Usuario, error) {
	usuario := scheamas.Usuario{}
	if err := db.First(&usuario, id).Error; err != nil {
		return scheamas.Usuario{}, err
	}
	if request.Nome != "" {
		usuario.Nome = request.Nome
	}
	if request.Documento != "" {
		usuario.Documento = request.Documento
	}
	if request.Email != "" {
		usuario.Email = request.Email
	}
	if request.Endereco != "" {
		usuario.Endereco = request.Endereco
	}
	if request.Telefone != "" {
		usuario.Telefone = request.Telefone
	}
	if err := db.Save(&usuario).Error; err != nil {
		return scheamas.Usuario{}, err
	}
	return usuario, nil
}
