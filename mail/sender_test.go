package mail

import (
	"testing"

	"github.com/dj-yacine-flutter/gojo/conf"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	config, err := conf.Load("..", "gojo")
	require.NoError(t, err)

	sender := NewGmailSender(config.Email.Name, config.Email.Address, config.Email.Password)

	subject := "A test email"
	content := `
			<h1>Hello world</h1>
			<p>This is a test message from DJ Yacine </p>
			`
	to := []string{"0dj.yacine0@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)

}
