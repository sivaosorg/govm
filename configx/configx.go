package configx

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sivaosorg/govm/asterisk"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/mongodb"
	"github.com/sivaosorg/govm/mysql"
	"github.com/sivaosorg/govm/postgres"
	"github.com/sivaosorg/govm/rabbitmqx"
	"github.com/sivaosorg/govm/redisx"
	"github.com/sivaosorg/govm/timex"
	"github.com/sivaosorg/govm/utils"
	"gopkg.in/yaml.v2"
)

func NewKeyCmtConfig() *CommentedConfig {
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

func NewKeysConfig() *KeysConfig {
	k := &KeysConfig{}
	return k
}

func (k *KeysConfig) SetAsterisk(value asterisk.AsteriskConfig) *KeysConfig {
	k.Asterisk = value
	return k
}

func (k *KeysConfig) SetAsteriskCursor(value *asterisk.AsteriskConfig) *KeysConfig {
	k.Asterisk = *value
	return k
}

func (k *KeysConfig) SetMongodb(value mongodb.MongodbConfig) *KeysConfig {
	k.Mongodb = value
	return k
}

func (k *KeysConfig) SetMongodbCursor(value *mongodb.MongodbConfig) *KeysConfig {
	k.Mongodb = *value
	return k
}

func (k *KeysConfig) SetMySql(value mysql.MysqlConfig) *KeysConfig {
	k.MySql = value
	return k
}

func (k *KeysConfig) SetMySqlCursor(value *mysql.MysqlConfig) *KeysConfig {
	k.MySql = *value
	return k
}

func (k *KeysConfig) SetPostgres(value postgres.PostgresConfig) *KeysConfig {
	k.Postgres = value
	return k
}

func (k *KeysConfig) SetPostgresCursor(value *postgres.PostgresConfig) *KeysConfig {
	k.Postgres = *value
	return k
}

func (k *KeysConfig) SetRabbitMq(value rabbitmqx.RabbitMqConfig) *KeysConfig {
	k.RabbitMq = value
	return k
}

func (k *KeysConfig) SetRabbitMqCursor(value *rabbitmqx.RabbitMqConfig) *KeysConfig {
	k.RabbitMq = *value
	return k
}

func (k *KeysConfig) SetRedis(value redisx.RedisConfig) *KeysConfig {
	k.Redis = value
	return k
}

func (k *KeysConfig) SetRedisCursor(value *redisx.RedisConfig) *KeysConfig {
	k.Redis = *value
	return k
}

func (k *KeysConfig) Json() string {
	return utils.ToJson(k)
}

func GetKeysDefaultConfig() *KeysConfig {
	k := NewKeysConfig()
	k.SetAsterisk(*asterisk.GetAsteriskConfigSample().SetEnabled(false))
	k.SetMongodb(*mongodb.GetMongodbConfigSample().SetEnabled(false))
	k.SetMySql(*mysql.GetMysqlConfigSample().SetEnabled(false))
	k.SetPostgres(*postgres.GetPostgresConfigSample().SetEnabled(false))
	k.SetRabbitMq(*rabbitmqx.GetRabbitMqConfigSample().SetEnabled(false))
	k.SetRedis(*redisx.GetRedisConfigSample().SetEnabled(false))
	return k
}

func (KeysConfig) WriteDefaultConfig() {
	_, err := os.OpenFile(FilenameDefaultConf, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while creating new filename: %s", err, FilenameDefaultConf)
		return
	}
	m := NewKeyCmtConfig()
	m.SetData(GetKeysDefaultConfig())
	m.SetComment(map[string]string{
		"asterisk": fmt.Sprintf("################################\n%s\n%s\n################################", "Asterisk Server Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mongodb":  fmt.Sprintf("################################\n%s\n%s\n################################", "Mongodb Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mysql":    fmt.Sprintf("################################\n%s\n%s\n################################", "MySQL Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"postgres": fmt.Sprintf("################################\n%s\n%s\n################################", "Postgres Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"rabbitmq": fmt.Sprintf("################################\n%s\n%s\n################################", "RabbitMQ Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"redis":    fmt.Sprintf("################################\n%s\n%s\n################################", "Redis Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
	})
	err = CreateConfigWithComments[KeysConfig](filepath.Join(".", FilenameDefaultConf), *m)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while writing keys default configs", err)
	}
	logger.Infof("View file keys default config: %s", FilenameDefaultConf)
}

func (KeysConfig) ReadDefaultConfig() {
	keys, err := ReadConfig[KeysConfig](filepath.Join(".", FilenameDefaultConf))
	if err != nil {
		logger.Errorf("ReadDefaultConfig(), an error occurred while reading keys default configs: %s", err, FilenameDefaultConf)
		return
	}
	logger.Infof("%+v", keys)
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
