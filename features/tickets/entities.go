package tickets

type TicketCore struct {
	ID        string
	EventID   string
	NameClass string
	Total     uint
	Price     uint
}

type TicketDetailCore struct {
	ID            string
	TicketID      string
	TransactionID string
	UseStatus     bool
}
