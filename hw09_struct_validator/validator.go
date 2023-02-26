package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	tag string = "validate"
)

var (
	ErrIsNotStruct           = errors.New("is not struct")
	ErrValueLessMin          = errors.New("value less min")
	ErrValueMoreMax          = errors.New("value more max")
	ErrValueNotInRange       = errors.New("value not in range")
	ErrValueLenLess          = errors.New("len string value less")
	ErrValueNotInRegexp      = errors.New("value not in regexp")
	ErrValueNotInRangeString = errors.New("value not in range string")
	ErrValidationTag         = errors.New("tag with an error")
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	errorString := strings.Builder{}
	for i, vv := range v {
		errorString.WriteString(vv.Field)
		errorString.WriteString(" - ")
		errorString.WriteString(vv.Err.Error())
		if i < len(v)-1 {
			errorString.WriteString(", ")
		}
	}
	return errorString.String()
}

func newValidationError(f string, e error) ValidationError {
	return ValidationError{
		Field: f,
		Err:   e,
	}
}

func checkString(rs reflect.StructField, v string) (fValidationErrors ValidationErrors, err error) {
	tagValue, ok := rs.Tag.Lookup(tag)
	if !ok {
		return fValidationErrors, nil
	}

	for _, tmp := range strings.Split(tagValue, "|") {
		tmpRule := strings.Split(tmp, ":")
		ruleK := tmpRule[0]
		ruleV := tmpRule[1]
		switch ruleK {
		case "len":
			ruleValue, err := strconv.Atoi(ruleV)
			if err != nil {
				return nil, err
			}
			if len(v) != ruleValue {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueLenLess))
			}
		case "regexp":
			matched, err := regexp.Match(ruleV, []byte(v))
			if err != nil {
				return nil, err
			}
			if !matched {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueNotInRegexp))
			}
		case "in":
			if !strings.Contains(ruleV, v) {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueNotInRangeString))
			}
		}
	}
	return fValidationErrors, nil
}

func checkInt(rs reflect.StructField, v int) (fValidationErrors ValidationErrors, err error) {
	tagValue, ok := rs.Tag.Lookup(tag)
	if !ok {
		return fValidationErrors, nil
	}
	for _, tmp := range strings.Split(tagValue, "|") {
		tmpRule := strings.Split(tmp, ":")
		ruleK, ruleV := tmpRule[0], tmpRule[1]
		switch ruleK {
		case "min":
			ruleValue, err := strconv.Atoi(ruleV)
			if err != nil {
				return nil, err
			}
			if v < ruleValue {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueLessMin))
			}
		case "max":
			ruleValue, err := strconv.Atoi(ruleV)
			if err != nil {
				return nil, err
			}
			if v > ruleValue {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueMoreMax))
			}
		case "in":
			ruleIn := strings.Split(ruleV, ",")
			if len(ruleIn) != 2 {
				return nil, ErrValidationTag
			}
			rmin, _ := strconv.Atoi(ruleIn[0])
			rmax, _ := strconv.Atoi(ruleIn[1])
			if rmin > v || v > rmax {
				fValidationErrors = append(fValidationErrors, newValidationError(rs.Name, ErrValueNotInRange))
			}
		}
	}
	return fValidationErrors, nil
}

func checkSlice(rs reflect.StructField, a interface{}) (fValidationErrors ValidationErrors, err error) {
	stringSlice, ok := a.([]string)
	if ok {
		for _, v := range stringSlice {
			vErr, err := checkString(rs, v)
			if err != nil {
				return nil, err
			}
			if vErr != nil {
				return vErr, err
			}
		}
	}

	intSlice, ok := a.([]int)
	if ok {
		for _, v := range intSlice {
			vErr, err := checkInt(rs, v)
			if err != nil {
				return nil, err
			}
			if vErr != nil {
				return vErr, err
			}
		}
	}
	return fValidationErrors, nil
}

func Validate(v interface{}) error {
	var sValidationErrors ValidationErrors
	getIn := reflect.ValueOf(v)

	if getIn.Kind().String() != "struct" {
		return ErrIsNotStruct
	}

	for i := 0; i < getIn.NumField(); i++ {
		switch getIn.Field(i).Kind().String() {
		case "int":
			vErr, err := checkInt(getIn.Type().Field(i), int(getIn.Field(i).Int()))
			if err != nil {
				return err
			}
			sValidationErrors = append(sValidationErrors, vErr...)
		case "string":
			vErr, err := checkString(getIn.Type().Field(i), getIn.Field(i).String())
			if err != nil {
				return err
			}
			sValidationErrors = append(sValidationErrors, vErr...)
		case "slice":
			if getIn.Field(i).CanInterface() {
				vErr, err := checkSlice(getIn.Type().Field(i), getIn.Field(i).Interface())
				if err != nil {
					return err
				}
				sValidationErrors = append(sValidationErrors, vErr...)
			}
		default:
		}
	}

	if len(sValidationErrors) > 0 {
		return sValidationErrors
	}
	return nil
}
