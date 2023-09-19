package builder

import (
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
	m.Result = value
	return m
}

func (m *MapBuilder) Add(key string, value interface{}) *MapBuilder {
	if utils.IsEmpty(key) {
		log.Panicf("Invalid key")
	}
	m.Result[key] = value
	return m
}

func (m *MapBuilder) Build() map[string]interface{} {
	return m.Result
}

func (m *MapBuilder) Size() int {
	return len(m.Result)
}

func (m *MapBuilder) Remove(key string) *MapBuilder {
	delete(m.Result, key)
	return m
}

func (m *MapBuilder) Contains(key string) bool {
	_, ok := m.Result[key]
	return ok
}

func (m *MapBuilder) Merge(value map[string]interface{}) *MapBuilder {
	if len(value) == 0 {
		return m
	}
	for k, v := range value {
		m.Result[k] = v
	}
	return m
}

func (m *MapBuilder) Json() string {
	return utils.ToJson(m.Result)
}

func (m *MapBuilder) Get(key string) (interface{}, bool) {
	v, ok := m.Result[key]
	return v, ok
}

func (m *MapBuilder) Keys() []string {
	if len(m.Result) == 0 {
		return []string{}
	}
	keys := make([]string, 0, len(m.Result))
	for k := range m.Result {
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
	if _, ok := m.Result[key]; ok {
		m.Result[key] = value
	}
	return m
}

func (m *MapBuilder) DeepMerge(values map[string]interface{}) *MapBuilder {
	coltx.DeepMergeMap(m.Result, values)
	return m
}

func (m *MapBuilder) Filter(callback func(key string, value interface{}) bool) *MapBuilder {
	filtered := make(map[string]interface{})
	for key, value := range m.Result {
		if callback(key, value) {
			filtered[key] = value
		}
	}
	m.SetResult(filtered)
	return m
}

func (m *MapBuilder) FilterWith(callback func(key string, value interface{}) bool) map[string]interface{} {
	filtered := make(map[string]interface{})
	for key, value := range m.Result {
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
		if value, exists := m.Result[key]; exists {
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
	pairs := make([]KeyValuePair, 0, len(m.Result))
	for key, value := range m.Result {
		pairs = append(pairs, *NewKeyValuePair().SetKey(key).SetValue(value))
	}
	return pairs
}

func (m *MapBuilder) DeserializeJson(jsonString string) (*MapBuilder, error) {
	err := utils.UnmarshalFromString(jsonString, &m.Result)
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
	value, ok := m.Result[key]
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
	for _, value := range m.Result {
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
	for _, value := range m.Result {
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
	value, ok := m.Result[key]
	if !ok {
		return false
	}
	_, is := value.([]interface{})
	return is
}

func (m *MapBuilder) IsObject(key string) bool {
	value, ok := m.Result[key]
	if !ok {
		return false
	}
	_, is := value.(map[string]interface{})
	return is
}

func (m *MapBuilder) ToBJSON(path string) bjson.BJsonContext {
	return bjson.Get(m.Json(), path)
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
		mb.Result = value.(map[string]interface{})
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
