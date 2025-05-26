output "repository_id" {
  description = "The ID of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.id
}

output "repository_name" {
  description = "The name of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.name
}

output "repository_location" {
  description = "The location of the Artifact Registry repository"
  value       = google_artifact_registry_repository.repo.location
}

output "repository_format" {
  description = "The format of the Artifact Registry repository (e.g., DOCKER, MAVEN, etc.)"
  value       = google_artifact_registry_repository.repo.format
}

output "repository_url" {
  description = "The full URL of the Artifact Registry repository"
  value       = "${google_artifact_registry_repository.repo.location}-${lower(google_artifact_registry_repository.repo.format)}.pkg.dev/${google_artifact_registry_repository.repo.project}/${google_artifact_registry_repository.repo.name}"
}
