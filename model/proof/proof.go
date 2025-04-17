package proof

import "gradeControleApi/model/discipline"

type Proof struct {
	Id         int
	Name       string
	Discipline discipline.Discipline
}
