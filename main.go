package main

import (
	"fmt"
	"reflect"
)

func pop(arr interface{}, isFront bool) (interface{},error) {
	arrReflected := reflect.Indirect(reflect.ValueOf(arr))

	if arr == nil {
		return nil, fmt.Errorf("nil input")
	}

	array,ok := arrReflected.Interface().([]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert to array")	
	}

	if len(array) == 0{
		return nil,fmt.Errorf("empty array/slice")
	}

	if isFront {
		return array[1:],nil
	} else {
		return array[:len(array) - 1],nil
	}

}

func pop1(arr interface{}, isFront bool) (interface{},error) {
	arrReflected := reflect.Indirect(reflect.ValueOf(arr))

	if arr == nil {
		return nil, fmt.Errorf("nil input")
	}

	var out []interface{}

	switch arrReflected.Kind() {
		
		case reflect.Slice :
			if isFront{
				for i:= 1; i < arrReflected.Len(); i++ {
					out = append(out,arrReflected.Index(i).Interface())
				}
			}else {
				for i:= 0; i < arrReflected.Len() - 1; i++ {
					out = append(out,arrReflected.Index(i).Interface())
				}

			}
		default:
			return nil,fmt.Errorf("not a array")

	}
	return out,nil

}

func main() {

	arr := []interface{}{[]interface{}{1,2},[]interface{}{"a","b","c"},[]interface{}{4,5,6}}
	arr1 := []int{1,2,3}
	fmt.Println(pop(arr,false))
	fmt.Println(pop1(arr1,true))

}