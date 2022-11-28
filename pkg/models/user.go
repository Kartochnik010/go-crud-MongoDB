package models

import (
	"context"
	"errors"
	"go-crud-MongoDB/pkg/env"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Login      string             `json:"login,omitempty"`
	Password   string             `json:"password,omitempty"`
	Pfp        string             `json:"pfp,omitempty"`
	IsVerified bool               `json:"isVerified,omitempty"`
	Categories []string           `json:"categories,omitempty"`
	About      string             `json:"about,omitempty"`
	TimeStamp  time.Time          `json:"timestamp" bson:"timestamp"`
}

// for now simply checking for same values
// Couldn't figure out how to create a unique field in MongoDB
func (u *User) Validate() error {
	users, err := GetAllUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Login == u.Login {
			return errors.New("User already exists")
		}
	}
	return nil
}

func CreateUser(user User) (*mongo.InsertOneResult, error) {
	collection := env.DB.Database("novye").Collection("users")
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetAllUsers() ([]User, error) {
	collection := env.DB.Database("novye").Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		env.ErrorLog.Println(err)
		return []User{}, err
	}

	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		env.ErrorLog.Println(err)
		return []User{}, err
	}
	return users, nil
}

func GetUserById(query string) (User, error) {
	collection := env.DB.Database("novye").Collection("users")
	var user User
	id, err := primitive.ObjectIDFromHex(query)
	if err != nil {
		env.ErrorLog.Println(err)
		return user, err
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		env.ErrorLog.Println(err)
		return User{}, err
	}
	return user, nil
}
func GetUserByLogin(query string) (User, error) {
	collection := env.DB.Database("novye").Collection("users")
	var user User
	err := collection.FindOne(context.TODO(), bson.M{"login": query}).Decode(&user)
	if err != nil {
		env.ErrorLog.Println(err)
		return User{}, err
	}
	return user, nil
}

// drop user by id
func DeleteUserByLogin(query string) (int, error) {
	collection := env.DB.Database("novye").Collection("users")
	cnt, err := collection.DeleteOne(context.TODO(), bson.M{"login": query})
	return int(cnt.DeletedCount), err
}

func UpdateUserByLogin(query string, newUser User) *mongo.SingleResult {
	collection := env.DB.Database("novye").Collection("users")

	res := collection.FindOneAndUpdate(context.TODO(), bson.M{"login": query}, bson.M{"$set": newUser})
	return res
}
