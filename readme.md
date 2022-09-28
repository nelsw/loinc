# LOINC Web Service ![](https://img.shields.io/badge/LOINC-v2.73.0-blue)

A project for accessing the LOINC Core Code Table values through a serverless API.

Learn more about LOINC.org [here](https://loinc.org/).

## AWS (Stack)
A concise for stack including
- API Gateway (HTTP v2)
- Lambda (go 1.x)
- DynamoDB (Hash & Range)

## Serverless.io (sls)
- Creates AWS resources defined in `serverless.yml`
- Generates AWS Cloud Formation Scripts
- Deploys AWS Lambda Functions  
- Tears down development resources

### How To
- [DL GO](https://go.dev/dl/)
- Locate project in shell
- Execute `make test` (will open browser with coverage results)
- Execute `make clean` (tidy the dir)
- Execute `make deploy` (AWS resources available)
- Enjoy `./.serverless` contents
- Execute `make scrub` (removes aforementioned dir)
