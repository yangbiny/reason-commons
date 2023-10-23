package common

type Function func(interface{}) interface{}

type Supplier func() bool

type Predicate func(interface{}) bool

type Consumer func(interface{})

type FilterFunc func(interface{}) int8
