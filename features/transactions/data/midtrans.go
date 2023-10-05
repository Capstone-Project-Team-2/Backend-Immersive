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

/*
"payment_type": "bank_transfer",
  "transaction_details": {
      "order_id": "order-101",
      "gross_amount": 44000
  },
  "bank_transfer":{
      "bank": "bca"
  }
*/
