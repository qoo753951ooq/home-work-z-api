package trade

import (
	"home-work-z-api/model"
	"home-work-z-api/model/dao"
	"home-work-z-api/model/vo"
	"home-work-z-api/util"
	"strconv"
)

func GetOrderList(starttime, endtime, top string) []*vo.OrderVO {

	order := make([]*vo.OrderVO, 0)

	order = getOrderList(starttime, endtime)

	if top == model.Empty_string {
		return order
	}

	if count, err := strconv.Atoi(top); err == nil && len(order) != 0 {
		lastCount := util.GetTopLastLenByCount(count, len(order))
		order = order[model.Zero_value:lastCount]
	}

	return order
}

func GetOrder(id int64) vo.OrderVO {

	orderData := dao.GetOrderById(id)

	if orderData.Id == model.Zero_value {
		return vo.OrderVO{}
	}

	order := getOrder(orderData)
	return order
}

func getOrderList(starttime, endtime string) []*vo.OrderVO {

	order := make([]*vo.OrderVO, 0)

	var dbDatas []*dao.Order

	if starttime != model.Empty_string && endtime != model.Empty_string {
		dbDatas = dao.GetOrderByDate(starttime, endtime)
	} else {
		dbDatas = dao.GetOrder()
	}

	for _, orderData := range dbDatas {
		o := getOrder(*orderData)
		order = append(order, &o)
	}

	return order
}

func getOrder(data dao.Order) vo.OrderVO {

	var o vo.OrderVO

	o.Id = data.Id
	o.Buy = data.Buy
	o.Sell = data.Sell
	o.Quantity = data.Quantity
	o.Market_price = data.Market_price
	o.Limit_price = data.Limit_price
	o.Date = data.Date.Format(util.DateFormat)

	return o
}
