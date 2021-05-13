package gorm

import (
	"sync"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"tasquest.com/server/application/security"
	"tasquest.com/server/commons"
)

var IsUsersGormRepositoryInstanced *sync.Once
var usersGormRepositoryInstance *UsersGormRepository

type UsersGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UsersGormRepository {
	IsUsersGormRepositoryInstanced.Do(func() {
		usersGormRepositoryInstance = &UsersGormRepository{db: db}
		usersGormRepositoryInstance.migrate()
	})

	return usersGormRepositoryInstance
}

func (repo UsersGormRepository) migrate() {
	err := usersGormRepositoryInstance.db.AutoMigrate(security.User{})

	if err != nil {
		commons.IrrecoverableFailure("Failed to run SQL Migrations for the Users Entity", err)
	}
}

func (repo UsersGormRepository) FindByID(id uuid.UUID) (security.User, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"id = ?": id})
}

func (repo UsersGormRepository) FindByEmail(email string) (security.User, error) {
	return repo.FindOneByFilter(commons.SqlFilter{"email = ?": email})
}

func (repo UsersGormRepository) FindOneByFilter(filter commons.SqlFilter) (security.User, error) {
	var user security.User

	tx := repo.db.First(&user, filter.ToFormattedQuery())

	if tx.Error != nil {
		return security.User{}, tx.Error
	}

	return user, nil
}

func (repo UsersGormRepository) FindAllByFilter(filter commons.SqlFilter) ([]security.User, error) {
	var users []security.User

	tx := repo.db.Find(&users, filter.ToFormattedQuery())

	if tx.Error != nil {
		return []security.User{}, tx.Error
	}

	return users, nil
}

func (repo UsersGormRepository) Save(user security.User) (security.User, error) {
	tx := repo.db.Create(&user)

	if tx.Error != nil {
		return security.User{}, tx.Error
	}

	return user, nil
}
