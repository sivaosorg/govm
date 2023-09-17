package configx

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sivaosorg/govm/asterisk"
	"github.com/sivaosorg/govm/bot/slack"
	"github.com/sivaosorg/govm/bot/telegram"
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

func NewMultiTenantKeysConfig() *MultiTenancyKeysConfig {
	m := &MultiTenancyKeysConfig{}
	m.SetUsableDefault(false)
	return m
}

func NewClusterMultiTenancyKeysConfig() *ClusterMultiTenancyKeysConfig {
	c := &ClusterMultiTenancyKeysConfig{}
	return c
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

func (m *MultiTenancyKeysConfig) SetKey(value string) *MultiTenancyKeysConfig {
	if utils.IsEmpty(value) {
		log.Panicf("Invalid key: %+v", value)
	}
	m.Key = value
	return m
}

func (m *MultiTenancyKeysConfig) SetConfig(value KeysConfig) *MultiTenancyKeysConfig {
	m.Config = value
	return m
}

func (m *MultiTenancyKeysConfig) SetConfigCursor(value *KeysConfig) *MultiTenancyKeysConfig {
	m.Config = *value
	return m
}

func (m *MultiTenancyKeysConfig) SetUsableDefault(value bool) *MultiTenancyKeysConfig {
	m.IsUsableDefault = value
	return m
}

func (m *MultiTenancyKeysConfig) Json() string {
	return utils.ToJson(m)
}

func (c *ClusterMultiTenancyKeysConfig) SetClusters(values []MultiTenancyKeysConfig) *ClusterMultiTenancyKeysConfig {
	c.Clusters = values
	return c
}

func (c *ClusterMultiTenancyKeysConfig) AppendClusters(values ...MultiTenancyKeysConfig) *ClusterMultiTenancyKeysConfig {
	c.Clusters = append(c.Clusters, values...)
	return c
}

func (c *ClusterMultiTenancyKeysConfig) Json() string {
	return utils.ToJson(c)
}

func GetKeysDefaultConfig() *KeysConfig {
	k := NewKeysConfig()
	k.SetAsterisk(*asterisk.GetAsteriskConfigSample().SetEnabled(false))
	k.SetMongodb(*mongodb.GetMongodbConfigSample().SetEnabled(false))
	k.SetMySql(*mysql.GetMysqlConfigSample().SetEnabled(false))
	k.SetPostgres(*postgres.GetPostgresConfigSample().SetEnabled(false))
	k.SetRabbitMq(*rabbitmqx.GetRabbitMqConfigSample().SetEnabled(false))
	k.SetRedis(*redisx.GetRedisConfigSample().SetEnabled(false))
	k.SetTelegram(*telegram.GetTelegramConfigSample().SetEnabled(false))
	k.SetSlack(*slack.GetSlackConfigSample().SetEnabled(false))
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
		"telegram": fmt.Sprintf("################################\n%s\n%s\n################################", "Telegram Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"slack":    fmt.Sprintf("################################\n%s\n%s\n################################", "Slack Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
	})
	err = CreateConfigWithComments[KeysConfig](filepath.Join(".", FilenameDefaultConf), *m)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while writing keys default configs", err)
	}
	logger.Infof("View file keys default config: %s", FilenameDefaultConf)
}

func (MultiTenancyKeysConfig) WriteDefaultConfig() {
	_, err := os.OpenFile(FilenameDefaultMultiTenantConf, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while creating new filename: %s", err, FilenameDefaultMultiTenantConf)
		return
	}
	m := NewKeyCmtConfig()
	mt := NewMultiTenantKeysConfig()
	mt.SetKey("tenant_1")
	mt.SetConfig(*GetKeysDefaultConfig())
	m.SetData(mt)
	err = CreateConfigWithComments[MultiTenancyKeysConfig](filepath.Join(".", FilenameDefaultMultiTenantConf), *m)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while writing keys default multi-tenant configs", err)
	}
	logger.Infof("View file keys default multi-tenant config: %s", FilenameDefaultMultiTenantConf)
}

