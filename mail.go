package emails

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/smtp"
	"strings"
)

type MailSender struct {
	Identity  string
	Host      string
	Port      int
	FromName  string
	FromEmail string
	Username  string
	Password  string
	auth      smtp.Auth
}

func NewMailSender(host string, port int, fromname, fromemail, username, password string) (*MailSender, error) {
	mail := MailSender{
		Identity:  "",
		Host:      host,
		Port:      port,
		FromName:  fromname,
		FromEmail: fromemail,
		Username:  username,
		Password:  password,
	}
	auth := smtp.PlainAuth("", username, password, host)
	mail.auth = auth
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return &mail, err
	}
	defer func() {
		client.Quit()
		client.Close()
	}()

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		VerifyConnection: func(cs tls.ConnectionState) error {
			opts := x509.VerifyOptions{
				DNSName:       cs.ServerName,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err := cs.PeerCertificates[0].Verify(opts)
			return err
		},
	}
	err = client.StartTLS(tlsConfig)
	if err != nil {
		return &mail, err
	}
	err = client.Auth(auth)

	return &mail, err
}

func (ms *MailSender) SendMail(to []string, subject, msg string) error {
	return smtp.SendMail(fmt.Sprintf("%s:%d", ms.Host, ms.Port), ms.auth, ms.FromEmail, to, ms.composemsg(to, subject, msg))
}

func (ms *MailSender) composemsg(to []string, subject, msg string) []byte {
	var byteMsg strings.Builder
	// From
	byteMsg.WriteString(fmt.Sprintf("From: \"%s\" <%s>\r\n", ms.FromName, ms.FromEmail))
	// To
	byteMsg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ", ")))
	// Subject
	byteMsg.WriteString(fmt.Sprintf("Subject: %s\r\n\r\n", subject))
	// Content
	byteMsg.WriteString(fmt.Sprintf("%s\r\n", msg))

	return []byte(byteMsg.String())
}
