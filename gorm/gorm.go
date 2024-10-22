package gorm

import (
	"errors"
	"fmt"
	"godemo/tools"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义User模型，绑定users表，ORM库操作数据库，需要定义一个struct类型和MYSQL表进行绑定或者叫映射，struct字段和MYSQL表字段一一对应
// 在这里User类型可以代表mysql users表
type User struct {
	ID int64 // 主键
	//通过在字段后面的标签说明，定义golang字段和表字段的关系
	//例如 `gorm:"column:username"` 标签说明含义是: Mysql表的列名（字段名)为username
	//这里golang定义的Username变量和MYSQL表字段username一样，他们的名字可以不一样。
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	CreateTime int64 `gorm:"column:createtime"`
}

// 设置表名，可以通过给struct类型定义 TableName函数，返回当前struct绑定的mysql表名是什么
func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}

// 商品
type Food struct {
	Id    int
	Title string
	Price float32
	Stock int
	Type  int
	//mysql datetime, date类型字段，可以和golang time.Time类型绑定， 详细说明请参考：gorm连接数据库章节。
	CreateTime time.Time
}

// 为Food绑定表名
func (v Food) TableName() string {
	return "foods"
}

type Order struct {
	Id     int64
	foodId int
	//mysql datetime, date类型字段，可以和golang time.Time类型绑定， 详细说明请参考：gorm连接数据库章节。
	CreateTime time.Time
}

// 为Food绑定表名
func (v Order) TableName() string {
	return "order"
}

