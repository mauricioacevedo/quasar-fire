package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/quasar-fire/services"
	"github.com/quasar-fire/structs"
)

func PostTopSecret(ctx echo.Context) error {

	var (
		secretRequest = structs.TopSecretRequest{}
	)

	if err := ctx.Bind(&secretRequest); err != nil {
		log.Printf("Theres an error with the tx: %s", err.Error())

		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	log.Println("[START] PostTopSecret with..", secretRequest)

	res, err := services.TopSecretService(secretRequest)

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
