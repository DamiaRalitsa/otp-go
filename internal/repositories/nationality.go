package repositories

import (
	"bookingtogo/internal/domain"
	"bookingtogo/pkg/postgres"
	"fmt"
	"log"
)

type NationalityRepository struct {
	db postgres.DatabaseHandlerFunc
}

func NewNationalityRepository(db postgres.DatabaseHandlerFunc) *NationalityRepository {
	return &NationalityRepository{db}
}

func (r *NationalityRepository) FetchAll() ([]domain.Nationality, error) {
	var results []domain.Nationality

	query := `
		SELECT
			nationality_id,
			nationality_name,
			nationality_code
		FROM Nationality
		ORDER BY nationality_name ASC
	`

	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving nationalities: %v\n", err)
		return []domain.Nationality{}, err
	}

	if results == nil {
		results = []domain.Nationality{}
	}

	return results, nil
}

func (r *NationalityRepository) FetchByID(id int) (*domain.Nationality, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	var results []domain.Nationality

	query := `
		SELECT
			nationality_id,
			nationality_name,
			nationality_code
		FROM Nationality
		WHERE nationality_id = $1
	`

	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving nationality by ID: %v\n", err)
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no nationality found with id: %d", id)
	}

	return &results[0], nil
}

func (r *NationalityRepository) Insert(n *domain.Nationality) error {
	if n == nil {
		return fmt.Errorf("nationality cannot be nil")
	}

	query := `
		INSERT INTO Nationality (
			nationality_id,
			nationality_name,
			nationality_code
		) VALUES ($1, $2, $3)
	`

	err := r.db(nil, true, query, n.ID, n.Name, n.Code)
	if err != nil {
		log.Printf("Error inserting nationality: %v\n", err)
		return err
	}

	return nil
}

func (r *NationalityRepository) Update(n *domain.Nationality) error {
	if n == nil {
		return fmt.Errorf("nationality cannot be nil")
	}

	if n.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		UPDATE Nationality
		SET nationality_name = $2, nationality_code = $3
		WHERE nationality_id = $1
	`

	err := r.db(nil, true, query, n.ID, n.Name, n.Code)
	if err != nil {
		log.Printf("Error updating nationality: %v\n", err)
		return err
	}

	return nil
}

func (r *NationalityRepository) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		DELETE FROM Nationality
		WHERE nationality_id = $1
	`

	err := r.db(nil, true, query, id)
	if err != nil {
		log.Printf("Error deleting nationality: %v\n", err)
		return err
	}

	return nil
}
