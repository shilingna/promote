package main

import (
	"database/sql"
	"fmt"
)

var dbDemo *sql.DB

type User struct {
	id    int
	Name  string
	Phone string
}

// 定义一个全局变量
var u User

// 初始化数据库连接
func init() {
	dbDemo, _ = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/fw_go")
	err := dbDemo.Ping()
	if err != nil {
		fmt.Printf("连接 failed，err：%v\n", err)
	}
}

// queryRow 单行查询
func queryRow() {
	// 确保QueryRow之后调用Scan方法，否则持有的数据链接不会被释放
	err := dbDemo.QueryRow("select id,name,phone,from `user_go` where id =?").Scan(&u.id, &u.Name, &u.Phone)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s phone:%s\n", u.id, u.Name, u.Phone)
}

// queryMultiRow 多行查询
func queryMultiRow() {
	rows, err := dbDemo.Query("select id,name,phone from `user_go` where id > ?", 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&u.id, &u.Name, &u.Phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s phone:%s\n", u.id, u.Name, u.Phone)
	}
}

// insertRow 插入语句
func insertRow() {
	ret, err := dbDemo.Exec("insert into user_go(name,phone) values (?,?)", "刘备", 13966666666)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId() // 获取新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert sucess,the id is %d.\n", id)
}

// updateRow 单行更新
func updateRow() {
	ret, err := dbDemo.Exec("update user_go set name=? where id = ?", "关羽“，3")
	if err != nil {
		fmt.Printf("update failed,err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed ,err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

// deleteRow 删除数据
func deleteRow() {
	ret, err := dbDemo.Exec("deleteById from user_go wherr id = ?", 2)
	if err != nil {
		fmt.Printf("deleteById failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("deleteId success, affected rows:%d\n", n)
}

// prepareQuery 预处理查询
func prepareQuery() {
	stmt, err := dbDemo.Prepare("select id,name,phone,from `user_go` where id > ?")
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	// 循环读取结果集中的数据
	for rows.Next() {
		err := rows.Scan(&u.id, &u.Name, &u.Phone)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s phone:%s\n", u.id, u.Name, &u.Phone)
	}
}

// prepareInsert 预处理插入示例
func prepareInsert() {
	stmt, err := dbDemo.Prepare("insert into user_go (name,phone) values (?,?);")
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("张飞", 13719049801)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("刘备", 123123123123)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success")
}

// transaction 事务
func transaction() {
	tx, err := dbDemo.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user_go set name='张飞' where id=?", 1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	_, err = tx.Exec("update user_go set name='吕布' where id =?", 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err %v\n", err)
		return
	}
	fmt.Println("exec transaction success!")
}

func main() {
	/*queryRow()
	queryMultiRow()
	insertRow()
	updateRow()
	deleteRow()
	prepareInsert()
	prepareQuery()*/
	transaction()
}
