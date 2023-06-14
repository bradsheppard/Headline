include {
  path = find_in_parent_folders()
}

dependency "networking" {
  config_path = "..//network"
}

terraform {
  source = "../../../..//modules/kubernetes"
}

inputs = {
  name = "headline"
  zone = "us-central1-b"
  vpc_id = dependency.networking.outputs.vpc_id
  subnet_id = dependency.networking.outputs.subnet_id
  machine_type = "n1-standard-2"
}
