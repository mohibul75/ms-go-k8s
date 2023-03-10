package repository

import (
	"github.com/mohibul75/ms-go-k8s/authentication/models"
	"github.com/mohibul75/ms-go-k8s/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface{

	Save(user *models.User) error
	GetById(id string) (user *models.User , err error)
	GetByGmail(email string) (user *models.User, err error)
	GetAll(email string)(users []*models.User, err error)
	Update(user *models.User) error
	Delete(id string) error
}

type usersRepository struct{
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository{
	return &usersRepository{c: conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error{
	return r.c.Insert(user)
}

func (r *usersRepository) GetById(id string) (user *models.User , err error){
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByGmail(email string) (user *models.User, err error){
	err = r.c.Find(bson.M{"email":email}).One(&user)
	return user, err
}
 
func (r *usersRepository) GetAll(email string)(users []*models.User, err error){
	err= r.c.Find(bson.M{}).One(&users)
	return users, err
}

func (r *usersRepository) Update(user *models.User) error{
	return r.c.UpdateId(user.Id, user)
}

func (r *usersRepository) Delete(id string) error{
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
