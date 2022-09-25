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

type OrderPostVO struct {
	Buy          bool    `form:"buy" default:"false" binding:"required"`
	Sell         bool    `form:"sell" default:"false" binding:"required"`
	Quantity     int64   `form:"quantity" default:"0" binding:"required"`
	Market_price float64 `form:"market_price" default:"0.0"`
	Limit_price  float64 `form:"limit_price" default:"0.0"`
	Date         string  `form:"date" default:"訂單日期(yyyy-MM-dd)" binding:"required"`
}

type OrderPutVO struct {
	Buy          bool    `json:"buy" example:"false"`
	Sell         bool    `json:"sell" example:"true"`
	Quantity     int64   `json:"quantity" example:"0" binding:"required"`
	Market_price float64 `json:"market_price" example:"0.0"`
	Limit_price  float64 `json:"limit_price" example:"0.0"`
	Date         string  `json:"date" example:"訂單日期(yyyy-MM-dd)" binding:"required"`
}
