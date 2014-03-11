package jsonutils

import (
  "strings"
  "strconv"
  "fmt"
  "reflect"
  "github.com/spkane/go-utils/strutils"
)

func JsonJoin(values []string) string {
  s := []string{values[0], values[1], " : ", values[2], ",\n"}
  return strings.Join(s, "")
}


func JsonBuild(input interface{}, debug bool) string {
  m := input.(map[string]interface{})

  json := "{\n"

  for k, v := range m {
    switch vv := v.(type) {
    case string:
      if debug == true {fmt.Println("k:", k, "v:", v, "vv:", vv, "type:", reflect.TypeOf(v) )}
      json += JsonJoin([]string{"  ", strutils.QuoteString(k), strutils.QuoteString(vv)})
    case int64:
      if debug == true {fmt.Println("k:", k, "v:", v, "vv:", vv, "type:", reflect.TypeOf(v) )}
      json += JsonJoin([]string{"  ", strutils.QuoteString(k), strconv.FormatInt(vv, 1)})
    case float64:
      if debug == true {fmt.Println("k:", k, "v:", v, "vv:", vv, "type:", reflect.TypeOf(v) )}
      json += JsonJoin([]string{"  ", strutils.QuoteString(k), strconv.FormatFloat(vv, 'f', -1, 64)})
    case bool:
      if debug == true {fmt.Println("k:", k, "v:", v, "vv:", vv, "type:", reflect.TypeOf(v) )}
      json += JsonJoin([]string{"  ", strutils.QuoteString(k), strconv.FormatBool(vv)})
    default:
      if debug == true {fmt.Println("k:", k, "v:", v, "vv:", vv, "type:", reflect.TypeOf(v) )}
      json += JsonJoin([]string{"  ", strutils.QuoteString(k), strutils.QuoteString("***UNKNOWN***")})
   }
  }

  json = strutils.TrimSuffix(json, ",\n")
  json += "\n}"

  return json
}
