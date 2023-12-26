// imageprocessing/imageprocessing.go
package imageprocessing

import (
	_ "image/jpeg" // Import the JPEG package for image.Decode
	_ "image/png"  // Import the PNG package for image.Decode
	"io"
	"os"
)

// CopyImage copies the original image to a new file.
func CopyImage(inputPath, outputPath string) error {
	// Open the input image file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Copy the content of the input file to the output file
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}
