package repositories

type UserDb struct {
	Db *Db
}

func NewUserDb(db *Db) *UserDb {
	return &UserDb{
		Db: db,
	}
}

// CreateUser creates a new user.
func (userDb *UserDb) CreateUser() {
	// Implement the logic for creating a user.
}

// GetUser retrieves a user by their ID.
func (userDb *UserDb) GetUser() {
	// Implement the logic for retrieving a user.
}
