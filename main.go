package main

import (
	"github.com/op/go-logging"
	"github.ibm.com/cyhuang/whatsappbot/constant"
	"github.ibm.com/cyhuang/whatsappbot/handler"
	"github.ibm.com/cyhuang/whatsappbot/logger"
	"log"
	"os"
	"strings"
	"time"
)

var LOGGER = logging.MustGetLogger("whats-app-bot")

func main() {
	f, err := logger.SetupLogging("./logger/log.txt", LOGGER)
	if err != nil {
		log.Fatal(LOGGER, err, "Unable to set up logging")
	}
	defer f.Close()

	bot := handler.Bot{}

	err = bot.Connect()
	if err != nil {
		LOGGER.Errorf("Create WhatsApp connection failed: %s", err.Error())
		os.Exit(0)
	}

	err = bot.Login()
	if err != nil {
		LOGGER.Errorf("Failed to login: %s", err.Error())
		os.Exit(0)
	}
	LOGGER.Infof("\tSuccessfully login!")

	for {
		dt := time.Now()
		theTime := dt.Format("01-02 15:04")
		hour := strings.Split(strings.Split(theTime, " ")[1], ":")[0]
		LOGGER.Debugf("Today is %s, It's %s o'clock now", strings.Split(theTime, " ")[0], hour)
		if hour == "10" {
			date := strings.Split(strings.Split(theTime, " ")[0], "-")[0]

			switch date {
			case "1":
				LOGGER.Infof("*** Time to pay the fibre fee, send message to the group ***")
				err = bot.SendMessage(constant.PayInternetFeeMsg, "98103612@s.whatsapp.net")
				if err != nil {
					LOGGER.Errorf("Failed to send message: %s", err.Error())
					os.Exit(0)
				}
			case "14":
				LOGGER.Infof("*** Time to pay the rental fee, send message to the group ***")
				err = bot.SendMessage(constant.PayRentalMsg, "98103612@s.whatsapp.net")
				if err != nil {
					LOGGER.Errorf("Failed to send message: %s", err.Error())
					os.Exit(0)
				}
			default:
				LOGGER.Infof("Normal date, nothing to notify")
			}
		}
		time.Sleep(1 * time.Hour)
	}
}
