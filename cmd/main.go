package main

import (
	"fmt"
	"memogen/internal/service/image"
)

func main() {
	imageService := image.NewService()

	resultPath, err := imageService.DrawText(
		"/Users/almaz/GolandProjects/TestProjects/memogen/meme.jpg",
		"Опять",
		"Не опять а снова")

	if err != nil {
		panic(err)
	}

	fmt.Println(resultPath)
}
