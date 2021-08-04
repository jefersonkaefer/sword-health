package repositories

import (
	"fmt"
	"sword-health/users/application/data_model"
	"sword-health/users/domain"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type UserRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func (UserRepository) New(redis *redis.Client, db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:    db,
		redis: redis,
	}
}

func (r *UserRepository) Add(u *domain.UserModel) bool {

	r.db.
		Where(
			data_model.User{
				Email: u.GetEmail(),
			},
		).
		FirstOrCreate(u.GetDataModel())
	fmt.Println(u)
	return r.db.RowsAffected > 0
}

func (r *UserRepository) FindByEmail(email string) *domain.UserModel {

	var dataModel data_model.User

	r.db.Where(
		data_model.User{
			Email: email},
	).Find(&dataModel)

	return (domain.UserModel{}).Load(&dataModel)
}

func (r *UserRepository) FindOne(userId int) (task *domain.UserModel, err error) {

	var dataModel data_model.User

	err = r.db.Where(
		data_model.User{
			ID: uint(userId),
		},
	).Find(&dataModel).Error

	if err != nil {
		return task, err
	}

	return (domain.UserModel{}).Load(&dataModel), err
}
