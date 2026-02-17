package producer

import (
	"fmt"
	"worker_pools_concurrency_prj/internal/core/domain"
)

func GenerateLog() {
	log := domain.LogTask{
		ID:      1,
		Level:   4,
		Message: "deneme",
	}
	fmt.Printf("log %+v", log)
}
