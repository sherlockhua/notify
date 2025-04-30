package common

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

// columnNamesCache 用于缓存结构体类型对应的列名列表
var columnNamesCache = make(map[reflect.Type][]string)
var cacheMutex sync.RWMutex

// 获取结构体中所有gorm column名称
func getGormColumnNames(model interface{}) ([]string, error) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or pointer to struct")
	}
	// 先从缓存中查找
	cacheMutex.RLock()
	if columnNames, exists := columnNamesCache[t]; exists {
		cacheMutex.RUnlock()
		return columnNames, nil
	}
	cacheMutex.RUnlock()
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag == "" {
			continue
		}

		// 解析gorm标签
		for _, part := range strings.Split(tag, ";") {
			if strings.HasPrefix(part, "column:") {
				column := strings.TrimPrefix(part, "column:")
				columns = append(columns, column)
				break
			}
		}
	}

	// 将结果存入缓存
	cacheMutex.Lock()
	columnNamesCache[t] = columns
	cacheMutex.Unlock()

	return columns, nil
}

func IsValidColumnName(s interface{}, columnName string) (bool, error) {
	columnNames, err := getGormColumnNames(s)
	if err != nil {
		return false, err
	}
	for _, name := range columnNames {
		if name == columnName {
			return true, nil
		}
	}
	return false, nil
}
