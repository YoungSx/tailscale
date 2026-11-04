// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package netmapcache

// TODO(cmol): Add compiler flags when having more stores

func NewStore(dir string) (Store, error) {
	return FileStore(dir), nil
}
