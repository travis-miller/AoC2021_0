package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failed to open input.txt file: %v", err)
	}
	defer f.Close()

	var imgAlg [512]bool
	var imgInput [][]bool
	firstLine := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		if firstLine {
			for i, c := range scanner.Text() {
				if c == '#' {
					imgAlg[i] = true
				}
			}
			firstLine = false
			continue
		}
		var imgRow []bool
		for _, c := range scanner.Text() {
			imgRow = append(imgRow, c == '#')
		}

		imgInput = append(imgInput, imgRow)
	}

	img := image{center: imgInput}
	for i := 0; i < 50; i++ {
		img = padImage(img)
		img = refineImg(img, imgAlg)
	}

	litNum := 0
	for _, r := range img.center {
		for _, p := range r {
			if p {
				litNum++
			}
		}
	}
	fmt.Printf("There are %d lit pixels.\n", litNum)
}

type image struct {
	center     [][]bool
	background bool
}

func padImage(img image) image {
	var paddedImg [][]bool
	height, width := len(img.center), len(img.center[0])
	for y := 0; y < height+4; y++ {
		var paddedRow []bool
		for x := 0; x < width+4; x++ {
			if y < 2 || y-2 >= height || x < 2 || x-2 >= width {
				paddedRow = append(paddedRow, img.background)
			} else {
				paddedRow = append(paddedRow, img.center[y-2][x-2])
			}
		}
		paddedImg = append(paddedImg, paddedRow)
	}
	return image{paddedImg, img.background}
}

func refineImg(img image, imgAlg [512]bool) image {
	var refImg [][]bool
	for y := 1; y < len(img.center)-1; y++ {
		var row []bool
		for x := 1; x < len(img.center[0])-1; x++ {
			var refPixel [9]bool
			refPixel[8] = img.center[y-1][x-1]
			refPixel[7] = img.center[y-1][x]
			refPixel[6] = img.center[y-1][x+1]
			refPixel[5] = img.center[y][x-1]
			refPixel[4] = img.center[y][x]
			refPixel[3] = img.center[y][x+1]
			refPixel[2] = img.center[y+1][x-1]
			refPixel[1] = img.center[y+1][x]
			refPixel[0] = img.center[y+1][x+1]
			newPixel := imgAlg[convRefPixelToNum(refPixel)]
			row = append(row, newPixel)
		}
		refImg = append(refImg, row)
	}
	refBG := false
	if img.background {
		refBG = imgAlg[511]
	} else {
		refBG = imgAlg[0]
	}
	return image{refImg, refBG}
}

func convRefPixelToNum(pixel [9]bool) int {
	num := 0
	for i, n := range pixel {
		if n {
			num += int(math.Pow(2, float64(i)))
		}
	}
	return num
}

func printImg(img image) {
	var sb strings.Builder
	for _, r := range img.center {
		for _, p := range r {
			if p {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}
