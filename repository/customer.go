package repository

type Customer struct {
	CustomerId  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zip_code"`
	Status      int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(id int) (*Customer, error)
}

func (r CustomerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}

	sql := "SELECT customer_id, name, date_of_birth, city, zip_code, status from customers"
	err := r.db.Select(&customers, sql)

	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r CustomerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	sql := "SELECT customer_id, name, date_of_birth, city, zip_code, status from customers where customer_id=?"
	err := r.db.Get(&customer, sql, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
