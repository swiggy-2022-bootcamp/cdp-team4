package domain

// data model of payment record which is going to store in database
type Payment struct {
	Id          string // unique identier of the record
	Amount      int16  // amount to be paid
	Currency    string // type of the currency
	Status      string // it can done, pending or failed
	OrderID     string // order associated to this payment
	UserID      string // user associated to this payment
	Method      string // type/mode of the payment such as debit card, credit card, etc
	Description string // description message
	VPA         string //  Virtual Payment Address (VPA) is a unique ID created
	//  by the user to send or receive money through UPI.
	Notes []string // extra/optional notes data to store
}

type PaymentMethod struct {
	Id      string   // unique identier of the record
	Method  []string // methods/modes of payment supported for the user
	Agree   string   // all methods are verified
	Comment string   // extra comments that case be stored
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
