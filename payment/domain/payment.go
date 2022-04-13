package domain

type Payment struct {
	Id          string
	Amount      int16
	Currency    string
	Status      string
	OrderID     string
	UserID      string
	Method      string
	Description string
	VPA         string
	Notes       []string
}

type PaymentDynamoRepository interface {
	Insert(Payment) (bool, error)
	FindById(string) (*Payment, error)
	FindByUserID(string) ([]*Payment, error)
	UpdateStatus(string, string) (*Payment, error)
	DeleteByID(string) (bool, error)
}
