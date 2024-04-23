package main

import (
	"fmt"
	"supamod/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("hello there!", *st)

	file, err := st.Upload("newfile.txt", []byte("it s a me, mario!"))
	if err != nil {
		panic(err)
	}

	restoredFile, err := st.FileById(file.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("file restored", restoredFile)
}
