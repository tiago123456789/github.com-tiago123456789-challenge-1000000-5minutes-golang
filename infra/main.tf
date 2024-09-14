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

resource "aws_instance" "challenge_1000000_5minutes" {
  ami           = "ami-0e86e20dae9224db8"
  instance_type = "t2.micro"

  tags = {
    Name = "challenge_1000000_5minutes"
  }

  key_name = "challenge-1000000-5minutes-golang"
}

# Output the instance ID
output "instance_id" {
  value = aws_instance.challenge_1000000_5minutes.id
}

# Output the public IP address
output "public_ip" {
  value = aws_instance.challenge_1000000_5minutes.public_ip
}

output "queue_url" {
  value = aws_sqs_queue.challenge_1000000_5minutes.url
}
