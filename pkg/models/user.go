package models

import (
	"context"
	"go-crud-MongoDB/pkg/env"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func CreateUser(user User) error {
	collection := env.DB.Database("novye").Collection("users")

	// indexName, err := db.CreateUniqueIndex(collection, "login")
	// if err != nil {
	// 	fmt.Println("Couldn't create unique index")
	// 	return err
	// }
	// fmt.Printf("Unique index %s created\n", indexName)
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
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
