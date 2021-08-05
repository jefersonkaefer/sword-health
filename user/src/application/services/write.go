package services

import (
	"sword-health/user/application/data_model"
	"sword-health/user/application/dto"
	"sword-health/user/application/repositories"
	"sword-health/user/domain"
)

type WriteService struct {
	userRepository *repositories.UserRepository
}

func (WriteService) New(repository *repositories.UserRepository) *WriteService {
	return &WriteService{
		userRepository: repository,
	}
}

func (us *WriteService) Create(userRequest dto.UserCreateDTO) (user *data_model.User, err error) {

	userModel, err := domain.Create(
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Email,
		userRequest.Password,
		userRequest.Role,
	)

	return us.userRepository.Add(&userModel)

}
