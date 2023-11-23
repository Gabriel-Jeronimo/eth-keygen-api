resource "aws_iam_role" "apiSQS" {
  name = "apigateway_sqsa"

  assume_role_policy = jsonencode({
    Version : "2012-10-17",
    Statement : [
      {
        Action : "sts:AssumeRole",
        Principal : {
          Service : "apigateway.amazonaws.com"
        },
        Effect : "Allow",
        Sid : ""
      }
    ]
  })
}

data "template_file" "gateway_policy" {
  template = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogGroups",
          "logs:DescribeLogStreams",
          "logs:PutLogEvents",
          "logs:GetLogEvents",
          "logs:FilterLogEvents"
        ],
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "sqs:GetQueueUrl",
          "sqs:ChangeMessageVisibility",
          "sqs:ListDeadLetterSourceQueues",
          "sqs:SendMessageBatch",
          "sqs:PurgeQueue",
          "sqs:ReceiveMessage",
          "sqs:SendMessage",
          "sqs:GetQueueAttributes",
          "sqs:CreateQueue",
          "sqs:ListQueueTags",
          "sqs:ChangeMessageVisibilityBatch",
          "sqs:SetQueueAttributes"
        ],
        "Resource" : "${aws_sqs_queue.queue.arn}"
      },
      {
        "Effect" : "Allow",
        "Action" : "sqs:ListQueues",
        "Resource" : "*"
      }
    ]
    }
  )
}

resource "aws_iam_policy" "api_policy" {
  name = "api-sqs-cloudwatch-policya"

  policy = data.template_file.gateway_policy.rendered
}

resource "aws_iam_role_policy_attachment" "api_exec_role" {
  role       = aws_iam_role.apiSQS.name
  policy_arn = aws_iam_policy.api_policy.arn
}


data "template_file" "lambda_policy" {
  template = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : [
          "sqs:DeleteMessage",
          "sqs:ReceiveMessage",
          "sqs:GetQueueAttributes"
        ],
        "Resource" : "${aws_sqs_queue.queue.arn}",
        "Effect" : "Allow"
      },
      {
        "Action" : [
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:PutLogEvents"
        ],
        "Resource" : "arn:aws:logs:*:*:*",
        "Effect" : "Allow"
      }
    ]
    }
  )
}

resource "aws_iam_policy" "lambda_sqs_policy" {
  name        = "lambda_policy_dba"
  description = "IAM policy for lambda being invoked by SQS"

  policy = data.template_file.lambda_policy.rendered
}

resource "aws_iam_role" "lambda_exec_role" {
  name = "${var.name}-lambda-db"
  assume_role_policy = jsonencode({
    Version : "2012-10-17"
    Statement : [
      {
        Action : "sts:AssumeRole"
        Principal : {
          Service : "lambda.amazonaws.com"
        },
        Effect : "Allow",
        Sid : ""
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_role_policy" {
  role = aws_iam_role.lambda_exec_role.name

  policy_arn = aws_iam_policy.lambda_sqs_policy.arn
}
