remote_state {
  backend = "gcs"
  generate = {
    path = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
  config = {
    bucket = "headline-infra-remote-state"
    prefix = "${path_relative_to_include()}"
    project = "headline-infrae012d38d"
    location = "us-central1"
  }
}

generate "provider" {
  path = "providers.tf"
  if_exists = "overwrite_terragrunt"

  contents = <<EOF
provider google {
  project = "headline-infrae012d38d"
}
provider google-beta {
  project = "headline-infrae012d38d"
}
EOF
}
