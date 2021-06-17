import datetime
import pytz
from datetime import date
import time

def get_jst_today():
    ### 現在の時刻 (JST)
    jst_now = datetime.datetime.now(pytz.timezone('Asia/Tokyo'))
    ### 今日の日付 (JST)
    jst_today = datetime.datetime.combine(jst_now, datetime.time(0, 0, 0))
    return jst_today
    
def get_from_timestamp(jst_today):
    ### 昨日 AM 0:00 ~
    jst_yesterday = datetime.datetime.combine(jst_today - datetime.timedelta(days = 1), datetime.time(0, 0, 0))
    from_ts = int(time.mktime(jst_yesterday.timetuple()))
    return from_ts
    
def get_to_timestamp(from_ts):
    ### 昨日 ~ PM 23:59
    to_ts = from_ts + (60 * 60 * 24) - 1
    return to_ts
    


if __name__ == "__main__":
    jst_today = get_jst_today()
    from_ts = get_from_timestamp(jst_today)
    to_ts = get_to_timestamp(from_ts)

    print(jst_today)
    print(from_ts)
    print(to_ts)

    print(datetime.datetime.fromtimestamp(from_ts))
    print(datetime.datetime.fromtimestamp(to_ts))
