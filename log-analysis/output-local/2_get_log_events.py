import boto3
import json

logGroupName = '/ecs/cta-staging'
logStreamName = 'ecs/api/82943fd92872427dbb0930196f4593ca'
client = boto3.client('logs', aws_access_key_id="[AWS_ACCESS_KEY_ID]", aws_secret_access_key="[AWS_SECRET_ACCESS_KEY]", region_name="ap-northeast-1")

### ログを取得
response = client.get_log_events(
    logGroupName=logGroupName,
    logStreamName=logStreamName,
    startFromHead=True
)
print(response)
