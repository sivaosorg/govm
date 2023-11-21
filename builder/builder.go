package builder

import (
	"fmt"
	"log"
	"reflect"
	"sort"

	"github.com/sivaosorg/govm/bjson"
	"github.com/sivaosorg/govm/coltx"
	"github.com/sivaosorg/govm/utils"
)

func NewMapBuilder() *MapBuilder {
	m := &MapBuilder{}
	m.SetResult(make(map[string]interface{}))
	return m
}

func (m *MapBuilder) SetResult(value map[string]interface{}) *MapBuilder {
	m.result = value
	return m
}

func (m *MapBuilder) Add(key string, value interface{}) *MapBuilder {
	if utils.IsEmpty(key) {
		log.Panicf("Invalid key")
	}
	m.result[key] = value
	return m
}

func (m *MapBuilder) Build() map[string]interface{} {
	return m.result
}

func (m *MapBuilder) Size() int {
	return len(m.result)
}

func (m *MapBuilder) Remove(key string) *MapBuilder {
	delete(m.result, key)
	return m
}

func (m *MapBuilder) Contains(key string) bool {
	_, ok := m.result[key]
	return ok
}

func (m *MapBuilder) Merge(value map[string]interface{}) *MapBuilder {
	if len(value) == 0 {
		return m
	}
	for k, v := range value {
		m.result[k] = v
	}
	return m
}

func (m *MapBuilder) Json() string {
	return utils.ToJson(m.result)
}

func (m *MapBuilder) Get(key string) (interface{}, bool) {
	v, ok := m.result[key]
	return v, ok
}

