package vo

type OrderVO struct {
	Id           int64   `json:"id"`
	Buy          bool    `json:"buy"`
	Sell         bool    `json:"sell"`
	Quantity     int64   `json:"quantity"`
	Market_price float64 `json:"market_price"`
	Limit_price  float64 `json:"limit_price"`
	Date         string  `json:"date"`
}
