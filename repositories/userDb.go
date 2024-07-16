package repositories

import (
	"rest-api/models"
	"time"
)

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
	user.CreatedAt = time.Now()
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

func (userDb *UserDb) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := userDb.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (userDb *UserDb) DeleteUser(id int) error {
	return userDb.Db.Delete(&models.User{}, id).Error
}

func (userDb *UserDb) UpdateUser(user *models.User) error {
	return userDb.Db.Save(user).Error
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

func (userDb *UserDb) SaveOrder(order *models.Order) error {
	return userDb.Db.Create(order).Error
}

func (userDb *UserDb) GetOrdersByUserID(userID int) ([]models.Order, error) {
	var orders []models.Order
	if err := userDb.Db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (userDb *UserDb) AddFavoriteRecipe(userID int, recipeID int64) error {
	// user has a field called FavoriteRecipes []int with the ids of the recipes
	user, err := userDb.GetUser(userID)
	if err != nil {
		return err
	}
	user.FavoriteRecipes = append(user.FavoriteRecipes, recipeID)
	return userDb.Db.Save(user).Error
}

func (userDb *UserDb) RemoveFavoriteRecipe(userID int, recipeID int64) error {
	// user has a field called FavoriteRecipes []int with the ids of the recipes
	user, err := userDb.GetUser(userID)
	if err != nil {
		return err
	}
	for i, id := range user.FavoriteRecipes {
		if id == recipeID {
			user.FavoriteRecipes = append(user.FavoriteRecipes[:i], user.FavoriteRecipes[i+1:]...)
			break
		}
	}
	return userDb.Db.Save(user).Error
}

func (userDb *UserDb) GetFavoriteRecipes(userID int) ([]models.Recipe, error) {
	var recipes []models.Recipe
	user, err := userDb.GetUser(userID)
	if err != nil {
		return nil, err
	}
	if len(user.FavoriteRecipes) == 0 {
		return nil, nil
	}
	for _, id := range user.FavoriteRecipes {
		recipe := &models.Recipe{}
		if err := userDb.Db.First(&recipe, id).Error; err != nil {
			continue
		}
		recipes = append(recipes, *recipe)
	}
	return recipes, nil
}

func (userDb *UserDb) AddReview(review *models.Review) error {
	//check if already has review on that recipe if so update it
	var existingReview models.Review
	if err := userDb.Db.Where("user_id = ? AND recipe_id = ?", review.UserID, review.RecipeID).First(&existingReview).Error; err != nil {
		return userDb.Db.Create(review).Error
	}
	existingReview.Rating = review.Rating
	existingReview.Comment = review.Comment
	return userDb.Db.Save(existingReview).Error
}

func (userDb *UserDb) GetReviewsByUserID(userID int) ([]models.Review, error) {
	var reviews []models.Review
	if err := userDb.Db.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (userDb *UserDb) UpdateReview(review *models.Review) error {
	return userDb.Db.Save(review).Error
}

func (userDb *UserDb) RemoveReview(reviewID int) error {
	return userDb.Db.Delete(&models.Review{}, reviewID).Error
}

func (userDb *UserDb) GetRecipes(userID int) ([]models.Recipe, error) {
	orders, err := userDb.GetOrdersByUserID(userID)
	if err != nil {
		return nil, err
	}
	recipes := []models.Recipe{}
	for _, order := range orders {
		recipe := &models.Recipe{}
		if order.RecipeID != 0 {
			if err := userDb.Db.First(&recipe, order.RecipeID).Error; err != nil {
				return nil, err
			}
		}
		recipe.Price = order.Price
		recipe.OrderStatus = order.Status
		recipes = append(recipes, *recipe)
	}

	return recipes, nil
}

func (userDb *UserDb) GetOrder(orderID int) (*models.Order, error) {
	var order models.Order
	if err := userDb.Db.First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (userDb *UserDb) RemoveOrder(userId, recipeId int) error {
	// Change the status of the order to Declined
	var order models.Order
	if err := userDb.Db.Where("user_id = ? AND recipe_id = ? AND status = ?", userId, recipeId, "In Progress").First(&order).Error; err != nil {
		return err
	}
	return userDb.Db.Model(&order).Update("status", "Declined").Error
}

func (userDb *UserDb) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := userDb.Db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
