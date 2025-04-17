package discipline

import (
	"context"
	"fmt"
	"gradeControleApi/configuration/database"
	"gradeControleApi/model/discipline"
)

func NewDisciplineRepository() DisciplineRepository {
	return &DisciplineRepositoryImpl{}
}

type DisciplineRepositoryImpl struct {
}

func (r *DisciplineRepositoryImpl) Create(discipline discipline.Discipline) error {
	query := "INSERT INTO discipline (name) VALUES ($1)"

	_, err := database.DB.Exec(context.Background(), query, discipline.Name)

	if err != nil {
		return fmt.Errorf("failed to insert discipline:  %s", err)
	}
	fmt.Println("Disciplina criada com sucesso!")
	return nil
}

func (r *DisciplineRepositoryImpl) FindAll() ([]discipline.Discipline, error) {

	query := "SELECT id, name FROM discipline"
	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("Error in findAll: %w", err)
	}
	defer rows.Close()
	var disciplines []discipline.Discipline

	for rows.Next() {
		var d discipline.Discipline
		if err := rows.Scan(&d.Id, &d.Name); err != nil {
			return nil, err
		}
		disciplines = append(disciplines, d)
	}

	return disciplines, err
}

func (r *DisciplineRepositoryImpl) FindById(id int) (*discipline.Discipline, error) {
	query := "SELECT id, name FROM discipline d WHERE d.id = $1"

	row := database.DB.QueryRow(context.Background(), query, id)

	var d discipline.Discipline

	if err := row.Scan(&d.Id, &d.Name); err != nil {
		return nil, fmt.Errorf("discipline not found: %w", err)
	}
	return &d, nil
}

func (r *DisciplineRepositoryImpl) Update(discipline discipline.Discipline) (*discipline.Discipline, error) {
	query := "UPDATE discipline SET name = $1 WHERE id = $2"
	cmdTag, err := database.DB.Exec(context.Background(), query, discipline.Name, discipline.Id)

	if err != nil {
		return nil, fmt.Errorf("Error in update discipline: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, fmt.Errorf("No discipline found with id: %d", err)
	}
	return nil, nil
}

func (r *DisciplineRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM discipline WHERE id = $1"
	cmdTag, err := database.DB.Exec(context.Background(), query, id)

	if err != nil {
		return fmt.Errorf("failed to delete discipline: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("No discipline found with id %d", id)
	}
	return nil
}
