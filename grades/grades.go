package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

type GradeType int

const (
	Quiz GradeType = iota
	Test
	Project
)

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score
	}

	return result / float32(len(s.Grades))
}

type Students []Student

var (
	students Students
	mutex    sync.Mutex
)

func (s Students) GetByID(id int) (*Student, error) {
	for i, student := range s {
		if i == id {
			return &student, nil
		}
	}
	return nil, fmt.Errorf("no student with id: %v", id)
}
