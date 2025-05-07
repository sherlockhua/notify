package entity

import (
	"time"
)

// PaymentReceipt 实体
type PaymentReceipt struct {
	rawData       string
	transactionID string
	productID     string
	purchaseDate  time.Time
	expiryDate    *time.Time
}

func NewPaymentReceipt(rawData string) *PaymentReceipt {
	return &PaymentReceipt{
		rawData: rawData,
	}
}

/*
func (r *PaymentReceipt) ParseFromAppleResponse(response AppleValidationResponse) error {
	if len(response.LatestReceiptInfo) == 0 {
		return ErrInvalidReceipt
	}

	latest := response.LatestReceiptInfo[0]

	expiresMs, err := strconv.ParseInt(latest.ExpiresDateMs, 10, 64)
	if err != nil {
		return err
	}
	expiry := time.Unix(0, expiresMs*int64(time.Millisecond))

	r.transactionID = latest.TransactionID
	r.productID = latest.ProductID
	r.expiryDate = &expiry

	return nil
}

func (r *PaymentReceipt) GetExpiryDate() (time.Time, error) {
	if r.expiryDate == nil {
		return time.Time{}, ErrInvalidReceipt
	}
	return *r.expiryDate, nil
}*/
