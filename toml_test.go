package _toml

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type config struct {
	Version  string      `toml:"version"`
	Database []*Database `toml:"database"`
}

type Database struct {
	Host     string   `toml:"host"`
	Password string   `toml:"password"`
	Tables   []string `toml:"tables"`
	Task     *Task    `toml:"task"`
}

type Task struct {
	Name string `toml:"name"`
}

func TestMarshal(t *testing.T) {
	
	var a = config{
		Version: "1.0.0",
		Database: []*Database{
			{
				Host:     "127.0.0.1",
				Password: "123456",
				Tables:   []string{"user", "password"},
				Task: &Task{
					Name: "a",
				},
			},
			{
				Host:     "127.0.0.2",
				Password: "123456",
				Tables:   []string{"user", "password"},
				Task: &Task{
					Name: "b",
				},
			},
		},
	}
	s, err := Marshal(a)
	require.NoError(t, err)
	fmt.Println(string(s))
}
