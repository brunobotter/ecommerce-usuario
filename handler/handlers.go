package handler

import (
	"fmt"
	"net/http"

	"github.com/brunobotter/ecommerce-usuario/service"
	"github.com/brunobotter/ecommerce-usuario/vo"
	"github.com/gin-gonic/gin"
)

func CreateUsuarioHandler(ctx *gin.Context) {
	request := vo.CreateUsuarioRequest{}
	ctx.Bind(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := service.CreateUsuario(request)
	if err != nil {
		logger.Errorf("Error creating user: %v", err.Error())
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "Create user", user)
}

func DeleteUsuarioHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, vo.ErrParamIsRequired("id", "Param").Error())
		return
	}

	usuario, err := service.DeleteUsuario(id)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting usuario with id: %s not found", id))
		return
	}

	vo.SendSuccess(ctx, "delete usuario", usuario)
}

func ListUsuairoHandler(ctx *gin.Context) {
	usuarios, err := service.ListUsuarios()
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, "error listing usuarios!")
		return
	}

	vo.SendSuccess(ctx, "list-usuarios", usuarios)
}

func ShowUsuarioHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	logger.Debugf("id %s", id)
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, vo.ErrParamIsRequired("id", "Param").Error())
		return
	}

	usuario, err := service.ShowUsuario(id)
	if err != nil {
		vo.SendError(ctx, http.StatusNotFound, "usuario não encontrado!")
		return
	}

	vo.SendSuccess(ctx, "exibindo-usuario", usuario)
}

func UpdateUsuarioHandler(ctx *gin.Context) {
	request := vo.UpdateUsuarioRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		vo.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id := ctx.Param("id")
	logger.Infof("id %s", id)
	if id == "" {
		vo.SendError(ctx, http.StatusBadRequest, vo.ErrParamIsRequired("id", "string").Error())
		return
	}

	usuario, err := service.UpdateUsuario(id, request)
	if err != nil {
		vo.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	vo.SendSuccess(ctx, "update-usuario", usuario)
}
