package main

import "fmt"

func main() {
	two()
}

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) { // note that type of `s` is the pointer of Student
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}

func one() {
	var s1 Student = Student{"Tim", []int{70, 90, 80, 85}, 19}
	var s2 Student = Student{"Joe", []int{80, 90, 80, 86, 59, 90, 93, 100}, 18}
	fmt.Println(s1.getAge()) // => 19
	s1.setAge(10)
	fmt.Println(s1.getAge()) // => 10

	average := s1.getAverageGrade()
	average2 := s2.getAverageGrade()
	fmt.Println(average)
	fmt.Println(average2)
}

func two() {
	var s1 Student = Student{"Tim", []int{70, 90, 80, 85}, 19}
	maxGrade := s1.getMaxGrade()
	fmt.Println(maxGrade)
}
