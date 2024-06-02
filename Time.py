import datetime
import time

# Get the current date and time
now = datetime.datetime.now()

# Check if today is Sunday or Monday
if now.weekday() in [6, 0]:  # 6 for Sunday, 0 for Monday
    print("BSE closed no data found !")
    exit()

# Check if the current time is 00:00:00
if now.hour == 0 and now.minute == 0 and now.second == 0:
    print("The current time is 00:00:00.")
else:
    # Calculate the time difference from the current time to 00:00:00
    time_diff = datetime.datetime(now.year, now.month, now.day, 0, 0, 0) - now

    # If the current time is after 00:00:00, the time difference will be negative
    # Use abs() to get the absolute value of the time difference
    sleep_duration = abs(time_diff.total_seconds())

    # Wait until 00:00:00
    time.sleep(sleep_duration)
print("Proceed with data fetching")
