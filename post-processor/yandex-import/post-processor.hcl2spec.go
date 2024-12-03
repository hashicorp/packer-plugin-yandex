// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package yandeximport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Endpoint              *string           `mapstructure:"endpoint" required:"false" cty:"endpoint" hcl:"endpoint"`
	ServiceAccountKeyFile *string           `mapstructure:"service_account_key_file" required:"false" cty:"service_account_key_file" hcl:"service_account_key_file"`
	Token                 *string           `mapstructure:"token" required:"true" cty:"token" hcl:"token"`
	MaxRetries            *int              `mapstructure:"max_retries" cty:"max_retries" hcl:"max_retries"`
	FolderID              *string           `mapstructure:"folder_id" required:"true" cty:"folder_id" hcl:"folder_id"`
	ServiceAccountID      *string           `mapstructure:"service_account_id" required:"true" cty:"service_account_id" hcl:"service_account_id"`
	ImageName             *string           `mapstructure:"image_name" required:"false" cty:"image_name" hcl:"image_name"`
	ImageDescription      *string           `mapstructure:"image_description" required:"false" cty:"image_description" hcl:"image_description"`
	ImageFamily           *string           `mapstructure:"image_family" required:"false" cty:"image_family" hcl:"image_family"`
	ImageLabels           map[string]string `mapstructure:"image_labels" required:"false" cty:"image_labels" hcl:"image_labels"`
	ImageMinDiskSizeGb    *int              `mapstructure:"image_min_disk_size_gb" required:"false" cty:"image_min_disk_size_gb" hcl:"image_min_disk_size_gb"`
	ImageProductIDs       []string          `mapstructure:"image_product_ids" required:"false" cty:"image_product_ids" hcl:"image_product_ids"`
	ImagePooled           *bool             `mapstructure:"image_pooled" required:"false" cty:"image_pooled" hcl:"image_pooled"`
	SkipCreateImage       *bool             `mapstructure:"skip_create_image" required:"false" cty:"skip_create_image" hcl:"skip_create_image"`
	ImagePCITopology      *string           `mapstructure:"image_pci_topology" required:"false" cty:"image_pci_topology" hcl:"image_pci_topology"`
	Bucket                *string           `mapstructure:"bucket" required:"false" cty:"bucket" hcl:"bucket"`
	ObjectName            *string           `mapstructure:"object_name" required:"false" cty:"object_name" hcl:"object_name"`
	SkipClean             *bool             `mapstructure:"skip_clean" required:"false" cty:"skip_clean" hcl:"skip_clean"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"endpoint":                   &hcldec.AttrSpec{Name: "endpoint", Type: cty.String, Required: false},
		"service_account_key_file":   &hcldec.AttrSpec{Name: "service_account_key_file", Type: cty.String, Required: false},
		"token":                      &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"max_retries":                &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"folder_id":                  &hcldec.AttrSpec{Name: "folder_id", Type: cty.String, Required: false},
		"service_account_id":         &hcldec.AttrSpec{Name: "service_account_id", Type: cty.String, Required: false},
		"image_name":                 &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":          &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_family":               &hcldec.AttrSpec{Name: "image_family", Type: cty.String, Required: false},
		"image_labels":               &hcldec.AttrSpec{Name: "image_labels", Type: cty.Map(cty.String), Required: false},
		"image_min_disk_size_gb":     &hcldec.AttrSpec{Name: "image_min_disk_size_gb", Type: cty.Number, Required: false},
		"image_product_ids":          &hcldec.AttrSpec{Name: "image_product_ids", Type: cty.List(cty.String), Required: false},
		"image_pooled":               &hcldec.AttrSpec{Name: "image_pooled", Type: cty.Bool, Required: false},
		"skip_create_image":          &hcldec.AttrSpec{Name: "skip_create_image", Type: cty.Bool, Required: false},
		"image_pci_topology":         &hcldec.AttrSpec{Name: "image_pci_topology", Type: cty.String, Required: false},
		"bucket":                     &hcldec.AttrSpec{Name: "bucket", Type: cty.String, Required: false},
		"object_name":                &hcldec.AttrSpec{Name: "object_name", Type: cty.String, Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
	}
	return s
}
