module aws_lambda_bench

go 1.17

replace github.com/xxgreg/floats => ../floats

require (
	github.com/aws/aws-lambda-go v1.27.1
	github.com/xxgreg/floats v0.0.0-00010101000000-000000000000
	gonum.org/v1/gonum v0.9.3
)
