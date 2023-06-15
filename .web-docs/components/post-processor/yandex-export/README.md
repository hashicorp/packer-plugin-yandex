Type: `yandex-export`
Artifact BuilderId: `packer.post-processor.yandex-export`

The Yandex.Cloud Compute Image Exporter post-processor exports the resultant image
from a yandex build as a qcow2 file to Yandex Object Storage.

The exporter uses the same Yandex.Cloud folder and
authentication credentials as the yandex build that produced the image.
A temporary VM is started in the folder using these credentials. The VM
mounts the built image as a secondary disk, then dumps the image in qcow2 format.
The VM then uploads the file to the provided Yandex Object Storage `paths` using the same
credentials.

As such, assigned Service Account must have write permissions to the Yandex Object Storage
`paths`. A new temporary static access keys from assigned Service Account used to upload
image.

Also, you should configure [ssh communicator](/packer/docs/communicators/ssh). Default `ssh_username` to `ubuntu`.

## Configuration

### Required:

#### Access

<!-- Code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; DO NOT EDIT MANUALLY -->

- `token` (string) - [OAuth token](https://cloud.yandex.com/docs/iam/concepts/authorization/oauth-token)
  or [IAM token](https://cloud.yandex.com/docs/iam/concepts/authorization/iam-token)
  to use to authenticate to Yandex.Cloud. Alternatively you may set
  value by environment variable `YC_TOKEN`.

<!-- End of code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; -->


#### Export

<!-- Code generated from the comments of the Config struct in post-processor/yandex-export/post-processor.go; DO NOT EDIT MANUALLY -->

- `paths` ([]string) - List of paths to Yandex Object Storage where exported image will be uploaded.
  Please be aware that use of space char inside path not supported.
  Also this param support [build](/packer/docs/templates/legacy_json_templates/engine) template function.
  Check available template data for [Yandex](/packer/integrations/hashicorp/yandex#build-template-data) builder.
  Paths to Yandex Object Storage where exported image will be uploaded.

<!-- End of code generated from the comments of the Config struct in post-processor/yandex-export/post-processor.go; -->


#### Common

<!-- Code generated from the comments of the CloudConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `folder_id` (string) - The folder ID that will be used to launch instances and store images.
  Alternatively you may set value by environment variable `YC_FOLDER_ID`.
  To use a different folder for looking up the source image or saving the target image to
  check options 'source_image_folder_id' and 'target_image_folder_id'.

<!-- End of code generated from the comments of the CloudConfig struct in builder/yandex/common_config.go; -->


<!-- Code generated from the comments of the ExchangeConfig struct in post-processor/yandex-export/config.go; DO NOT EDIT MANUALLY -->

- `service_account_id` (string) - Service Account ID with proper permission to modify an instance, create and attach disk and
  make upload to specific Yandex Object Storage paths.

<!-- End of code generated from the comments of the ExchangeConfig struct in post-processor/yandex-export/config.go; -->


### Optional:

#### Access

<!-- Code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; DO NOT EDIT MANUALLY -->

- `endpoint` (string) - Non standard API endpoint. Default is `api.cloud.yandex.net:443`.

- `service_account_key_file` (string) - Path to file with Service Account key in json format. This
  is an alternative method to authenticate to Yandex.Cloud. Alternatively you may set environment variable
  `YC_SERVICE_ACCOUNT_KEY_FILE`.

- `max_retries` (int) - The maximum number of times an API request is being executed.

<!-- End of code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; -->


#### Export

<!-- Code generated from the comments of the Config struct in post-processor/yandex-export/post-processor.go; DO NOT EDIT MANUALLY -->

- `source_image_folder_id` (string) - The ID of the folder containing the source image. Default `standard-images`.

- `source_image_family` (string) - The source image family to start export process. Default `ubuntu-1604-lts`.
  Image must contains utils or supported package manager: `apt` or `yum` -
  requires `root` or `sudo` without password.
  Utils: `qemu-img`, `aws`. The `qemu-img` utility requires `root` user or
  `sudo` access without password.

- `source_image_id` (string) - The source image ID to use to create the new image from. Just one of a source_image_id or
  source_image_family must be specified.

- `source_disk_extra_size` (int) - The extra size of the source disk in GB. This defaults to `0GB`.
  Requires `losetup` utility on the instance.
  > **Careful!** Increases payment cost.
  > See [perfomance](https://cloud.yandex.com/docs/compute/concepts/disk#performance).

- `storage_endpoint` (string) - StorageEndpoint custom Yandex Object Storage endpoint to upload image, Default `storage.yandexcloud.net`.

- `storage_endpoint_autoresolve` (bool) - StorageEndpointAutoresolve auto resolve storage endpoint via YC Public API ListEndpoints call. Option has
  precedence over 'storage_endpoint' option.

- `storage_region` (string) - StorageRegion custom Yandex Object region. Default `ru-central1`

<!-- End of code generated from the comments of the Config struct in post-processor/yandex-export/post-processor.go; -->


#### Common

<!-- Code generated from the comments of the CommonConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `serial_log_file` (string) - File path to save serial port output of the launched instance.

- `state_timeout` (duration string | ex: "1h5m2s") - The time to wait for instance state changes.
  Defaults to `5m`.

<!-- End of code generated from the comments of the CommonConfig struct in builder/yandex/common_config.go; -->


#### Instance

<!-- Code generated from the comments of the InstanceConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `instance_cores` (int) - The number of cores available to the instance.

- `instance_gpus` (int) - The number of GPU available to the instance.

- `instance_mem_gb` (int) - The amount of memory available to the instance, specified in gigabytes.

- `instance_name` (string) - The name assigned to the instance.

- `platform_id` (string) - Identifier of the hardware platform configuration for the instance. This defaults to `standard-v2`.

- `labels` (map[string]string) - Key/value pair labels to apply to the launched instance.

- `metadata` (map[string]string) - Metadata applied to the launched instance.

- `metadata_from_file` (map[string]string) - Metadata applied to the launched instance.
  The values in this map are the paths to the content files for the corresponding metadata keys.

- `preemptible` (bool) - Launch a preemptible instance. This defaults to `false`.

<!-- End of code generated from the comments of the InstanceConfig struct in builder/yandex/common_config.go; -->


#### Disk

<!-- Code generated from the comments of the DiskConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `disk_name` (string) - The name of the disk, if unset the instance name
  will be used.

- `disk_size_gb` (int) - The size of the disk in GB. This defaults to 10/100GB.

- `disk_type` (string) - Specify disk type for the launched instance. Defaults to `network-ssd`.

- `disk_labels` (map[string]string) - Key/value pair labels to apply to the disk.

<!-- End of code generated from the comments of the DiskConfig struct in builder/yandex/common_config.go; -->


#### Network

<!-- Code generated from the comments of the NetworkConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `subnet_id` (string) - The Yandex VPC subnet id to use for
  the launched instance. Note, the zone of the subnet must match the
  zone in which the VM is launched.

- `zone` (string) - The name of the zone to launch the instance.  This defaults to `ru-central1-a`.

- `security_group_ids` ([]string) - Security group ids for network interface of the instance.

- `use_ipv4_nat` (bool) - If set to true, then launched instance will have external internet
  access.

- `use_ipv6` (bool) - Set to true to enable IPv6 for the instance being
  created. This defaults to `false`, or not enabled.
  
  -> **Note**: Usage of IPv6 will be available in the future.

- `use_internal_ip` (bool) - If true, use the instance's internal IP address
  instead of its external IP during building.

<!-- End of code generated from the comments of the NetworkConfig struct in builder/yandex/common_config.go; -->


## Basic Example

The following example builds a Compute image in the folder with id `b1g8jvfcgmitdrslcn86`, with an
Service Account whose keyfile is `account.json`. After the image build, a temporary VM
will be created to export the image as a qcow2 file to
`s3://packer-export/my-exported-image.qcow2` and
`s3://packer-export/image-number-two.qcow2`. `keep_input_artifact` is true, so the
source Compute image won't be deleted after the export.

In order for this example to work, the service account associated with builder
must have write access to both `s3://packer-export/my-exported-image.qcow2` and
`s3://packer-export/image-number-two.qcow2` and get permission to modify temporary instance
(create new disk, attach to instance, etc).

```json
{
  "builders": [
    {
      "type": "yandex",
      "folder_id": "b1g8jvfcgmitdrslcn86",
      "subnet_id": "e9bp6l8sa4q39yourxzq",
      "zone": "ru-central1-a",

      "source_image_family": "ubuntu-1604-lts",
      "ssh_username": "ubuntu",
      "use_ipv4_nat": true
    }
  ],
  "post-processors": [
    {
      "type": "yandex-export",
      "folder_id": "b1g8jvfcgmitdrslcn86",
      "subnet_id": "e9bp6l8sa4q39yourxzq",

      "service_account_id": "ajeu0363240rrnn7xgen",

      "paths": [
        "s3://packer-export-bucket/my-exported-image.qcow2",
        "s3://packer-export-bucket/template-supported-get-{{build `ImageID` }}-right-here.qcow2"
      ],
      "keep_input_artifact": true
    }
  ]
}
```
