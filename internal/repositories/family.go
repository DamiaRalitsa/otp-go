package repositories

import (
	"fmt"
	"log"

	"bookingtogo/internal/domain"
	"bookingtogo/pkg/postgres"
)

type FamilyRepository struct {
	db postgres.DatabaseHandlerFunc
}

func NewFamilyRepository(db postgres.DatabaseHandlerFunc) *FamilyRepository {
	return &FamilyRepository{db}
}

func (r *FamilyRepository) FetchAllByCustomer(customerID int) ([]domain.Family, error) {
	if customerID <= 0 {
		return []domain.Family{}, fmt.Errorf("customer ID must be greater than 0")
	}

	var results []domain.Family

	query := `
		SELECT
			ft_id,
			cst_id,
			ft_relation,
			ft_name,
			ft_dob
		FROM family_list
		WHERE cst_id = $1
		ORDER BY ft_name ASC
	`

	err := r.db(&results, false, query, customerID)
	if err != nil {
		log.Printf("Error retrieving family members for customer %d: %v\n", customerID, err)
		return []domain.Family{}, err
	}

	if results == nil {
		results = []domain.Family{}
	}

	return results, nil
}

func (r *FamilyRepository) FetchByID(id int) (*domain.Family, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	var results []domain.Family

	query := `
		SELECT
			ft_id,
			cst_id,
			ft_relation,
			ft_name,
			ft_dob
		FROM family_list
		WHERE ft_id = $1
	`

	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving family member by ID: %v\n", err)
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no family member found with id: %d", id)
	}

	return &results[0], nil
}

func (r *FamilyRepository) Insert(f *domain.Family) error {
	if f == nil {
		return fmt.Errorf("family cannot be nil")
	}

	query := `
		INSERT INTO family_list (
			ft_id,
			cst_id,
			ft_relation,
			ft_name,
			ft_dob
		) VALUES ($1, $2, $3, $4, $5)
	`

	err := r.db(nil, true, query,
		f.ID,
		f.CustomerID,
		f.Relation,
		f.Name,
		f.DOB,
	)
	if err != nil {
		log.Printf("Error inserting family member: %v\n", err)
		return err
	}

	return nil
}

func (r *FamilyRepository) Update(f *domain.Family) error {
	if f == nil {
		return fmt.Errorf("family cannot be nil")
	}

	if f.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		UPDATE family_list
		SET cst_id = $2, ft_relation = $3, ft_name = $4, ft_dob = $5
		WHERE ft_id = $1
	`

	err := r.db(nil, true, query,
		f.ID,
		f.CustomerID,
		f.Relation,
		f.Name,
		f.DOB,
	)
	if err != nil {
		log.Printf("Error updating family member: %v\n", err)
		return err
	}

	return nil
}

func (r *FamilyRepository) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		DELETE FROM family_list
		WHERE ft_id = $1
	`

	err := r.db(nil, true, query, id)
	if err != nil {
		log.Printf("Error deleting family member: %v\n", err)
		return err
	}

	return nil
}
