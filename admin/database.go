package admin

type Database interface {
	FindAll(interface{}) error
	FindByID(interface{}, string) error
	Create(interface{}) error
	Update(interface{}) error
	Delete(interface{}, string) error
}
