package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/shaileshhb/restapi/app/domain/entity"
	"github.com/shaileshhb/restapi/app/domain/repository"
)

type Repositories struct {
	Repository *repository.GormRepository
	db         *gorm.DB
}

func DBConnenction(DBDriver, DBHost, DBPort, DBUser, DBName, DBPassword string) (*Repositories, error) {
	DBURL := fmt.Sprintf("driver=%s host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBDriver, DBHost, DBPort, DBUser, DBName, DBPassword)
	db, err := gorm.Open(DBDriver, getConnectionString(DBHost, DBPort, DBUser, DBName, DBPassword))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(DBURL + " successfully connected")

	return &Repositories{
		Repository: repository.NewGormRepository(),
		db:         db,
	}, nil
}

//This migrate all tables
func (repo *Repositories) Automigrate() error {

	return repo.db.AutoMigrate(
		&entity.User{}, &entity.Student{},
		&entity.Book{}, &entity.BookIssue{}).Error
}

//closes the  database connection
func (repo *Repositories) Close() error {
	return repo.db.Close()
}

func getConnectionString(DBHost, DBPort, DBUser, DBName, DBPassword string) (url string) {
	url += DBUser
	url += ":"
	url += DBPassword
	url += "@tcp("
	url += DBHost
	url += ":"
	url += DBPort
	url += ")/"
	url += DBName
	url += "?charset=utf8&parseTime=true"
	return
}
