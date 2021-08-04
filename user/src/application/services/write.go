package services

import (
	"sword-health/users/application/dto"
	"sword-health/users/application/repositories"
	"sword-health/users/domain"
)

type WriteService struct {
	userRepository *repositories.UserRepository
}

func (WriteService) New(repository *repositories.UserRepository) *WriteService {
	return &WriteService{
		userRepository: repository,
	}
}

func (us *WriteService) Create(user dto.UserCreateDTO) {

	userModel := domain.Create(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Role,
	)

	us.userRepository.Add(&userModel)

}
