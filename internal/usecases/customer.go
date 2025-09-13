package usecases

import (
	"bookingtogo/internal/domain"
	"bookingtogo/internal/repositories"
	"bookingtogo/pkg/postgres"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type CustomerUsecase struct {
	Repo            *repositories.CustomerRepository
	NationalityRepo *repositories.NationalityRepository
}

func NewCustomerUsecase() *CustomerUsecase {
	databaseHandler := postgres.NewDatabase(postgres.DbDetails).CreateDatabaseHandler()
	repo := repositories.NewCustomerRepository(databaseHandler)
	nationalityRepository := repositories.NewNationalityRepository(databaseHandler)
	return &CustomerUsecase{
		Repo:            repo,
		NationalityRepo: nationalityRepository,
	}
}

func (uc *CustomerUsecase) GetAll() ([]domain.Customer, error) {
	return uc.Repo.FetchAll()
}

func (uc *CustomerUsecase) GetByID(id int) (*domain.Customer, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ID must be greater than 0")
	}

	customer, err := uc.Repo.FetchByID(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (uc *CustomerUsecase) Create(c *domain.Customer) error {
	if c == nil {
		return fmt.Errorf("customer cannot be nil")
	}

	// Validation
	if err := uc.validateCustomer(c); err != nil {
		return err
	}

	return uc.Repo.Insert(c)
}

func (uc *CustomerUsecase) Update(c *domain.Customer) error {
	if c == nil {
		return fmt.Errorf("customer cannot be nil")
	}

	if c.ID <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if customer exists
	_, err := uc.Repo.FetchByID(c.ID)
	if err != nil {
		return fmt.Errorf("customer not found: %v", err)
	}

	// Validation
	if err := uc.validateCustomer(c); err != nil {
		return err
	}

	return uc.Repo.Update(c)
}

func (uc *CustomerUsecase) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("ID must be greater than 0")
	}

	// Check if customer exists
	_, err := uc.Repo.FetchByID(id)
	if err != nil {
		return fmt.Errorf("customer not found: %v", err)
	}

	return uc.Repo.Delete(id)
}

func (uc *CustomerUsecase) validateCustomer(c *domain.Customer) error {
	// Validate and clean name
	name := strings.TrimSpace(c.Name)
	if name == "" {
		return fmt.Errorf("customer name cannot be empty")
	}
	c.Name = name

	// Validate nationality
	if c.NationalityID <= 0 {
		return fmt.Errorf("nationality ID must be greater than 0")
	}

	// Check if nationality exists
	_, err := uc.NationalityRepo.FetchByID(c.NationalityID)
	if err != nil {
		return fmt.Errorf("invalid nationality ID: %v", err)
	}

	// Validate date of birth
	if c.DOB.IsZero() {
		return fmt.Errorf("date of birth cannot be empty")
	}

	// Check if DOB is not in the future
	if c.DOB.After(time.Now()) {
		return fmt.Errorf("date of birth cannot be in the future")
	}

	// Check minimum age (assuming 0 years old minimum)
	if time.Since(c.DOB).Hours() < 0 {
		return fmt.Errorf("invalid date of birth")
	}

	// Validate phone number
	phoneNumber := strings.TrimSpace(c.PhoneNumber)
	if phoneNumber == "" {
		return fmt.Errorf("phone number cannot be empty")
	}

	// Basic phone number validation (adjust regex based on your requirements)
	phoneRegex := regexp.MustCompile(`^[\d\-\+\(\)\s]+$`)
	if !phoneRegex.MatchString(phoneNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	c.PhoneNumber = phoneNumber

	// Validate email
	email := strings.TrimSpace(c.Email)
	if email == "" {
		return fmt.Errorf("email cannot be empty")
	}

	// Basic email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	c.Email = strings.ToLower(email)

	return nil
}
