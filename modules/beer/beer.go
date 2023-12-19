package beer

import "time"

type Beer struct {
	BeerID          int       `json:"beer_id" db:"beer_id"`
	BeerName        string    `json:"beer_name" db:"beer_name"`
	BeerType        string    `json:"beer_type" db:"beer_type"`
	BeerDescription string    `json:"beer_description" db:"beer_description"`
	BeerImage       string    `json:"beer_image" db:"beer_image"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type BeerData struct {
	BeerName        string `json:"beer_name" db:"beer_name"`
	BeerType        string `json:"beer_type" db:"beer_type"`
	BeerDescription string `json:"beer_description" db:"beer_description"`
	BeerImage       string `json:"beer_image" db:"beer_image"`
}

type BeerResponse struct {
	BeerID          int    `json:"beer_id" db:"beer_id"`
	BeerName        string `json:"beer_name" db:"beer_name"`
	BeerType        string `json:"beer_type" db:"beer_type"`
	BeerDescription string `json:"beer_description" db:"beer_description"`
	BeerImage       string `json:"beer_image" db:"beer_image"`
}
