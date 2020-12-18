package schedulers

import (
	"fmt"
	"time"
)

func TestJob() {
	t := time.Now()
	fmt.Println("Time's up! @", t.UTC())
}