package utils

import "github.com/rs/xid"

func GenLogId() string {
	guid := xid.New()
	return guid.String()
}
