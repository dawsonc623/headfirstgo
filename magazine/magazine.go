package magazine

type Address struct {
	City       string
	PostalCode string
	State      string
	Street     string
}

type Employee struct {
	HomeAddress Address
	Name        string
	Salary      float64
}

type Subscriber struct {
	Active      bool
	HomeAddress Address
	Name        string
	Rate        float64
}
