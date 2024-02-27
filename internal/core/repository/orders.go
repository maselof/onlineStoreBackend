package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	"onlineStoreBackend/entity/dbmodel"
	"onlineStoreBackend/entity/response"
)

type OrdersRepository struct {
}

func (repo OrdersRepository) CreateOrder(tx *gorm.DB, userID int) (result dbmodel.Order, err error) {
	query, args, err := squirrel.
		Insert("public.orders").
		Columns("user_id").
		Values(userID).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo OrdersRepository) PostOrderItems(tx *gorm.DB, products []response.OrdersResponse) (result []response.OrdersResponse, err error) {
	selectBuilder := squirrel.
		Insert("public.order_items").
		Columns("order_id",
			"product_id",
			"amount")

	for _, v := range products {
		selectBuilder = selectBuilder.Values(
			v.ID,
			v.Products,
			v.Amount,
		)
	}

	query, args, err := selectBuilder.Suffix("RETURNING *").ToSql()

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo OrdersRepository) GetOrdersByUserID(tx *gorm.DB, userID int) (result []response.OrdersResponse, err error) {
	query, args, err := squirrel.
		Select("oi.order_id", "oi.product_id", "oi.amount").
		From("order_items oi").
		Join("public.orders o ON o.id = oi.order_id AND o.user_id = ?", userID).
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo OrdersRepository) DeleteOrderItems(tx *gorm.DB, orderID int) (result []response.OrdersResponse, err error) {
	query, args, err := squirrel.
		Delete("public.order_items").
		Where(squirrel.Eq{"order_id": orderID}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo OrdersRepository) DeleteOrder(tx *gorm.DB, orderID int) (err error) {
	query, args, err := squirrel.
		Delete("public.orders").
		Where(squirrel.Eq{"id": orderID}).
		ToSql()
	if err != nil {
		return err
	}

	return tx.Exec(query, args...).Error
}

func (repo OrdersRepository) DeleteOrderItemByItemID(tx *gorm.DB, itemID int) (err error) {
	query, args, err := squirrel.
		Delete("public.order_items").
		Where(squirrel.Eq{"product_id": itemID}).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return err
	}

	return tx.Exec(query, args...).Error
}
