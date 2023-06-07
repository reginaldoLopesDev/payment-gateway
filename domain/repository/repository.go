package entity

type TransactionRepository interface {
	Insert(id string, account float64, amount float64, status string, errorMessage string) error
}
