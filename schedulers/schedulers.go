package schedulers

import (
	"context"
	"fmt"
	"github.com/Kvikikpt/wow-traider/db"
	"github.com/carlescere/scheduler"
	"log"
	"time"
)

func testJob() {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())

	var greeting string
	err := db.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)
}

func InitSchedulers() {
	_, _ = scheduler.Every(60).Seconds().NotImmediately().Run(testJob)
}