package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/**
连接数据库的function
*/
func initDB() (err error) {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db"

	// 1.1 连接MySQL数据库（注意不能使用 := ）
	// // 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	// 因为上面已经定义了全局变量 用:=将会只能成为从局部变量
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s failed, err: %v \n", dsn, err)
		return
	}

	// 1.2 尝试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err: %v, \n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
	return
}

/**
2.0
定义一个结构体来存储
*/
type user struct {
	id   int
	name string
	age  int
}

/**
3.0
演示查询操作 之单行查询
单行查询db.QueryRow()执行一次查询，并期望返回最多一行结果（即Row）。
QueryRow总是返回非nil的值，直到返回值的Scan方法被调用时，才会返回被延迟的错误。（如：未找到结果）
*/
func queryRowDemo() {
	// 3.1 定义一个sql语句 ?表示需要传入的参数
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 3.2 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	// 传入了参数1
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

/**
4.0 多行查询
*/
// 查询多条数据示例
func queryMultiRowDemo() {
	// 4.1 查询语句
	sqlStr := "select id, name, age from user where id > ?"
	// 4.2 Query
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 0.0 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		// 4.2 读取数据
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

/**
插入数据 5.0
*/
func insertRowDemo() {
	// 5.1 sql
	sqlStr := "insert into user(name, age) values (?,?)"
	// 5.2 Exec执行一次命令（包括查询、删除、更新、插入等），
	//返回的Result是对已执行的SQL命令的总结。参数args表示query中的占位参数。
	ret, err := db.Exec(sqlStr, "王五", 31)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	// 5.3 拿到执行sql之后的返回值并取新插入数据的id
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

/**
6.0 更新数据
*/
// 更新数据
func updateRowDemo() {
	// 6.1 sql
	sqlStr := "update user set age=? where id = ?"
	// 6.2 Exec 及其参数
	ret, err := db.Exec(sqlStr, 39, 3)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	// 6.2 影响行数
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

/**
7.0 删除数据
*/
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库初始化失败")
	}
	//
	//queryRowDemo()
	//
	//queryMultiRowDemo()
	//
	//insertRowDemo()

	//updateRowDemo()

	deleteRowDemo()

}
