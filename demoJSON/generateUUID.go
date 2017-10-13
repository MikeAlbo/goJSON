package main

import (
	"io"
	"fmt"
	"crypto/rand"
)

// this package will generate a new random 16 byte user id
// ref: https://play.golang.org/p/4FkNSiUDMg

func NewUUID() (string, error)  {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
