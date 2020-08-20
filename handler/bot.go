package handler

import (
	"encoding/gob"
	"github.com/Rhymen/go-whatsapp"
	"github.com/op/go-logging"
	"github.com/skip2/go-qrcode"
	"os"
	"time"
)

var LOGGER = logging.MustGetLogger("whats-app-bot")

type Bot struct {
	Connection *whatsapp.Conn
}

func (op *Bot) Connect() error {
	LOGGER.Infof("\tConnecting to WhatsApp...")
	conn, err := whatsapp.NewConn(90 * time.Second)
	if err != nil {
		return err
	}

	op.Connection = conn
	LOGGER.Infof("\tSuccessfully connect to WhatsApp!")
	return nil
}

func (op *Bot) Login() error {
	LOGGER.Infof("\tLogin")
	sess, err := readSessionFromFileSystem()
	if err != nil {
		LOGGER.Warningf("No session found from file system: %s", err.Error())
		sess, err := op.createSession()
		if err != nil {
			LOGGER.Errorf("Can't create new session: %s", err.Error())
			return err
		}

		err = writeSessionToFileSystem(sess)
		if err != nil {
			LOGGER.Errorf("Can't write session into file system: %s", err.Error())
			return err
		}

		return nil
	}

	sess, err = op.Connection.RestoreWithSession(sess)
	if err != nil {
		LOGGER.Warningf("Can't restore previous session: %s", err.Error())
		_, err := op.createSession()
		if err != nil {
			LOGGER.Errorf("Can't create new session: %s", err.Error())
			return err
		}

		err = writeSessionToFileSystem(sess)
		if err != nil {
			LOGGER.Errorf("Can't write session into file system: %s", err.Error())
			return err
		}

		return nil
	}

	err = writeSessionToFileSystem(sess)
	if err != nil {
		LOGGER.Errorf("Can't write session into file system: %s", err.Error())
		return err
	}

	return nil
}

func (op *Bot) createSession() (whatsapp.Session, error) {
	//op.Connection.AddHandler(WhatsAppMessageHandler{})

	qrChan := make(chan string)
	go func() {
		err := qrcode.WriteFile(<-qrChan, qrcode.Medium, 256, "login.png")
		if err != nil {
			LOGGER.Error("Failed to encode the qr code string: %s", err.Error())
			return
		}
	}()

	return op.Connection.Login(qrChan)
}

func (op *Bot) SendMessage(msg, receiver string) error {
	LOGGER.Infof("\tSending message ...")
	text := whatsapp.TextMessage{
		Info: whatsapp.MessageInfo{
			RemoteJid: receiver,
		},
		Text: msg,
	}

	_, err := op.Connection.Send(text)
	if err != nil {
		return err
	}

	LOGGER.Infof("\tSuccessfully sent message!")
	return nil
}

func readSessionFromFileSystem() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open("./session/waSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSessionToFileSystem(session whatsapp.Session) error {
	file, err := os.Create( "./session/waSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}
