package service

import (
	"home-work-z-api/model/vo"
	"home-work-z-api/service/trade"
)

func GetOrderList(starttime, endtime, top string) []*vo.OrderVO {
	o := trade.GetOrderList(starttime, endtime, top)
	return o
}

func GetOrder(id int64) vo.OrderVO {
	o := trade.GetOrder(id)
	return o
}

func PostOrder(body vo.OrderPostVO) (vo.OrderVO, error) {
	o, err := trade.AddOrder(body)
	return o, err
}
