package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type Line []image.Point

func main() {
	inputFile, _ := os.Open("input.txt")

	defer inputFile.Close()
	fileScanner := bufio.NewScanner(inputFile)

	var line string
	topLeft := image.Point{1000, 0}
	botRight := image.Point{0, 0}
	lines := make([]Line, 0)
	for fileScanner.Scan() {
		line = fileScanner.Text()
		strCoordinates := strings.Split(line, " -> ")
		rocks := make(Line, len(strCoordinates))

		for i, v := range strCoordinates {
			strX, strY, _ := strings.Cut(v, ",")
			x, _ := strconv.ParseInt(strX, 0, 0)
			y, _ := strconv.ParseInt(strY, 0, 0)
			rocks[i] = image.Point{int(x), int(y)}

			if x > int64(botRight.X) {
				botRight.X = int(x)
			}
			if x < int64(topLeft.X) {
				topLeft.X = int(x)
			}
			if y > int64(botRight.Y) {
				botRight.Y = int(y)
			}
			if y < int64(topLeft.Y) {
				topLeft.Y = int(y)
			}

		}
		lines = append(lines, rocks)

	}
	topLeft.X -= 5
	botRight.X += 5
	botRight.Y += 5
	fmt.Println(topLeft, botRight)

	drawImage(lines, topLeft, botRight)

}

func drawImage(lines []Line, topLeft, botRight image.Point) {

	img := image.NewRGBA(image.Rectangle{topLeft, botRight})

	// Draw rock segments

	for _, v := range lines {
		var p1, p2 image.Point
		for i := 0; i < len(v)-1; i++ {
			p1 = v[i]
			p2 = v[i+1]

			minX, minY := min(p1.X, p2.X), min(p1.Y, p2.Y)
			maxX, maxY := max(p1.X, p2.X), max(p1.Y, p2.Y)

			for x := minX; x <= maxX; x++ {
				for y := minY; y <= maxY; y++ {
					img.Set(x, y, color.Black)
				}
			}
		}
	}

	dropSand(img)

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func dropSand(img *image.RGBA) {
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	empty := color.RGBA{0, 0, 0, 0}
	sandCount := 0
	sandAbyss := false

	for {
		sand := image.Point{500, 0}

		for {
			img.Set(sand.X, sand.Y, red)

			if sand.Y-2 > img.Bounds().Max.Y {
				// sand fell off image
				sandAbyss = true
				break
			}

			if img.At(sand.X, sand.Y+1) == empty || img.At(sand.X, sand.Y+1) == red {
				sand.Y++
				continue
			} else if img.At(sand.X-1, sand.Y+1) == empty || img.At(sand.X-1, sand.Y+1) == red {
				sand.X--
				sand.Y++
				continue
			} else if img.At(sand.X+1, sand.Y+1) == empty || img.At(sand.X+1, sand.Y+1) == red {
				sand.X++
				sand.Y++
				continue
			}

			img.Set(sand.X, sand.Y, blue)
			sandCount++

			if sand.X == 500 && sand.Y == 0 {
				// sand fell off image
				sandAbyss = true
				break
			}

			break
		}

		if sandAbyss || sandCount > 60000 {
			break
		}

	}

	fmt.Println("Sand Count :", sandCount)
}
