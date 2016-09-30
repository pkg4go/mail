package mail

import . "github.com/pkg4go/assert"
import "testing"

func TestFmtMsg(t *testing.T) {
	a := A{t}

	msg := fmtMsg("Hi", "hello world!", "test@mail.com")
	result := "To: test@mail.com\r\n" +
		"Subject: Hi\r\n" +
		"\r\nhello world!\r\n"

	a.Equal(string(msg[:]), result)
}
