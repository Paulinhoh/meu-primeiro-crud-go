package controller

import (
	rest_err "github.com/Paulinhoh/meu-primeiro-crud-go/src/configuration"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("voce chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
