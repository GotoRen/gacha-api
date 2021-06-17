import datetime
import time
import boto3

logGroupName = '/ecs/cta-staging'
logStreamNamePrefix = 'ecs/api'
client = boto3.client('logs', aws_access_key_id="[AWS_ACCESS_KEY_ID]", aws_secret_access_key="[AWS_SECRET_ACCESS_KEY]", region_name="ap-northeast-1")

def get_from_ts():
    from_ts = time.mktime(datetime.date.today().timetuple())
    return int(from_ts)

    #today = datetime.date.today()
    #yesterday = datetime.datetime.combine(today - datetime.timedelta(days = 1), datetime.time(0, 0, 0))
    #timestamp = time.mktime(yesterday.timetuple())
    #return int(timestamp)
    
def get_to_ts(from_ts):
    to_ts = int(from_ts) + (60 * 60 * 24) - 1
    return int(to_ts)

    #to_ts = from_ts + (60 * 60 * 24) - 1
    #return to_ts



if __name__ == '__main__':
    from_ts = get_from_ts()
    print(from_ts)
    to_ts = get_to_ts(from_ts)
    print(to_ts)

    
    response = client.filter_log_events(
        logGroupName=logGroupName,
        logStreamNamePrefix=logStreamNamePrefix,
        startTime=from_ts,
        endTime=to_ts,
    )
    print(response)




    # response = client.describe_log_streams(
    #     logGroupName=logGroupName,
    #     logStreamNamePrefix=logStreamNamePrefix,
    #     orderBy='LogStreamName',
    #     descending=True
    # )

    # print(response)
