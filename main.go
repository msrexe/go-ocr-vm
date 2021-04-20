package main

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/Nr90/imgsim"
)

func main() {
	imagePath := "test2.png"

	if isDifferent(imagePath) {
		strings, err := recognizeText(imagePath)
		if err != nil {
			log.Panic(err)
		}
		for _, value := range strings {
			fmt.Println(value)
		}
	}
}

func isDifferent(imagePath string) bool {
	imgfile, err := os.Open(imagePath)
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()
	img, err := png.Decode(imgfile)
	if err != nil {
		panic(err)
	}
	ahash := imgsim.AverageHash(img)
	lastHash, err := ioutil.ReadFile("lastHash")
	if err != nil {
		panic(err)
	}
	if string(lastHash) != "" {
		_ = ioutil.WriteFile("lasthash", []byte(ahash.String()), 0644)
		return !(ahash.String() == string(lastHash))
	}
	_ = ioutil.WriteFile("lasthash", []byte(ahash.String()), 0644)
	return true
}

func recognizeText(imagePath string) ([]string, error) {
	cmd := exec.Command("tesseract", imagePath, "-")
	text, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return strings.Split(string(text), " "), nil
}
