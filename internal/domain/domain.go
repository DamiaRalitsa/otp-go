package domain

import "time"

// =====================
// Nationality
// =====================
type Nationality struct {
	ID   int    `json:"id" db:"nationality_id"`     // nationality_id
	Name string `json:"name" db:"nationality_name"` // nationality_name
	Code string `json:"code" db:"nationality_code"` // nationality_code
}

// Nationality Usecase Interface
type NationalityUsecase interface {
	GetAll() ([]Nationality, error)
	GetByID(id int) (*Nationality, error)
	Create(n *Nationality) error
	Update(n *Nationality) error
	Delete(id int) error
}

// Nationality Repository Interface
type NationalityRepository interface {
	FetchAll() ([]Nationality, error)
	FetchByID(id int) (*Nationality, error)
	Insert(n *Nationality) error
	Update(n *Nationality) error
	Delete(id int) error
}

// =====================
// Customer
// =====================
type Customer struct {
	ID            int       `json:"id" db:"cst_id"`                         // cst_id
	NationalityID int       `json:"nationality_id" db:"nationality_id"`     
	Name          string    `json:"name" db:"cst_name"`                     // cst_name
	DOB           time.Time `json:"dob" db:"cst_dob"`                       // cst_dob
	PhoneNumber   string    `json:"phone_number" db:"cst_phonenum"`         // cst_phonenum
	Email         string    `json:"email" db:"cst_email"`                   // cst_email
}

// Customer Usecase Interface
type CustomerUsecase interface {
	GetAll() ([]Customer, error)
	GetByID(id int) (*Customer, error)
	Create(c *Customer) error
	Update(c *Customer) error
	Delete(id int) error
}

// Customer Repository Interface
type CustomerRepository interface {
	FetchAll() ([]Customer, error)
	FetchByID(id int) (*Customer, error)
	Insert(c *Customer) error
	Update(c *Customer) error
	Delete(id int) error
}

// =====================
// Family
// =====================
type Family struct {
	ID         int       `json:"id" db:"ft_id"`                   // ft_id
	CustomerID int       `json:"customer_id" db:"cst_id"`         // cst_id
	Relation   string    `json:"relation" db:"ft_relation"`       // ft_relation
	Name       string    `json:"name" db:"ft_name"`               // ft_name
	DOB        time.Time `json:"dob" db:"ft_dob"`                 // ft_dob
}

// Family Usecase Interface
type FamilyUsecase interface {
	GetAllByCustomer(customerID int) ([]Family, error)
	GetByID(id int) (*Family, error)
	Create(f *Family) error
	Update(f *Family) error
	Delete(id int) error
}

// Family Repository Interface
type FamilyRepository interface {
	FetchAllByCustomer(customerID int) ([]Family, error)
	FetchByID(id int) (*Family, error)
	Insert(f *Family) error
	Update(f *Family) error
	Delete(id int) error
}
