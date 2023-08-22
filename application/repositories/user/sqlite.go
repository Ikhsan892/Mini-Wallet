package user

import (
	"context"
	"errors"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"wallet/application/entity"
	"wallet/application/models"
)

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

type UserSqliteRepository struct {
	conn *gorm.DB
}

func NewUserSqliteRepository(conn *gorm.DB) *UserSqliteRepository {
	return &UserSqliteRepository{conn}
}

func toAggregate(m models.User) entity.User {
	return entity.User{Id: m.Id, Token: m.Token}
}

func (u UserSqliteRepository) CreateUserAccount(ctx context.Context, userId string) (entity.User, error) {
	param := models.User{Id: userId, Token: ulid.Make().String()}

	if err := u.conn.Model(&models.User{}).WithContext(ctx).Create(&param).Error; err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return entity.User{}, ErrUserAlreadyRegistered
		}
		return entity.User{}, err
	}

	return toAggregate(param), nil
}

func (u UserSqliteRepository) IsTokenValid(ctx context.Context, token string) bool {
	if err := u.conn.Model(&models.User{}).WithContext(ctx).Where("token = ?", token).First(&models.User{}).Error; err != nil {
		return false
	}

	return true
}

func (u UserSqliteRepository) GetUserAccountFromToken(ctx context.Context, token string) (entity.User, error) {
	var data models.User

	if err := u.conn.Model(&models.User{}).WithContext(ctx).Where("token = ?", token).First(&data).Error; err != nil {
		return entity.User{}, err
	}

	return toAggregate(data), nil

}
