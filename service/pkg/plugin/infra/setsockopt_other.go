//go:build !linux
// +build !linux

package infra

import "syscall"

func SoMarkControl(c syscall.RawConn) error {
	return nil
}
func BindControl(c syscall.RawConn, laddr string, lport uint32) error {
	return nil
}