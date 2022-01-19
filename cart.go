package shopping_cart

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type ShoppingCart struct {
	RowId    string
	Id       uint32
	Spec     string
	Num      uint32
	IsSelect uint8
}

type Cart struct{}

var storage Storage

func New(s Storage) *Cart {
	storage = s
	return new(Cart)
}

func (t Cart) generateRowId(id uint32, spec string) string {
	str := strconv.Itoa(int(id)) + spec
	md5 := md5.New()
	md5.Write([]byte(str))
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func (t Cart) Save(data *ShoppingCart) error {
	rowId := t.generateRowId(data.Id, data.Spec)
	data.RowId = rowId
	return storage.Save(rowId, data)
}

func (t Cart) Clean() error {
	return storage.Clean()
}

func (t Cart) CleanOne(rowId string) error {
	return storage.CleanOne(rowId)
}

func (t Cart) Edit(rowId string, num uint32, isSelect uint8) error {
	data, err := storage.Get(rowId)
	if err != nil {
		return err
	}
	data.Num = num
	data.IsSelect = isSelect
	return storage.Save(rowId, data)
}

func (t Cart) Get(rowId string) (*ShoppingCart, error) {
	return storage.Get(rowId)
}

func (t Cart) GetAll() ([]ShoppingCart, error) {
	return storage.GetAll()
}
