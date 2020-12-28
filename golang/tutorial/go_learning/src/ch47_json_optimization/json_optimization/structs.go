package json_optimization

type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad       []int  `json:"pay_load"`
}

type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