func gormDemo() {
	db := tools.GetDB()
	if db == nil {
		log.Println("get db is nil")
		return
	}
	//create(db)
	//query(db)
	//update(db)
	delete(db)
	transactionAuto(db)
	transactionManual(db)
}
func transactionManual(db *gorm.DB) {
	//例子：
	food := Food{}
	// 开启事务
	tx := db.Begin()

	//在事务中执行数据库操作，使用的是tx变量，不是db。

	//库存减一
	//等价于: UPDATE `foods` SET `stock` = stock - 1  WHERE `foods`.`id` = '2' and stock > 0
	//RowsAffected用于返回sql执行后影响的行数
	rowsAffected := tx.Model(&food).Where("stock > 0").Update("stock", gorm.Expr("stock - 1")).RowsAffected
	if rowsAffected == 0 {
		//如果更新库存操作，返回影响行数为0，说明没有库存了，结束下单流程
		//这里回滚作用不大，因为前面没成功执行什么数据库更新操作，也没什么数据需要回滚。
		//这里就是举个例子，事务中可以执行多个sql语句，错误了可以回滚事务
		tx.Rollback()
		return
	}
	order := Order{}
	err := tx.Create(order).Error

	//保存订单失败，则回滚事务
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

// 自动事务
// 通过db.Transaction函数实现事务，如果闭包函数返回错误，则回滚事务。
func transactionAuto(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&Food{Title: "Giraffe"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&Food{Title: "Lion"}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}
func delete(db *gorm.DB) {
	//例子：
	food := Food{}
	//先查询一条记录, 保存在模型变量food
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	db.Where("id = ?", 2).Take(&food)

	//删除food对应的记录，通过主键Id标识记录
	//等价于： DELETE from `foods` where id=2;
	db.Delete(&food)

	//等价于：DELETE from `foods` where (`type` = 5);
	db.Where("type = ?", 5).Delete(&Food{})
}
func update(db *gorm.DB) {
	//提示: 相当于根据主键id，更新所有模型字段值。
	food := Food{}
	//先查询一条记录, 保存在模型变量food
	//等价于: SELECT * FROM `foods`  WHERE (id = '2') LIMIT 1
	db.Where("id = ?", 2).Take(&food)

	//修改food模型的值
	food.Price = 100

	//等价于: UPDATE `foods` SET `title` = '可乐', `type` = '0', `price` = '100', `stock` = '26', `create_time` = '2018-11-06 11:12:04'  WHERE `foods`.`id` = '2'
	db.Save(&food)

	//例子1:
	//更新food模型对应的表记录
	//等价于: UPDATE `foods` SET `price` = '25'  WHERE `foods`.`id` = '2'
	db.Model(&food).Update("price", 25)
	//通过food模型的主键id的值作为where条件，更新price字段值。

	//例子2:
	//上面的例子只是更新一条记录，如果我们要更全部记录怎么办？
	//等价于: UPDATE `foods` SET `price` = '25'
	db.Model(&Food{}).Update("price", 25)
	//注意这里的Model参数，使用的是Food{}，新生成一个空白的模型变量，没有绑定任何记录。
	//因为Food{}的id为空，gorm库就不会以id作为条件，where语句就是空的

	//例子3:
	//根据自定义条件更新记录，而不是根据主键id
	//等价于: UPDATE `foods` SET `price` = '25'  WHERE (create_time > '2018-11-06 20:00:00')
	db.Model(&Food{}).Where("create_time > ?", "2018-11-06 20:00:00").Update("price", 25)

	//提示： 通过结构体变量更新字段值, gorm库会忽略零值字段。就是字段值等于0, nil, "", false这些值会被忽略掉，不会更新。如果想更新零值，可以使用map类型替代结构体。

	//例子1：
	//通过结构体变量设置更新字段
	updataFood := Food{
		Price: 120,
		Title: "柠檬雪碧",
	}

	//根据food模型更新数据库记录
	//等价于: UPDATE `foods` SET `price` = '120', `title` = '柠檬雪碧'  WHERE `foods`.`id` = '2'
	//Updates会忽略掉updataFood结构体变量的零值字段, 所以生成的sql语句只有price和title字段。
	db.Model(&food).Updates(&updataFood)

	//例子2:
	//根据自定义条件更新记录，而不是根据模型id
	updataFood = Food{
		Stock: 120,
		Title: "柠檬雪碧",
	}

	//设置Where条件，Model参数绑定一个空的模型变量
	//等价于: UPDATE `foods` SET `stock` = '120', `title` = '柠檬雪碧'  WHERE (price > '10')
	db.Model(&Food{}).Where("price > ?", 10).Updates(&updataFood)

	//例子3:
	//如果想更新所有字段值，包括零值，就是不想忽略掉空值字段怎么办？
	//使用map类型，替代上面的结构体变量

	//定义map类型，key为字符串，value为interface{}类型，方便保存任意值
	data := make(map[string]interface{})
	data["stock"] = 0 //零值字段
	data["price"] = 35

	//等价于: UPDATE `foods` SET `price` = '35', `stock` = '0'  WHERE (id = '2')
	db.Model(&Food{}).Where("id = ?", 2).Updates(data)

	//orm提供了Expr函数用于设置表达式

	//等价于: UPDATE `foods` SET `stock` = stock + 1  WHERE `foods`.`id` = '2'
	db.Model(&food).Update("stock", gorm.Expr("stock + 1"))
}
func query(db *gorm.DB) {
	//定义接收查询结果的结构体变量
	food := Food{}
	//定义一个用户，并初始化数据
	f := Food{
		Title:      "apple",
		Price:      4.5,
		Stock:      12,
		Type:       1,
		CreateTime: time.Now(),
	}

	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('rain','123456','1540824823')
	if err := db.Debug().Create(&f).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	log.Println("id:", f.Id)
	//等价于：SELECT * FROM `foods`   LIMIT 1
	err := db.Take(&food).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("查询不到数据")
	} else if err != nil {
		//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
		fmt.Println("查询失败", err)
	}
	log.Println("take food:", food)
	//等价于：SELECT * FROM `foods`   ORDER BY `foods`.`id` ASC LIMIT 1
	db.First(&food)
	///等价于：SELECT * FROM `foods`   ORDER BY `foods`.`id` DESC LIMIT 1
	//语义上相当于返回最后一条记录
	db.Last(&food)
	//因为Find返回的是数组，所以定义一个商品数组用来接收结果
	var foods []Food

	//等价于：SELECT * FROM `foods`
	db.Find(&foods)

	//商品标题数组
	var titles []string

	//返回所有商品标题
	//等价于：SELECT title FROM `foods`
	//Pluck提取了title字段，保存到titles变量
	//这里Model函数是为了绑定一个模型实例，可以从里面提取表名。
	db.Model(&Food{}).Pluck("title", &titles)
	log.Println("pluck:", titles)

	//例子1:
	//等价于: SELECT * FROM `foods`  WHERE (id = '10') LIMIT 1
	//这里问号(?), 在执行的时候会被10替代
	db.Where("id = ?", 10).Take(&food)

	//例子2:
	// in 语句
	//等价于: SELECT * FROM `foods`  WHERE (id in ('1','2','5','6')) LIMIT 1
	//args参数传递的是数组
	db.Where("id in (?)", []int{1, 2, 5, 6}).Take(&food)

	//例子3:
	//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00' and create_time <= '2018-11-06 23:59:59')
	//这里使用了两个问号(?)占位符，后面传递了两个参数替换两个问号。
	db.Where("create_time >= ? and create_time <= ?", "2018-11-06 00:00:00", "2018-11-06 23:59:59").Find(&foods)

	//例子4:
	//like语句
	//等价于: SELECT * FROM `foods`  WHERE (title like '%可乐%')
	db.Where("title like ?", "%可乐%").Find(&foods)

	//例子1:
	//等价于: SELECT id,title FROM `foods`  WHERE `foods`.`id` = '1' AND ((id = '1')) LIMIT 1
	db.Select("id,title").Where("id = ?", 1).Take(&food)

	//这种写法是直接往Select函数传递数组，数组元素代表需要选择的字段名
	db.Select([]string{"id", "title"}).Where("id = ?", 1).Take(&food)
	fmt.Println("food id and title:", food)
	//例子2:
	//可以直接书写聚合语句
	//等价于: SELECT count(*) as total FROM `foods`
	total := []int{}

	//Model函数，用于指定绑定的模型，这里生成了一个Food{}变量。目的是从模型变量里面提取表名，Pluck函数我们没有直接传递绑定表名的结构体变量，gorm库不知道表名是什么，所以这里需要指定表名
	//Pluck函数，主要用于查询一列值
	db.Model(&Food{}).Select("count(*) as total").Pluck("total", &total)

	fmt.Println("total:", total[0])

	//例子:
	//等价于: SELECT * FROM `foods`  WHERE (create_time >= '2018-11-06 00:00:00') ORDER BY create_time desc
	db.Where("create_time >= ?", "2018-11-06 00:00:00").Order("create_time desc").Find(&foods)

	//等价于: SELECT * FROM `foods` ORDER BY create_time desc LIMIT 10 OFFSET 0
	db.Order("create_time desc").Limit(10).Offset(0).Find(&foods)

	//定义一个用户，并初始化数据
	f1 := Food{
		Title:      "watermelon",
		Price:      23.5,
		Stock:      6,
		Type:       30,
		CreateTime: time.Now(),
	}

	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('rain','123456','1540824823')
	if err := db.Debug().Create(&f1).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	log.Println("f1 id:", f1.Id)
	//例子:
	var totalCount int64 = 0
	//等价于: SELECT count(*) FROM `foods`
	//这里也需要通过model设置模型，让gorm可以提取模型对应的表名
	db.Model(Food{}).Count(&totalCount)
	fmt.Println("totalCount:", totalCount)

	//例子:
	//统计每个商品分类下面有多少个商品
	//定一个Result结构体类型，用来保存查询结果
	type Result struct {
		Type  int
		Total int
	}

	var results []Result
	//提示：Group函数必须搭配Select函数一起使用
	//等价于: SELECT type, count(*) as  total FROM `foods` GROUP BY type HAVING (total > 0)
	db.Model(Food{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)
	fmt.Println("type group count result:", results)
	//scan类似Find都是用于执行查询语句，然后把查询结果赋值给结构体变量，区别在于scan不会从传递进来的结构体变量提取表名.
	//这里因为我们重新定义了一个结构体用于保存结果，但是这个结构体并没有绑定foods表，所以这里只能使用scan查询函数。
}

func create(db *gorm.DB) {
	//定义一个用户，并初始化数据
	u := User{
		Username:   "rain",
		Password:   "123456",
		CreateTime: time.Now().Unix(),
	}

	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('rain','123456','1540824823')
	if err := db.Debug().Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	log.Println("id:", u.ID)

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	u = User{}
	//自动生成sql： SELECT * FROM `users`  WHERE (username = 'rain') LIMIT 1
	result := db.Debug().Where("username = ?", "rain").First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.Username, u.Password)

	//更新
	//自动生成Sql: UPDATE `users` SET `password` = '654321'  WHERE (username = 'tizi365')
	db.Model(&User{}).Where("username = ?", "rain").Update("password", "654321")
	result = db.Debug().Where("username = ?", "rain").First(&u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("找不到记录")
		return
	}
	//打印查询到的数据
	fmt.Println(u.Username, u.Password)
	//删除
	//自动生成Sql： DELETE FROM `users`  WHERE (username = 'rain')
	db.Where("username = ?", "rain").Delete(&User{})

	u.Username = "jack"
	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('jack','123456','1540824823')
	if err := db.Debug().Create(&u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
}
