package coltx

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

func Contains[T comparable](array []T, item T) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func MapContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func Filter[T any](list []T, condition func(T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range list {
		if condition(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Map[T any, U any](list []T, f func(T) U) []U {
	result := make([]U, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	return result
}

func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, s := range slices {
		totalLen += len(s)
	}
	result := make([]T, totalLen)
	i := 0
	for _, s := range slices {
		copy(result[i:], s)
		i += len(s)
	}
	return result
}

func Sum[T any](slice []T, transformer func(T) float64) float64 {
	sum := 0.0
	for _, item := range slice {
		sum += transformer(item)
	}
	return sum
}

func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func SliceToMap[T any, K comparable](slice []T, keyFunc func(T) K) map[K]T {
	result := make(map[K]T)
	for _, item := range slice {
		result[keyFunc(item)] = item
	}
	return result
}

func Reduce[T any, U any](slice []T, accumulator func(U, T) U, initialValue U) U {
	result := initialValue
	for _, item := range slice {
		result = accumulator(result, item)
	}
	return result
}

func IndexOf[T comparable](slice []T, item T) int {
	for i, value := range slice {
		if value == item {
			return i
		}
	}
	return -1
}

func Unique[T comparable](slice []T) []T {
	uniqueMap := make(map[T]bool)
	uniqueValues := make([]T, 0)
	for _, value := range slice {
		if _, found := uniqueMap[value]; !found {
			uniqueValues = append(uniqueValues, value)
			uniqueMap[value] = true
		}
	}
	return uniqueValues
}

func Flatten[T any](s []interface{}) []T {
	result := make([]T, 0)
	for _, v := range s {
		switch val := v.(type) {
		case []interface{}:
			result = append(result, Flatten[T](val)...)
		default:
			if _, ok := val.(T); ok {
				result = append(result, val.(T))
			}
		}
	}
	return result
}

func DeepEqual[T comparable](a, b T) bool {
	if !reflect.DeepEqual(a, b) {
		return false
	}
	return true
}

// This is a function named GroupBy that takes in a generic type T, and a comparable type K. It takes in a slice of type T, and a function named getKey that takes in a parameter of type T and returns a value of type K.
// The function then creates a new empty map, with the key of type K and the value of type slice of type T. It then loops through the input slice of type T, and gets the key for each item by calling the getKey function.
// It then appends the current item into the slice stored in the corresponding key in the map.
// In summary, this function groups the input slice by a key, which is determined by a function that maps each item in the slice to a key value. This can be useful for organizing and sorting data.
func GroupBy[T any, K comparable](slice []T, getKey func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, item := range slice {
		key := getKey(item)
		result[key] = append(result[key], item)
	}
	return result
}

func FlattenDeep(arr interface{}) []interface{} {
	result := make([]interface{}, 0)
	switch v := arr.(type) {
	case []interface{}:
		for _, val := range v {
			result = append(result, FlattenDeep(val)...)
		}
	case interface{}:
		result = append(result, v)
	}
	return result
}

func Join[T any](slice []T, separator string) string {
	result := ""
	for i, item := range slice {
		if i > 0 {
			result += separator
		}
		result += fmt.Sprintf("%v", item)
	}
	return result
}

func Reverse[T any](slice []T) []T {
	reversed := make([]T, len(slice))
	for i, j := 0, len(slice)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = slice[j], slice[i]
	}
	return reversed
}

func FindIndex[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

func MapToSlice[T any, U any](slice []T, mapper func(T) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, item := range slice {
		mappedSlice[i] = mapper(item)
	}
	return mappedSlice
}

func MergeMaps[K any, V any](maps ...map[interface{}]V) map[interface{}]V {
	mergedMap := make(map[interface{}]V)
	for _, m := range maps {
		for k, v := range m {
			mergedMap[k] = v
		}
	}
	return mergedMap
}

func FilterMap[K any, V any](m map[any]V, filter func(V) bool) map[any]V {
	filteredMap := make(map[any]V)
	for k, v := range m {
		if filter(v) {
			filteredMap[k] = v
		}
	}
	return filteredMap
}

func Chunk[T any](slice []T, chunkSize int) [][]T {
	if chunkSize <= 0 {
		return nil
	}
	var chunks [][]T
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}

func Values[K any, V any](m map[any]V) []V {
	values := make([]V, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}

func Shuffle[T any](slice []T) []T {
	shuffledSlice := make([]T, len(slice))
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(slice))
	for i, randIndex := range perm {
		shuffledSlice[i] = slice[randIndex]
	}
	return shuffledSlice
}

func CartesianProduct[T any](slices ...[]T) [][]T {
	n := len(slices)
	if n == 0 {
		return [][]T{{}}
	}
	if n == 1 {
		product := make([][]T, len(slices[0]))
		for i, item := range slices[0] {
			product[i] = []T{item}
		}
		return product
	}
	tailProduct := CartesianProduct(slices[1:]...)
	product := make([][]T, 0, len(slices[0])*len(tailProduct))
	for _, head := range slices[0] {
		for _, tail := range tailProduct {
			product = append(product, append([]T{head}, tail...))
		}
	}
	return product
}

func Sort[T any](slice []T, comparer func(T, T) bool) []T {
	sortedSlice := make([]T, len(slice))
	copy(sortedSlice, slice)
	sort.Slice(sortedSlice, func(i, j int) bool {
		return comparer(sortedSlice[i], sortedSlice[j])
	})
	return sortedSlice
}

func AllMatch[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

func AnyMatch[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

func Push[T any](slice []T, element T) []T {
	return append(slice, element)
}

// Remove the element on the last order slice
func Pop[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

// Add new element on the first order slice
func Unshift[T any](slice []T, element T) []T {
	return append([]T{element}, slice...)
}

// Remove the element on the first order slice
func Shift[T any](slice []T) []T {
	return slice[1:]
}

func AppendIfMissing[T comparable](slice []T, element T) []T {
	if !Contains(slice, element) {
		return append(slice, element)
	}
	return slice
}

// Get all elements common from 2 slices
func Intersect[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	result := []T{}
	for _, item := range slice1 {
		set[item] = true
	}
	for _, item := range slice2 {
		if set[item] {
			result = append(result, item)
		}
	}
	return result
}

func Difference[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]bool)
	result := []T{}
	for _, item := range slice1 {
		set[item] = true
	}
	for _, item := range slice2 {
		if !set[item] {
			result = append(result, item)
		}
	}
	for _, item := range slice1 {
		if !set[item] {
			result = append(result, item)
		}
	}
	return result
}

func JoinMapKeys[V any](m map[string]V, separator string) string {
	joined_keys := []string{}
	for key := range m {
		joined_keys = append(joined_keys, key)
	}
	return strings.Join(joined_keys, separator)
}

func DeepMergeMap(target, source map[string]interface{}) {
	for key, sourceValue := range source {
		if targetValue, exists := target[key]; exists {
			if sourceMap, sourceIsMap := sourceValue.(map[string]interface{}); sourceIsMap {
				if targetMap, targetIsMap := targetValue.(map[string]interface{}); targetIsMap {
					DeepMergeMap(targetMap, sourceMap)
				}
			} else {
				target[key] = sourceValue
			}
		} else {
			target[key] = sourceValue
		}
	}
}

func MergeMapsString(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}
