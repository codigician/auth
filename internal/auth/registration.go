package auth

type RegistrationInfo struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type User struct {
	ID             string
	Firstname      string
	Lastname       string
	Email          string
	HashedPassword string
}

type Repository interface {
	Save(u *User) error
}

type Registrator struct {
	repository Repository
}

func NewRegistrator(repo Repository) *Registrator {
	return &Registrator{
		repository: repo,
	}
}

func (r *Registrator) Register(ri RegistrationInfo) (*User, error) {
	u := User{
		ID:             "0",
		Firstname:      ri.Firstname,
		Lastname:       ri.Lastname,
		Email:          ri.Email,
		HashedPassword: ri.Password,
	}

	if err := r.repository.Save(&u); err != nil {
		return nil, err
	}

	return &u, nil
}
