The Yandex plugin is able to manage images for use with the Yandex Compute Cloud.

### Installation
To install this plugin add this code into your Packer configuration and run [packer init](/packer/docs/commands/init)

```hcl
packer {
  required_plugins {
    yandex = {
      version = "~> 1"
      source  = "github.com/hashicorp/yandex"
    }
  }
}
```

Alternatively, you can use `packer plugins install` to manage installation of this plugin.

```sh
packer plugins install github.com/hashicorp/yandex
```

### Components
#### Builders

- [yandex](/docs/builders/builder-name.mdx) - The builder is able to create images for use with Yandex Compute Cloud based on existing images.

### Post-processors

- [yandex-export](/packer/integration/hashicorp/yandex/latest/components/post-processor/yandex-export) - The export post-processor exports the resultant image from a Yandex 
  build as a qcow2 file to Yandex Object Storage.
- [yandex-import](/packer/integration/hashicorp/yandex/latest/components/post-processor/yandex-import) - The Import post-processor create new Compute Image from a qcow2 file.


