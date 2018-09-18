package gomail

import (
	"crypto/tls"
	"fmt"
	"mime"
	"net"
	"net/smtp"
	"strings"
)

// Stubbed out for tests.
var (
	netDialTimeout = net.DialTimeout
	tlsClient      = tls.Client
	smtpNewClient  = func(conn net.Conn, host string) (smtpClient, error) {
		return smtp.NewClient(conn, host)
	}
	bEncoding     = mimeEncoder{mime.BEncoding}
	qEncoding     = mimeEncoder{mime.QEncoding}
	lastIndexByte = strings.LastIndexByte
)

func addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
