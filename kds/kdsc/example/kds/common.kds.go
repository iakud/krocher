// Code generated by kds. DO NOT EDIT.
// source: TODO: Source File

package kds;
import (
	"google.golang.org/protobuf/types/known/emptypb"
)

type dirtyParentFunc_int32_empty_Map func()

func (f dirtyParentFunc_int32_empty_Map) invoke() {
	if f == nil {
		return
	}
	f()
}

type int32_empty_Map struct {
	syncable map[int32]struct{}

	dirtyParent dirtyParentFunc_int32_empty_Map
}

func (x *int32_empty_Map) Len() int {
	return len(x.syncable)
}

func (x *int32_empty_Map) Clear() {
	for k := range x.syncable {
		delete(x.syncable, k)
	}
}

func (x *int32_empty_Map) Get(k int32) (struct{}, bool) {
	v, ok := x.syncable[k]
	return v, ok
}

func (x *int32_empty_Map) Set(k int32, v struct{}) {
	x.syncable[k] = v
}

func (x *int32_empty_Map) Delete(k int32) {
	delete(x.syncable, k)
}

func (x *int32_empty_Map) All() func(yield func(int32, struct{}) bool) {
	return func(yield func(int32, struct{}) bool) {
		for k, v := range x.syncable {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (x *int32_empty_Map) Keys() []int32 {
	r := make([]int32, 0, len(x.syncable))
	for k := range x.syncable {
		r = append(r, k)
	}
	return r
}

func (x *int32_empty_Map) Values() []struct{} {
	r := make([]struct{}, 0, len(x.syncable))
	for _, v := range x.syncable {
		r = append(r, v)
	}
	return r
}

type dirtyParentFunc_int32_int32_Map func()

func (f dirtyParentFunc_int32_int32_Map) invoke() {
	if f == nil {
		return
	}
	f()
}

type int32_int32_Map struct {
	syncable map[int32]int32

	dirtyParent dirtyParentFunc_int32_int32_Map
}

func (x *int32_int32_Map) Len() int {
	return len(x.syncable)
}

func (x *int32_int32_Map) Clear() {
	for k := range x.syncable {
		delete(x.syncable, k)
	}
}

func (x *int32_int32_Map) Get(k int32) (int32, bool) {
	v, ok := x.syncable[k]
	return v, ok
}

func (x *int32_int32_Map) Set(k int32, v int32) {
	x.syncable[k] = v
}

func (x *int32_int32_Map) Delete(k int32) {
	delete(x.syncable, k)
}

func (x *int32_int32_Map) All() func(yield func(int32, int32) bool) {
	return func(yield func(int32, int32) bool) {
		for k, v := range x.syncable {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (x *int32_int32_Map) Keys() []int32 {
	r := make([]int32, 0, len(x.syncable))
	for k := range x.syncable {
		r = append(r, k)
	}
	return r
}

func (x *int32_int32_Map) Values() []int32 {
	r := make([]int32, 0, len(x.syncable))
	for _, v := range x.syncable {
		r = append(r, v)
	}
	return r
}

type dirtyParentFunc_int64_List func()

func (f dirtyParentFunc_int64_List) invoke() {
	if f == nil {
		return
	}
	f()
}

type int64_List struct {
	syncable []int64

	dirtyParent dirtyParentFunc_int64_List
}

func (x *int64_List) Len() int {
	return len(x.syncable)
}

func (x *int64_List) Get(i int) int64 {
	return x.syncable[i]
}

func (x *int64_List) Set(i int, v int64) {
	x.syncable[i] = v
}

func (x *int64_List) Index(v int64) int {
	for i := range x.syncable {
		if v == x.syncable[i] {
			return i
		}
	}
	return -1
}

func (x *int64_List) Contains(v int64) bool {
	return x.Index(v) >= 0
}

func (x *int64_List) Append(v ...int64) {
	x.syncable = append(x.syncable, v...)
}

func (x *int64_List) Insert(i int, v ...int64) {
	// x.syncable = slices.Insert(x.syncable, i, v...)
}

func (x *int64_List) Delete(i, j int) {
	// x.syncable = slices.Delete(x.syncable, i, j)
}

func (x *int64_List) Replace(i, j int, v ...int64) {
	// x.syncable = slices.Replace(x.syncable, i, j, v...)
}

func (x *int64_List) All() func(yield func(int, int64) bool) {
	return func(yield func(int, int64) bool) {
		for i, v := range x.syncable {
			if !yield(i, v) {
				return
			}
		}
	}
}

func (x *int64_List) Values() []int64 {
	return append(x.syncable[:0:0], x.syncable...)
}