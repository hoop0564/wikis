package ch37_reflect

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	ID   string
	Name string `format:"normal"`
	Age  int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

func (e *Employee) UpdateName(name string) {
	e.Name = name
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f)) // int64 10
	t.Log(reflect.ValueOf(f).Type())             // int64
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("fail to get Name field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	age := 1
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(age)})
	t.Log("Updated Age: ", e)

	// 写模块、框架时好用！
	reflect.ValueOf(e).MethodByName("UpdateName").Call([]reflect.Value{reflect.ValueOf("John")})
	t.Log("Updated Age: ", e)
}

// 万能程序
func fillBySettings(st interface{}, setting map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr { // todo ??
		// elem 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if setting == nil {
		return errors.New("setting is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range setting {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st).Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Curry", "Age": 40}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
}
