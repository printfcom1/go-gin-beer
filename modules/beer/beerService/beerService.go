package beerService

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-api/modules/beer"
	"github.com/gin-api/modules/beer/beerRepository"
	"github.com/gin-api/pkg/imageHandler"
	"github.com/gin-api/pkg/logs"
	"github.com/gin-api/pkg/utils"
)

type BeerService interface {
	CreateBeerService(beer.BeerData) (*beer.BeerResponse, error)
	GetBeerService(string) ([]beer.Beer, error)
	UpdateBeerService(string, beer.BeerData) (*beer.BeerResponse, error)
	DeleteBeerService(string) (*string, error)
}

type beerService struct {
	beerRepo beerRepository.BeerRepository
	logs     logs.LogsDB
}

func NewBeerService(beerRepo beerRepository.BeerRepository, logs logs.LogsDB) beerService {
	return beerService{beerRepo: beerRepo, logs: logs}
}

func (s *beerService) CreateBeerService(beerData beer.BeerData) (*beer.BeerResponse, error) {
	log := logs.Logs{
		Title:        "LogError",
		Descriptions: "Create beer failed",
		TimeStemp:    time.Now(),
		ErrorMessage: "",
	}

	beerDB, err := s.beerRepo.CreateBeer(beerData)
	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()

	}

	beerRes := beer.BeerResponse{
		BeerID:          beerDB.BeerID,
		BeerName:        beerDB.BeerName,
		BeerType:        beerDB.BeerType,
		BeerDescription: beerDB.BeerDescription,
		BeerImage:       beerDB.BeerImage,
	}

	log.Title = "LogSuccess"
	log.Descriptions = "Create beer success"
	s.logs.WriteLog(log)
	return &beerRes, nil
}

func (s *beerService) GetBeerService(name string) ([]beer.Beer, error) {
	beer, err := s.beerRepo.GetBeer(name)
	if err != nil {
		return nil, utils.NewInternalServerError()
	}
	return beer, nil
}

func (s *beerService) UpdateBeerService(id string, beerData beer.BeerData) (*beer.BeerResponse, error) {
	log := logs.Logs{
		Title:        "LogError",
		Descriptions: "Update beer failed",
		TimeStemp:    time.Now(),
		ErrorMessage: "",
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	getBeer, err := s.beerRepo.GetBeerByID(idInt)

	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	err = imageHandler.DeleteImage(getBeer.BeerImage)
	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	beerDB, err := s.beerRepo.UpdateBeer(idInt, beerData)
	if err != nil {
		if err.Error() == "beer not found" {
			log.ErrorMessage = err.Error()
			s.logs.WriteLog(log)
			return nil, err
		}
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	beerRes := beer.BeerResponse{
		BeerID:          beerDB.BeerID,
		BeerName:        beerDB.BeerName,
		BeerType:        beerDB.BeerType,
		BeerDescription: beerDB.BeerDescription,
		BeerImage:       beerDB.BeerImage,
	}

	log.Title = "LogSuccess"
	log.Descriptions = "Update beer success"
	s.logs.WriteLog(log)
	return &beerRes, nil
}

func (s *beerService) DeleteBeerService(id string) (*string, error) {
	log := logs.Logs{
		Title:        "LogError",
		Descriptions: "Delete beer failed",
		TimeStemp:    time.Now(),
		ErrorMessage: "",
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	idBeer, err := s.beerRepo.DeleteBeer(idInt)

	if *idBeer == 0 {
		log.ErrorMessage = "beer not found"
		s.logs.WriteLog(log)
		return nil, utils.NewNotFoundError("beer not found")
	}

	if err != nil {
		log.ErrorMessage = err.Error()
		s.logs.WriteLog(log)
		return nil, utils.NewInternalServerError()
	}

	log.Title = "LogSuccess"
	log.Descriptions = "Delete beer success"
	s.logs.WriteLog(log)

	messageRes := fmt.Sprintf("Delete beer id %d success", idInt)

	return &messageRes, nil
}
