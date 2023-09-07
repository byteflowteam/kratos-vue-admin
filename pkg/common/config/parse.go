package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"

	"github.com/byteflowteam/kratos-vue-admin/pkg/common/constant"
	"github.com/byteflowteam/kratos-vue-admin/pkg/utils"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project.
	Root = filepath.Join(filepath.Dir(b), "../../..")
)

const (
	FileName             = "config.yaml"
	NotificationFileName = "notification.yaml"
	ENV                  = "CONFIG_NAME"
	DefaultFolderPath    = "../config/"
	ConfKey              = "conf"
)

func GetOptionsByNotification(cfg NotificationConf) utils.Options {
	opts := utils.NewOptions()
	if cfg.UnreadCount {
		opts = utils.WithOptions(opts, utils.WithUnreadCount(true))
	}
	if cfg.OfflinePush.Enable {
		opts = utils.WithOptions(opts, utils.WithOfflinePush(true))
	}
	switch cfg.ReliabilityLevel {
	case constant.UnreliableNotification:
	case constant.ReliableNotificationNoMsg:
		opts = utils.WithOptions(opts, utils.WithHistory(true), utils.WithPersistent())
	}
	opts = utils.WithOptions(opts, utils.WithSendMsg(cfg.IsSendMsg))
	return opts
}

func (c *config) unmarshalConfig(config interface{}, configPath string) error {
	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}

func (c *config) initConfig(config interface{}, configName, configFolderPath string) error {
	if configFolderPath == "" {
		configFolderPath = DefaultFolderPath
	}
	configPath := filepath.Join(configFolderPath, configName)
	defer func() {
		fmt.Println("use config", configPath)
	}()
	_, err := os.Stat(configPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		configPath = filepath.Join(Root, "config", configName)
	} else {
		Root = filepath.Dir(configPath)
	}
	return c.unmarshalConfig(config, configPath)
}

func InitConfig(configFolderPath string) error {
	err := Config.initConfig(&Config, FileName, configFolderPath)
	if err != nil {
		return err
	}
	err = Config.initConfig(&Config.Notification, NotificationFileName, configFolderPath)
	if err != nil {
		return err
	}
	return nil
}

func EncodeConfig() []byte {
	buf := bytes.NewBuffer(nil)
	if err := yaml.NewEncoder(buf).Encode(Config); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func init() {
	InitConfig(DefaultFolderPath)
}
