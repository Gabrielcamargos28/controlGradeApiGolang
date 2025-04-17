package grade

import "gradeControleApi/model/proof"

type Grade struct {
	Id    int
	Score float64
	Proof proof.Proof
}
