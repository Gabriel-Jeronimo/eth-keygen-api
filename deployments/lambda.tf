data "archive_file" "lambda" {
  source_dir  = "src/"
  output_path = "src/lambda.zip"
  type        = "zip"
}

resource "aws_lambda_function" "lambda_sqs" {
  function_name    = "lambda"
  handler          = "main"
  role             = aws_iam_role.lambda_exec_role.arn
  runtime          = "go1.x"
  filename         = data.archive_file.lambda.output_path
  source_code_hash = data.archive_file.lambda.output_base64sha256
}


