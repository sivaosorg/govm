package coltx

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/sivaosorg/govm/utils"
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

// MapString2Table returns a map as a table without borders.
func MapString2Table(data map[string]string) string {
	var builder strings.Builder
	maxKeyLength := 0
	for key := range data {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}
	for key, value := range data {
		fmt.Fprintf(&builder, "%-*s   %s\n", maxKeyLength, key, value)
	}
	return builder.String()
}

// MapToTable returns a map as a table without borders.
func Map2Table(data map[string]interface{}) string {
	var builder strings.Builder
	maxKeyLength := 0
	for key := range data {
		if len(key) > maxKeyLength {
			maxKeyLength = len(key)
		}
	}
	for key, value := range data {
		fmt.Fprintf(&builder, "%-*s   %s\n", maxKeyLength, key, utils.ToJson(value))
	}
	return builder.String()
}

// IndexExists checks if the given index is within the valid range of the slice.
func IndexExists[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

// Iterate generic function to iterate over a collection (slice, array, map)
func Iterate(collection interface{}, callback func(index int, value interface{})) {
	v := reflect.ValueOf(collection)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			callback(i, v.Index(i).Interface())
		}
	} else if v.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, key := range keys {
			callback(-1, key.Interface())
			callback(-1, v.MapIndex(key).Interface())
		}
	}
}

// Map generic function to map a collection (slice, array, map) using a mapping function
func MapZ(collection interface{}, mapper func(value interface{}) interface{}) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(mapper(v.Index(0).Interface()))), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			mappedValue := mapper(v.Index(i).Interface())
			result = reflect.Append(result, reflect.ValueOf(mappedValue))
		}
	} else if v.Kind() == reflect.Map {
		keys := v.MapKeys()
		for _, key := range keys {
			mappedKey := mapper(key.Interface())
			mappedValue := mapper(v.MapIndex(key).Interface())
			result = reflect.Append(result, reflect.ValueOf(mappedKey))
			result = reflect.Append(result, reflect.ValueOf(mappedValue))
		}
	}

	return result.Interface()
}

// Filter generic function to filter a collection (slice, array) using a filter function
func FilterZ(collection interface{}, predicate func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if predicate(item) {
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}

	return result.Interface()
}

// Reduce generic function to reduce a collection (slice, array) to a single value using a reducer function
func ReduceZ(collection interface{}, reducer func(acc interface{}, value interface{}) interface{}, initialValue interface{}) interface{} {
	v := reflect.ValueOf(collection)
	accumulator := initialValue

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			accumulator = reducer(accumulator, v.Index(i).Interface())
		}
	}

	return accumulator
}

// Find generic function to find an element in a collection (slice, array) based on a condition
func Find(collection interface{}, predicate func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if predicate(item) {
				return item
			}
		}
	}

	return nil
}

// All generic function to check if all elements in a collection (slice, array) satisfy a condition
func All(collection interface{}, condition func(value interface{}) bool) bool {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if !condition(v.Index(i).Interface()) {
				return false
			}
		}
		return true
	}

	return false
}

// Any generic function to check if any element in a collection (slice, array) satisfies a condition
func Any(collection interface{}, condition func(value interface{}) bool) bool {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if condition(v.Index(i).Interface()) {
				return true
			}
		}
		return false
	}

	return false
}

// Count generic function to count the number of elements in a collection (slice, array) that satisfy a condition
func Count(collection interface{}, condition func(value interface{}) bool) int {
	v := reflect.ValueOf(collection)
	count := 0

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if condition(v.Index(i).Interface()) {
				count++
			}
		}
	}

	return count
}

// Remove generic function to remove elements from a collection (slice) that satisfy a condition
func Remove(collection interface{}, condition func(value interface{}) bool) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if !condition(item) {
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}

	return result.Interface()
}

// Sort generic function to sort a collection (slice, array)
func SortZ(collection interface{}, less func(i, j int) bool) {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		sort.SliceStable(collection, func(i, j int) bool {
			return less(i, j)
		})
	}
}

// Reverse generic function to reverse the elements of a collection (slice, array)
func ReverseZ(collection interface{}) {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		length := v.Len()
		for i := 0; i < length/2; i++ {
			j := length - i - 1
			vi := v.Index(i).Interface()
			vj := v.Index(j).Interface()
			v.Index(i).Set(reflect.ValueOf(vj))
			v.Index(j).Set(reflect.ValueOf(vi))
		}
	}
}

// Unique generic function to remove duplicate elements from a collection (slice, array)
func UniqueZ(collection interface{}) interface{} {
	v := reflect.ValueOf(collection)
	uniqueMap := make(map[interface{}]struct{})
	result := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if _, found := uniqueMap[item]; !found {
				uniqueMap[item] = struct{}{}
				result = reflect.Append(result, reflect.ValueOf(item))
			}
		}
	}

	return result.Interface()
}

// Contains generic function to check if a collection (slice, array) contains an element
func ContainsZ(collection interface{}, element interface{}) bool {
	v := reflect.ValueOf(collection)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			if reflect.DeepEqual(v.Index(i).Interface(), element) {
				return true
			}
		}
	}

	return false
}

