package customer

var name = "customer"

func GetCustomer() string {
	return name
}

func Sum(result *int) int {
	a := 10
	b := 20
	*result = a + b
	return *result
}
