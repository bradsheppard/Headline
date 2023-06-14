terraform {
  source = "../../..//modules/project"
}

inputs = {
  billing_account = "01BED7-340A56-12F88F"
  name = "headline-infra"
}

remote_state {
  backend = "local"
  generate = {
    path = "backend.tf"
    if_exists = "overwrite_terragrunt"
  }
  config = {}
}
