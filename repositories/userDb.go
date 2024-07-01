package repositories

import "rest-api/models"

type UserDb struct {
	Db *Db
}

func NewUserDb(db *Db) *UserDb {
	return &UserDb{
		Db: db,
	}
}

// CreateUser creates a new user.
func (userDb *UserDb) CreateUser(user *models.User) error {
	return userDb.Db.Create(user).Error
}

// GetUser retrieves a user by their ID.
func (userDb *UserDb) GetUser(id int) (*models.User, error) {
	var user models.User
	if err := userDb.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (userDb *UserDb) CheckLogin(login *models.LoginResp) error {
	var user models.User
	if err := userDb.Db.Where("password = ?  AND email = ?", login.Password, login.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (userDb *UserDb) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := userDb.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
