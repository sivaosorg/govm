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
	"github.com/sivaosorg/govm/cookies"
	"github.com/sivaosorg/govm/corsx"
	"github.com/sivaosorg/govm/logger"
	"github.com/sivaosorg/govm/mongodb"
	"github.com/sivaosorg/govm/mysql"
	"github.com/sivaosorg/govm/postgres"
	"github.com/sivaosorg/govm/rabbitmqx"
	"github.com/sivaosorg/govm/redisx"
	"github.com/sivaosorg/govm/server"
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
	k.SetServer(*server.GetServerSample())
	k.SetAsterisk(*asterisk.GetAsteriskConfigSample().SetEnabled(false))
	k.SetMongodb(*mongodb.GetMongodbConfigSample().SetEnabled(false))
	k.SetMySql(*mysql.GetMysqlConfigSample().SetEnabled(false))
	k.SetPostgres(*postgres.GetPostgresConfigSample().SetEnabled(false))
	k.SetRabbitMq(*rabbitmqx.GetRabbitMqConfigSample().SetEnabled(false))
	k.SetRedis(*redisx.GetRedisConfigSample().SetEnabled(false))
	k.SetTelegram(*telegram.GetTelegramConfigSample().SetEnabled(false))
	k.SetSlack(*slack.GetSlackConfigSample().SetEnabled(false))
	k.SetCors(*corsx.GetCorsConfigSample().SetEnabled(false))
	k.SetCookie(*cookies.GetCookieConfigSample().SetEnabled(false))
	k.AppendTelegramSeekers(*telegram.GetMultiTenantTelegramConfigSample())
	k.AppendSlackSeekers(*slack.GetMultiTenantSlackConfigSample())
	k.AppendAsteriskSeekers(*asterisk.GetMultiTenantAsteriskConfigSample())
	k.AppendMongodbSeekers(*mongodb.GetMultiTenantMongodbConfigSample())
	k.AppendMySqlSeekers(*mysql.GetMultiTenantMysqlConfigSample())
	k.AppendPostgresSeekers(*postgres.GetMultiTenantPostgresConfigSample())
	k.AppendRabbitMqSeekers(*rabbitmqx.GetMultiTenantRabbitMqConfigSample())
	k.AppendRedisSeekers(*redisx.GetMultiTenantRedisConfigSample())
	k.AppendCookieSeekers(*cookies.GetMultiTenantCookieConfigSample())
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
		"server":           fmt.Sprintf("################################\n%s\n%s\n################################", "Server Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"cookie":           fmt.Sprintf("################################\n%s\n%s\n################################", "Cookie Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"asterisk":         fmt.Sprintf("################################\n%s\n%s\n################################", "Asterisk Server Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mongodb":          fmt.Sprintf("################################\n%s\n%s\n################################", "Mongodb Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mysql":            fmt.Sprintf("################################\n%s\n%s\n################################", "MySQL Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"postgres":         fmt.Sprintf("################################\n%s\n%s\n################################", "Postgres Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"rabbitmq":         fmt.Sprintf("################################\n%s\n%s\n################################", "RabbitMQ Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"redis":            fmt.Sprintf("################################\n%s\n%s\n################################", "Redis Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"telegram":         fmt.Sprintf("################################\n%s\n%s\n################################", "Telegram Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"slack":            fmt.Sprintf("################################\n%s\n%s\n################################", "Slack Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"cors":             fmt.Sprintf("################################\n%s\n%s\n################################", "Cors Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"telegram-seekers": fmt.Sprintf("################################\n%s\n%s\n################################", "Telegram Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"slack-seekers":    fmt.Sprintf("################################\n%s\n%s\n################################", "Slack Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"asterisk-seekers": fmt.Sprintf("################################\n%s\n%s\n################################", "Asterisk Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mongodb-seekers":  fmt.Sprintf("################################\n%s\n%s\n################################", "Mongodb Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"mysql-seekers":    fmt.Sprintf("################################\n%s\n%s\n################################", "MySQL Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"postgres-seekers": fmt.Sprintf("################################\n%s\n%s\n################################", "Postgres Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"rabbitmq-seekers": fmt.Sprintf("################################\n%s\n%s\n################################", "RabbitMQ Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"redis-seekers":    fmt.Sprintf("################################\n%s\n%s\n################################", "Redis Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
		"cookie-seekers":   fmt.Sprintf("################################\n%s\n%s\n################################", "Cookie Seekers Config", timex.With(time.Now()).Format(timex.DateTimeFormYearMonthDayHourMinuteSecond)),
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

func (ClusterMultiTenancyKeysConfig) ReadCurrentConfig() (ClusterMultiTenancyKeysConfig, error) {
	keys, err := ReadConfig[ClusterMultiTenancyKeysConfig](filepath.Join(".", FilenameDefaultClusterMultiTenantConf))
	return *keys, err
}

func (ClusterMultiTenancyKeysConfig) ReadCurrentConfigWith(filename string) (ClusterMultiTenancyKeysConfig, error) {
	keys, err := ReadConfig[ClusterMultiTenancyKeysConfig](filepath.Join(".", filename))
	return *keys, err
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

func (c *ClusterMultiTenancyKeysConfig) AllowedUsableDefault() bool {
	if len(c.Clusters) == 0 {
		return true
	}
	counter := 0
	for _, v := range c.Clusters {
		if v.IsUsableDefault {
			counter++
		}
	}
	return counter <= 1
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

func (k *KeysConfig) SetCors(value corsx.CorsConfig) *KeysConfig {
	k.Cors = value
	return k
}

func (k *KeysConfig) SetCorsCursor(value *corsx.CorsConfig) *KeysConfig {
	k.Cors = *value
	return k
}

func (k *KeysConfig) SetServer(value server.Server) *KeysConfig {
	k.Server = value
	return k
}

func (k *KeysConfig) SetServerCursor(value *server.Server) *KeysConfig {
	k.Server = *value
	return k
}

func (k *KeysConfig) SetCookie(value cookies.CookieConfig) *KeysConfig {
	k.Cookie = value
	return k
}

func (k *KeysConfig) SetCookieCursor(value *cookies.CookieConfig) *KeysConfig {
	k.Cookie = *value
	return k
}

func (k *KeysConfig) SetTelegramSeekers(values []telegram.MultiTenantTelegramConfig) *KeysConfig {
	k.TelegramSeekers = values
	return k
}

func (k *KeysConfig) AppendTelegramSeekers(values ...telegram.MultiTenantTelegramConfig) *KeysConfig {
	k.TelegramSeekers = append(k.TelegramSeekers, values...)
	return k
}

func (k *KeysConfig) SetSlackSeekers(values []slack.MultiTenantSlackConfig) *KeysConfig {
	k.SlackSeekers = values
	return k
}

func (k *KeysConfig) AppendSlackSeekers(values ...slack.MultiTenantSlackConfig) *KeysConfig {
	k.SlackSeekers = append(k.SlackSeekers, values...)
	return k
}

func (k *KeysConfig) SetAsteriskSeekers(values []asterisk.MultiTenantAsteriskConfig) *KeysConfig {
	k.AsteriskSeekers = values
	return k
}

func (k *KeysConfig) AppendAsteriskSeekers(values ...asterisk.MultiTenantAsteriskConfig) *KeysConfig {
	k.AsteriskSeekers = append(k.AsteriskSeekers, values...)
	return k
}

func (k *KeysConfig) SetMongodbSeekers(values []mongodb.MultiTenantMongodbConfig) *KeysConfig {
	k.MongodbSeekers = values
	return k
}

func (k *KeysConfig) AppendMongodbSeekers(values ...mongodb.MultiTenantMongodbConfig) *KeysConfig {
	k.MongodbSeekers = append(k.MongodbSeekers, values...)
	return k
}

func (k *KeysConfig) SetMySqlSeekers(values []mysql.MultiTenantMysqlConfig) *KeysConfig {
	k.MySqlSeekers = values
	return k
}

func (k *KeysConfig) AppendMySqlSeekers(values ...mysql.MultiTenantMysqlConfig) *KeysConfig {
	k.MySqlSeekers = append(k.MySqlSeekers, values...)
	return k
}

func (k *KeysConfig) SetPostgresSeekers(values []postgres.MultiTenantPostgresConfig) *KeysConfig {
	k.PostgresSeekers = values
	return k
}

func (k *KeysConfig) AppendPostgresSeekers(values ...postgres.MultiTenantPostgresConfig) *KeysConfig {
	k.PostgresSeekers = values
	return k
}

func (k *KeysConfig) SetRabbitMqSeekers(values []rabbitmqx.MultiTenantRabbitMqConfig) *KeysConfig {
	k.RabbitMqSeekers = values
	return k
}

func (k *KeysConfig) AppendRabbitMqSeekers(values ...rabbitmqx.MultiTenantRabbitMqConfig) *KeysConfig {
	k.RabbitMqSeekers = append(k.RabbitMqSeekers, values...)
	return k
}

func (k *KeysConfig) SetRedisSeekers(values []redisx.MultiTenantRedisConfig) *KeysConfig {
	k.RedisSeekers = values
	return k
}

func (k *KeysConfig) AppendRedisSeekers(values ...redisx.MultiTenantRedisConfig) *KeysConfig {
	k.RedisSeekers = append(k.RedisSeekers, values...)
	return k
}

func (k *KeysConfig) SetCookieSeekers(values []cookies.MultiTenantCookieConfig) *KeysConfig {
	k.CookieSeekers = values
	return k
}

func (k *KeysConfig) AppendCookieSeekers(values ...cookies.MultiTenantCookieConfig) *KeysConfig {
	k.CookieSeekers = append(k.CookieSeekers, values...)
	return k
}

func (k *KeysConfig) AvailableTelegramSeekers() bool {
	return len(k.TelegramSeekers) > 0
}

func (k *KeysConfig) AvailableSlackSeekers() bool {
	return len(k.SlackSeekers) > 0
}

func (k *KeysConfig) AvailableAsteriskSeekers() bool {
	return len(k.AsteriskSeekers) > 0
}

func (k *KeysConfig) AvailableMongodbSeekers() bool {
	return len(k.MongodbSeekers) > 0
}

func (k *KeysConfig) AvailableMySqlSeekers() bool {
	return len(k.MySqlSeekers) > 0
}

func (k *KeysConfig) AvailablePostgresSeekers() bool {
	return len(k.PostgresSeekers) > 0
}

func (k *KeysConfig) AvailableRabbitMqSeekers() bool {
	return len(k.RabbitMqSeekers) > 0
}

func (k *KeysConfig) AvailableRedisSeekers() bool {
	return len(k.RedisSeekers) > 0
}

func (k *KeysConfig) AvailableCookieSeekers() bool {
	return len(k.CookieSeekers) > 0
}

func (k *KeysConfig) FindTelegramSeeker(key string) (telegram.MultiTenantTelegramConfig, error) {
	return telegram.NewClusterMultiTenantTelegramConfig().SetClusters(k.TelegramSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindSlackSeeker(key string) (slack.MultiTenantSlackConfig, error) {
	return slack.NewClusterMultiTenantSlackConfig().SetClusters(k.SlackSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindAsteriskSeeker(key string) (asterisk.MultiTenantAsteriskConfig, error) {
	return asterisk.NewClusterMultiTenantAsteriskConfig().SetClusters(k.AsteriskSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindMongodbSeeker(key string) (mongodb.MultiTenantMongodbConfig, error) {
	return mongodb.NewClusterMultiTenantMongodbConfig().SetClusters(k.MongodbSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindMySqlSeeker(key string) (mysql.MultiTenantMysqlConfig, error) {
	return mysql.NewClusterMultiTenantMysqlConfig().SetClusters(k.MySqlSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindPostgresSeeker(key string) (postgres.MultiTenantPostgresConfig, error) {
	return postgres.NewClusterMultiTenantPostgresConfig().SetClusters(k.PostgresSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindRabbitMqSeeker(key string) (rabbitmqx.MultiTenantRabbitMqConfig, error) {
	return rabbitmqx.NewClusterMultiTenantRabbitMqConfig().SetClusters(k.RabbitMqSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindRedisSeeker(key string) (redisx.MultiTenantRedisConfig, error) {
	return redisx.NewClusterMultiTenantRedisConfig().SetClusters(k.RedisSeekers).FindClusterBy(key)
}

func (k *KeysConfig) FindCookieSeeker(key string) (cookies.MultiTenantCookieConfig, error) {
	return cookies.NewClusterMultiTenantCookieConfig().SetClusters(k.CookieSeekers).FindClusterBy(key)
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
