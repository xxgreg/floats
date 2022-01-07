clean:
	rm -f build/vmath_x86.zip
	rm -f build/vmath_arm64.zip

build: clean
	GOARCH=amd64 GOOS=linux go build -o build/vmath_x86 .
	zip -jm build/vmath_x86.zip build/vmath_x86
	GOARCH=arm64 GOOS=linux go build -o build/bootstrap .
	zip -jm build/vmath_arm64.zip build/bootstrap

build_noasm: clean
	GOARCH=amd64 GOOS=linux go build -tags noasm -o build/vmath_x86 .
	zip -jm build/vmath_x86.zip build/vmath_x86
	GOARCH=arm64 GOOS=linux go build -tags noasm -o build/bootstrap .
	zip -jm build/vmath_arm64.zip build/bootstrap

deploy:
	aws lambda update-function-code --profile greg-personal --function-name vmath_x86 --zip-file fileb://build/vmath_x86.zip
	aws lambda update-function-code --profile greg-personal --function-name vmath_arm64 --zip-file fileb://build/vmath_arm64.zip

invoke_x86:
	aws lambda invoke --profile greg-personal --function-name vmath_x86 response.json \
	--log-type Tail --query 'LogResult' --output text |  base64 -d

invoke_arm64:
	aws lambda invoke --profile greg-personal --function-name vmath_arm64 response.json \
	--log-type Tail --query 'LogResult' --output text |  base64 -d

run:
	go run . local

run_noasm:
	go run -tags noasm . local
