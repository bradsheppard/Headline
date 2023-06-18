resource "google_artifact_registry_repository" "repository" {
  location      = var.location
  repository_id = var.name
  format        = "DOCKER"
  description   = var.description
}

