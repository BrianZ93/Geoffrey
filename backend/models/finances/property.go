// property.go
package models

import "time"

type PropertyUse int

const (
	Residential PropertyUse = iota
	Investment
	SecondHome
	VacationHome
)

type PropertyType int

const (
	House PropertyType = iota
	Townhouse
	Condo
	MultiUnit
	Commercial
	Land
)

type LeaseType int

const (
	MonthToMonth LeaseType = iota
	FixedPeriod
)

type Income struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Memo        string    `json:"memo"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

type Expense struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Memo        string    `json:"memo"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

type Address struct {
	ID           string `json:"id"`
	StreetNumber int    `json:"streetNumber"`
	StreetName   string `json:"streetName"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          int    `json:"zip"`
}

type Tenant struct {
	ID              string    `json:"id"`
	PropertyId      string    `json:"propertyId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Address         Address   `json:"address"`
	LeaseType       LeaseType `json:"leaseType"`
	LeaseStartDate  time.Time `json:"leaseStartDate"`
	LeaseEndDate    time.Time `json:"leaseEndDate"`
	Rent            float64   `json:"rent"`
	SecurityDeposit float64   `json:"securityDeposit"`
	OtherIncome     float64   `json:"otherIncome"`
}

type Property struct {
	ID            string       `json:"id"`
	PurchaseDate  time.Time    `json:"purchaseDate"`
	PurchasePrice float64      `json:"purchasePrice"`
	Address       Address      `json:"address"`
	PropertyUse   PropertyUse  `json:"propertyUse"`
	PropertyType  PropertyType `json:"propertyType"`
	Mortgages     []Mortgage   `json:"mortgages"`
	Tenants       []Tenant     `json:"tenants"`
	Income        []Income     `json:"income"`
	Expenses      []Expense    `json:"expenses"`
}
