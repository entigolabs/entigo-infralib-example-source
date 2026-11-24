variable "prefix" {
  type        = string
  description = "Name prefix of the environment, set by the agent."
}

variable "greeting" {
  type        = string
  default     = "Hello"
  description = "The word the output greets with, for example Hello."
}
