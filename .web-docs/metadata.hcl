# For full specification on the configuration of this file visit:
# https://github.com/hashicorp/integration-template#metadata-configuration
integration {
  name = "Yandex"
  description = "The Yandex plugin is able to manage images for use with the Yandex Compute Cloud."
  identifier = "packer/hashicorp/yandex"
  component {
    type = "builder"
    name = "Yandex Compute"
    slug = "yandex"
  }
  component {
    type = "post-processor"
    name = "Yandex Export"
    slug = "yandex-export"
  }
  component {
    type = "post-processor"
    name = "Yandex Import"
    slug = "yandex-import"
  }
}
