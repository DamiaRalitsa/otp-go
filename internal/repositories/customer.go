package repositories

import (
	"bookingtogo/internal/domain"
	"bookingtogo/pkg/postgres"
	"fmt"
	"log"
)

type CustomerRepository struct {
	db postgres.DatabaseHandlerFunc
}

func NewCustomerRepository(db postgres.DatabaseHandlerFunc) *CustomerRepository {
	return &CustomerRepository{db}
}

func (r *CustomerRepository) FetchAll() ([]domain.Customer, error) {
	var results []domain.Customer

	query := `
		SELECT
			cst_id,
			nationality_id,
			cst_name,
			cst_dob,
			cst_phonenum,
			cst_email
		FROM Customer
		ORDER BY cst_name ASC
	`

	err := r.db(&results, false, query)
	if err != nil {
		log.Printf("Error retrieving customers: %v\n", err)
		return []domain.Customer{}, err
	}

	if results == nil {
		results = []domain.Customer{}
	}

	return results, nil
}

func (r *CustomerRepository) FetchByID(id int) (*domain.Customer, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	var results []domain.Customer

	query := `
		SELECT
			cst_id,
			nationality_id,
			cst_name,
			cst_dob,
			cst_phonenum,
			cst_email
		FROM Customer
		WHERE cst_id = $1
	`

	err := r.db(&results, false, query, id)
	if err != nil {
		log.Printf("Error retrieving customer by ID: %v\n", err)
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no customer found with id: %d", id)
	}

	return &results[0], nil
}

func (r *CustomerRepository) Insert(c *domain.Customer) error {
	if c == nil {
		return fmt.Errorf("customer cannot be nil")
	}

	query := `
		INSERT INTO Customer (
			cst_id,
			nationality_id,
			cst_name,
			cst_dob,
			cst_phonenum,
			cst_email
		) VALUES ($1, $2, $3, $4, $5, $6)
	`

	err := r.db(nil, true, query,
		c.ID,
		c.NationalityID,
		c.Name,
		c.DOB,
		c.PhoneNumber,
		c.Email,
	)
	if err != nil {
		log.Printf("Error inserting customer: %v\n", err)
		return err
	}

	return nil
}

func (r *CustomerRepository) Update(c *domain.Customer) error {
	if c == nil {
		return fmt.Errorf("customer cannot be nil")
	}

	if c.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		UPDATE Customer
		SET nationality_id = $2, cst_name = $3, cst_dob = $4, cst_phonenum = $5, cst_email = $6
		WHERE cst_id = $1
	`

	err := r.db(nil, true, query,
		c.ID,
		c.NationalityID,
		c.Name,
		c.DOB,
		c.PhoneNumber,
		c.Email,
	)
	if err != nil {
		log.Printf("Error updating customer: %v\n", err)
		return err
	}

	return nil
}

func (r *CustomerRepository) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	query := `
		DELETE FROM Customer
		WHERE cst_id = $1
	`

	err := r.db(nil, true, query, id)
	if err != nil {
		log.Printf("Error deleting customer: %v\n", err)
		return err
	}

	return nil
}
