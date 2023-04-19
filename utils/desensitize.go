package utils

import "strings"

func DesnsitizeEmail(email string) string {
	at := strings.Index(email, "@")
	if at > 1 {
		e := email[:at]
		if len(e) > 4 {
			e = e[:2] + "****" + e[len(e)-2:]
		} else if len(e) > 2 {
			e = e[:1] + "****" + e[len(e)-1:]
		}
		return e + email[at:]
	}

	return ""
}
