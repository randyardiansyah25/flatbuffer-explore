package object

type Request struct {
	DateStart string
	DateEnd   string
}

type ArchiveItem struct {
	Id                int64
	DateTrans         string
	TransactionAmount float64
	Description       string
	Status            int
}

type HistoryItem struct {
	Id           int64
	DateTrans    string
	Description  string
	DebetAmount  float64
	CreditAmount float64
	Balance      float64
}

type Status struct {
	Code    string
	Message string
}

type Response[T any] struct {
	Response Status
	Data     T
}
