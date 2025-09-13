package usecases

import (
	"bookingtogo/internal/domain"
	"bookingtogo/internal/repositories"
	"bookingtogo/pkg/postgres"
	"fmt"
	"strings"
	"time"
)

type FamilyUsecase struct {
	Repo         *repositories.FamilyRepository
	CustomerRepo *repositories.CustomerRepository
}

func NewFamilyUsecase() *FamilyUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := repositories.NewFamilyRepository(databaseHandler)
	customerRepository := repositories.NewCustomerRepository(databaseHandler)
	return &FamilyUsecase{
		Repo:         repo,
		CustomerRepo: customerRepository,
	}
}

func (uc *FamilyUsecase) GetAllByCustomer(customerID int) ([]domain.Family, error) {
	if customerID <= 0 {
		return []domain.Family{}, fmt.Errorf("customer ID must be greater than 0")
	}

	// Check if customer exists
	_, err := uc.CustomerRepo.FetchByID(customerID)
	if err != nil {
		return []domain.Family{}, fmt.Errorf("customer not found: %v", err)
	}

	return uc.Repo.FetchAllByCustomer(customerID)
}

func (uc *FamilyUsecase) GetByID(id int) (*domain.Family, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	family, err := uc.Repo.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return family, nil
}

func (uc *FamilyUsecase) Create(f *domain.Family) error {
	if f == nil {
		return fmt.Errorf("family cannot be nil")
	}

	// Validation
	if err := uc.validateFamily(f); err != nil {
		return err
	}

	return uc.Repo.Insert(f)
}

func (uc *FamilyUsecase) Update(f *domain.Family) error {
	if f == nil {
		return fmt.Errorf("family cannot be nil")
	}

	if f.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if family member exists
	_, err := uc.Repo.FetchByID(f.ID)
	if err != nil {
		return fmt.Errorf("family member not found: %v", err)
	}

	// Validation
	if err := uc.validateFamily(f); err != nil {
		return err
	}

	return uc.Repo.Update(f)
}

func (uc *FamilyUsecase) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if family member exists
	_, err := uc.Repo.FetchByID(id)
	if err != nil {
		return fmt.Errorf("family member not found: %v", err)
	}

	return uc.Repo.Delete(id)
}

func (uc *FamilyUsecase) validateFamily(f *domain.Family) error {
	// Validate customer ID
	if f.CustomerID <= 0 {
		return fmt.Errorf("customer ID must be greater than 0")
	}

	// Check if customer exists
	_, err := uc.CustomerRepo.FetchByID(f.CustomerID)
	if err != nil {
		return fmt.Errorf("invalid customer ID: %v", err)
	}

	// Validate and clean relation
	relation := strings.TrimSpace(f.Relation)
	if relation == "" {
		return fmt.Errorf("family relation cannot be empty")
	}

	// Validate relation type (adjust based on your business rules)
	validRelations := map[string]bool{
		"spouse":  true,
		"child":   true,
		"parent":  true,
		"sibling": true,
		"other":   true,
	}

	relationLower := strings.ToLower(relation)
	if !validRelations[relationLower] {
		return fmt.Errorf("invalid family relation. Valid relations are: spouse, child, parent, sibling, other")
	}
	f.Relation = relationLower

	// Validate and clean name
	name := strings.TrimSpace(f.Name)
	if name == "" {
		return fmt.Errorf("family member name cannot be empty")
	}
	f.Name = name

	// Validate date of birth
	if f.DOB.IsZero() {
		return fmt.Errorf("date of birth cannot be empty")
	}

	// Check if DOB is not in the future
	if f.DOB.After(time.Now()) {
		return fmt.Errorf("date of birth cannot be in the future")
	}

	// Check minimum age (assuming 0 years old minimum)
	if time.Since(f.DOB).Hours() < 0 {
		return fmt.Errorf("invalid date of birth")
	}

	return nil
}
