package utils

//db_init 初始化数据库
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func From_Db_close() {
	defer db.Close()
	fmt.Println("数据库已关闭")
}
func From_Db_iniit() (err error) {
	dsn := "root:wangshuai1@tcp(127.0.0.1:3306)/goMysql1"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	//测试链接数据库
	err = db.Ping()
	if err != nil {
		//fmt.Printf("db ping failed err:%v \n", err
		return err
	}
	return err
}

//Db_Select数据库查询方法
func From_Db_Select(n int) {
	sqlStr := `select id,name,age from user where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("%v", err)
	}
	//在使用完这个方法后，要记得关闭数据库
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("err is :%v\n", err)
			return
		}
		fmt.Printf("查询id:%v,查询名称为：%v，查询年龄为：%v\n", u.id, u.name, u.age)
	}

	//return &u.id, &u.name, &u.age
}

//Db_edit数据库更改数据
func From_Db_edit(age, id int) {
	//先写sql语句
	sqlStr := `update user set age =? where id = ?;`
	ret, err := db.Exec(sqlStr, age, id)
	if err != nil {
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		return
	}
	fmt.Println("n", n)

}

//Db_Insert数据库插入数据
func From_Db_Insert(name string, age int) {
	//先写sql语句
	sqlStr := `insert into user (name,age) values (?,?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("db exec failed err:%v", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("ret.lastinsertid failed err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

//Db_Deleat数据库删除数据
func From_Db_Deleat(id int) {
	//写删除的sql语句
	sqlStr := `delete from user where id =?;`
	rows, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	n, err := rows.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("id", n)
}
