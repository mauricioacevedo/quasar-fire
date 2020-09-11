package controllers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/quasar-fire/services"
	"github.com/quasar-fire/structs"
)

// PostTopSecret retrieves location and message from a config input
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

// PostTopSecretSplit load configuration for one satellite
func PostTopSecretSplit(ctx echo.Context) error {

	var (
		secretRequest = structs.TopSecretRequestSplit{}
		satelliteName = ctx.Param("satellite_name")
	)

	if err := ctx.Bind(&secretRequest); err != nil {
		log.Printf("Theres an error with the tx: %s", err.Error())

		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	//TODO: scheme validation
	log.Println("[START] PostTopSecretSplit with..", secretRequest)

	res, err := services.TopSecretSplitPostService(satelliteName, secretRequest)

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

// TopSecretSplitReset leave with no distance and message to the topsecret_split satellites
func TopSecretSplitReset(ctx echo.Context) error {

	log.Println("[START] TopSecretSplitReset")

	res, err := services.TopSecretSplitResetService()

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

// GetTopSecretSplit get location and posible message
func GetTopSecretSplit(ctx echo.Context) error {

	log.Println("[START] GetTopSecretSplit..")

	res, err := services.TopSecretSplitGetService()

	if err != nil {
		return ctx.String(http.StatusNotFound, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
