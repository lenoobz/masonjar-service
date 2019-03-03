/**
 * Create by Le Trong on 02/Mar/2019
 */

package utilities

import (
	"encoding/json"
)

// PrettyPrint print a struct with pretty indentation
func PrettyPrint(v interface{}) (s string, err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		return string(b), nil
	}
	return "", err
}
