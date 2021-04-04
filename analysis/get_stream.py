import boto3

logGroupName = '[CWL LogGroup Name]'
logStreamNamePrefix = '[CWL LogStream NamePrefix]'

client = boto3.client('logs', aws_access_key_id="[AWS_ACCESS_KEY_ID]", aws_secret_access_key="[AWS_SECRET_ACCESS_KEY]", region_name="ap-northeast-1")

### ログストリームを取得
response = client.describe_log_streams(
    logGroupName=logGroupName,
    logStreamNamePrefix=logStreamNamePrefix,
    orderBy='LogStreamName',
    descending=True
)

print(response)