package core

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/spf13/viper"
)

type Notification_Config struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
}

func notification_GetConfig() (config Notification_Config, err error) {
	// get current os directory path
	pwd, _ := os.Getwd()

	//fmt.Println(pwd)
	viper.AddConfigPath(pwd + "/config/")
	viper.SetConfigName("notifications")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = viper.Unmarshal(&config)
	//spew.Dump(config)

	return config, err
}

func Notification_Emergency(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + SystemHostname
	messageTitle = "[" + ApplicationName() + "] Notification - " + messageTitle + " - " + SystemHostname

	// NOTE Notification Message & Title
	message := Notification_New(messageTitle, messageBody, pushover.PriorityEmergency)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}

func Notification_New(title string, body string, priority int) *pushover.Message {
	port := ApplicationHTTPPort()

	return &pushover.Message{
		Message:     body,
		Title:       title,
		Priority:    priority,
		URL:         "http://" + SystemHostname + ":" + port + "/",
		URLTitle:    SystemHostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SystemHostname, ".", "_"),
		CallbackURL: "http://" + SystemHostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}
}

func Notification_Normal(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + SystemHostname
	messageTitle = "[" + ApplicationName() + "] Notification - " + messageTitle + " - " + SystemHostname

	// NOTE Notification Message & Title
	message := Notification_New(messageTitle, messageBody, pushover.PriorityNormal)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}

func Notification_URL(messageTitle string, messageBody string, url string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + SystemHostname
	messageTitle = "[" + ApplicationName() + "] Notification - " + messageTitle + " - " + SystemHostname

	// NOTE Notification Message & Title
	message := Notification_New(messageTitle, messageBody, pushover.PriorityNormal)
	message.URL = message.URL + url

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}

func Notification_High(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + SystemHostname
	messageTitle = "[" + ApplicationName() + "] Notification - " + messageTitle + " - " + SystemHostname

	// NOTE Notification Message
	message := Notification_New(messageTitle, messageBody, pushover.PriorityHigh)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}

func Notification_Low(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + SystemHostname
	messageTitle = "[" + ApplicationName() + "] Notification - " + messageTitle + " - " + SystemHostname

	// NOTE Notification Message
	message := Notification_New(messageTitle, messageBody, pushover.PriorityLow)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}
}
