package logic

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	storedArray []float64
	mu          sync.RWMutex
)

type Result struct {
	NegSum      float64 `json:"neg_sum"`
	Multiplying float64 `json:"multiplying"`
}

func InitArray(size int) []float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]float64, size)
	for i := range arr {
		arr[i] = math.Round((r.Float64()*20-10)*10) / 10
	}
	mu.Lock()
	storedArray = arr
	mu.Unlock()
	return arr
}

func GetStoredArray() []float64 {
	mu.RLock()
	defer mu.RUnlock()
	return append([]float64(nil), storedArray...)
}

func ProcessArray() Result {
	mu.RLock()
	defer mu.RUnlock()

	negSum := 0.0
	mul := 1.0

	for _, n := range storedArray {
		if n < 0 {
			negSum += n
		}
		mul *= n
	}

	return Result{
		NegSum:      negSum,
		Multiplying: mul,
	}
}
