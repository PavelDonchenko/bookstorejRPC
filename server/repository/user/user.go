package repository

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/PavelDonchenko/bookstorejRPC/models"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db   *gorm.DB
	user model.User
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

func (u *UserRepo) GetAll() ([]model.User, error) {
	users := []model.User{}
	err := u.db.Model(&model.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
func (u *UserRepo) GetOne(id uint32) (model.User, error) {
	err := u.db.Model(&model.User{}).Where("id = ?", id).Take(&u.user).Error
	if err != nil {
		return model.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return model.User{}, errors.New("User Not Found")
	}

	return u.user, nil
}
func (u *UserRepo) Create(user model.User) (model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, err
}
func (u *UserRepo) Update(id uint32) (model.User, error) {
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db := u.db.Model(&model.User{}).Where("id = ?", id).Take(&model.User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.user.Password,
			"nickname":   u.user.Nickname,
			"email":      u.user.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return model.User{}, err
	}

	err = db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&u.user).Error
	if err != nil {
		return model.User{}, err
	}

	return u.user, nil
}
func (u *UserRepo) Delete(id uint32) (bool, error) {
	db := u.db.Debug().Model(&model.User{}).Where("id = ?", id).Take(&model.User{}).Delete(&model.User{})
	result := true
	if db.Error != nil {
		result = false
		return result, db.Error
	}
	return result, nil
}
