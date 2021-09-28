package usecase

import (
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/internal/infrastructure/persistence"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/body"
	"github.com/rakutenshortintern2021-D-utopia/D-4_2/pkg/domain/model"
)

// GetUserList get a user by ID.
func GetUser(id uint64) model.User {
	userPersistence := persistence.NewUserPersistence()

	user, error := userPersistence.FindUserByID(id)

	if error != nil {
		panic(error)
	}
	return user
}

func UpdateUser(id uint64, requestBody body.PutUserRequest) model.User {
	userPersistence := persistence.NewUserPersistence()

	attributes := map[string]interface{}{
		"Name":     requestBody.Name,
		"Email":    requestBody.Email,
		"Password": requestBody.Password,
		"RoomID":   requestBody.RoomID,
	}

	user, error := userPersistence.UpdateUser(id, attributes)
	if error != nil {
		panic(error)
	}

	error = userPersistence.SaveTags(id, requestBody.TagIDs)
	if error != nil {
		panic(error)
	}

	return user
}

func AddUser(name string, email string, password string) model.User {
	userPersistence := persistence.NewUserPersistence()

	newUser := model.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	user, error := userPersistence.CreateUser(newUser)

	if error != nil {
		panic(error)
	}
	return user
}

func VerifyUser(email string, password string) model.User {
	userPersistence := persistence.NewUserPersistence()

	user, err := userPersistence.FindUserByUserInfo(email, password)
	if err != nil {
		panic(err)
	}

	return user
}
