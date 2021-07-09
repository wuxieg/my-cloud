package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct {
	id   int
	name string
}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	fmt.Println(Student{})
}
