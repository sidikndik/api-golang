package controller

import (
	"api/app/lib"
	"api/app/model"
	"api/app/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllUsers retrieves a list of all users
func GetAllUsers(c *fiber.Ctx) error {
    // Logic to retrieve all users from the database
    // Return the list of users as a JSON response
    // Anda perlu menentukan tipe data slice untuk menyimpan hasil
    db := services.DB.WithContext(c.UserContext())
    var users []model.User

    if err := db.Find(&users).Error; err != nil {
        return lib.ErrorInternal(c)
    }

    return lib.OK(c,users)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *fiber.Ctx) error {
    // Logic to retrieve a user by ID from the database
    // Return the user data as a JSON response
    db := services.DB.WithContext(c.UserContext())
    id, _:= uuid.Parse(c.Params("id"))

    var data model.User
    result := db.Model(&data).First(&data,"id = ?", id)

    if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

    return lib.OK(c, data)
}

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
    // Logic to create a new user in the database
    // Return the created user data as a JSON response

    api := new(model.User)
    if err := lib.BodyParser(c, api); nil != err{
        return lib.ErrorBadRequest(c, err)
    }

    db := services.DB.WithContext(c.UserContext())
    var data model.User
	lib.Merge(api, &data)
	data.ID = lib.GetXUserID(c)


	if err := db.Create(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

    return lib.Created(c)
}

// UpdateUser updates an existing user by ID
func UpdateUser(c *fiber.Ctx) error {
    // Logic to update an existing user by ID in the database
    // Return the updated user data as a JSON response
    api := new(model.User)
    if err := lib.BodyParser(c, api); nil != err{
        return lib.ErrorBadRequest(c, err)
    }

    db := services.DB.WithContext(c.UserContext())
	id, _ := uuid.Parse(c.Params("id"))

    var data model.User
    result := db.Model(&data).First(&data,"id = ?", id)

    if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

	lib.Merge(api, &data)
    if err := db.Model(&data).Updates(&data).Error; nil != err {
		return lib.ErrorConflict(c, err)
	}

    return lib.OK(c, data)
}

// DeleteUser deletes a user by ID
func DeleteUser(c *fiber.Ctx) error {
    // Logic to delete a user by ID from the database
    // Return a success response as a JSON response
    db := services.DB.WithContext(c.UserContext())
    var data model.User

    result1 := db.Model(&data).Where("id = ?", c.Params("id")).First(&data)
	if result1.RowsAffected < 1 {
		return lib.ErrorNotFound(c)
	}

    db.Delete(&data)
    
    return lib.OK(c)
}
