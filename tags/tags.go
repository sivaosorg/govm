package tags

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func IsFieldZero(f reflect.Value) bool {
	zero := reflect.Zero(f.Type()).Interface()
	return reflect.DeepEqual(f.Interface(), zero)
}

func IsNoTraverseType(v reflect.Value) bool {
	if !IsStruct(v) {
		return false
	}
	t := DeepTypeOf(v)
	_, found := NoTraverseTypes[t]
	return found
}

func ValidateCopyField(f reflect.StructField, sfv, dfv reflect.Value) error {
	if !dfv.IsValid() {
		return ErrorFieldNotExists
	}
	if ConversionExists(sfv.Type(), dfv.Type()) {
		return nil
	}
	if (sfv.Kind() != dfv.Kind()) && !IsInterface(dfv) {
		return fmt.Errorf("Field: '%v', src [%v] & dst [%v] kind didn't match",
			f.Name,
			sfv.Kind(),
			dfv.Kind(),
		)
	}
	_sfv := DeepTypeOf(sfv)
	_dfv := DeepTypeOf(dfv)
	if (_sfv.Kind() == reflect.Slice || _sfv.Kind() == reflect.Map) && _sfv.Kind() == _dfv.Kind() && ConversionExists(_sfv.Elem(), _dfv.Elem()) {
		return nil
	}
	if (_sfv != _dfv) && !IsInterface(dfv) {
		return fmt.Errorf("Field: '%v', src [%v] & dst [%v] type didn't match",
			f.Name,
			_sfv,
			_dfv,
		)
	}
	return nil
}

func ModelFields(v reflect.Value) []reflect.StructField {
	v = indirect(v)
	t := v.Type()
	var fs []reflect.StructField
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Only exported fields of a struct can be accessed.
		// So, non-exported fields will be ignored
		if f.PkgPath == "" {
			fs = append(fs, f)
		}
	}
	return fs
}

func StructValue(s interface{}) (reflect.Value, error) {
	if s == nil {
		return reflect.Value{}, errors.New("Invalid input <nil>")
	}
	sv := indirect(valueOf(s))
	if !IsStruct(sv) {
		return reflect.Value{}, errors.New("Input is not a struct")
	}
	return sv, nil
}

func GetField(sv reflect.Value, name string) (reflect.Value, error) {
	field := sv.FieldByName(name)
	if !field.IsValid() {
		return reflect.Value{}, fmt.Errorf("Field: '%v', does not exists", name)
	}
	return field, nil
}

func ZeroOf(f reflect.Value) reflect.Value {
	v := reflect.Zero(f.Type())
	if f.Kind() == reflect.Ptr {
		return v
	}
	return indirect(valueOf(v.Interface()))
}

func DeepTypeOf(v reflect.Value) reflect.Type {
	if IsInterface(v) {
		if !IsFieldZero(v) {
			v = valueOf(v.Interface())
		}
	}
	return v.Type()
}

func valueOf(i interface{}) reflect.Value {
	return reflect.ValueOf(i)
}

func indirect(v reflect.Value) reflect.Value {
	return reflect.Indirect(v)
}

func isPointer(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr
}

func IsStruct(v reflect.Value) bool {
	if IsInterface(v) {
		v = valueOf(v.Interface())
	}
	pv := indirect(v)
	if pv.Kind() == reflect.Invalid {
		return false
	}
	return pv.Kind() == reflect.Struct
}

func IsInterface(v reflect.Value) bool {
	return v.Kind() == reflect.Interface
}

func ExtractType(x interface{}) reflect.Type {
	return reflect.TypeOf(x).Elem()
}

func ConversionExists(sourceType reflect.Type, destType reflect.Type) bool {
	if _, ok := ReflectConverters[sourceType]; !ok {
		return false
	}
	if _, ok := ReflectConverters[sourceType][destType]; !ok {
		return false
	}
	return true
}

