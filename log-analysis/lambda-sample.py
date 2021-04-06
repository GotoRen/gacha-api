import boto3
import datetime
import time

# ロググループを検索する際のクエリ―
PREFIX = 'test-'
# ログを保管するs3バケット
S3_BUCKET = 'test_bucket'
# ログを保管するs3のディレクトリ
S3_DIR = 'logs'

def main(event,context):
    '''
    メインで呼び出される関数
    '''
    # boto3のクライアント
    client = boto3.client('logs')
    # ロググループのリストを取得
    log_groups = get_log_group_list(client)
    # ログ内容をs3に保管
    move_logs(client,log_groups)

def get_log_group_list(client):
    '''
    ロググループの情報のリストを取得
    '''
    should_continue = True
    next_token = None
    log_groups=[]
    # 一度で全て取り切れない場合もあるので、繰り返し取得する
    while should_continue:
        if next_token == None:
            # 初回のリクエストの場合
            response = client.describe_log_groups(
                logGroupNamePrefix=PREFIX,
                limit=50
            )
        else:
            # 二回目以降のリクエストの場合
            response = client.describe_log_groups(
                logGroupNamePrefix=PREFIX,
                limit=50,
                nextToken=next_token
            )
        # 取得した結果を、リストに追加
        for log in response['logGroups']:
            log_groups.append(log)
        # またリクエストを投げるべきかどうかを判定
        if 'nextToken' in response.keys():
            next_token = response['nextToken']
        else:
            should_continue = False
    return log_groups

def create_export_task(client,log_groups):
    '''
    ログ内容をs3に移動
    '''
    # 現在時刻を取得し、UNIX timeに変換
    time_now = datetime.datetime.now()
    unix_time_now = int(time_now.timestamp())
    # ロググループの数だけ繰り返し
    for log in log_groups:
        for x in range(20):
            try:
                response = client.create_export_task(
                    fromTime=0,
                    to=unix_time_now,
                    logGroupName=log['logGroupName'],
                    destination=S3_BUCKET,
                    destinationPrefix=S3_DIR
                )
            except:
                # 既にタスクがある場合は、ちょっと待ってから再実行
                time.sleep(20)
                continue