func (ClusterMultiTenancyKeysConfig) WriteDefaultConfig() {
	_, err := os.OpenFile(FilenameDefaultClusterMultiTenantConf, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while creating new filename: %s", err, FilenameDefaultClusterMultiTenantConf)
		return
	}
	m := NewKeyCmtConfig()
	c := NewClusterMultiTenancyKeysConfig()
	c.AppendClusters(
		*NewMultiTenantKeysConfig().
			SetKey("tenant_1").
			SetConfig(*GetKeysDefaultConfig()),
		*NewMultiTenantKeysConfig().
			SetKey("tenant_2").
			SetConfig(*GetKeysDefaultConfig()),
		*NewMultiTenantKeysConfig().
			SetKey("tenant_3").
			SetConfig(*GetKeysDefaultConfig()),
		*NewMultiTenantKeysConfig().
			SetKey("tenant_4").
			SetConfig(*GetKeysDefaultConfig()))
	m.SetData(c)
	err = CreateConfigWithComments[ClusterMultiTenancyKeysConfig](filepath.Join(".", FilenameDefaultClusterMultiTenantConf), *m)
	if err != nil {
		logger.Errorf("WriteDefaultConfig(), an error occurred while writing keys default cluster multi-tenant configs", err)
	}
	logger.Infof("View file keys default cluster multi-tenant config: %s", FilenameDefaultClusterMultiTenantConf)
}

func (KeysConfig) ReadDefaultConfig() {
	keys, err := ReadConfig[KeysConfig](filepath.Join(".", FilenameDefaultConf))
	if err != nil {
		logger.Errorf("ReadDefaultConfig(), an error occurred while reading keys default configs: %s", err, FilenameDefaultConf)
		return
	}
	logger.Infof("%+v", keys)
}

func (MultiTenancyKeysConfig) ReadDefaultConfig() {
	keys, err := ReadConfig[MultiTenancyKeysConfig](filepath.Join(".", FilenameDefaultMultiTenantConf))
	if err != nil {
		logger.Errorf("ReadDefaultConfig(), an error occurred while reading keys default multi-tenant configs: %s", err, FilenameDefaultMultiTenantConf)
		return
	}
	logger.Infof("%+v", keys)
}

func (ClusterMultiTenancyKeysConfig) ReadDefaultConfig() {
	keys, err := ReadConfig[ClusterMultiTenancyKeysConfig](filepath.Join(".", FilenameDefaultClusterMultiTenantConf))
	if err != nil {
		logger.Errorf("ReadDefaultConfig(), an error occurred while reading keys default cluster multi-tenant configs: %s", err, FilenameDefaultClusterMultiTenantConf)
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

func (c *ClusterMultiTenancyKeysConfig) FindClusterBy(key string) (MultiTenancyKeysConfig, error) {
	if len(c.Clusters) == 0 {
		return *NewMultiTenantKeysConfig(), fmt.Errorf("No multi-tenant cluster")
	}
	if utils.IsEmpty(key) {
		return *NewMultiTenantKeysConfig(), fmt.Errorf("Invalid key")
	}
	if len(c.Clusters) == 1 {
		return c.Clusters[0], nil
	}
	for _, v := range c.Clusters {
		if v.Key == key {
			return v, nil
		}
	}
	return *NewMultiTenantKeysConfig(), fmt.Errorf("The multi-tenant cluster not found")
}

func (k *KeysConfig) SetTelegram(value telegram.TelegramConfig) *KeysConfig {
	k.Telegram = value
	return k
}

func (k *KeysConfig) SetTelegramCursor(value *telegram.TelegramConfig) *KeysConfig {
	k.Telegram = *value
	return k
}

func (k *KeysConfig) SetSlack(value slack.SlackConfig) *KeysConfig {
	k.Slack = value
	return k
}

func (k *KeysConfig) SetSlackCursor(value *slack.SlackConfig) *KeysConfig {
	k.Slack = *value
	return k
}
