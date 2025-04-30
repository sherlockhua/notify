package entity

// ProductID 值对象
type ProductID string

// Product 值对象
type Product struct {
	id            ProductID
	name          string
	productType   ProductType
	autoRenewable bool
}

type ProductType string

const (
	TypeConsumable    ProductType = "consumable"
	TypeNonConsumable ProductType = "non_consumable"
	TypeSubscription  ProductType = "subscription"
)

func NewProduct(id ProductID, name string, productType ProductType, autoRenewable bool) Product {
	return Product{
		id:            id,
		name:          name,
		productType:   productType,
		autoRenewable: autoRenewable && productType == TypeSubscription,
	}
}

func (p Product) IsSubscription() bool {
	return p.productType == TypeSubscription
}

func (p Product) AutoRenewable() bool {
	return p.autoRenewable
}
