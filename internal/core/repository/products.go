package repository

import (
	"github.com/Masterminds/squirrel"
	"gorm.io/gorm"
	"onlineStoreBackend/entity/dbmodel"
	"onlineStoreBackend/entity/request"
)

type ProductsRepository struct {
}

func (repo ProductsRepository) PostProduct(tx *gorm.DB, request request.ProductRequest) (result dbmodel.Products, err error) {
	query, args, err := squirrel.
		Insert("public.products").
		Columns("name", "price").
		Values(request.Name, request.Price).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo ProductsRepository) GetAllProducts(tx *gorm.DB) (result []dbmodel.Products, err error) {
	query, args, err := squirrel.
		Select("*").
		From("public.products").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo ProductsRepository) GetProductsByIds(tx *gorm.DB, ids []int) (result []dbmodel.Products, err error) {
	query, args, err := squirrel.
		Select("*").
		From("public.products").
		Where(squirrel.Eq{"id": ids}).
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo ProductsRepository) DeleteProductByID(tx *gorm.DB, id int) (result dbmodel.Products, err error) {
	query, args, err := squirrel.
		Delete("public.products").
		Where(squirrel.Eq{"id": id}).
		Suffix("RETURNING *").
		ToSql()

	return result, tx.Raw(query, args...).Scan(&result).Error
}
