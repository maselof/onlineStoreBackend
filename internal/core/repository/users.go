package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/gorm"
	"onlineStoreBackend/entity/response"
)

type UsersRepository struct{}

func (repo UsersRepository) PostUser(tx *gorm.DB, name string) (result response.UsersResponse, err error) {
	query, args, err := squirrel.
		Insert("public.users").
		Columns("name").
		Values(name).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo UsersRepository) DeleteUser(tx *gorm.DB, id int) (result response.UsersResponse, err error) {
	query, args, err := squirrel.
		Delete("public.users").
		Where("id = ?", id).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo UsersRepository) GetUsers(tx *gorm.DB) (result []response.UsersResponse, err error) {
	query, args, err := squirrel.
		Select("*").
		From("public.users").
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}

func (repo UsersRepository) GetUserByID(tx *gorm.DB, userID int) (result []response.UsersResponse, err error) {
	query, args, err := squirrel.
		Select("*").
		From("public.users").
		Where(squirrel.Eq{"id": userID}).
		ToSql()
	if err != nil {
		return result, err
	}

	return result, tx.Raw(query, args...).Scan(&result).Error
}