// AddNoTraverseType method adds the Go Lang type into global `NoTraverseTypeList`.
// The type(s) from list is considered as "No Traverse" type
// for model mapping process. See also `RemoveNoTraverseType()` method.
//
//	model.AddNoTraverseType(time.Time{}, &time.Time{}, os.File{}, &os.File{})
//
// Default NoTraverseTypeList: time.Time{}, &time.Time{}, os.File{}, &os.File{},
// http.Request{}, &http.Request{}, http.Response{}, &http.Response{}
func AddNoTraverseType(i ...interface{}) {
	for _, v := range i {
		t := reflect.TypeOf(v)
		if _, ok := NoTraverseTypes[t]; ok {
			continue
		}
		NoTraverseTypes[t] = true
	}
}

// RemoveNoTraverseType method is used to remove Go Lang type from the `NoTraverseTypeList`.
// See also `AddNoTraverseType()` method.
//
//	model.RemoveNoTraverseType(http.Request{}, &http.Request{})
func RemoveNoTraverseType(i ...interface{}) {
	for _, v := range i {
		t := reflect.TypeOf(v)
		if _, ok := NoTraverseTypes[t]; ok {
			delete(NoTraverseTypes, t)
		}
	}
}

// AddConversion method allows registering a custom `Converter` into the global `converterMap`
// by supplying pointers of the target types.
func AddConversion(in interface{}, out interface{}, converter TagConverter) {
	srcType := ExtractType(in)
	targetType := ExtractType(out)
	AddConversionByType(srcType, targetType, converter)
}

// AddConversionByType allows registering a custom `Converter` into global `converterMap` by types.
func AddConversionByType(sourceType reflect.Type, targetType reflect.Type, converter TagConverter) {
	if _, ok := ReflectConverters[sourceType]; !ok {
		ReflectConverters[sourceType] = map[reflect.Type]TagConverter{}
	}
	ReflectConverters[sourceType][targetType] = converter
}

// RemoveConversion registered conversions
func RemoveConversion(in interface{}, out interface{}) {
	sourceType := ExtractType(in)
	targetType := ExtractType(out)
	if _, ok := ReflectConverters[sourceType]; !ok {
		return
	}
	if _, ok := ReflectConverters[sourceType][targetType]; !ok {
		return
	}
	delete(ReflectConverters[sourceType], targetType)
}

// IsZero method returns `true` if all the exported fields in a given `struct`
// are zero value otherwise `false`. If input is not a struct, method returns `false`.
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	// Field is ignored by processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
//
// A "model" tag value with the option of "no_traverse"; library will not traverse
// inside the struct object. However, the field value will be evaluated whether
// it's zero value or not.
//
//	Example:
//
// Field is not traversed but value is evaluated/processed
//
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,no_traverse"`
//	Region		BookLocale	`defined:",no_traverse"`
func IsZero(s interface{}) bool {
	if s == nil {
		return true
	}
	sv, err := StructValue(s)
	if err != nil {
		return false
	}
	fields := ModelFields(sv)
	for _, f := range fields {
		fv := sv.FieldByName(f.Name)
		tag := NewTag(f.Tag.Get(TagName))
		if tag.isOmitField() {
			continue
		}
		if IsStruct(fv) {
			if IsNoTraverseType(fv) || tag.isNoTraverse() {
				if !IsFieldZero(fv) {
					return false
				}
				continue
			}
			if !IsZero(fv.Interface()) {
				return false
			}
			continue
		}
		if !IsFieldZero(fv) {
			return false
		}
	}
	return true
}

// IsZeroInFields method verifies the value for the given list of field names against
// given struct. Method returns `Field Name` and `true` for the zero value field.
// Otherwise method returns empty `string` and `false`.
//
// Note:
// [1] This method doesn't traverse nested and embedded `struct`, instead it just evaluates that `struct`.
// [2] If given field is not exists in the struct, method moves on to next field
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	Field is ignored by go-model processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
func IsZeroInFields(s interface{}, names ...string) (string, bool) {
	if s == nil || len(names) == 0 {
		return "", true
	}
	sv, err := StructValue(s)
	if err != nil {
		return "", false
	}
	for _, name := range names {
		fv := sv.FieldByName(name)
		if !fv.IsValid() {
			continue
		}
		if IsFieldZero(fv) {
			return name, true
		}
	}
	return "", false
}

