// Copyright 2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ants

import (
	tp "github.com/henrylee2cn/teleport"
)

// Linker linker for client.
type Linker interface {
	// Select selects a service address by URI path.
	Select(uriPath string) (addr string, rerr *tp.Rerror)
	// EventDel pushs service node offline notification.
	EventDel() <-chan string
	// Void voids this address.
	Void(addr string)
	// Close closes the linker.
	Close()
}

// static linker

// NewStaticLinker creates a static linker.
func NewStaticLinker(srvAddr string) Linker {
	return &staticLinker{
		srvAddr: srvAddr,
		ch:      make(chan string),
	}
}

type staticLinker struct {
	srvAddr string
	ch      chan string
}

// Select selects a service address by URI path.
func (d *staticLinker) Select(string) (string, *tp.Rerror) {
	return d.srvAddr, nil
}

// EventDel pushs service node offline notification.
func (d *staticLinker) EventDel() <-chan string {
	return d.ch
}

// Void voids this address.
func (d *staticLinker) Void(string) {}

// Close closes the linker.
func (d *staticLinker) Close() {
	close(d.ch)
}
