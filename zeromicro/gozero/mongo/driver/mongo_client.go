package driver

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

func initModel() *mon.Model {
	return mon.MustNewModel("mongodb://root:123456@127.0.0.1:27017", "gozero_test", "User")
}

func insert() {
	conn := initModel()
	ctx := context.Background()
	u := &User{
		ID:       primitive.ObjectID{},
		Username: "rain",
		Password: "123456",
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	// insert one
	result, err := conn.InsertOne(ctx, u)
	if err != nil {
		panic(err)
	}
	log.Println("Insert result id:", result.InsertedID)
}
func query() {
	conn := initModel()
	ctx := context.Background()
	// 查询
	var newUser User
	oid, err := primitive.ObjectIDFromHex("672d8c553b14e31778ccfa92")
	err = conn.FindOne(ctx, &newUser, bson.M{"_id": oid})
	if err != nil {
		panic(err)
	}
	log.Println("query result:", newUser)
}
func update() {
	conn := initModel()
	ctx := context.Background()
	var newUser User
	// 更新
	newUser.Username = "mark"
	oid, err := primitive.ObjectIDFromHex("672d8c553b14e31778ccfa92")
	updateResult, err := conn.ReplaceOne(ctx, bson.M{"_id": oid}, newUser)
	if err != nil {
		panic(err)
	}
	log.Println("update result:", updateResult.UpsertedCount)

}
func delete() {
	conn := initModel()
	ctx := context.Background()
	// 删除
	oid, err := primitive.ObjectIDFromHex("672d8c553b14e31778ccfa92")
	deleteCount, err := conn.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		panic(err)
	}
	log.Println("deleteCount:", deleteCount)
}
func transaction() {
	conn := initModel()
	// 使用副本集的 MongoDB 才支持事务
	session, err := conn.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.Background())

	// 执行事务操作
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		ctx := context.Background()
		u := &User{
			ID:       primitive.ObjectID{},
			Username: "rain",
			Password: "123456",
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		}
		// insert one
		result, err := conn.InsertOne(ctx, u)
		if err != nil {
			return nil, err
		}
		log.Println("Insert result id:", result.InsertedID)
		if err != nil {
			return nil, err
		}
		u = &User{
			ID:       primitive.ObjectID{},
			Username: "jamel",
			Password: "123456",
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		}
		// insert one
		result, err = conn.InsertOne(ctx, u)
		if err != nil {
			return nil, err
		}
		log.Println("Insert result id:", result.InsertedID)
		return nil, nil
	}

	// 启动事务
	_, err = session.WithTransaction(context.Background(), callback)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	} else {
		fmt.Println("Transaction committed successfully")
	}
}
