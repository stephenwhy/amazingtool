package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// TransformSlice 对 slice 进行 map 操作
func TransformSlice[T any, R any](slice []T, fn func(T) R) []R {
	res := make([]R, len(slice))
	for i, v := range slice {
		res[i] = fn(v)
	}
	return res
}

func Int64SliceToStr(slice []int64, sep string) string {
	// 将 int64 切片转换为 string 切片
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(strSlice, sep)
}

func IntSliceToStr(slice []int, sep string) string {
	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprintf("%d", v)
	}

	return strings.Join(strSlice, sep)
}

func StrToInt64Slice(str string, sep string) []int64 {
	ids := strings.Split(str, sep)
	res := make([]int64, len(ids))
	for i, id := range ids {
		res[i], _ = strconv.ParseInt(id, 10, 64)
	}
	return res
}

func StrToIntSlice(str string, sep string) []int {
	ids := strings.Split(str, sep)
	res := make([]int, len(ids))
	for i, id := range ids {
		res[i], _ = strconv.Atoi(id)
	}
	return res
}

// IfAElseB 三元运算符
func IfAElseB[T any](condition func() bool, A, B T) T {
	if condition() {
		return A
	} else {
		return B
	}
}

// PaginateBySlice slice分页
func PaginateBySlice[T any](original []T, page int, pageSize int) []T {
	var (
		total = len(original)
		start = (page - 1) * pageSize
		end   = page * pageSize
	)
	if start > total {
		return make([]T, 0)
	}
	if end > total {
		end = total
	}
	return original[start:end]
}

func StrLike(str string) string {
	return fmt.Sprintf("%%%s%%", str)
}
