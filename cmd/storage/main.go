package main

import (
	"fmt"
	"supamod/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("hello there!", *st)
}
