package main

import (
	"fmt"
)

type Student struct {
	name  string
	score int
}

func (s Student) String() string {
	return fmt.Sprintf("%s (%d)", s.name, s.score)
}

type Students []Student

func (ss Students) AvgScore() float64 {
	total := 0
	for _, s := range ss {
		total += s.score
	}
	return float64(total) / float64(len(ss))
}

func (ss Students) MinScore() Student {
	minScore := ss[0]
	for _, s := range ss {
		if s.score < minScore.score {
			minScore = s
		}
	}
	return minScore
}

func (ss Students) MaxScore() Student {
	maxScore := ss[0]
	for _, s := range ss {
		if s.score > maxScore.score {
			maxScore = s
		}
	}
	return maxScore
}

func main() {
	students := make(Students, 5)
	for i := 0; i < 5; i++ {
		var name string
		var score int
		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scanln(&score)
		students[i] = Student{name: name, score: score}
	}
	fmt.Printf("Average Score: %.2f\n", students.AvgScore())
	fmt.Printf("Min Score of Students: %s\n", students.MinScore())
	fmt.Printf("Max Score of Students: %s\n", students.MaxScore())
}
