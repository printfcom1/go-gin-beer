package main

import (
	"github.com/gin-api/modules/beer/beerHandler"
	"github.com/gin-api/modules/beer/beerRepository"
	"github.com/gin-api/modules/beer/beerService"
	"github.com/gin-api/pkg/database"
	"github.com/gin-api/pkg/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	logsDB := database.InitDataBaseMongoDB()
	logs := logs.NewwriteLogsDB(logsDB)

	BeerDB := database.InitDatabaseMariaDB()
	beerRepository := beerRepository.NewCostomerRepositoryDB(BeerDB)
	beerService := beerService.NewBeerService(&beerRepository, logs)
	beerHandler := beerHandler.NewBeerHandler(&beerService, logs)

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/beer", beerHandler.CreatedBeer)
	router.GET("/beer", beerHandler.GetBeer)
	router.PUT("/beer/:id", beerHandler.UpdateBeer)
	router.DELETE("/beer/:id", beerHandler.DeleteBeer)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
