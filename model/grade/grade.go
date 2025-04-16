package grade

import "gradeControleApi/model/test"

type Grade struct {
	Id    int
	Score float64
	Test  test.Test
}
