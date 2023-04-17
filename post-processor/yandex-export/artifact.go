// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package yandexexport

import (
	"fmt"
)

const BuilderId = "packer.post-processor.yandex-export"

type Artifact struct {
	paths []string
	urls  []string
}

func (*Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Id() string {
	return a.urls[0]
}

func (a *Artifact) Files() []string {
	pathsCopy := make([]string, len(a.paths))
	copy(pathsCopy, a.paths)
	return pathsCopy
}

func (a *Artifact) String() string {
	return fmt.Sprintf("Exported artifacts in: %s", a.paths)
}

func (*Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
	return nil
}
