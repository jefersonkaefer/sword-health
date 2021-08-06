package repositories

import (
	"errors"
	"sword-health/user/application/data_model"
	"sword-health/user/domain"

	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
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

func (r *UserRepository) Add(u *domain.UserModel) (user *data_model.User, err error) {
	user = u.GetDataModel()

	err = r.db.Save(user).Error
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		err = errors.New("This user already exists.")
	}
	return user, err
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
