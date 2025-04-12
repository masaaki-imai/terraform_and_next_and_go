package database

import (
	"L-cart/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedOrganization creates a test organization if it doesn't exist
func SeedUser() (*models.User, error) {
	var user models.User

	// Check if test organization exists
	if err := DB.Where("email = ?", "test@example.com").First(&user).Error; err == nil {
		return &user, nil
	}

	password, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create new test organization
	user = models.User{
		Email:         "test@example.com",
		Password:      string(password),
		Name:          "Test User",
		CompanyName:   "Test Company",
		PhoneNumber:   1234567890,
		Address1:      "123 Test St",
		Address2:      "Apt 1",
		Address3:      "Anytown",
		PostCode1:     12345,
		LastTimeLogin: nil,
	}

	if err := DB.Create(&user).Error; err != nil {
		log.Printf("Failed to create test user: %v", err)
		return nil, err
	}

	return &user, nil
}
