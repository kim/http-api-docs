package docs

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type JsonFormatter struct{}

func (js *JsonFormatter) GenerateEndpointBlock(endp *Endpoint) string {
	if endp.ResponseTyp != nil {
		return "\"" + endpointName(endp) + "Response\":" + responseJsonish(endp.ResponseTyp, 0)
	} else {
		return ""
	}
}

func responseJsonish(t reflect.Type, i int) string {
	// Aux function
	insertIndent := func(i int) string {
		buf := new(bytes.Buffer)
		for j := 0; j < i; j++ {
			buf.WriteRune(' ')
		}
		return buf.String()
	}

	countExported := func(t reflect.Type) int {
		if t.Kind() != reflect.Struct {
			return 0
		}

		count := 0

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if f.Name[0:1] == strings.ToUpper(f.Name[0:1]) {
				count++
			}
		}
		return count
	}

	result := new(bytes.Buffer)
	if i > MaxIndent { // 5 levels is enough. Infinite loop failsafe
		return insertIndent(i) + "...\n"
	}

	switch t.Kind() {
	case reflect.Invalid:
		result.WriteString("null\n")
	case reflect.Interface:
		result.WriteString(insertIndent(i) + "{}\n")
	case reflect.Ptr:
		if _, ok := t.MethodByName("String"); ok && countExported(t.Elem()) == 0 {
			return responseJsonish(reflect.TypeOf(""), i)
		}
		return responseJsonish(t.Elem(), i)
	case reflect.Map:
		result.WriteString(insertIndent(i) + "{\n")
		result.WriteString(insertIndent(i+IndentLevel) + fmt.Sprintf(`"<%s>": `, t.Key().Kind()))
		result.WriteString(responseJsonish(t.Elem(), i+IndentLevel))
		result.WriteString(insertIndent(i) + "}\n")
	case reflect.Struct:
		if _, ok := t.MethodByName("String"); ok && countExported(t) == 0 {
			return responseJsonish(reflect.TypeOf(""), i)
		}
		result.WriteString(insertIndent(i) + "{\n")
		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)
			result.WriteString(fmt.Sprintf(insertIndent(i+IndentLevel)+"\"%s\": ", f.Name))
			result.WriteString(responseJsonish(f.Type, i+IndentLevel))
		}
		result.WriteString(insertIndent(i) + "}\n")
	case reflect.Slice:
		result.WriteString("[\n")
		result.WriteString(responseJsonish(t.Elem(), i+IndentLevel))
		result.WriteString(insertIndent(i) + "]\n")

	case reflect.Bool:
		result.WriteString(insertIndent(i) + "true,")

	case reflect.Int:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Uintptr:
		result.WriteString(insertIndent(i) + "42,\n")

	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		fallthrough
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		result.WriteString(insertIndent(i) + "42.69,\n")

	case reflect.String:
		result.WriteString(insertIndent(i) + "\"hello world\",\n")

	default:
		result.WriteString(insertIndent(i) + "\"<" + t.Kind().String() + ">\"\n")

	}

	return result.String()
}

// Unneeded

func (js *JsonFormatter) GenerateIntro() string {
	return ""
}

func (js *JsonFormatter) GenerateIndex(endp []*Endpoint) string {
	return ""
}

func (js *JsonFormatter) GenerateArgumentsBlock(args []*Argument, opts []*Argument) string {
	return ""
}

func (js *JsonFormatter) GenerateBodyBlock(args []*Argument) string {
	return ""
}

func (js *JsonFormatter) GenerateResponseBlock(response string) string {
	return ""
}

func (js *JsonFormatter) GenerateExampleBlock(endp *Endpoint) string {
	return ""
}
