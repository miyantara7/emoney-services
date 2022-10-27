package biller

type Biller interface {
	ListBiller(interface{}) (interface{}, error)
	DetailBiller(interface{}) (interface{}, error)
}
