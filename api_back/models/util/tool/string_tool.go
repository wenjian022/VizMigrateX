package tool

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"text/template"
	"time"
)

var letters = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

var randSeqMs sync.Mutex // 在多线程中防止随机因子一样导致随机字符串重复

// RandSeq 生成随机字符串
//
//	@Description:  生成随机字符串
//	@param n 生成几个字符串
//	@return string
func RandSeq(n int) string {
	// 使用 Mutex 保护对 rand.Seed 的调用
	randSeqMs.Lock()
	defer func() { randSeqMs.Unlock() }()

	b := make([]rune, n)
	rand.New(rand.NewSource(time.Now().Unix()))
	// 随机因子
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// DeleteStringElement
//
//	@Description: 删除切片下标
//	@param slice
//	@param index
//	@return []int
func DeleteStringElement(slice *[]string, index int) {
	// 检查下标是否合法
	if index < 0 || index >= len(*slice) {
		return
	}

	// 将切片元素向前移动覆盖待删除元素
	copy((*slice)[index:], (*slice)[index+1:])

	// 调整切片长度
	*slice = (*slice)[:len(*slice)-1]
}

// NowTimeString
//
//	@Description: 获取当前时间字符串
func NowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// IsElementStr
//
//	@Description: 判断元素是否存在在列表中
//	@param items 列表
//	@param Item 元素
//	@return bool
func IsElementStr(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

// MD5
//
//	@Description: md5
//	@param by
//	@return string
func MD5(by []byte) string {
	return fmt.Sprintf("%x", md5.Sum(by))
}

// IsTimedTaskFormat
//
//	@Description: 定时任务格式判断
//	@param timerSting 定时器
//	@return bool
func IsTimedTaskFormat(timerSting string) bool {
	_, err := cron.ParseStandard(timerSting)
	return err == nil
}

// StringToErr
//
//	@Description: 字符串转error
//	@param str
//	@return error
func StringToErr(str string) error {
	return errors.New(str)
}

// NumericalJudgment
//
//	@Description: 数值判断
//	@return bool
func NumericalJudgment(key int, comparator string, value int) bool {
	switch comparator {
	case "==":
		return key == value
	case "!=":
		return key != value
	case ">=":
		return key >= value
	case "<=":
		return key <= value
	case ">":
		return key > value
	case "<":
		return key < value
	default:
		return false
	}
}

// StringJudgment
//
//	@Description: 字符串判断
//	@return bool
func StringJudgment(key, comparator, value string) bool {
	switch comparator {
	case "==":
		return key == value
	case "!=":
		return key != value
	default:
		return false
	}
}

// StringToInt64
//
//	@Description: 字符串转int64
//	@param str
//	@return int64
func StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// ToString
//
//	@Description: interface的字符串值
//	@param value
//	@return string
func ToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

// TemplateToString
//
//	@Description: 模板转字符串
//	@param _template
//	@param m
//	@return string
func TemplateToString(_template string, m map[string]string) string {
	t := template.Must(template.New("push").Parse(_template))
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "push", m); err != nil {
		return "模板生成失败: " + err.Error()
	}
	return buf.String()
}

// ColumnNumberToName
//
//	@Description: 将整数转换为Excel的列数
//	@param num
//	@return string
//	@return error
func ColumnNumberToName(num int) (string, error) {
	if num < 1 || num > 16384 {
		return "", StringToErr("列号必须大于或等于16384且小于或等于1")
	}
	var col string
	for num > 0 {
		col = string(rune((num-1)%26+65)) + col
		num = (num - 1) / 26
	}
	return col, nil
}

// JsonToMD5
//
//	@Description: json转md5
//	@param data
//	@return string
func JsonToMD5(data interface{}) string {
	by, _ := json.Marshal(data)
	return MD5(by)
}

// IsStructEmpty
//
//	@Description: 空结构体判断
//	@param s
//	@return bool
func IsStructEmpty(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

// CalculationDate
//
//	@Description: 计算日期
//	@param day 天数，可以使用(-15还在15， - 就是15天的日期)
//	@return time.Time 时间
func CalculationDate(day int) time.Time {
	// 获取当前时间
	currentTime := time.Now()

	return currentTime.AddDate(0, 0, day)
}
