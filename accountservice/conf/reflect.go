/**
 * Create by Le Trong on 02/Mar/2019
 */

package conf

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

const tagPrefix = "viper"

func populateConfig(config *Config) (*Config, error) {
	err := recursivelySet(reflect.ValueOf(config), "")
	if err != nil {
		return nil, err
	}

	return config, nil
}

func recursivelySet(val reflect.Value, prefix string) error {
	if val.Kind() != reflect.Ptr {
		return errors.New("Unknown reflect kind")
	}

	// Dereference
	val = reflect.Indirect(val)
	if val.Kind() != reflect.Struct {
		return errors.New("Unknown reflect kind")
	}

	// Grab the type for this instance
	vType := reflect.TypeOf(val.Interface())

	// Go through child fields
	for i := 0; i < val.NumField(); i++ {
		thisField := val.Field(i)
		thisType := vType.Field(i)
		tag := prefix + getTag(thisType)

		switch thisField.Kind() {
		case reflect.Struct:
			if err := recursivelySet(thisField.Addr(), tag+"."); err != nil {
				return err
			}
		case reflect.Int:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			// You can only set with an int64 -> int
			configVal := int64(viper.GetInt(tag))
			thisField.SetInt(configVal)
		case reflect.String:
			thisField.SetString(viper.GetString(tag))
		case reflect.Bool:
			thisField.SetBool(viper.GetBool(tag))
		default:
			return fmt.Errorf("unexpected type detected ~ aborting: %s", thisField.Kind())
		}
	}

	return nil
}

func getTag(field reflect.StructField) string {
	// Check if maybe we have a special magic tag
	tag := field.Tag
	if tag != "" {
		for _, prefix := range []string{tagPrefix, "mapstructure", "json"} {
			if v := tag.Get(prefix); v != "" {
				return v
			}
		}
	}

	return field.Name
}
