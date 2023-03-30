package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog/log"
)

func FilePathDir(root string) ([]string, []uint64, []time.Time, error) {
	var files []string
	var sizes []uint64
	var times []time.Time
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			name := info.Name()
			size := uint64(info.Size())
			date := info.ModTime()

			files = append(files, name)
			sizes = append(sizes, size)
			times = append(times, date)
		}
		return nil
	})
	return files, sizes, times, err
}

func main() {
	var fileDir string

	log.Info().Msg("Insert file directory")

	fmt.Scanln(&fileDir)

	files, sizes, times, err := FilePathDir(fileDir)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	for _, file := range files {
		fmt.Println(file)
	}
	for _, size := range sizes {
		fmt.Println(humanize.Bytes(size))
	}
	for _, date := range times {
		fmt.Println(date)
	}

}
