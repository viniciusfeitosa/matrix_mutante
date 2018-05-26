package main

import (
	"flag"
	"github.com/viniciusfeitosa/matrix_mutante/db"
	"github.com/viniciusfeitosa/matrix_mutante/models"
	"os"
)

func main() {
	var numWorkers int
	db := &db.DB{Enable: true}
	flag.StringVar(&db.Address, "redis_address", os.Getenv("APP_RD_ADDRESS"), "Redis Address")
	flag.StringVar(&db.Auth, "redis_auth", os.Getenv("APP_RD_AUTH"), "Redis Auth")
	flag.StringVar(&db.DB, "redis_db_name", os.Getenv("APP_RD_DBNAME"), "Redis DB name")
	flag.IntVar(&db.MaxIdle, "redis_max_idle", 10, "Redis Max Idle")
	flag.IntVar(&db.MaxActive, "redis_max_active", 10, "Redis Max Active")
	flag.IntVar(&db.IdleTimeoutSecs, "redis_timeout", 60, "Redis timeout in seconds")
	flag.IntVar(&numWorkers, "num_workers", 1, "Number of workers to consume queue")
	flag.Parse()

	db.Pool = db.NewDBPool()

	go models.NewStats(db).JoinStatsWorker(numWorkers)

	a := new(app)
	a.Initialize(db)
	a.initializeRoutes()
	a.Run(":" + os.Getenv("PORT"))
}
