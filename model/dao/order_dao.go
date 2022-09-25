package dao

import (
	"fmt"
	db "home-work-z-api/database"
	"time"
)

type Order struct {
	Id           int64 `gorm:"<-:create"`
	Buy          bool
	Sell         bool
	Quantity     int64
	Market_price float64
	Limit_price  float64
	Date         time.Time
	Created_at   time.Time `gorm:"autoCreateTime;<-:create"`
	Updated_at   time.Time `gorm:"autoUpdateTime;<-:update"`
}

func GetOrder() []*Order {
	var order []*Order

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().Find(&order)
	}

	return order
}

func GetOrderByDate(starttime, endtime string) []*Order {
	var order []*Order

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().
			Where(`date >= ? AND date < ?`, starttime, endtime).
			Order("date asc").
			Find(&order)
	}

	return order
}

func GetOrderById(id int64) Order {
	var order Order

	if mariaDB, err := db.GetMariaDB(); err == nil {
		mariaDB.Debug().First(&order, id)
	}

	return order
}

func InsertOrder(data Order) (int64, error) {

	mariaDB, err := db.GetMariaDB()

	if err != nil {
		return data.Id, err
	}

	result := mariaDB.Debug().Create(&data)
	fmt.Printf("Insert %s RowsAffected: %d\n", "order", result.RowsAffected)
	return data.Id, result.Error
}

func UpdateOrder(id int64, data Order) error {

	mariaDB, err := db.GetMariaDB()

	if err != nil {
		return err
	}

	result := mariaDB.Debug().Model(&Order{}).Where("id = ?", id).Updates(&data)
	fmt.Printf("Update %s RowsAffected: %d\n", "order", result.RowsAffected)
	return result.Error
}

func DeleteOrder(id int64) error {

	mariaDB, err := db.GetMariaDB()

	if err != nil {
		return err
	}

	result := mariaDB.Debug().Where("id = ?", id).Delete(&Order{})
	fmt.Printf("Delete %s RowsAffected: %d\n", "order", result.RowsAffected)
	return result.Error
}
