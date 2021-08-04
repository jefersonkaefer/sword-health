package services

import (
	"sword-health/users/application/repositories"
	"sword-health/users/domain"
)

type ReadService struct {
	userRepository *repositories.UserRepository
}

func (ReadService) New(repository *repositories.UserRepository) *ReadService {
	return &ReadService{
		userRepository: repository,
	}
}

func (us *ReadService) FindByEmail(email string) *domain.UserModel {

	return us.userRepository.FindByEmail(email)

}

func (us *ReadService) FindOne(userId int) (*domain.UserModel, error) {

	return us.userRepository.FindOne(userId)
}
