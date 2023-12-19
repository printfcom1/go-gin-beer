package beerHandler

import (
	"net/http"
	"time"

	"github.com/gin-api/modules/beer"
	"github.com/gin-api/modules/beer/beerService"
	"github.com/gin-api/pkg/imageHandler"
	"github.com/gin-api/pkg/logs"
	"github.com/gin-api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type beerHandler struct {
	beerHand beerService.BeerService
	logs     logs.LogsDB
}

func NewBeerHandler(beerHand beerService.BeerService, logs logs.LogsDB) beerHandler {
	return beerHandler{beerHand: beerHand, logs: logs}
}

func (h *beerHandler) CreatedBeer(c *gin.Context) {
	jsonDataStr := c.PostForm("data")
	beerData, err := utils.JsonConvert(jsonDataStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	fileName, err := imageHandler.CreateImage(beerData.BeerName, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Upload image failed"})
		log := logs.Logs{
			Title:        "LogError",
			Descriptions: "Upload image failed",
			ErrorMessage: err.Error(),
			TimeStemp:    time.Now(),
		}
		h.logs.WriteLog(log)
		return
	}

	beer := beer.BeerData{
		BeerName:        beerData.BeerName,
		BeerType:        beerData.BeerType,
		BeerDescription: beerData.BeerDescription,
		BeerImage:       *fileName,
	}

	beerRes, err := h.beerHand.CreateBeerService(beer)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, *beerRes)
}

func (h *beerHandler) GetBeer(c *gin.Context) {
	name := c.Query("name")
	beers, err := h.beerHand.GetBeerService(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, beers)
}

func (h *beerHandler) UpdateBeer(c *gin.Context) {
	id := c.Param("id")

	jsonDataStr := c.PostForm("data")
	beerData, err := utils.JsonConvert(jsonDataStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	fileName, err := imageHandler.CreateImage(beerData.BeerName, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Upload image failed"})
		log := logs.Logs{
			Title:        "LogError",
			Descriptions: "Upload image failed",
			ErrorMessage: err.Error(),
			TimeStemp:    time.Now(),
		}
		h.logs.WriteLog(log)
		return
	}

	beer := beer.BeerData{
		BeerName:        beerData.BeerName,
		BeerType:        beerData.BeerType,
		BeerDescription: beerData.BeerDescription,
		BeerImage:       *fileName,
	}

	beerRes, err := h.beerHand.UpdateBeerService(id, beer)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *beerRes)
}

func (h *beerHandler) DeleteBeer(c *gin.Context) {
	id := c.Param("id")
	res, err := h.beerHand.DeleteBeerService(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"message": *res})
}
