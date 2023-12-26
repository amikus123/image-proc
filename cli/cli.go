package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/amikus1223/image-proc/imageprocessing"
)

func ParseFlags() (filenames []string, outputDir string) {
	var dir string
	var fileNamesS string

	flag.StringVar(&dir, "d", "generated", "Output dir")
	flag.StringVar(&fileNamesS, "f", "", "Images to compress seperated by commas")
	flag.Parse()

	fileNames := strings.Split(fileNamesS, ",")

	return fileNames, outputDir
}

func ProcessImage(inputPath, outputPath string) error {
	err := imageprocessing.CopyImage(inputPath, outputPath)
	if err != nil {
		return err
	}

	fmt.Println("Image successfully copied to", outputPath)
	return nil
}
