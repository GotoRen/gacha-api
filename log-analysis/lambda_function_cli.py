import boto3
import os
import datetime
import pytz
from datetime import date
import time

# CloudWatch Logs
target_loggroup_name = '/ecs/cta-staging'

# S3
dst_s3_bucket_name = 'cta-staging-api-logs'
dst_s3_prefix = 'export_by_lambda_cli'

def get_jst_today():
    # 現在の時刻 (JST)
    jst_now = datetime.datetime.now(pytz.timezone('Asia/Tokyo'))
    # 今日の日付 (JST)
    jst_today = datetime.datetime.combine(jst_now, datetime.time(0, 0, 0))
    return jst_today
    
def get_from_timestamp(jst_today):
    # 昨日 AM 0:00 ~
    jst_yesterday = datetime.datetime.combine(jst_today - datetime.timedelta(days = 1), datetime.time(0, 0, 0))
    from_ts = int(time.mktime(jst_yesterday.timetuple()))
    return from_ts
    
def get_to_timestamp(from_ts):
    # 昨日 ~ PM 23:59
    to_ts = from_ts + (60 * 60 * 24) - 1
    return to_ts

def lambda_handler():
   jst_today = get_jst_today()
   from_ts = get_from_timestamp(jst_today)
   to_ts = get_to_timestamp(from_ts)
   client = boto3.client('logs', aws_access_key_id="AKIA6D57QWWUOZHRQEWP", aws_secret_access_key="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", region_name="ap-northeast-1")
   
   response = client.create_export_task(
       logGroupName      = target_loggroup_name,
       fromTime          = from_ts * 1000,
       to                = to_ts * 1000,
       destination       = dst_s3_bucket_name,
       destinationPrefix = dst_s3_prefix
   )
   return response



if __name__ == "__main__":

   jst_today = get_jst_today()
   print(jst_today)
   from_ts = get_from_timestamp(jst_today)
   print(from_ts)
   to_ts = get_to_timestamp(from_ts)
   print(to_ts)

   print('Timestamp: from_ts %s, to_ts %s', from_ts, to_ts)
   print('Timestamp: from_ts %s, to_ts %s', datetime.datetime.fromtimestamp(from_ts), datetime.datetime.fromtimestamp(to_ts))

   print('出力結果：')
   print(lambda_handler())
   