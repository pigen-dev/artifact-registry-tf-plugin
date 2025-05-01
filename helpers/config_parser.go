package helpers

import (
	"gopkg.in/yaml.v2"
)


func YamlConfigParser(in map[string] interface{}, output interface{}) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(out, output)
	if err != nil {
		return err
	}
	return nil
}