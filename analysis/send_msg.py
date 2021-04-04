import boto3

logGroupNamePrefix = '[CWL LogGroup NamePrefix]'
client = boto3.client('logs', aws_access_key_id="[AWS_ACCESS_KEY_ID]", aws_secret_access_key="[AWS_SECRET_ACCESS_KEY]", region_name="ap-northeast-1")

### ロググループを取得
response = client.describe_log_groups(
    logGroupNamePrefix=logGroupNamePrefix
)

print(response)