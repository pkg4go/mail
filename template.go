package mail

func fmtMsg(title, content, to string) []byte {
	return []byte("To: " + to + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" + content + "\r\n")
}
