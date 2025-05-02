package accounts

type Account struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	IsActive  string `json:"is_active"`
	IsDeleted string `json:"is_deleted"`
}
