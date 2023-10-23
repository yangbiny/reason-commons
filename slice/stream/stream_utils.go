package stream

import "github.com/yangbiny/reason-commons/common"

func Just[REQ []*interface{}](slice REQ) (Stream, error) {
	return DefaultStream[REQ]{
		slice: slice,
	}, nil
}

type Stream interface {
	// Distinct 去除重复item
	Distinct(keyFunc common.Function) Stream
	// Filter 按条件过滤item
	Filter(filterFunc common.Predicate) Stream
	// Group 分组
	Group(fn common.Function) Stream
	// Head 返回前n个元素
	Head(n int64) Stream
	// Tail 返回后n个元素
	Tail(n int64) Stream
	// First 获取第一个元素
	First() interface{}
	// Last 获取最后一个元素
	Last() interface{}
	// Map 转换对象
	Map(fn common.Function) Stream
	// Merge 合并item到slice生成新的stream
	Merge() Stream
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
	result []interface{}
}

func (stream DefaultStream[REQ]) Distinct(keyFunc common.Function) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Filter(filterFunc common.Predicate) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Group(fn common.Function) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Head(n int64) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Tail(n int64) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) First() interface{} {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Last() interface{} {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Map(fn common.Function) Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Merge() Stream {
	//TODO implement me
	panic("implement me")
}

func (stream DefaultStream[REQ]) Reverse() Stream {
	//TODO implement me
	panic("implement me")
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
