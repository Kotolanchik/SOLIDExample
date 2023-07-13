package main

type UserService struct {
	repository UserRepository
}

func (u *UserService) RegisterUser(user User) error {
	if err := user.Validate(); err != nil {
		return err
	}

	if err := u.repository.Save(user); err != nil {
		return err
	}

	return nil
}

type UserRepository struct {
}

func (r UserRepository) Save(user User) error {
	return nil
}

type User struct {
}

func (u User) Validate() error {
	return nil
}
