package discipline

import "gradeControleApi/model/discipline"

type DisciplineRepository interface {
	Create(discipline discipline.Discipline) error
	FindAll() ([]discipline.Discipline, error)
	FindById(id int) (*discipline.Discipline, error)
	Update(discipline discipline.Discipline) (*discipline.Discipline, error)
	Delete(id int) error
}
