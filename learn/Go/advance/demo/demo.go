package demo

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"path"
	"reflect"
)

type MyJobList struct {
	Items []MyJob
}

func (jl *MyJobList) Info() {

}

type Object interface {
	Info()
}

type MyJob struct {
}

type Codec interface {
	Encode(obj interface{}) ([]byte, error)
	Decode(data []byte, out Object) (Object, error)
}

type EtcdStorage struct {
	client *clientv3.Client
	codec  Codec
	prefix string
}

func (s *EtcdStorage) GetList(ctx context.Context, key string, listObj Object) error {
	listPtr, err := GetItemsPtr(listObj)
	if err != nil {
		panic(err)
	}

	v, err := EnforcePtr(listPtr)
	if err != nil || v.Kind() != reflect.Slice {
		panic(fmt.Errorf("list error"))
	}

	newItemFunc := getNewItemFunc(listObj, v)

	key = path.Join(s.prefix, key)

	resp, err := s.client.Get(ctx, key)
	if err != nil {
		return err
	}

	glowSlice(v, len(resp.Kvs))

	for i, kv := range resp.Kvs {

	}
}

func getNewItemFunc(listObj Object, v reflect.Value) func() Object {
	elem := v.Type().Elem()
	return func() Object {
		return reflect.New(elem).Interface().(Object)
	}
}

func glowSlice(v reflect.Value, maxCapacity int, sizes ...int) {
	cap := v.Cap()
	max := cap
	for _, size := range sizes {
		if size > max {
			max = size
		}
	}
	if len(sizes) == 0 || max > maxCapacity {
		max = maxCapacity
	}
	if max <= cap {
		return
	}
	if v.Len() > 0 {
		extra := reflect.MakeSlice(v.Type(), 0, max)
		reflect.Copy(extra, v)
		v.Set(extra)
	} else {
		extra := reflect.MakeSlice(v.Type(), 0, max)
		v.Set(extra)
	}
}

func appendListItem(v reflect.Value, data []byte, codec Codec, newItemFunc func() Object) error {
	obj, err := codec.Decode(data, newItemFunc())
	if err != nil {
		return err
	}
	v.Set(reflect.Append(v, reflect.ValueOf(obj).Elem()))
	return nil
}

func GetItemsPtr(list Object) (interface{}, error) {
	obj, err := getItemsPtr(list)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func getItemsPtr(list Object) (interface{}, error) {
	v, err := EnforcePtr(list)
	if err != nil {
		panic(err)
	}

	items := v.FieldByName("Items")
	if !items.IsValid() {
		panic(fmt.Errorf("invalid"))
	}

	switch items.Kind() {
	case reflect.Slice:
		return items.Addr().Interface(), nil
	default:
		return nil, fmt.Errorf("not slice")
	}
	return nil, nil
}

func EnforcePtr(obj interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		//if v.Kind() == reflect.Invalid {
		//	return reflect.Value{}, fmt.Errorf("")
		//}
		return reflect.Value{}, fmt.Errorf("not pointer")
	}
	if v.IsNil() {
		return reflect.Value{}, fmt.Errorf("nil pointer")
	}
	return v.Elem(), nil
}