// HasZero method returns `true` if any one of the exported fields in a given
// `struct` is zero value otherwise `false`. If input is not a struct, method
// returns `false`.
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	// Field is ignored by go-model processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
//
// A "defined" tag value with the option of "no_traverse"; library will not traverse
// inside the struct object. However, the field value will be evaluated whether
// it's zero value or not.
//
//	Example:
//
// Field is not traversed but value is evaluated/processed
//
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,no_traverse"`
//	Region		BookLocale	`defined:",no_traverse"`
func HasZero(s interface{}) bool {
	if s == nil {
		return true
	}
	sv, err := StructValue(s)
	if err != nil {
		return false
	}
	fields := ModelFields(sv)
	for _, f := range fields {
		fv := sv.FieldByName(f.Name)
		tag := NewTag(f.Tag.Get(TagName))
		if tag.isOmitField() {
			continue
		}
		if IsStruct(fv) {
			if IsNoTraverseType(fv) || tag.isNoTraverse() {
				if IsFieldZero(fv) {
					return true
				}
				continue
			}
			if HasZero(fv.Interface()) {
				return true
			}
			continue
		}
		if IsFieldZero(fv) {
			return true
		}
	}
	return false
}

// Copy method copies all the exported field values from source `struct` into destination `struct`.
// The "Name", "Type" and "Kind" is should match to qualify a copy. One exception though;
// if the destination field type is "interface{}" then "Type" and "Kind" doesn't matter,
// source value gets copied to that destination field.
//
//	Example:
//
//	src := SampleStruct { /* source struct field values go here */ }
//	dst := SampleStruct {}
//
//	errs := model.Copy(&dst, src)
//	if errs != nil {
//		fmt.Println("Errors:", errs)
//	}
//
// Note:
// [1] Copy process continues regardless of the case it qualifies or not. The non-qualified field(s)
// gets added to '[]error' that you will get at the end.
// [2] Two dimensional slice type is not supported yet.
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	// Field is ignored while processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
//
// A "defined" tag value with the option of "omitempty"; library will not copy those values
// into destination struct object. It may be handy for partial put or patch update
// request scenarios; if you don't want to copy empty/zero value into destination object.
//
//	Example:
//
//	// Field is not copy into 'dst' if it's empty/zero value
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,omitempty"`
//	Region		BookLocale	`defined:",omitempty,no_traverse"`
//
// A "defined" tag value with the option of "no_traverse"; library will not traverse
// inside the struct object. However, the field value will be evaluated whether
// it's zero value or not, and then copied to the destination object accordingly.
//
//	Example:
//
// Field is not traversed but value is evaluated/processed
//
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,no_traverse"`
//	Region		BookLocale	`defined:",no_traverse"`
func Copy(target, source interface{}) []error {
	var errs []error
	if source == nil || target == nil {
		return append(errs, errors.New("Source or Destination is nil"))
	}
	sv := valueOf(source)
	dv := valueOf(target)
	if !IsStruct(sv) || !IsStruct(dv) {
		return append(errs, errors.New("Source or Destination is not a struct"))
	}
	if !isPointer(dv) {
		return append(errs, errors.New("Destination struct is not a pointer"))
	}
	if IsZero(source) {
		return append(errs, errors.New("Source struct is empty"))
	}
	errs = doCopy(dv, sv)
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// Clone method creates a clone of given `struct` object.
// So all field values you get in the result.
//
//	Example:
//	input := SampleStruct { /* input struct field values go here */ }
//
//	clonedObj := tags.Clone(input)
//
//	fmt.Printf("\nCloned Object: %#v\n", clonedObj)
//
// Note:
// [1] Two dimensional slice type is not supported yet.
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	// Field is ignored while processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
//
// A "defined" tag value with the option of "omitempty"; library will not clone those values
// into result struct object.
//
//	Example:
//
//	Field is not cloned into 'result' if it's empty/zero value
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,omitempty"`
//	Region		BookLocale	`defined:",omitempty,no_traverse"`
//
// A "defined" tag value with the option of "no_traverse"; library will not traverse
// inside the struct object. However, the field value will be evaluated whether
// it's zero value or not, and then cloned to the result accordingly.
//
//	Example:
//
// Field is not traversed but value is evaluated/processed
//
//	ArchiveInfo	BookArchive	`defined:"archiveInfo,no_traverse"`
//	Region		BookLocale	`defined:",no_traverse"`
func Clone(s interface{}) (interface{}, error) {
	sv, err := StructValue(s)
	if err != nil {
		return nil, err
	}
	st := DeepTypeOf(sv)
	dv := reflect.New(st)
	doCopy(dv, sv)
	return dv.Interface(), nil
}

// Map method converts all the exported field values from the given `struct`
// into `map[string]interface{}`. In which the keys of the map are the field names
// and the values of the map are the associated values of the field.
//
//	Example:
//
//	src := SampleStruct { /* source struct field values go here */ }
//
//	err := model.Map(src)
//	if err != nil {
//		fmt.Println("Error:", err)
//	}
//
// Note:
// [1] Two dimensional slice type is not supported yet.
//
// The default 'Key Name' string is the struct field name. However, it can be
// changed in the struct field's tag value via "defined" tag.
//
//	Example:
//
//	// Now field 'Key Name' is customized
//	BookTitle	string	`defined:"bookTitle"`
//
// A "defined" tag with the value of "-" is ignored by library for processing.
//
//	Example:
//
//	// Field is ignored while processing
//	BookCount	int	`defined:"-"`
//	BookCode	string	`defined:"-"`
//
// A "defined" tag value with the option of "omitempty"; library will not include those values
// while converting to map[string]interface{}. If it's empty/zero value.
//
//	Example:
//
//	Field is not included in result map if it's empty/zero value
//	ArchivedDate	time.Time	`defined:"archivedDate,omitempty"`
//	Region		BookLocale	`defined:",omitempty,no_traverse"`
//
// A "defined" tag value with the option of "no_traverse"; library will not traverse
// inside the struct object. However, the field value will be evaluated whether
// it's zero value or not, and then added to the result map accordingly.
//
//	Example:
//
//	Field is not traversed but value is evaluated/processed
//	ArchivedDate	time.Time	`defined:"archivedDate,no_traverse"`
//	Region		BookLocale	`defined:",no_traverse"`
func Map(s interface{}) (map[string]interface{}, error) {
	sv, err := StructValue(s)
	if err != nil {
		return nil, err
	}
	return doMap(sv), nil
}

// Fields method returns the exported struct fields from the given `struct`.
//
//	Example:
//
//	src := SampleStruct { /* source struct field values go here */ }
//
//	fields, _ := model.Fields(src)
//	for _, f := range fields {
//		tag := newTag(f.Tag.Get("model"))
//		fmt.Println("Field Name:", f.Name, "Tag Name:", tag.Name, "Tag Options:", tag.Options)
//	}
func Fields(s interface{}) ([]reflect.StructField, error) {
	sv, err := StructValue(s)
	if err != nil {
		return nil, err
	}
	return ModelFields(sv), nil
}

// Kind method returns `reflect.Kind` for the given field name from the `struct`.
//
//	Example:
//
//	src := SampleStruct {
//		BookCount      int         `json:"-"`
//		BookCode       string      `json:"-"`
//		ArchiveInfo    BookArchive `json:"archive_info,omitempty"`
//		Region         BookLocale  `json:"region,omitempty"`
//	}
//
//	fieldKind, _ := model.Kind(src, "ArchiveInfo")
//	fmt.Println("Field kind:", fieldKind)
func Kind(s interface{}, name string) (reflect.Kind, error) {
	sv, err := StructValue(s)
	if err != nil {
		return reflect.Invalid, err
	}
	fv, err := GetField(sv, name)
	if err != nil {
		return reflect.Invalid, err
	}
	return fv.Type().Kind(), nil
}

// Get method returns a field value from `struct` by field name.
//
//	Example:
//
//	src := SampleStruct {
//		BookCount      int         `json:"-"`
//		BookCode       string      `json:"-"`
//		ArchiveInfo    BookArchive `json:"archive_info,omitempty"`
//		Region         BookLocale  `json:"region,omitempty"`
//	}
//
//	value, err := model.Get(src, "ArchiveInfo")
//	fmt.Println("Field Value:", value)
//	fmt.Println("Error:", err)
//
// Note: Get method does not honor model tag annotations. Get simply access
// value on exported fields.
func Get(s interface{}, name string) (interface{}, error) {
	sv, err := StructValue(s)
	if err != nil {
		return nil, err
	}
	fv, err := GetField(sv, name)
	if err != nil {
		return nil, err
	}
	return fv.Interface(), nil
}

// Set method sets a value into field on struct by field name.
//
//	Example:
//
//	src := SampleStruct {
//		BookCount      int         `json:"-"`
//		BookCode       string      `json:"-"`
//		ArchiveInfo    BookArchive `json:"archive_info,omitempty"`
//		Region         BookLocale  `json:"region,omitempty"`
//	}
//
//	bookLocale := BookLocale {
//		Locale: "en-US",
//		Language: "en",
//		Region: "US",
//	}
//
//	err := model.Set(&src, "Region", bookLocale)
//	fmt.Println("Error:", err)
//
// Note: Set method does not honor model tag annotations. Set simply given
// value by field name on exported fields.
func Set(s interface{}, name string, value interface{}) error {
	if s == nil {
		return errors.New("Invalid input <nil>")
	}
	sv := valueOf(s)
	if isPointer(sv) {
		sv = sv.Elem()
	} else {
		return errors.New("Destination struct is not a pointer")
	}
	fv, err := GetField(sv, name)
	if err != nil {
		return err
	}
	if !fv.CanSet() {
		return fmt.Errorf("Field: %v, cannot be settable", name)
	}
	tv := valueOf(value)
	if isPointer(tv) {
		tv = tv.Elem()
	}
	if (fv.Kind() != tv.Kind()) || fv.Type() != tv.Type() {
		return fmt.Errorf("Field: %v, type/kind did not match", name)
	}
	fv.Set(tv)
	return nil
}

func init() {
	NoTraverseTypes = map[reflect.Type]bool{}
	ReflectConverters = map[reflect.Type]map[reflect.Type]TagConverter{}
	// Default NoTraverseTypeList
	// --------------------------
	// Auto No Traverse struct list for not traversing Deep Level
	// However, field value will be evaluated/processed by go-model library
	AddNoTraverseType(
		time.Time{},
		&time.Time{},
		os.File{},
		&os.File{},
		http.Request{},
		&http.Request{},
		http.Response{},
		&http.Response{},
	)
}

func doCopy(dv, sv reflect.Value) []error {
	dv = indirect(dv)
	sv = indirect(sv)
	fields := ModelFields(sv)
	var errs []error
	for _, f := range fields {
		sfv := sv.FieldByName(f.Name)
		tag := NewTag(f.Tag.Get(TagName))
		if tag.isOmitField() {
			continue
		}
		noTraverse := (IsNoTraverseType(sfv) || tag.isNoTraverse())
		var isValue bool
		if IsStruct(sfv) && !noTraverse {
			isValue = !IsZero(sfv.Interface())
		} else {
			isValue = !IsFieldZero(sfv)
		}
		dfv := dv.FieldByName(f.Name)
		err := ValidateCopyField(f, sfv, dfv)
		if err != nil {
			if err != ErrorFieldNotExists {
				errs = append(errs, err)
			}
			continue
		}
		if !isValue {
			// field value is zero and check 'omitempty' option present
			// then don't copy into destination struct
			// otherwise copy to dst
			if !tag.isOmitEmpty() {
				dfv.Set(ZeroOf(dfv))
			}
			continue
		}
		if dfv.CanSet() {
			if IsStruct(sfv) {
				v, innerErrs := copyVal(dfv.Type(), sfv, noTraverse)
				errs = append(errs, innerErrs...)
				dfv.Set(v)
			} else {
				v, err := copyVal(dfv.Type(), sfv, false)
				errs = append(errs, err...)
				dfv.Set(v)
			}
		}
	}

	return errs
}

func doMap(sv reflect.Value) map[string]interface{} {
	sv = indirect(sv)
	fields := ModelFields(sv)
	m := map[string]interface{}{}
	for _, f := range fields {
		fv := sv.FieldByName(f.Name)
		tag := NewTag(f.Tag.Get(TagName))
		if tag.isOmitField() {
			continue
		}
		keyName := f.Name
		if !utils.IsEmpty(tag.Name) {
			keyName = tag.Name
		}
		noTraverse := (IsNoTraverseType(fv) || tag.isNoTraverse())
		var isValue bool
		if IsStruct(fv) && !noTraverse {
			isValue = !IsZero(fv.Interface())
		} else {
			isValue = !IsFieldZero(fv)
		}
		if !isValue {
			if !tag.isOmitEmpty() {
				m[keyName] = ZeroOf(fv).Interface()
			}
			continue
		}
		if IsStruct(fv) {
			if noTraverse {
				// This is struct kind and it's present in NoTraverseTypes or
				// has 'no_traverse' tag option.
				// however will take care of field value
				m[keyName] = mapVal(fv, true).Interface()
			} else {
				fmv := doMap(fv)
				if f.Anonymous {
					for k, v := range fmv {
						m[k] = v
					}
				} else {
					m[keyName] = fmv
				}
			}
			continue
		}
		m[keyName] = mapVal(fv, false).Interface()
	}
	return m
}

func copyVal(dt reflect.Type, f reflect.Value, _noTraverse bool) (reflect.Value, []error) {
	var (
		_pointer bool
		nf       reflect.Value
		errs     []error
	)
	if ConversionExists(f.Type(), dt) && !_noTraverse {
		res, err := ReflectConverters[f.Type()][dt](f)
		if err != nil {
			errs = append(errs, err)
		}
		return res, errs
	}

	if IsInterface(f) {
		f = valueOf(f.Interface())
	}
	if isPointer(f) {
		_pointer = true
		f = f.Elem()
	}
	switch f.Kind() {
	case reflect.Struct:
		if _noTraverse {
			nf = f
		} else {
			nf = reflect.New(f.Type())
			doCopy(nf, f)
			nf = nf.Elem()
		}
	case reflect.Map:
		if dt.Kind() == reflect.Ptr {
			dt = dt.Elem()
		}
		nf = reflect.MakeMap(dt)

		for _, key := range f.MapKeys() {
			ov := f.MapIndex(key)

			cv := reflect.New(dt.Elem()).Elem()
			v, err := copyVal(dt.Elem(), ov, IsNoTraverseType(ov))
			if len(err) > 0 {
				errs = append(errs, err...)
			} else {
				cv.Set(v)
				nf.SetMapIndex(key, cv)
			}
		}
	case reflect.Slice:
		if f.Type() == TypeOfBytes {
			nf = f
		} else {
			if dt.Kind() == reflect.Ptr {
				dt = dt.Elem()
			}
			nf = reflect.MakeSlice(dt, f.Len(), f.Cap())

			for i := 0; i < f.Len(); i++ {
				ov := f.Index(i)

				cv := reflect.New(dt.Elem()).Elem()
				v, err := copyVal(dt.Elem(), ov, IsNoTraverseType(ov))
				if len(err) > 0 {
					errs = append(errs, err...)
				} else {
					cv.Set(v)
					nf.Index(i).Set(cv)
				}
			}
		}
	default:
		nf = f
	}
	if _pointer {
		o := reflect.New(nf.Type())
		o.Elem().Set(nf)
		return o, errs
	}
	return nf, errs
}

func mapVal(f reflect.Value, _noTraverse bool) reflect.Value {
	var (
		_pointer bool
		nf       reflect.Value
	)
	if IsInterface(f) {
		f = valueOf(f.Interface())
	}
	if isPointer(f) {
		_pointer = true
		f = f.Elem()
	}
	switch f.Kind() {
	case reflect.Struct:
		if _noTraverse {
			nf = f
		} else {
			nf = valueOf(doMap(f))
		}
	case reflect.Map:
		nmv := map[string]interface{}{}
		for _, key := range f.MapKeys() {
			_key := fmt.Sprintf("%v", key.Interface())
			mv := f.MapIndex(key)
			nv := mapVal(mv, IsNoTraverseType(mv))
			nmv[_key] = nv.Interface()
		}

		nf = valueOf(nmv)
	case reflect.Slice:
		if f.Type() == TypeOfBytes {
			nf = f
		} else {
			if f.Len() > 0 {
				fsv := f.Index(0)
				if IsStruct(fsv) {
					nf = reflect.MakeSlice(reflect.SliceOf(TypeOfInterface), f.Len(), f.Cap())
				} else {
					nf = reflect.MakeSlice(f.Type(), f.Len(), f.Cap())
				}
				for i := 0; i < f.Len(); i++ {
					sv := f.Index(i)
					var dv reflect.Value
					if IsStruct(sv) {
						dv = reflect.New(TypeOfInterface).Elem()
					} else {
						dv = reflect.New(sv.Type()).Elem()
					}
					dv.Set(mapVal(sv, IsNoTraverseType(sv)))
					nf.Index(i).Set(dv)
				}
			}
		}
	default:
		nf = f
	}
	if _pointer {
		o := reflect.New(nf.Type())
		o.Elem().Set(nf)
		return o
	}
	return nf
}

// Tag method returns the exported struct field `Tag` value from the given struct.
//
//	Example:
//
//	src := SampleStruct {
//		BookCount      int         `json:"-"`
//		BookCode       string      `json:"-"`
//		ArchiveInfo    BookArchive `json:"archive_info,omitempty"`
//		Region         BookLocale  `json:"region,omitempty"`
//	}
//
//	tag, _ := model.Tag(src, "ArchiveInfo")
//	fmt.Println("Tag Value:", tag.Get("json"))
//
//	// Output:
//	Tag Value: archive_info,omitempty
func Tag(s interface{}, name string) (reflect.StructTag, error) {
	sv, err := StructValue(s)
	if err != nil {
		return "", err
	}
	if fv, ok := sv.Type().FieldByName(name); ok {
		return fv.Tag, nil
	}
	return "", fmt.Errorf("Field: '%v', does not exists", name)
}

// Tags method returns the exported struct fields `Tag` value from the given struct.
//
//	Example:
//
//	src := SampleStruct {
//		BookCount      int         `json:"-"`
//		BookCode       string      `json:"-"`
//		ArchiveInfo    BookArchive `json:"archive_info,omitempty"`
//		Region         BookLocale  `json:"region,omitempty"`
//	}
//
//	tags, _ := model.Tags(src)
//	fmt.Println("Tags:", tags)
func Tags(s interface{}) (map[string]reflect.StructTag, error) {
	sv, err := StructValue(s)
	if err != nil {
		return nil, err
	}
	tags := map[string]reflect.StructTag{}
	fields := ModelFields(sv)
	for _, f := range fields {
		tags[f.Name] = f.Tag
	}
	return tags, nil
}

func NewTag(modelTag string) *TagConfig {
	t := TagConfig{}
	values := strings.Split(modelTag, ",")
	t.Name = values[0]
	t.Options = strings.Join(values[1:], ",")
	return &t
}

func (t *TagConfig) isOmitField() bool {
	return t.Name == OmitField
}

func (t *TagConfig) isOmitEmpty() bool {
	return t.isExists(OmitEmpty)
}

func (t *TagConfig) isNoTraverse() bool {
	return t.isExists(NoTraverse)
}

func (t *TagConfig) isExists(opt string) bool {
	return strings.Contains(t.Options, opt)
}
