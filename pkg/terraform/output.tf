output "artifact_registry_repository_id" {
  description = "The ID of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.id
}

output "artifact_registry_repository_name" {
  description = "The name of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.name
}

output "artifact_registry_repository_location" {
  description = "The location of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.location
}

output "artifact_registry_repository_format" {
  description = "The format of the Artifact Registry repository (e.g., DOCKER, MAVEN, etc.)"
  value       = google_artifact_registry_repository.repo.format
}

output "artifact_registry_repository_url" {
  description = "The full URL of the Artifact Registry repository"
  value       = "https://${google_artifact_registry_repository.repo.location}-docker.pkg.dev/${google_artifact_registry_repository.repo.project}/${google_artifact_registry_repository.repo.name}"
}
