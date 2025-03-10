package todo

// объект для хранения информации о задаче пользователя
type Lists struct {
	ID          int    `json:"-"`
	Header      string `json:"header"`
	Description string `json:"desc"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
