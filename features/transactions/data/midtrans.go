package data

import (
	"crypto/sha512"
	"encoding/hex"
)

type DataMidtrans struct {
	PaymentType       string            `json:"payment_type"`
	TransactionDetail TransactionDetail `json:"transaction_details"`
	BankTransfer      BankTransfer      `json:"bank_transfer"`
}

type TransactionDetail struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float64 `json:"gross_amount"`
}

type BankTransfer struct {
	Bank string `json:"bank"`
}

type MidtransResponse struct {
	StatusCode     string           `json:"status_code"`
	TransactionID  string           `json:"transaction_id"`
	VirtualAccount []VirtualAccount `json:"va_numbers"`
	ExpiredTime    string           `json:"expiry_time"`
}

type VirtualAccount struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

func CheckStatus(transactionStatus, fraudStatus string) string {
	var res string
	if transactionStatus == "capture" {
		if fraudStatus == "accept" {
			res = "Paid"
		}
	} else if transactionStatus == "settlement" {
		res = "Paid"
	} else if transactionStatus == "cancel" ||
		transactionStatus == "deny" ||
		transactionStatus == "expire" {
		res = "Failed"
	} else if transactionStatus == "pending" {
		res = "Pending"
	}
	return res
}

func CheckSignatureKey(SignatureKey, orderId, statusCode, grossAmount, serverKey string) bool {
	encrypt := sha512.New()
	data := orderId + statusCode + grossAmount + serverKey
	encrypt.Write([]byte(data))
	hash := encrypt.Sum(nil)
	str := hex.EncodeToString(hash)
	return SignatureKey == str
}
