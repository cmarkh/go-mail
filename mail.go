package mail

import (
	"errors"
	"net"
	"net/smtp"
	"strconv"
	"strings"
)

var defaultAccount struct {
	set bool
	Account
}

type Account struct {
	Name     string //display name of account user, ie John Smith
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

// Send formats the headers and sends the email
func (a Account) Send(subject, body string, to ...string) (err error) {
	if len(to) == 0 {
		return errors.New("recipients cannot be blank")
	}

	msg := "From: " + a.Name + " <" + a.Email + ">" + "\r\n" +
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

// SetAsDefault sets a default account so the Send function can be used without specifying an account
func (a Account) SetAsDefault() {
	defaultAccount.Account = a
	defaultAccount.set = true
}

func Send(subject, body string, to ...string) (err error) {
	if !defaultAccount.set {
		return errors.New("no account set. use the SendFromAccount function to specify a global default or use account.Send")
	}
	return defaultAccount.Send(subject, body, to...)
}
