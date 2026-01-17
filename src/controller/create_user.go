package controller

import (
	"fmt"

	rest_err "github.com/Paulinhoh/meu-primeiro-crud-go/src/configuration"
	"github.com/Paulinhoh/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds. error=%s", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
