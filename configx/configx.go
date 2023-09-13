package configx

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sivaosorg/govm/utils"
	"gopkg.in/yaml.v2"
)

func NewCommentedConfig() *CommentedConfig {
	c := &CommentedConfig{}
	return c
}

func (c *CommentedConfig) SetData(value interface{}) *CommentedConfig {
	c.Data = value
	return c
}

func (c *CommentedConfig) SetComment(value FieldCommentConfig) *CommentedConfig {
	c.Comments = value
	return c
}

func (c *CommentedConfig) Json() string {
	return utils.ToJson(c)
}

func ReadConfig[T any](path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var cfg T
	err = yaml.Unmarshal(config, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func CreateConfig[T any](path string, data *T) error {
	config, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, config, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func CreateConfigWithComments[T any](path string, data CommentedConfig) error {
	bytes, err := _marshal(data.Data, data.Comments)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, bytes, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func _marshal(data interface{}, comments FieldCommentConfig) ([]byte, error) {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	for field, comment := range comments {
		for i, line := range lines {
			if strings.Contains(line, field+":") {
				if strings.Contains(comment, "\n") {
					commentLines := strings.Split(comment, "\n")
					for j := len(commentLines) - 1; j >= 0; j-- {
						commentLine := fmt.Sprintf("# %s", commentLines[j])
						lines = insertStringAt(lines, i, commentLine)
					}
				} else {
					commentLine := fmt.Sprintf("# %s", comment)
					lines = insertStringAt(lines, i, commentLine)
				}
				break
			}
		}
	}
	c := strings.Join(lines, "\n")
	return []byte(c), nil
}

func insertStringAt(slice []string, index int, value string) []string {
	return append(slice[:index], append([]string{value}, slice[index:]...)...)
}
