provider "aws" {
  region = "us-east-1"  
  profile = "tiago"
}

# Define the SQS Queue
resource "aws_sqs_queue" "challenge_1000000_5minutes" {
  name = "challenge_1000000_5minutes"
  
  delay_seconds               = 0
  message_retention_seconds   = 345600
  visibility_timeout_seconds  = 30
}

output "queue_url" {
  value = aws_sqs_queue.challenge_1000000_5minutes.url
}
