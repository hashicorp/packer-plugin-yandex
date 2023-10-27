// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc struct-markdown

package yandex

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/yandex-cloud/go-sdk/iamkey"
)

type keyType int

const (
	Undefined keyType = iota
	File
	Content
)

const (
	defaultEndpoint   = "api.cloud.yandex.net:443"
	defaultMaxRetries = 3
)

// AccessConfig is for common configuration related to Yandex.Cloud API access
type AccessConfig struct {
	// Non standard API endpoint. Default is `api.cloud.yandex.net:443`.
	Endpoint string `mapstructure:"endpoint" required:"false"`
	// Contains either a path to or the contents of the Service Account file in JSON format.
	// This can also be specified using environment variable `YC_SERVICE_ACCOUNT_KEY_FILE`.
	// You can read how to create service account key file [here](https://cloud.yandex.com/docs/iam/operations/iam-token/create-for-sa#keys-create).
	ServiceAccountKeyFile string `mapstructure:"service_account_key_file" required:"false"`
	// [OAuth token](https://cloud.yandex.com/docs/iam/concepts/authorization/oauth-token)
	// or [IAM token](https://cloud.yandex.com/docs/iam/concepts/authorization/iam-token)
	// to use to authenticate to Yandex.Cloud. Alternatively you may set
	// value by environment variable `YC_TOKEN`.
	Token string `mapstructure:"token" required:"true"`
	// The maximum number of times an API request is being executed.
	MaxRetries int `mapstructure:"max_retries"`

	saKeyType keyType
}

func (c *AccessConfig) Prepare(ctx *interpolate.Context) []error {
	var errs []error

	if c.MaxRetries == 0 {
		c.MaxRetries = defaultMaxRetries
	}

	if c.Endpoint == "" {
		c.Endpoint = defaultEndpoint
	}

	// provision config by OS environment variables
	if c.Token == "" {
		c.Token = os.Getenv("YC_TOKEN")
	}

	if c.ServiceAccountKeyFile == "" {
		c.ServiceAccountKeyFile = os.Getenv("YC_SERVICE_ACCOUNT_KEY_FILE")
	}

	if c.Token != "" && c.ServiceAccountKeyFile != "" {
		errs = append(errs, errors.New("one of token or service account key file must be specified, not both"))
	}

	if c.Token != "" {
		packersdk.LogSecretFilter.Set(c.Token)
	}

	if c.ServiceAccountKeyFile != "" {
		// if ServiceAccountKeyFile is file path
		if _, err := os.Stat(c.ServiceAccountKeyFile); err == nil {
			if _, err := iamkey.ReadFromJSONFile(c.ServiceAccountKeyFile); err != nil {
				errs = append(errs, fmt.Errorf("fail to read service account key file: %s", err))
			}
			c.saKeyType = File
		} else {
			// else check for a valid json data value
			var f map[string]interface{}
			if err := json.Unmarshal([]byte(c.ServiceAccountKeyFile), &f); err != nil {
				errs = append(errs, fmt.Errorf("JSON in %q are not valid: %s", c.ServiceAccountKeyFile, err))
			}
			c.saKeyType = Content
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
