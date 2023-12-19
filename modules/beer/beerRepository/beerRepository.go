package beerRepository

import (
	"github.com/gin-api/modules/beer"
	"github.com/gin-api/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type BeerRepository interface {
	CreateBeer(beer.BeerData) (*beer.Beer, error)
	GetBeer(string) ([]beer.Beer, error)
	GetBeerByID(int) (*beer.Beer, error)
	UpdateBeer(int, beer.BeerData) (*beer.Beer, error)
	DeleteBeer(int) (*int, error)
}

type beerRepositoryDB struct {
	db *sqlx.DB
}

func NewBeerRepositoryDB(db *sqlx.DB) beerRepositoryDB {
	return beerRepositoryDB{db: db}
}

func (r *beerRepositoryDB) CreateBeer(beerData beer.BeerData) (*beer.Beer, error) {
	query := "INSERT INTO beers (beer_name, beer_type, beer_description, beer_image) VALUES (:beer_name, :beer_type, :beer_description, :beer_image)"

	result, err := r.db.NamedExec(query, beerData)
	if err != nil {
		return nil, err
	}
	beerID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	beerRes := beer.Beer{
		BeerID:          int(beerID),
		BeerName:        beerData.BeerName,
		BeerType:        beerData.BeerType,
		BeerDescription: beerData.BeerDescription,
		BeerImage:       beerData.BeerImage,
	}

	return &beerRes, nil
}

func (r *beerRepositoryDB) GetBeer(name string) ([]beer.Beer, error) {
	var beers []beer.Beer
	query := "SELECT * FROM beers"
	if name != "" {
		query = "SELECT * FROM beers WHERE beer_name =?"
		err := r.db.Select(&beers, query, name)
		if err != nil {
			return nil, err
		}
	} else {
		err := r.db.Select(&beers, query)
		if err != nil {
			return nil, err
		}
	}

	return beers, nil
}

func (r *beerRepositoryDB) GetBeerByID(id int) (*beer.Beer, error) {
	var beerRes beer.Beer
	query := "SELECT * FROM beers WHERE beer_id = ?"
	err := r.db.Get(&beerRes, query, id)
	if err != nil {
		return nil, err
	}

	return &beerRes, nil
}
func (r *beerRepositoryDB) UpdateBeer(id int, beerData beer.BeerData) (*beer.Beer, error) {

	query := "UPDATE beers SET beer_name = ?, beer_type = ?, beer_description = ?, beer_image = ? WHERE beer_id = ?"

	result, err := r.db.Exec(query, beerData.BeerName, beerData.BeerType, beerData.BeerDescription, beerData.BeerImage, id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, utils.NewNotFoundError("beer not found")
	}

	beerRes := beer.Beer{
		BeerID:          id,
		BeerName:        beerData.BeerName,
		BeerType:        beerData.BeerType,
		BeerDescription: beerData.BeerDescription,
		BeerImage:       beerData.BeerImage,
	}

	return &beerRes, nil
}

func (r *beerRepositoryDB) DeleteBeer(id int) (*int, error) {
	query := "DELETE FROM beers WHERE beer_id = ?"
	res, err := r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	idRes, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	idInt := int(idRes)
	return &idInt, nil
}
