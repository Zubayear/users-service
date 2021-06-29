package repo

import (
	"fmt"
	"log"
	"os"

	"github.com/Zubayear/bruce-almighty/external/entity"
	"github.com/Zubayear/bruce-almighty/external/entity/config"
	_ "github.com/go-sql-driver/mysql"
	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entity.User) (string, error)
	GetUser(id string) (string, error)
}

type database struct {
	dbConn *gorm.DB
}

func readYAML() *config.YAMLConfig {
	f, err := os.Open("application.yml")
	if err != nil {
		log.Fatalf("os.Open() failed with '%s'\n", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)

	var yamlFile config.YAMLConfig
	err = dec.Decode(&yamlFile)
	if err != nil {
		log.Fatalf("dec.Decode() failed with '%s'\n", err)
	}
	return &yamlFile
}

func NewUserRepository() UserRepository {
	y := readYAML()
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true", y.Databse.Username, y.Databse.Password, y.Databse.TableName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return &database{
		dbConn: db,
	}
}

func (d *database) CreateUser(user *entity.User) (string, error) {
	result := d.dbConn.Create(&user)
	if err := result.Error; err != nil {
		return "", err
	}
	return user.Name, nil
}

func (d *database) GetUser(id string) (string, error) {
	var user *entity.User
	result := d.dbConn.First(&user, id)
	if err := result.Error; err != nil {
		return "", err
	}
	return user.Name, nil
}
