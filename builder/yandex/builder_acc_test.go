package yandex

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

const InstanceMetadataAddr = "169.254.169.254"

func TestBuilderAcc_basic(t *testing.T) {
	testAccPreCheck(t)
	testCase := &acctest.PluginTestCase{
		Name:     "yandex_basic_test",
		Template: testBuilderAccBasic,
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

func TestBuilderAcc_instanceSA(t *testing.T) {
	testAccPreCheckInstanceSA(t)
	testCase := &acctest.PluginTestCase{
		Name:     "yandex_basic_test_instance_sa",
		Template: testBuilderAccInstanceSA,
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			if buildCommand.ProcessState != nil {
				if buildCommand.ProcessState.ExitCode() != 0 {
					return fmt.Errorf("Bad exit code. Logfile: %s", logfile)
				}
			}
			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("YC_TOKEN"); v == "" {
		t.Skip("YC_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("YC_FOLDER_ID"); v == "" {
		t.Skip("YC_FOLDER_ID must be set for acceptance tests")
	}
}

func testAccPreCheckInstanceSA(t *testing.T) {
	if v := os.Getenv("YC_FOLDER_ID"); v == "" {
		t.Skip("YC_FOLDER_ID must be set for acceptance tests")
	}

	client := resty.New()

	_, err := client.R().SetHeader("Metadata-Flavor", "Google").Get(tokenUrl())
	if err != nil {
		t.Fatalf("error get Service Account token assignment: %s", err)
	}

}

const testBuilderAccBasic = `
{
	"builders": [{
		"type": "test",
        "source_image_family": "ubuntu-1804-lts",
		"use_ipv4_nat": "true",
		"ssh_username": "ubuntu"
	}]
}
`

const testBuilderAccInstanceSA = `
{
	"builders": [{
		"type": "test",
        "source_image_family": "ubuntu-1804-lts",
		"use_ipv4_nat": "true",
		"ssh_username": "ubuntu"
	}]
}
`

func tokenUrl() string {
	return fmt.Sprintf("http://%s/computeMetadata/v1/instance/service-accounts/default/token", InstanceMetadataAddr)
}
