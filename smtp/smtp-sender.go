package smtp

import "io"

type smtpSender struct {
	smtpClient
	d *Dialer
}

func (s *smtpSender) Send(from string, to []string, msg io.WriterTo) error {
	if err := s.Mail(from); err != nil {
		if err == io.EOF {
			// This is probably due to a timeout, so reconnect and try again.
			sc, dErr := s.d.Dial()
			if dErr == nil {
				if sender, ok := sc.(*smtpSender); ok {
					*s = *sender
					return s.Send(from, to, msg)
				}
			}
		}
		return err
	}

	for _, addr := range to {
		if err := s.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := s.Data()
	if err != nil {
		return err
	}

	if _, err = msg.WriteTo(w); err != nil {
		w.Close()
		return err
	}

	return w.Close()
}

func (c *smtpSender) Close() error {
	return c.Quit()
}
