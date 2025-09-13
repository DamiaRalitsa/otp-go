package usecases

import (
	"bookingtogo/internal/domain"
	"bookingtogo/internal/repositories"
	"bookingtogo/pkg/postgres"
	"fmt"
	"strings"
)

type NationalityUsecase struct {
	Repo *repositories.NationalityRepository
}

func NewNationalityUsecase() *NationalityUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := repositories.NewNationalityRepository(databaseHandler)
	return &NationalityUsecase{
		Repo: repo,
	}
}

func (uc *NationalityUsecase) GetAll() ([]domain.Nationality, error) {
	return uc.Repo.FetchAll()
}

func (uc *NationalityUsecase) GetByID(id int) (*domain.Nationality, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	nationality, err := uc.Repo.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return nationality, nil
}

func (uc *NationalityUsecase) Create(n *domain.Nationality) error {
	if n == nil {
		return fmt.Errorf("nationality cannot be nil")
	}

	// Validation
	name := strings.TrimSpace(n.Name)
	code := strings.TrimSpace(n.Code)

	if name == "" {
		return fmt.Errorf("nationality name cannot be empty")
	}

	if code == "" {
		return fmt.Errorf("nationality code cannot be empty")
	}

	if len(code) > 2 {
		return fmt.Errorf("nationality code must be maximum 2 characters")
	}

	// Set cleaned values
	n.Name = name
	n.Code = strings.ToUpper(code)

	return uc.Repo.Insert(n)
}

func (uc *NationalityUsecase) Update(n *domain.Nationality) error {
	if n == nil {
		return fmt.Errorf("nationality cannot be nil")
	}

	if n.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if nationality exists
	_, err := uc.Repo.FetchByID(n.ID)
	if err != nil {
		return fmt.Errorf("nationality not found: %v", err)
	}

	// Validation
	name := strings.TrimSpace(n.Name)
	code := strings.TrimSpace(n.Code)

	if name == "" {
		return fmt.Errorf("nationality name cannot be empty")
	}

	if code == "" {
		return fmt.Errorf("nationality code cannot be empty")
	}

	if len(code) > 2 {
		return fmt.Errorf("nationality code must be maximum 2 characters")
	}

	// Set cleaned values
	n.Name = name
	n.Code = strings.ToUpper(code)

	return uc.Repo.Update(n)
}

func (uc *NationalityUsecase) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if nationality exists
	_, err := uc.Repo.FetchByID(id)
	if err != nil {
		return fmt.Errorf("nationality not found: %v", err)
	}

	return uc.Repo.Delete(id)
}
