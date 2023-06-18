include {
  path = find_in_parent_folders()
}

terraform {
  source = "../../../..//modules/registry"
}

inputs = {
  location = "us-central1"
  name = "headline"
  description = "headline docker registry"
}

