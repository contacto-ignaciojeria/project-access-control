package domain

type Customer struct {
	CustomerID
	Email    string
	Password string
}

type CustomerID struct {
	Value string
}