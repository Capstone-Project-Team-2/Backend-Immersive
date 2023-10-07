package data

import "crypto/sha512"

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

func CheckSignatureKey(SignatureKey, orderId, statusCode, grossAmount, serverKey string) (string, bool) {
	encrypt := sha512.New()
	data := orderId + statusCode + grossAmount + serverKey
	encrypt.Write([]byte(data))
	hash := encrypt.Sum(nil)
	if SignatureKey == string(hash) {
		return string(hash), true
	}
	return string(hash), false
}

/*
{
"transaction_time": "2021-06-23 11:53:34",
  "transaction_status": "settlement",
  "transaction_id": "9aed5972-5b6a-401e-894b-a32c91ed1a3a",
  "status_message": "midtrans payment notification",
  "status_code": "200",
  "signature_key": "fe5f725ea770c451017e9d6300af72b830a668d2f7d5da9b778ec2c4f9177efe5127d492d9ddfbcf6806ea5cd7dc1a7337c674d6139026b28f49ad0ea1ce5107",
  "settlement_time": "2021-06-23 11:53:34",
  "payment_type": "bank_transfer",
  "payment_amounts": [],
  "order_id": "bri-va-01",
  "merchant_id": "G141532850",
  "gross_amount": "300000.00",
  "fraud_status": "accept",
  "currency": "IDR"
}

{
  "status_code": "201",
  "status_message": "Success, Bank Transfer transaction is created",
  "transaction_id": "eb80333f-2ea0-42ac-ab39-82e01eed1161",
  "order_id": "order108",
  "merchant_id": "G790509189",
  "gross_amount": "40000.00",
  "currency": "IDR",
  "payment_type": "bank_transfer",
  "transaction_time": "2023-10-05 19:09:17",
  "transaction_status": "pending",
  "fraud_status": "accept",
  "va_numbers": [
      {
          "bank": "bca",
          "va_number": "09189878
      }
  ],
  "expiry_time": "2023-10-06 19:09:17"
}


"payment_type": "bank_transfer",
  "transaction_details": {
      "order_id": "order-101",
      "gross_amount": 44000
  },
  "bank_transfer":{
      "bank": "bca"
  }
*/
