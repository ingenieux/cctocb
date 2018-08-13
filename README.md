# README

Automates CodeCommit Changes (Push/Tag/Release) to CodeBuild

Features:
  * Handling Gitflow
  * Notifying Build State Changes to Slack

## Instalation

Requires binutils (for strip), Make, dep, serverless and upx.

```
$ make clean dep build
```

Then deploy:

```
$ serverless deploy -s prod -r us-east-1
Serverless: Packaging service...
Serverless: Excluding development dependencies...
Serverless: Creating Stack...
Serverless: Checking Stack create progress...
.....
Serverless: Stack create finished...
Serverless: Uploading CloudFormation file to S3...
Serverless: Uploading artifacts...
Serverless: Uploading service .zip file to S3 (2.77 MB)...
Serverless: Validating template...
Serverless: Updating Stack...
Serverless: Checking Stack update progress...
...............
Serverless: Stack update finished...
Service Information
service: cc2cb
stage: prod
region: us-east-1
stack: cc2cb-prod
api keys:
  None
endpoints:
  None
functions:
  cctocb_handler: cc2cb-prod-cctocb_handler
```

And wire up the event to the lambda, by hooking "Repository State Change" for CodeCommit to the Lambda. e.g.:

https://youtu.be/COIk-WRj76Y

