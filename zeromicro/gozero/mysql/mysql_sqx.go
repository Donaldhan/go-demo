package mysql

import (
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"time"
)

func initConn() sqlx.SqlConn {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 需要自行将 dsn 中的 host，账号 密码配置正确
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test2?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	return conn
}
func insert() {
	conn := initConn()
	r, err := conn.ExecCtx(context.Background(), "insert into user (type, name) values (?, ?)", 1, "test")
	if err != nil {
		panic(err)
	}
	id, err := r.LastInsertId()
	fmt.Println("insert:", id)
}

type User struct {
	Id       int64          `db:"id"`
	Name     sql.NullString `db:"name"` // The username
	Type     int64          `db:"type"` // The user type, 0:normal,1:vip, for test golang keyword
	CreateAt sql.NullTime   `db:"create_at"`
	UpdateAt time.Time      `db:"update_at"`
}

func query() {
	conn := initConn()
	var u User
	query := "select id, name, type, create_at, update_at from user where id=?"
	err := conn.QueryRowCtx(context.Background(), &u, query, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("query:", u)
}

func update() {
	conn := initConn()
	_, err := conn.ExecCtx(context.Background(), "update user set type = ? where name = ?", 1, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
}
func delete() {
	conn := initConn()
	_, err := conn.ExecCtx(context.Background(), "delete from user where `id` = ?", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func transaction() {
	conn := initConn()
	err := conn.TransactCtx(context.Background(), func(ctx context.Context, session sqlx.Session) error {
		r, err := session.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 1, "test")
		if err != nil {
			return err
		}
		id, err := r.LastInsertId()
		fmt.Println("insert:", id)
		r, err = session.ExecCtx(ctx, "insert into user (id, name) values (?, ?)", 2, "test02")
		if err != nil {
			return err
		}
		id, err = r.LastInsertId()
		fmt.Println("insert:", id)
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
