package strutils

import (
  "strings"
)

func TrimSuffix(s, suffix string) string {
  if strings.HasSuffix(s, suffix) {
    s = s[:len(s)-len(suffix)]
  }
  return s
}

func QuoteString(str string) string {
  s := []string{"\"", str, "\""}
  return strings.Join(s, "")
}
