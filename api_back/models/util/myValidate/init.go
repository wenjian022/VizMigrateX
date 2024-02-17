package myValidate

import (
	"dataxWeb/models/util"
	"dataxWeb/models/util/lg"
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

type CustomizeError struct {
	Name     string // 名称
	Tag      string // tag
	ErrorStr string // 告警内容
}

// Add 添加告警
func (receiver CustomizeError) Add() {
	customizeErrors = append(customizeErrors, receiver)
}

// showError
func showError(name string, tag string) (string, bool) {
	for _, customizeError := range customizeErrors {
		if customizeError.Name == name || customizeError.Tag == tag {
			return customizeError.ErrorStr, true
		}
	}
	return "", false
}

var customizeErrors []CustomizeError
var trans ut.Translator

func init() {
	//修改gin框架中的Validator属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator("zh")
		if !ok {
			lg.Error(errors.New("使用中文翻译器失败"))
			return
		}
		if err := zhTranslations.RegisterDefaultTranslations(v, trans); err != nil {
			lg.Error(errors.New("加载翻译器失败"))
		}
		return
	}
	return
}

// translationValidate 翻译错误信息
func translationValidate(fields validator.FieldError) string {
	if s, ok := showError(fields.Namespace(), fields.Tag()); ok {
		return s
	}
	return fields.Translate(trans)
}

// ErrResp validator 错误校验
func ErrResp(err error) util.Code {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return util.UtilError(err.Error()) // 翻译校验错误提示
	}
	for _, fieldError := range errs {
		return util.UtilError(translationValidate(fieldError))
	}
	return util.UtilError(err.Error()) // 翻译校验错误提示
}
