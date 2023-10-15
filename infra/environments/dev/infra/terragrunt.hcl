remote_state {
  backend = "gcs"
  generate = {
    path = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
  config = {
    bucket = "headline-inf-remote-state"
    prefix = "${path_relative_to_include()}"
    project = "headline-inf624ade82"
    location = "us-central1"
  }
}

generate "provider" {
  path = "providers.tf"
  if_exists = "overwrite_terragrunt"

  contents = <<EOF
provider google {
  project = "headline-inf624ade82"
}
provider google-beta {
  project = "headline-inf624ade82"
}
EOF
}
