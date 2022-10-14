package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	model "github.com/PavelDonchenko/bookstorejRPC/server/models"
	"github.com/badoux/checkmail"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db   *gorm.DB
	user *model.User
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

var db *sql.DB

func (u *UserRepo) BeforeSave() error {
	hashedPassword, err := Hash(u.user.Password)
	if err != nil {
		return err
	}
	u.user.Password = string(hashedPassword)
	return nil
}

func (u *UserRepo) Prepare() {
	u.user.ID = 0
	u.user.Nickname = html.EscapeString(strings.TrimSpace(u.user.Nickname))
	u.user.Email = html.EscapeString(strings.TrimSpace(u.user.Email))
	u.user.CreatedAt = time.Now()
	u.user.UpdatedAt = time.Now()
}

func (u *UserRepo) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.user.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.user.Password == "" {
			return errors.New("Required Password")
		}
		if u.user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.user.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.user.Password == "" {
			return errors.New("Required Password")
		}
		if u.user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.user.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.user.Password == "" {
			return errors.New("Required Password")
		}
		if u.user.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.user.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *UserRepo) GetAllUsers(offset int, limit int) ([]model.User, error) {
	users := []model.User{}
	err := u.db.Model(&model.User{}).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (u *UserRepo) GetUser(id uint32) (*model.User, error) {
	result := &model.User{}
	err := u.db.Debug().Model(&model.User{}).Where("id = ?", id).Take(result).Error
	if err != nil {
		return &model.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &model.User{}, errors.New("User Not Found")
	}
	return result, err
}

func (u *UserRepo) CreateUser(user model.User) (model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

func (u *UserRepo) UpdateUser(user model.User) (*model.User, error) {
	//fmt.Println(user)
	//err := u.db.Model(&user).Where("email = ?", user.Email).Updates(&user).Error
	//fmt.Println(err)
	//return &user, err

	usertake := &model.User{}

	err := u.db.Debug().Model(&model.User{}).Where("email = ?", user.Email).Take(usertake).UpdateColumns(map[string]interface{}{
		"password":   user.Password,
		"nickname":   user.Nickname,
		"email":      user.Email,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		fmt.Println("erro: %v", err)
		return &model.User{}, err
	}
	return usertake, nil

	//u.db.First(user)
	//
	//usertake.Nickname = user.Nickname
	//usertake.Email = user.Email
	//usertake.Password = user.Password
	//u.db.Save(&usertake)
	//return usertake, nil
	//if err := u.db.First(&user).Error; err != nil {
	//	return user, err
	//}

	//result := &model.User{}
	//err := u.db.Debug().Model(&model.User{}).Where("id = ?", user.ID).Take(&result).Error
	//if err != nil {
	//	return &model.User{}, err
	//}
	//fmt.Println("take complete")
	//fmt.Println("update start ")

	//err := u.db.Model(&model.User{}).Updates(&user).Error
	//fmt.Println(err)
	//fmt.Println("update finished ")
	//return &user, err

	//fmt.Println("update start")
	//

	//err := u.db.Debug().Model(&model.User{}).Where("id = ?", user.ID).Take(&user).UpdateColumns(
	//	map[string]interface{}{
	//		"password":   u.user.Password,
	//		"nickname":   u.user.Nickname,
	//		"email":      u.user.Email,
	//		"updated_at": time.Now(),
	//	},
	//).Error
	//if err != nil {
	//	return &model.User{}, err
	//}
	//// This is the display the updated user
	//err = u.db.Debug().Model(&model.User{}).Where("id = ?", user.ID).Take(&user).Error
	//if err != nil {
	//	return &model.User{}, err
	//}
	//return &user, nil
}

func (u *UserRepo) DeleteUser(id uint32) (bool, error) {
	db := u.db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&model.User{}).Delete(&model.User{})
	result := true
	if db.Error != nil {
		result = false
		return result, db.Error
	}
	return result, nil
}
