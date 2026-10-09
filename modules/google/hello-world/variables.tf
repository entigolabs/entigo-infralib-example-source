variable "prefix" {
  type = string
}

variable "greeting" {
  type        = string
  default     = "Hello"
  description = "The word the output greets with, for example Hello."
}
