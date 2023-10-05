package data

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
	VirtualAccount []VirtualAccount `json:"va_number"`
	ExpiredTime    string           `json:"expiry_time"`
}

type VirtualAccount struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

/*
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
          "va_number": "09189878747"
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