// Difference generic function to compute the difference between two collections (slices, arrays)
func DifferenceZ(collection1 interface{}, collection2 interface{}) interface{} {
	v1 := reflect.ValueOf(collection1)
	result := reflect.MakeSlice(v1.Type(), 0, 0)

	if v1.Kind() == reflect.Slice || v1.Kind() == reflect.Array {
		for i := 0; i < v1.Len(); i++ {
			item := v1.Index(i).Interface()
			if !ContainsZ(collection2, item) {
				result = reflect.Append(result, v1.Index(i))
			}
		}
	}

	return result.Interface()
}

// Intersection generic function to compute the intersection of two collections (slices, arrays)
func Intersection(collection1 interface{}, collection2 interface{}) interface{} {
	v1 := reflect.ValueOf(collection1)
	result := reflect.MakeSlice(v1.Type(), 0, 0)

	if v1.Kind() == reflect.Slice || v1.Kind() == reflect.Array {
		for i := 0; i < v1.Len(); i++ {
			item := v1.Index(i).Interface()
			if ContainsZ(collection2, item) {
				result = reflect.Append(result, v1.Index(i))
			}
		}
	}

	return result.Interface()
}

// Slice generic function to extract a sub-collection from a collection (slice, array) based on start and end indices.
func Slice(collection interface{}, start, end int) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if start < 0 {
			start = 0
		}
		if end > v.Len() {
			end = v.Len()
		}

		for i := start; i < end; i++ {
			result = reflect.Append(result, v.Index(i))
		}
	}

	return result.Interface()
}

// SliceWithIndices generic function to extract a sub-collection from a collection (slice, array) based on a list of indices.
func SliceWithIndices(collection interface{}, indices []int) interface{} {
	v := reflect.ValueOf(collection)
	result := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for _, index := range indices {
			if index >= 0 && index < v.Len() {
				result = reflect.Append(result, v.Index(index))
			}
		}
	}

	return result.Interface()
}

// GroupBy generic function to group elements in a collection (slice, array) by a specified key function.
func GroupByZ(collection interface{}, keyFunc func(value interface{}) interface{}) map[interface{}][]interface{} {
	v := reflect.ValueOf(collection)
	groups := make(map[interface{}][]interface{})

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			key := keyFunc(item)
			if _, found := groups[key]; !found {
				groups[key] = make([]interface{}, 0)
			}
			groups[key] = append(groups[key], item)
		}
	}

	return groups
}

// Partition generic function to partition elements in a collection (slice, array) based on a condition.
func Partition(collection interface{}, condition func(value interface{}) bool) (interface{}, interface{}) {
	v := reflect.ValueOf(collection)
	truePartition := reflect.MakeSlice(v.Type(), 0, 0)
	falsePartition := reflect.MakeSlice(v.Type(), 0, 0)

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Interface()
			if condition(item) {
				truePartition = reflect.Append(truePartition, reflect.ValueOf(item))
			} else {
				falsePartition = reflect.Append(falsePartition, reflect.ValueOf(item))
			}
		}
	}

	return truePartition.Interface(), falsePartition.Interface()
}

// Zip generic function to combine elements from multiple collections (slices, arrays) into tuples.
func Zip(collections ...interface{}) []interface{} {
	minLength := -1

	for _, collection := range collections {
		v := reflect.ValueOf(collection)
		if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
			return nil
		}

		if minLength == -1 || v.Len() < minLength {
			minLength = v.Len()
		}
	}

	result := make([]interface{}, minLength)

	for i := 0; i < minLength; i++ {
		tuple := make([]interface{}, len(collections))
		for j, collection := range collections {
			v := reflect.ValueOf(collection)
			tuple[j] = v.Index(i).Interface()
		}
		result[i] = tuple
	}

	return result
}

// ReduceRight generic function to perform a right-to-left reduction on a collection (slice, array).
func ReduceRight(collection interface{}, reducer func(acc, value interface{}) interface{}, initialValue interface{}) interface{} {
	v := reflect.ValueOf(collection)
	accumulator := initialValue

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		for i := v.Len() - 1; i >= 0; i-- {
			accumulator = reducer(accumulator, v.Index(i).Interface())
		}
	}

	return accumulator
}

// RotateLeft generic function to cyclically rotate elements in a collection (slice, array) to the left by a specified number of positions.
func RotateLeft(collection interface{}, positions int) interface{} {
	v := reflect.ValueOf(collection)
	length := v.Len()

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if positions < 0 {
			positions = (positions%length + length) % length
		} else {
			positions = positions % length
		}

		result := reflect.MakeSlice(v.Type(), length, length)
		for i := 0; i < length; i++ {
			result.Index((i - positions + length) % length).Set(v.Index(i))
		}
		return result.Interface()
	}

	return collection
}

// RotateRight generic function to cyclically rotate elements in a collection (slice, array) to the right by a specified number of positions.
func RotateRight(collection interface{}, positions int) interface{} {
	v := reflect.ValueOf(collection)
	length := v.Len()

	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if positions < 0 {
			positions = (-positions%length + length) % length
		} else {
			positions = positions % length
		}

		result := reflect.MakeSlice(v.Type(), length, length)
		for i := 0; i < length; i++ {
			result.Index((i + positions) % length).Set(v.Index(i))
		}
		return result.Interface()
	}

	return collection
}
