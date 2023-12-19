package logs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Logs struct {
		Title        string    `json:"title"`
		Descriptions string    `json:"descriptions"`
		ErrorMessage string    `json:"errorMessage"`
		TimeStemp    time.Time `json:"timeStemp"`
	}

	LogsDB interface {
		WriteLog(Logs)
	}

	writeLogsDB struct {
		db *mongo.Database
	}
)

func NewwriteLogsDB(db *mongo.Database) writeLogsDB {
	return writeLogsDB{db: db}
}

func (l writeLogsDB) WriteLog(logs Logs) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := l.db.Collection("logs").InsertOne(ctx, logs)
	if err != nil {
		fmt.Println(err)
	}
}
