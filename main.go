package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)



func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/denglu?charset=utf8")
	//	db.SetMaxOpenConns(1000)
	if err != nil {
		fmt.Println(err)
		fmt.Println("打开数据库失败")
	} else {
		fmt.Println("打开数据库成功！")
	}
	defer db.Close()
	r := gin.Default()
	//登录
		r.POST("/login", func(c *gin.Context) {
      var user User
	 	err := c.Bind(&user)
	 if err != nil {
	 	fmt.Println(err)
 	}
	 	username := user.Username
	 password :=user.Password
	 if login(db,username,password){
	 		c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": username+"登录成功",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message":   "登录失败",
			})
		}
	})
	//注册
	r.POST("/register", func(c *gin.Context) {
		var user User
		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
		}
		username := user.Username
		password :=user.Password
		if insert(db,username,password){
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": username+"注册成功",
			})
		}else{
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "注册失败",
			})
		}
	})
	r.Run(":8080")
}

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" bdinding:"required"`
}

func insert(db *sql.DB , username string,password string ) bool {
	_, err := db.Exec("INSERT into user(username,password) values (?,?)", username, password)
	if (err != nil){
		fmt.Println("插入数据失败")
		fmt.Println(err)
		return false
	} else{
		fmt.Println("插入成功")
		return true
	}
	return true
}


func login(db *sql.DB ,username string,password string)  bool {

	stmt, err := db.Query("select * from user;")
	if err != nil {
	 fmt.Println(err)
	}
	defer stmt.Close()
	for stmt.Next(){
		var name string
		var word string
		var n int
		err := stmt.Scan(&n,&name,&word)
		if err != nil {
			fmt.Println(err)

		}

		if username == name&&password == word{
			return true
		}
	}

	return false
}
