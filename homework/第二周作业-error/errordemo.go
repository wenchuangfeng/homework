package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

var NotFound = errors.New("not found")

func dao(query string) error {
	err := mockError()

	if err == sql.ErrNoRows{
		// 把错误堆栈信息返回给业务层，并且带上查询语句
		return errors.Wrapf(NotFound, fmt.Sprintf("data not found, sql: %s", query))
	}

	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("query falied: %s", query))
	}

	return nil
}


func biz() error {
	err := dao("xxx")
	if errors.Is(err, NotFound){
		// 写业务代码
		return nil
	}
	if err != nil {
		// 写业务代码
	}

	return nil
}


