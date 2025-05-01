provider "google" {
  project = var.project_id
  region  = var.location
  
}

provider "google-beta" {
  project = var.project_id
  region  = var.location
}

terraform {
  backend "gcs" {}
}

resource "google_artifact_registry_repository" "repo" {
  location      = var.location
  repository_id = var.repo_id
  description   = var.description
  format        = var.format
  project = var.project_id
}