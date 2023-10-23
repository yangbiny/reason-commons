package stream

import "github.com/yangbiny/reason-commons/common"

func FromSlice[REQ []*interface{}](slice REQ) (Stream, error) {
	return DefaultStream[REQ]{
		slice: slice,
	}, nil
}

type Stream interface {
	// Distinct 去除重复item
	Distinct(keyFunc common.Function) Stream
	// Filter 按条件过滤item
	Filter(filterFunc common.Predicate) Stream
	// First 获取第一个元素
	First() interface{}
	// Last 获取最后一个元素
	Last() interface{}
	// Map 转换对象
	Map(fn common.Function) Stream
	// Reverse 反转
	Reverse() Stream
	// Sort 排序
	Sort(fn common.FilterFunc) Stream
	// Concat 聚合其他Stream
	Concat(streams ...Stream) Stream
	// AllMatch 检查是否全部匹配
	AllMatch(fn common.Predicate) bool
	// AnyMatch 检查是否存在至少一项匹配
	AnyMatch(fn common.Predicate) bool
	// NoneMatch 检查全部不匹配
	NoneMatch(fn common.Predicate) bool
	// Count 统计数量
	Count() int
	// ForEach 对每个元素执行操作
	ForEach(fn common.Consumer)
	// Done 清空stream
	Done()
}

type DefaultStream[REQ []*interface{}] struct {
	slice  REQ
	result []*interface{}
}

func (stream DefaultStream[REQ]) value() []*interface{} {
	if len(stream.result) == 0 {
		return stream.slice
	}
	return stream.result
}

func (stream DefaultStream[REQ]) Distinct(keyFunc common.Function) Stream {
	result := make(REQ, 0)
	occur := map[string]bool{}
	for _, item := range stream.value() {
		tmp := keyFunc(item)
		s := tmp.(string)
		if occur[s] {
			continue
		}
		occur[s] = true
		result = append(result, item)
	}
	stream.result = result
	return stream
}

func (stream DefaultStream[REQ]) Filter(filterFunc common.Predicate) Stream {
	var result = make(REQ, 0)
	for _, item := range stream.value() {
		if filterFunc(item) {
			result = append(result, item)
		}
	}
	stream.result = result
	return stream
}

func (stream DefaultStream[REQ]) First() interface{} {
	value := stream.value()
	if len(value) == 0 {
		return nil
	}
	return value[0]
}

func (stream DefaultStream[REQ]) Last() interface{} {
	value := stream.value()
	if len(value) == 0 {
		return nil
	}
	return value[len(value)-1]
}

func (stream DefaultStream[REQ]) Map(fn common.Function) Stream {
	result := make([]*interface{}, 0)
	for _, item := range stream.value() {
		tmp := fn(item)
		result = append(result, &tmp)
	}
	stream.result = result
	return stream
}

func (stream DefaultStream[REQ]) Reverse() Stream {
	value := stream.value()
	result := make([]*interface{}, 0)
	copy(result, value)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	stream.result = result
	return stream
}

func (stream DefaultStream[REQ]) Sort(fn common.FilterFunc) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Concat(streams ...Stream) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) AllMatch(fn common.Predicate) bool {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) AnyMatch(fn common.Predicate) bool {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) NoneMatch(fn common.Predicate) bool {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Count() int {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) ForEach(fn common.Consumer) {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Done() {
	//TODO implement me
	panic("implement me")
}
