// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !windows

// Package netutil provides network-related helpers.
package netutil

import (
	"errors"
	"syscall"
)

// IsAddrUnavailable reports whether err indicates that the address
// is unavailable (as opposed to a specific port being busy).
//
// ENODEV and EINVAL come back when a link-local address carries the zone of an
// interface that no longer exists (by index or by name). Like EADDRNOTAVAIL,
// they do not depend on the port, so retrying other ports cannot succeed.
func IsAddrUnavailable(err error) bool {
	var errno syscall.Errno
	if errors.As(err, &errno) {
		return errno == syscall.EADDRNOTAVAIL || errno == syscall.ENODEV || errno == syscall.EINVAL
	}

	return false
}
