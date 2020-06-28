package inject

import (
	"reflect"
	"unsafe"
)

const (
	InjectorTag = "auto"
)

var objs map[string]reflect.Value

func init() {
	objs = make(map[string]reflect.Value, 10)
}

// Register 注册对象
func Register(name string, v interface{}) {
	//todo:设置成锁操作，就需要使用容器来进行标准化
	objs[name] = reflect.ValueOf(v)
	//查找当前的结构来进行tag标准注册
}

// AutoRegister 注册对象
func AutoRegister(v interface{}) {
	rv := reflect.ValueOf(v)
	Register(rv.Type().String(), rv)
}

// Get 获取注册对象
func Get(key string) interface{} {
	v, ok := objs[key]
	if ok {
		return v.Interface()
	}
	return nil
}

// Remove 删除注册对象
func Remove(key string) {
	delete(objs, key)
}

//当前注入的逻辑本质上就是将这里设置的基本结构查找下面的tag来进行附着
func Inject() {
	for _, v := range objs {
		value := v
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
		for i := 0; i < value.NumField(); i++ {
			name := value.Type().Field(i).Tag.Get(InjectorTag)
			temp, ok := objs[name]
			if ok {
				field := value.Field(i)
				if field.CanSet() {
					field.Set(temp)
				} else {
					field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
					field.Set(temp)
				}
			}
		}
	}
}

