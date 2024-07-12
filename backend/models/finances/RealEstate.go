// debt.go
package models

import "time"

type DebtType int

const (
	Installment DebtType = iota
	Revolving
	Balloon
)

type PrincipleOnlyPayment struct {
	ID     string    `json:"id"`
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount"`
}

type EscrowBreakdown struct {
	ID              string  `json:"id"`
	PropertyTaxes   float64 `json:"propertyTaxes"`
	Interest        float64 `json:"interest"`
	AssociationDues float64 `json:"associationDues"`
}

type Mortgage struct {
	ID                    string                 `json:"id"`
	PropertyId            string                 `json:"propertyId"`
	DebtType              DebtType               `json:"debtType"`
	OriginationDate       time.Time              `json:"originationDate"`
	FirstPaymentDate      time.Time              `json:"firstPaymentDate"`
	Maturity              time.Time              `json:"maturity"`
	InterestRate          float64                `json:"interestRate"`
	LienholderPosition    int                    `json:"lienholderPosition"`
	Servicer              string                 `json:"servicer"`
	Address               string                 `json:"address"`
	PrincipleOnlyPayments []PrincipleOnlyPayment `json:"principleOnlyPayments"`
	Term                  int                    `json:"term"`
	Escrow                bool                   `json:"escrow"`
	EscrowBreakdown       EscrowBreakdown        `json:"escrowBreakdown"`
}
