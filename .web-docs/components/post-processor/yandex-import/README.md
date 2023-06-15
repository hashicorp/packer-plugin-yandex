Type: `yandex-import`
Artifact BuilderId: `packer.post-processor.yandex-import`

The Yandex.Cloud Compute Image Import post-processor create new Compute Image
from a qcow2 file. As Compute service support image creation from Storage service object
just before request to create its upload file into Storage service.

Assigned Service Account must have write permissions to the Yandex Object Storage.
A new temporary static access keys from assigned Service Account used to upload
file.

## Configuration

### Required:

#### Access

<!-- Code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; DO NOT EDIT MANUALLY -->

- `token` (string) - [OAuth token](https://cloud.yandex.com/docs/iam/concepts/authorization/oauth-token)
  or [IAM token](https://cloud.yandex.com/docs/iam/concepts/authorization/iam-token)
  to use to authenticate to Yandex.Cloud. Alternatively you may set
  value by environment variable `YC_TOKEN`.

<!-- End of code generated from the comments of the AccessConfig struct in builder/yandex/access_config.go; -->


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


#### Import

<!-- Code generated from the comments of the Config struct in post-processor/yandex-import/post-processor.go; DO NOT EDIT MANUALLY -->

- `bucket` (string) - The name of the bucket where the qcow2 file will be uploaded to for import.
  This bucket must exist when the post-processor is run.
  
  If import occurred after Yandex-Export post-processor, artifact already
  in storage service and first paths (URL) is used to, so no need to set this param.

- `object_name` (string) - The name of the object key in `bucket` where the qcow2 file will be copied to import.
  This is a [template engine](/packer/docs/templates/legacy_json_templates/engine).
  Therefore, you may use user variables and template functions in this field.

- `skip_clean` (bool) - Whether skip removing the qcow2 file uploaded to Storage
  after the import process has completed. Possible values are: `true` to
  leave it in the bucket, `false` to remove it. Default is `false`.

<!-- End of code generated from the comments of the Config struct in post-processor/yandex-import/post-processor.go; -->


#### Image

<!-- Code generated from the comments of the ImageConfig struct in builder/yandex/common_config.go; DO NOT EDIT MANUALLY -->

- `image_name` (string) - The name of the resulting image, which contains 1-63 characters and only
  supports lowercase English characters, numbers and hyphen. Defaults to
  `packer-{{timestamp}}`.

- `image_description` (string) - The description of the image.

- `image_family` (string) - The family name of the image.

- `image_labels` (map[string]string) - Key/value pair labels to apply to the image.

- `image_min_disk_size_gb` (int) - Minimum size of the disk that will be created from built image, specified in gigabytes.
  Should be more or equal to `disk_size_gb`.

- `image_product_ids` ([]string) - License IDs that indicate which licenses are attached to resulting image.

- `image_pooled` (bool) - When true, an image pool will be created for fast creation disks from the image.

- `skip_create_image` (bool) - Skip creating the image. Useful for setting to `true` during a build test stage. Defaults to `false`.

<!-- End of code generated from the comments of the ImageConfig struct in builder/yandex/common_config.go; -->


## Basic Example

TBD

```json
{
  "variables": {
    "token": "{{env `YC_TOKEN`}}"
  },
  "sensitive-variables": ["token"],
  "builders": [
    {
      "type": "file",
      "source": "xenial-server-cloudimg-amd64-disk1.img",
      "target": "test_artifact.qcow2"
    }
  ],
  "post-processors": [
    {
      "type": "yandex-import",
      "token": "{{user `token`}}",
      "folder_id": "b1g8jvfcgmitdrslcn86",
      "service_account_id": "ajeui8kdvg8qs44fbrbr",

      "bucket": "bucket1",

      "image_name": "my-first-imported-image-{{isotime \"02-Jan-06-03-04-05\" | lower }}",

      "keep_input_artifact": false
    }
  ]
}
```
