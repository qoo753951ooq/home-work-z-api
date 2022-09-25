package trade

import (
	"fmt"
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

func AddOrder(data vo.OrderPostVO) (vo.OrderVO, error) {

	util.AddOrderMutex.Lock()

	order := newInsertOrder(data)
	orderCreateId, err := dao.InsertOrder(order)

	if err != nil {
		util.AddOrderMutex.Unlock()
		return vo.OrderVO{}, err
	}

	util.AddOrderMutex.Unlock()
	return vo.OrderVO{Id: orderCreateId}, nil
}

func EditOrder(id int64, data vo.OrderPutVO) (string, error) {

	if editNews := dao.GetOrderById(id); editNews.Id == model.Zero_value {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	order := newUpdateOrder(data)

	if err := dao.UpdateOrder(id, order); err != nil {
		return model.Empty_string, err
	}

	return model.Ok_string, nil
}

func DeleteOrder(id int64) (string, error) {

	if deleteNews := dao.GetOrderById(id); deleteNews.Id == model.Zero_value {
		return model.Empty_string, fmt.Errorf("%s\n", model.Id_not_found_string)
	}

	if err := dao.DeleteOrder(id); err != nil {
		return model.Empty_string, err
	}

	return model.Ok_string, nil
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

func newInsertOrder(data vo.OrderPostVO) dao.Order {

	var order dao.Order

	order.Quantity = data.Quantity

	if data.Buy != false {
		order.Buy = true
	} else {
		order.Buy = false
	}

	if data.Sell != false {
		order.Sell = true
	} else {
		order.Sell = false
	}

	if data.Market_price != model.Zero_value {
		order.Market_price = data.Market_price
	}

	if data.Limit_price != model.Zero_value {
		order.Limit_price = data.Limit_price
	}

	if date, err := util.GetLocationDate(data.Date); err == nil {
		order.Date = date
	}

	return order
}

func newUpdateOrder(data vo.OrderPutVO) dao.Order {

	var order dao.Order

	order.Quantity = data.Quantity

	if data.Buy != false {
		order.Buy = true
	} else {
		order.Buy = false
	}

	if data.Sell != false {
		order.Sell = true
	} else {
		order.Sell = false
	}

	if data.Market_price != model.Zero_value {
		order.Market_price = data.Market_price
	}

	if data.Limit_price != model.Zero_value {
		order.Limit_price = data.Limit_price
	}

	if date, err := util.GetLocationDate(data.Date); err == nil {
		order.Date = date
	}

	return order
}
