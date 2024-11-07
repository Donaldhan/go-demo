package mysql

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"time"
)

func initModel() UserModel {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行将 dsn 中的 host，账号 密码配置正确
	dsn := "root:123456@tcp(127.0.0.1:3306)/gozero_test?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	_ = conn

	return NewUserModel(conn)
}

func mysqlInsert() {
	userModel := initModel()
	// 创建一个有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保请求完成后取消 context，释放资源
	// 插入数据
	user := &User{Name: "Alice2", Password: "123456", Type: 1, Gender: "男", Nickname: "Alice", Mobile: "123456789021", CreateAt: time.Now()}
	result, err := userModel.Insert(ctx, user)
	if err != nil {
		fmt.Println("Insert Error:", err)
	}
	id, err := result.LastInsertId()
	fmt.Println("Insert result id:", id)

}
func find() {
	userModel := initModel()
	// 创建一个有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保请求完成后取消 context，释放资源
	// 查询数据
	userResult, err := userModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Println("FindOne Error:", err)
	} else {
		fmt.Println("User FindOne:", userResult)
	}
	name := "Alice"
	userResult, err = userModel.FindOneByName(ctx, name)
	if err != nil {
		fmt.Println("FindOneByName Error:", err)
	} else {
		fmt.Println("User FindOneByName:", userResult)
	}
	mobile := "123456789021"
	userResult, err = userModel.FindOneByMobile(ctx, mobile)
	if err != nil {
		fmt.Println("FindOneByMobile Error:", err)
	} else {
		fmt.Println("User FindOneByMobile:", userResult)
	}
}
func update() {
	userModel := initModel()
	// 创建一个有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保请求完成后取消 context，释放资源
	// 查询数据
	userResult, err := userModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Println("FindOne Error:", err)
	} else {
		fmt.Println("User FindOne:", userResult)
	}
	userResult.Nickname = "AliceX"
	err = userModel.Update(ctx, userResult)
	if err != nil {
		fmt.Println("Update Error:", err)
	}
	userResultX, err := userModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Println("FindOne Error:", err)
	} else {
		fmt.Println("User FindOne:", userResultX)
	}

	userUpdate := &User{Id: userResult.Id, Name: "Alice2X"}
	err = userModel.Update(ctx, userUpdate)
	if err != nil {
		fmt.Println("Update Error:", err)
	}
	userResultX, err = userModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Println("FindOne Error:", err)
	} else {
		fmt.Println("User FindOne:", userResultX)
	}
}
func delete() {
	userModel := initModel()
	// 创建一个有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保请求完成后取消 context，释放资源
	// 查询数据
	userResult, err := userModel.FindOne(ctx, 1)
	if err != nil {
		fmt.Println("FindOne Error:", err)
	} else {
		fmt.Println("User FindOne:", userResult)
	}
	err = userModel.Delete(ctx, 1)
	if err != nil {
		fmt.Println("Delete Error:", err)
	}
}
