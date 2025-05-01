variable "repo_id" {
  description = "The ID of the repository."
  type        = string
}

variable "location" {
  description = "The location of the repository."
  type        = string
  
}

variable "description" {
  description = "The description of the repository."
  type        = string
  default     = "My Artifact Registry Repository"
  
}

variable "format" {
  description = "The format of the repository."
  type        = string
  default     = "docker"
}

variable "project_id" {
  description = "The ID of the project."
  type        = string
}