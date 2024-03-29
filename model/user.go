package model

import (
	"fmt"
	"log"
)

//insert 插入操作
func Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close() // 清理资源
	if err != nil {
		return 1, err
	}
	result, err := stmt.Exec(args...)
	if err != nil {
		return 1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 1, err
	}
	fmt.Printf("插入成功, ID为%v\n", id)
	return id, nil
}

// delete操作
func Delete(sql string, args ...interface{}) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()
	CheckErr(err, "SQL语句设置失败")
	result, err := stmt.Exec(args...)
	CheckErr(err, "参数添加失败")
	num, err := result.RowsAffected()
	CheckErr(err, "删除失败")
	fmt.Printf("删除成功，删除")
}

//update修改操作
func Update(sql string, args ...interface{}) {
	stmt, err := DB.prepare(sql)
	defer stmt.Close()
	CheckErr(err, "SQL语句设置失败")
	result, err := stmt.Exec(args...)
	CheckErr(err, "参数添加失败")
	num, err := result.RowsAffected()
	CheckErr(err, "修改失败")
	fmt.Printf("修改成功，修改行数为%d\n", num)
}

// CheckErr
func CheckErr(err error, msg string) {
	if nil != err {
		log.Panicln(msg, err)
	}
}
