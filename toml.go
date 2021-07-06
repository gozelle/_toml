package _toml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

func Unmarshal(p []byte, v interface{}) error {
	return toml.Unmarshal(p, v)
}

func UnmarshalText(p string, v interface{}) error {
	return toml.Unmarshal([]byte(p), v)
}

func UnmarshalFile(path string, v interface{}) error {
	a, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	return Unmarshal(a, v)
}

func Marshal(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func MarshalIdent(v interface{}, ident string) ([]byte, error) {
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	e.Indent = ident
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func MarshalText(v interface{}) (string, error) {
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	err := e.Encode(v)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func MarshalTextIdent(v interface{}, ident string) (string, error) {
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	e.Indent = ident
	err := e.Encode(v)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
