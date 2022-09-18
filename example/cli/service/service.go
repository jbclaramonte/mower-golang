package service

import (
	"log"
	"strconv"
	"strings"

	"github.com/jbclaramonte/mower-golang/core/domain"
)

func TranslateContentToCommands(fileContent string) (*domain.Garden, *domain.Mower, string) {
	lines := strings.Split(fileContent, "\n")
	log.Printf("lines: %v", lines)
	width, height := extractWidthAndHeight(strings.Trim(lines[0], " "))
	garden := domain.Garden{Width: width, Height: height}

	x, y, orientationStr := extractMowerPosition(strings.Trim(lines[1], " "))
	orientation, err := domain.GetOrientation(orientationStr)
	checkError(err)
	mower := domain.Mower{X: x, Y: y, Orientation: *orientation}

	commands := strings.Trim(lines[2], " ")

	return &garden, &mower, commands
}

func extractWidthAndHeight(line string) (width int, height int) {
	size := strings.Split(line, " ")
	width, _ = strconv.Atoi(size[0])
	height, _ = strconv.Atoi(size[1])
	return
}

func extractMowerPosition(line string) (x int, y int, orientation string) {
	split := strings.Split(line, " ")
	log.Printf("spliting %v to %v", line, split)
	x, _ = strconv.Atoi(split[0])
	y, _ = strconv.Atoi(split[1])
	orientation = split[2]
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
