package repositories

import (
	. "api/src/models"
	"encoding/json"

	"gopkg.in/validator.v2"
)

type UserRepository struct {
	Model User
	UserRepositoryInterface
}

type UserRepositoryInterface interface {
	All() []User
	Get(id uint64) (User, error)
	Create(params *json.Decoder) (User, error)
	Update(id uint64, params *json.Decoder) (User, error)
	Delete(id uint64) (bool, error)
}

func (r *UserRepository) All() []User {
	return r.Model.GetAll()
}

func (r *UserRepository) Get(id uint64) (User, error) {
	return r.Model.Get(id)
}

func (r *UserRepository) Create(params *json.Decoder) (User, error) {
	var user User

	err := params.Decode(&user)

	if err != nil {
		return user, err
	}

	err = validator.Validate(&user)

	if err != nil {
		return user, err
	}

	return r.Model.Create(user)
}

func (r *UserRepository) Update(id uint64, params *json.Decoder) (User, error) {
	var user User

	err := params.Decode(&user)

	if err != nil {
		return user, err
	}

	err = validator.Validate(&user)

	if err != nil {
		return user, err
	}

	return r.Model.Update(id, user)
}

func (r *UserRepository) Delete(id uint64) (bool, error) {
	return r.Model.Delete(id)
}
