import boto3
import json
from datetime import datetime
import os

logGroupName = '[CWL LogGroup Name]'
logStreamName = '[CWL LogStream Name]'
client = boto3.client('logs', aws_access_key_id="[AWS_ACCESS_KEY_ID]", aws_secret_access_key="[AWS_SECRET_ACCESS_KEY]", region_name="ap-northeast-1")

### ログストリーム情報を取得
response = client.describe_log_streams(
    logGroupName=logGroupName,
    logStreamNamePrefix=logStreamName,
    orderBy='LogStreamName',
    descending=True
)
print(response)

### 全てのストリームログを取得
for stream in response['logStreams']:

    stream_name = stream['logStreamName']
    print('StreamName: ', stream_name)

    creationTime = (datetime.fromtimestamp(int(str(stream['creationTime'])[:10]))).strftime('%Y-%m-%d %H:%M:%S')
    print('creationTime: ', creationTime)

    firstEventTimestamp = (datetime.fromtimestamp(int(str(stream['firstEventTimestamp'])[:10]))).strftime('%Y-%m-%d %H:%M:%S')
    print('firstEventTimestamp: ', firstEventTimestamp)

    lastEventTimestamp = (datetime.fromtimestamp(int(str(stream['lastEventTimestamp'])[:10]))).strftime('%Y-%m-%d %H:%M:%S')
    print('lastEventTimestamp: ', lastEventTimestamp)

    lastIngestionTime = (datetime.fromtimestamp(int(str(stream['lastIngestionTime'])[:10]))).strftime('%Y-%m-%d %H:%M:%S')
    print('lastIngestionTime: ', lastIngestionTime)


    ### 出力ディレクトリ
    dir_path = 'logs/'
    
    ### ファイル名：[ストリーム作成日時]_[ストリーム名].json（ '/' の部分は '-' に置き換える）
    file_name = '{}_{}.json'.format(creationTime, stream_name).replace('/', '-')
    
    # ログを取得
    logs = client.get_log_events(
        logGroupName=logGroupName,
        logStreamName=logStreamName,
        startFromHead=True
    )

    ### 出力
    body = logs['events']
    with open(os.path.join(dir_path, file_name), 'w') as f:
        for line in body:
            message = line['message'] + '\n'
            f.write(message)


