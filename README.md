# Emails(ender)
> A simple but useful email sender wrapper via SMTP protocol.

# Installation

```bash
go get -u github.com/collatzc/emails
```

# Quick Start

```go
mailSender, err := NewMailSender("smtphost", 587, "SenderName", "senderDisplayEmail", "accountEmail", "accountPassword")
if err != nil {
	log.Fatal("Unexpected err", err)
}

mailSender.SendMail([]string{"recipientEmail"}, "subject", "content")
```

# License
[The MIT License (MIT)](LICENSE)
