// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strutils

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"

	"github.com/hexya-erp/hexya/hexya/tools/logging"
)

var log *logging.Logger

func init() {
	log = logging.GetLogger("strutils")
}

// SnakeCase convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func SnakeCase(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}

// Title convert the given camelCase string to a title string.
// eg. MyHTMLData => My HTML Data
func Title(in string) string {

	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, ' ')
		}
		out = append(out, runes[i])
	}

	return string(out)
}

// GetDefaultString returns str if it is not an empty string or def otherwise
func GetDefaultString(str, def string) string {
	if str == "" {
		return def
	}
	return str
}

// StartsAndEndsWith returns true if the given string starts with prefix
// and ends with suffix.
func StartsAndEndsWith(str, prefix, suffix string) bool {
	return strings.HasPrefix(str, prefix) && strings.HasSuffix(str, suffix)
}

// MarshalToJSONString marshals the given data to its JSON representation and
// returns it as a string. It panics in case of error.
func MarshalToJSONString(data interface{}) string {
	if _, ok := data.(string); !ok {
		domBytes, err := json.Marshal(data)
		if err != nil {
			log.Panic("Unable to marshal given data", "error", err, "data", data)
		}
		return string(domBytes)
	}
	return data.(string)
}

// HumanSize returns the given size (in bytes) in a human readable format
func HumanSize(size int64) string {
	units := []string{"bytes", "KB", "MB", "GB"}
	s, i := float64(size), 0
	for s >= 1024 && i < len(units)-1 {
		s /= 1024
		i++
	}
	return fmt.Sprintf("%.2f %s", s, units[i])
}

// Substitute substitutes each occurrence of each key of mapping in str by the
// corresponding mapping value and returns the substituted string.
func Substitute(str string, mapping map[string]string) string {
	for key, val := range mapping {
		str = strings.Replace(str, key, val, -1)
	}
	return str
}
