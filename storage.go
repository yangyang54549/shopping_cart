package shopping_cart

type StorageTestDb struct {
	Db  interface{}
	Key string
}

type Storage interface {
	New(key interface{}, db interface{}) Storage
	Save(rowId string, data *ShoppingCart) error
	CleanOne(rowIds string) error
	Clean() error
	Get(rowId string) (*ShoppingCart, error)
	GetAll() ([]ShoppingCart, error)
}
