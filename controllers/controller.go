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

func PostTopSecretSplit(ctx echo.Context) error {

	var (
		secretRequest = structs.TopSecretRequestSplit{}
		satelliteName = ctx.Param("satellite_name")
	)

	if err := ctx.Bind(&secretRequest); err != nil {
		log.Printf("Theres an error with the tx: %s", err.Error())

		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	log.Println("[START] PostTopSecretSplit with..", secretRequest)

	res, err := services.TopSecretSplitPostService(satelliteName, secretRequest)

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func TopSecretSplitReset(ctx echo.Context) error {

	log.Println("[START] TopSecretSplitReset")

	res, err := services.TopSecretSplitResetService()

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func GetTopSecretSplit(ctx echo.Context) error {

	log.Println("[START] GetTopSecretSplit..")

	res, err := services.TopSecretSplitGetService()

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
