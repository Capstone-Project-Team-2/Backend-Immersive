package data

import (
	"bytes"
	"capstone-tickets/apps/config"
	"capstone-tickets/features/events"
	_eventData "capstone-tickets/features/events/data"
	"capstone-tickets/features/transactions"
	"capstone-tickets/helpers"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type transactionQuery struct {
	db *gorm.DB
}

func New(database *gorm.DB) transactions.TransactionDataInterface {
	return &transactionQuery{
		db: database,
	}
}

// GetAllPaymentMethode implements transactions.TransactionDataInterface.
func (r *transactionQuery) GetAllPaymentMethod() ([]transactions.PaymentMethodCore, error) {
	var payment []PaymentMethod
	tx := r.db.Find(&payment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var paymenCore = ListPaymentMethodModelToCore(payment)
	return paymenCore, nil
}

// Insert implements transactions.TransactionDataInterface.
func (r *transactionQuery) Insert(data transactions.TransactionCore, buyer_id string) error {
	//1. memulai transaksi
	tx := r.db.Begin()
	//2. pengecekan stok tiket
	var transactionModel = TransactionCoreToModel(data)

	var event _eventData.Event
	errTx := tx.Where("id=?", data.EventID).First(&event).Error
	if errTx != nil {
		return errTx
	}

	if event.ValidationStatus != "Valid" {
		return errors.New("event belum divalidasi")
	}

	var count = map[string]int{}
	var paymentTotal float64
	for i, v := range transactionModel.TicketDetail {
		var errGen error
		transactionModel.TicketDetail[i].ID, errGen = helpers.GenerateUUID()
		if errGen != nil {
			return errGen
		}

		transactionModel.TicketDetail[i].BuyerID = buyer_id
		_, exist := count[v.TicketID]
		if !exist {
			count[v.TicketID] = 1
		} else {
			count[v.TicketID] += 1
		}
	}

	for key, v := range count {
		var ticket _eventData.Ticket
		//3.pengecekan waktu penjualan tiket
		err := tx.Where("id =? and sell_start <= NOW() and sell_end > NOW()", key).First(&ticket).Error
		if err != nil {
			tx.Rollback()
			return errors.New("tiket tidak tersedia " + err.Error())
		}

		//4. hitung total pembayaran dan stok tiket
		paymentTotal = paymentTotal + (float64(ticket.Price) * float64(v))
		if ticket.Total < uint(v) {
			tx.Rollback()
			return errors.New("tiket tidak mencukupi" + err.Error())
		}
	}

	uuid, err := helpers.GenerateUUID()
	if err != nil {
		tx.Rollback()
		return err
	}

	var payment PaymentMethod
	err = tx.Where("id=?", transactionModel.PaymentMethod).First(&payment).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	transactionModel.ID = uuid
	transactionModel.PaymentTotal = paymentTotal + payment.ServiceFee
	transactionModel.BuyerID = buyer_id

	var bank = BankTransfer{
		Bank: transactionModel.PaymentMethod,
	}
	var trans = TransactionDetail{
		OrderID:     transactionModel.ID,
		GrossAmount: transactionModel.PaymentTotal,
	}
	var midtrans = DataMidtrans{
		PaymentType:       "bank_transfer",
		TransactionDetail: trans,
		BankTransfer:      bank,
	}

	jsonData, errMars := json.Marshal(midtrans)
	if errMars != nil {
		tx.Rollback()
		return errMars
	}

	//5. kirim ke midtrans
	request, errReq := http.NewRequest("POST", "https://api.sandbox.midtrans.com/v2/charge", bytes.NewBuffer(jsonData))
	if errReq != nil {
		tx.Rollback()
		return errReq
	}

	serverKey := "Basic " + base64.StdEncoding.EncodeToString([]byte(config.KEY_SERVER))

	request.Header.Add("Authorization", serverKey)
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	response, errResp := client.Do(request)
	if errResp != nil {
		tx.Rollback()
		return errResp
	}

	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		tx.Rollback()
		return errRead
	}

	var midtransresp MidtransResponse
	//7. ambil response midtrans
	json.Unmarshal(body, &midtransresp)

	transactionModel.VirtualAccount = midtransresp.VirtualAccount[0].VANumber
	transactionModel.TimeLimit = helpers.ParseTimeMidtrans(midtransresp.ExpiredTime)

	for key, v := range count {
		var ticket _eventData.Ticket
		err := tx.Where("id=?", key).First(&ticket).Error
		if err != nil {
			tx.Rollback()
			return err
		}

		//8. pengurangan ticket
		err = tx.Model(&_eventData.Ticket{}).Where("id = ?", key).Update("total", ticket.Total-uint(v)).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	//9. simpan ke database
	err = tx.Create(&transactionModel).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//10. commit transaksi
	defer tx.Commit()
	return nil
}

// Select implements transactions.TransactionDataInterface.
func (r *transactionQuery) Select(transaction_id, buyer_id string) (transactions.TransactionCore, events.EventCore, error) {
	var transaction Transaction
	var event _eventData.Event

	tx := r.db.Where("id=?", transaction_id).Preload("Buyer").First(&transaction)
	if tx.Error != nil {
		return transactions.TransactionCore{}, events.EventCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return transactions.TransactionCore{}, events.EventCore{}, errors.New("no row affected")
	}

	if buyer_id != transaction.BuyerID {
		return transactions.TransactionCore{}, events.EventCore{}, errors.New("unauthorize")
	}

	txevent := r.db.Where("id = ?", transaction.EventID).First(&event)
	if txevent.Error != nil {
		return transactions.TransactionCore{}, events.EventCore{}, txevent.Error
	}
	if txevent.RowsAffected == 0 {
		return transactions.TransactionCore{}, events.EventCore{}, errors.New("no row affected")
	}

	dataTrans := TransactionModelToCore(transaction)
	dataEvent := _eventData.EventModelToCore(event)
	return dataTrans, dataEvent, nil
}

// Update implements transactions.TransactionDataInterface.
func (r *transactionQuery) Update(input transactions.MidtransCallbackCore) error {
	tx := r.db.Begin()
	var err error

	var sign = CheckSignatureKey(input.SignatureKey, input.OrderID, input.StatusCode, input.GrossAmount, config.KEY_SERVER)
	if !sign {
		return errors.New("signature unvalid")
	}

	var transaksiUpdate Transaction

	transaksiUpdate.OrderID = input.TransactionID
	transaksiUpdate.PaymentStatus = CheckStatus(input.TransactionStatus, input.FraudStatus)

	if transaksiUpdate.PaymentStatus == "Failed" {
		var trans Transaction
		var count = map[string]int{}
		err = tx.Where("id=?", input.OrderID).Preload("TicketDetail").First(&trans).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		for _, v := range trans.TicketDetail {
			_, check := count[v.ID]
			if !check {
				count[v.ID] = 1
			} else {
				count[v.ID] += 1
			}
		}
		for key, v := range count {
			var ticket _eventData.Ticket
			err = tx.Where("id = ?", key).First(&ticket).Error
			if err != nil {
				tx.Rollback()
				return err
			}
			err = tx.Model(&_eventData.Ticket{}).Where("id = ?", key).Update("total", ticket.Total+uint(v)).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	err = tx.Model(&Transaction{}).Where("id=?", input.OrderID).Updates(&transaksiUpdate).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	defer tx.Commit()
	return nil
}

// GetAllTicketDetail implements transactions.TransactionDataInterface.
func (r *transactionQuery) GetAllTicketDetail(buyer_id string) ([]transactions.TicketDetailCore, error) {
	var ticketDetailModel []TicketDetail
	tx := r.db.Where("buyer_id = ?", buyer_id).Find(&ticketDetailModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var ticketDetailCore = TicketDetailModelToCore(ticketDetailModel)
	return ticketDetailCore, nil
}
