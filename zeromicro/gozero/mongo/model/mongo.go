package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"log"
	"time"
)

func initModel() UserModel {
	return NewUserModel("mongodb://root:123456@127.0.0.1:27017", "gozero_test2", "User")
}

func insert() {
	userModel := initModel()
	u := &User{
		ID:       primitive.ObjectID{},
		Username: "rain",
		Password: "123456",
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	err := userModel.Insert(context.Background(), u)
	if err != nil {
		panic(err)
	}
	log.Println("Insert result id:", u.ID)
}

func query() {
	userModel := initModel()
	newUser, err := userModel.FindOne(context.Background(), "672d9b3abe7e9ab144a616a9")
	if err != nil {
		panic(err)
	}
	log.Println("query result:", newUser)
}

func update() {
	userModel := initModel()
	oid, err := primitive.ObjectIDFromHex("672d9b3abe7e9ab144a616a9")
	if err != nil {
		log.Fatalln("ObjectIDFromHex error", err)
		return
	}
	u := &User{
		ID:       oid,
		Username: "mark",
		Password: "123456",
	}
	// 更新
	updateResult, err := userModel.Update(context.Background(), u)
	if err != nil {
		panic(err)
	}
	log.Println("update result:", updateResult.UpsertedCount)
}

func delete() {
	userModel := initModel()
	// 删除
	deleteCount, err := userModel.Delete(context.Background(), "672d9b3abe7e9ab144a616a9")
	if err != nil {
		log.Fatalln("Delete error", err)
		return
	}
	log.Println("deleteCount:", deleteCount)
}
