package gomail

import (
	"mime"
	"mime/quotedprintable"
)

var newQPWriter = quotedprintable.NewWriter

type mimeEncoder struct {
	mime.WordEncoder
}
