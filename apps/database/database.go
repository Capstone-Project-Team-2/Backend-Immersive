package database

import (
	"capstone-tickets/apps/config"
	adminModel "capstone-tickets/features/admins/data"
	buyerModel "capstone-tickets/features/buyers/data"
	eventModel "capstone-tickets/features/events/data"
	partnerModel "capstone-tickets/features/partners/data"
	refundModel "capstone-tickets/features/refund/data"
	transactionModel "capstone-tickets/features/transactions/data"
	volunteerModel "capstone-tickets/features/volunteers/data"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&adminModel.Admin{}, &buyerModel.Buyer{}, &partnerModel.Partner{},
		&eventModel.Event{}, &volunteerModel.Volunteer{}, &eventModel.Ticket{},
		&transactionModel.Transaction{}, &transactionModel.TicketDetail{}, &refundModel.Refund{},
		&transactionModel.PaymentMethod{})
}
