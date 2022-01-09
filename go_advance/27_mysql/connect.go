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

// 8. mysql预处理
// ### 什么是预处理？
//
//普通SQL语句执行过程：
//
//1. 客户端对SQL语句进行占位符替换得到完整的SQL语句。
//2. 客户端发送完整SQL语句到MySQL服务端
//3. MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//
//预处理执行过程：
//
//1. 把SQL语句分成两部分，命令部分与数据部分。
//2. 先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
//3. 然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
//4. MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//
//### 为什么要预处理？
//
//1. 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
//2. 避免SQL注入问题。
// 预处理查询示例
func prepareQueryDemo() {
	// 8.1 准备sql
	sqlStr := "select id, name, age from user where id > ?"
	// 8.2 预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	// 8.3 执行sql
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 8.4 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 9. sql注入
// sql注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}

}

// 10.事物
// 事务操作示例
func transactionDemo() {
	// 10.1 开启事务
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "Update user set age=30 where id=?"
	// 10.2 执行sql
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "Update user set age=40 where id=?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)

	// 3. 提交事务
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
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

	//deleteRowDemo()

	//prepareQueryDemo()

	// #"表示
	// select id, name, age from user where name='xxx' or 1=1 #'
	// 纯粹是字符串连接  然后#'是注释了后面多余的'
	sqlInjectDemo("xxx' or 1=1 #")
}
