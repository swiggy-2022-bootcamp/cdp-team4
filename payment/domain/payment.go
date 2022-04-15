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

type PaymentMethod struct {
	Id      string
	Method  []string
	Agree   string
	Comment string
}

type PaymentDynamoRepository interface {
	InsertPaymentRecord(Payment) (bool, error)
	FindPaymentRecordById(string) (*Payment, error)
	FindPaymentRecordByUserID(string) ([]Payment, error)
	UpdatePaymentRecord(string, string) (bool, error)
	UpdatePaymentMethods(string, string) (bool, error)
	DeletePaymentRecordByID(string) (bool, error)
	InsertPaymentMethod(PaymentMethod) (bool, error)
	GetPaymentMethods(string) ([]string, error)
}