func (m *MapBuilder) Keys() []string {
	if len(m.result) == 0 {
		return []string{}
	}
	keys := make([]string, 0, len(m.result))
	for k := range m.result {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (m *MapBuilder) Reset() *MapBuilder {
	m.SetResult(make(map[string]interface{}))
	return m
}

func (m *MapBuilder) Update(key string, value interface{}) *MapBuilder {
	if utils.IsEmpty(key) {
		return m
	}
	if _, ok := m.result[key]; ok {
		m.result[key] = value
	}
	return m
}

func (m *MapBuilder) DeepMerge(values map[string]interface{}) *MapBuilder {
	coltx.DeepMergeMap(m.result, values)
	return m
}

func (m *MapBuilder) Filter(callback func(key string, value interface{}) bool) *MapBuilder {
	filtered := make(map[string]interface{})
	for key, value := range m.result {
		if callback(key, value) {
			filtered[key] = value
		}
	}
	m.SetResult(filtered)
	return m
}

func (m *MapBuilder) FilterWith(callback func(key string, value interface{}) bool) map[string]interface{} {
	filtered := make(map[string]interface{})
	for key, value := range m.result {
		if callback(key, value) {
			filtered[key] = value
		}
	}
	return filtered
}

func (m *MapBuilder) IsEmpty() bool {
	return m.Size() == 0
}

func (m *MapBuilder) SubMap(keys []string) map[string]interface{} {
	subs := make(map[string]interface{})
	for _, key := range keys {
		if value, exists := m.result[key]; exists {
			subs[key] = value
		}
	}
	return subs
}

func NewKeyValuePair() *KeyValuePair {
	k := &KeyValuePair{}
	return k
}

func (k *KeyValuePair) SetKey(value string) *KeyValuePair {
	k.Key = value
	return k
}

func (k *KeyValuePair) SetValue(value interface{}) *KeyValuePair {
	k.Value = value
	return k
}

func (k *KeyValuePair) Json() string {
	return utils.ToJson(k)
}

func (m *MapBuilder) ToKeyValuePairs() []KeyValuePair {
	pairs := make([]KeyValuePair, 0, len(m.result))
	for key, value := range m.result {
		pairs = append(pairs, *NewKeyValuePair().SetKey(key).SetValue(value))
	}
	return pairs
}

func (m *MapBuilder) DeserializeJson(jsonString string) (*MapBuilder, error) {
	err := utils.UnmarshalFromString(jsonString, &m.result)
	if err != nil {
		return m, err
	}
	return m, nil
}

func (m *MapBuilder) DeserializeJsonI(value interface{}) (*MapBuilder, error) {
	v, ok := value.(MapBuilder)
	if ok {
		return m.DeserializeJson(v.Json())
	}
	return m.DeserializeJson(utils.ToJson(value))
}

func (m *MapBuilder) IsNumericValue(key string) bool {
	if utils.IsEmpty(key) {
		return false
	}
	value, ok := m.result[key]
	if !ok {
		return false
	}
	_, isFloat := value.(float64)
	_, isInt := value.(int)
	return isFloat || isInt
}

func (m *MapBuilder) MaxNumericValue() (float64, bool) {
	max := -1.0
	found := false
	for _, value := range m.result {
		if floatValue, ok := value.(float64); ok {
			if !found || floatValue > max {
				max = floatValue
				found = true
			}
		} else if intValue, ok := value.(int); ok {
			if !found || float64(intValue) > max {
				max = float64(intValue)
				found = true
			}
		}
	}
	return max, found
}

func (m *MapBuilder) MinNumericValue() (float64, bool) {
	min := -1.0
	found := false
	for _, value := range m.result {
		if floatValue, ok := value.(float64); ok {
			if !found || floatValue < min {
				min = floatValue
				found = true
			}
		} else if intValue, ok := value.(int); ok {
			if !found || float64(intValue) < min {
				min = float64(intValue)
				found = true
			}
		}
	}
	return min, found
}

func (m *MapBuilder) IsArray(key string) bool {
	value, ok := m.result[key]
	if !ok {
		return false
	}
	_, is := value.([]interface{})
	return is
}

func (m *MapBuilder) IsObject(key string) bool {
	value, ok := m.result[key]
	if !ok {
		return false
	}
	_, is := value.(map[string]interface{})
	return is
}

func (m *MapBuilder) ToBJSON(path string) bjson.BJsonContext {
	return bjson.Get(m.Json(), path)
}

// NewMatrixBuilder creates a new MatrixBuilder.
func NewMatrixBuilder() *MatrixBuilder {
	return &MatrixBuilder{}
}

// AddRow adds a new row to the matrix.
func (b *MatrixBuilder) AddRow(row ...interface{}) *MatrixBuilder {
	b.matrix = append(b.matrix, row)
	return b
}

// Build returns the final [][]interface{} matrix.
func (b *MatrixBuilder) Build() [][]interface{} {
	return b.matrix
}

// Clear resets the matrix to an empty state.
func (b *MatrixBuilder) Clear() *MatrixBuilder {
	b.matrix = nil
	return b
}

// Rows returns the number of rows in the matrix.
func (b *MatrixBuilder) Rows() int {
	return len(b.matrix)
}

// Cols returns the number of columns in the matrix.
func (b *MatrixBuilder) Cols() int {
	if len(b.matrix) == 0 {
		return 0
	}
	return len(b.matrix[0])
}

// PrintMatrix prints the matrix in a formatted way.
func (b *MatrixBuilder) PrintMatrix() {
	for _, row := range b.matrix {
		for _, val := range row {
			fmt.Printf("%-10v", val)
		}
		fmt.Println()
	}
}

// GetElement returns the value at the specified row and column.
func (b *MatrixBuilder) GetElement(row, col int) (interface{}, error) {
	if row < 0 || row >= len(b.matrix) || col < 0 || col >= len(b.matrix[0]) {
		return nil, fmt.Errorf("index out of bounds")
	}
	return b.matrix[row][col], nil
}

// SetElement sets the value at the specified row and column.
func (b *MatrixBuilder) SetElement(row, col int, value interface{}) error {
	if row < 0 || row >= len(b.matrix) || col < 0 || col >= len(b.matrix[0]) {
		return fmt.Errorf("index out of bounds")
	}
	b.matrix[row][col] = value
	return nil
}

// AddMatrix adds the given matrix to the current matrix.
func (b *MatrixBuilder) AddMatrix(other [][]interface{}) error {
	if len(b.matrix) != len(other) || len(b.matrix[0]) != len(other[0]) {
		return fmt.Errorf("matrix dimensions do not match")
	}
	for i := range b.matrix {
		for j := range b.matrix[i] {
			switch b.matrix[i][j].(type) {
			case int, float64, string:
				switch b.matrix[i][j].(type) {
				case int:
					b.matrix[i][j] = b.matrix[i][j].(int) + other[i][j].(int)
				case int32:
					b.matrix[i][j] = b.matrix[i][j].(int32) + other[i][j].(int32)
				case int64:
					b.matrix[i][j] = b.matrix[i][j].(int64) + other[i][j].(int64)
				case float32:
					b.matrix[i][j] = b.matrix[i][j].(float32) + other[i][j].(float32)
				case float64:
					b.matrix[i][j] = b.matrix[i][j].(float64) + other[i][j].(float64)
				case string:
					b.matrix[i][j] = fmt.Sprintf("%v%v", b.matrix[i][j], other[i][j])
				}
			default:
				return fmt.Errorf("unsupported type for addition")
			}
		}
	}
	return nil
}

// Transpose transposes the matrix, switching its rows and columns.
func (b *MatrixBuilder) Transpose() *MatrixBuilder {
	transposed := NewMatrixBuilder()
	// Initialize transposed matrix with the correct dimensions
	for j := 0; j < b.Cols(); j++ {
		transposed.matrix = append(transposed.matrix, make([]interface{}, b.Rows()))
	}
	// Populate transposed matrix
	for i := 0; i < b.Rows(); i++ {
		for j := 0; j < b.Cols(); j++ {
			transposed.matrix[j][i] = b.matrix[i][j]
		}
	}
	return transposed
}

// Clone creates a deep copy of the matrix.
func (b *MatrixBuilder) Clone() *MatrixBuilder {
	cloned := NewMatrixBuilder()
	// Populate cloned matrix with the same values
	for _, row := range b.matrix {
		cloned.matrix = append(cloned.matrix, append([]interface{}{}, row...))
	}
	return cloned
}

// asMap recursively converts an interface{} to a MapBuilder if it's a map or struct.
// Returns the MapBuilder representation or nil if not a map or struct.
func asMap(value interface{}) (*MapBuilder, bool) {
	if value == nil {
		return nil, false
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Map && v.Type().Key().Kind() == reflect.String {
		mb := NewMapBuilder()
		mb.result = value.(map[string]interface{})
		return mb, true
	}
	if v.Kind() == reflect.Struct {
		// It's a struct, convert it to a MapBuilder
		mb := NewMapBuilder()
		_type := v.Type()
		for i := 0; i < v.NumField(); i++ {
			fName := _type.Field(i).Name
			fValue := v.Field(i).Interface()
			// Recursively convert fieldValue to MapBuilder
			fields, _ := asMap(fValue)
			// Merge the field MapBuilder into the parent MapBuilder
			if fields != nil && !fields.IsEmpty() {
				mb.Merge(map[string]interface{}{fName: fields.Build()})
			} else {
				mb.Merge(map[string]interface{}{fName: fValue})
			}
		}
		if mb.IsEmpty() {
			return nil, false
		}
		return mb, true
	}
	return nil, false
}
