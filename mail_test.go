package emails

import "log"

func ExampleSendMail() {
	mailSender, err := NewMailSender("smtphost", 587, "SenderName", "senderDisplayEmail", "accountEmail", "accountPassword")
	if err != nil {
		log.Fatal("Unexpected err", err)
	}

	mailSender.SendMail([]string{"recipientEmail"}, "subject", "content")
}
