package mail

import (
	"errors"
	"net"
	"net/smtp"
	"strconv"
	"strings"
)

type Account struct {
	Name     string //full name of account user
	Email    string
	Password string
	Host     string //ie smtp.gmail.com
	Port     int    //ie 587
}

// NewGmail returns an account struct with a default server and port for Gmail
func NewGmail(name, email, password string) (a Account) {
	return Account{name, email, password,
		"smtp.gmail.com",
		587,
	}
}

func (a Account) Send(subject, body string, to ...string) (err error) {
	if len(to) == 0 {
		return errors.New("recipients cannot be blank")
	}

	msg := "From: " + a.Name + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html\r\n\r\n" + body

	err = smtp.SendMail(net.JoinHostPort(a.Host, strconv.Itoa(a.Port)),
		smtp.PlainAuth("", a.Email, a.Password, a.Host),
		a.Email, to, []byte(msg))
	if err != nil {
		return
	}

	return
}
