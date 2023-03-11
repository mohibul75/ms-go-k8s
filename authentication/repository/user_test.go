package repository

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/mohibul75/ms-go-k8s/authentication/models"
	"github.com/mohibul75/ms-go-k8s/db"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Panicln(err)
	}
}

func TestUsersRepositorySave(t *testing.T) {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	assert.NoError(t, err)
	defer conn.Close()

	id := bson.NewObjectId()

	user := &models.User{
		Id: id,
		Name: "TEST",
		Email: fmt.Sprintf("%s@email.test",id.Hex()),
		Password: "123456789",
		Created: time.Now(),
		Updated: time.Now(),
	}

	r := NewUsersRepository(conn)
	err = r.Save(user)
	assert.NoError(t, err)

	found, err := r.GetById(user.Id.Hex())
	assert.NoError(t, err)
	assert.Equal(t, user.Id, found.Id)
	assert.Equal(t,user.Name, found.Name)
	assert.Equal(t, user.Email, found.Email)
	assert.Equal(t, user.Password, found.Password)
}
