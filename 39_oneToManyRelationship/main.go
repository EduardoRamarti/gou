package main

import "fmt"

type Course struct {
	Titulo string
	Videos []Video
}

type Video struct {
	Titulo string
	Course Course
}

func main() {

	video1 := Video{Titulo: "intro"}
	video2 := Video{Titulo: "install"}

	videos := []Video{video1, video2}

	curso := Course{Titulo: "Go Course", Videos: videos}

	video1.Course = curso
	video2.Course = curso

	fmt.Println(video1.Course.Titulo)

	for _, video := range curso.Videos {
		fmt.Println(video.Titulo)
	}
}
