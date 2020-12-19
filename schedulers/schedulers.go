package schedulers

import (
	"context"
	"fmt"
	"github.com/Kvikikpt/wow-traider/db"
	"log"
	"time"
)

func TestJob() {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())

	var greeting string
	err := db.Pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting)
}