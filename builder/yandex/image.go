// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package yandex

import (
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
)

type Image struct {
	ID            string
	FolderID      string
	Labels        map[string]string
	Licenses      []string
	MinDiskSizeGb int
	Name          string
	Description   string
	Family        string
	SizeGb        int
	Os            *compute.Os
}

func convert(image *compute.Image) *Image {
	return &Image{
		ID:            image.Id,
		Labels:        image.Labels,
		Licenses:      image.ProductIds,
		Name:          image.Name,
		Family:        image.Family,
		Description:   image.Description,
		FolderID:      image.FolderId,
		MinDiskSizeGb: toGigabytes(image.MinDiskSize),
		SizeGb:        toGigabytes(image.StorageSize),
		Os:            image.Os,
	}
}
