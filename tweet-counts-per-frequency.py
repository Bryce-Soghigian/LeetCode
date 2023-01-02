from collections import defaultdict
from bisect import insort_left, bisect_left

class TweetCounts:

    def __init__(self):

        # dictionary:
        # key   : tweetname
        # value : list of timestamps
        self.tweet_timestamp = defaultdict( list )            

        
        
    def recordTweet(self, tweetName: str, time: int) -> None:
        
        # insert new timestamp for tweetName, and kepp in order
        insort_left( self.tweet_timestamp[tweetName], time )

        
        
    def getTweetCountsPerFrequency(self, freq: str, tweetName: str, startTime: int, endTime: int) -> List[int]:
        
        arr_timestamp = self.tweet_timestamp[tweetName]
        
        unit_size = None
        
        # compute unit size from given frequency
        if freq == 'minute':
            unit_size = 60
        elif freq == 'hour':
            unit_size = 60 * 60
        elif freq == 'day':
            unit_size = 60 * 60 * 24
        
        # compute total number of intervals
        num_of_interval = (endTime - startTime)//unit_size + 1
        
        # list for output
        result = [ 0 for i in range(num_of_interval) ]
        
        # find the index of first timestamp in query range
        index = bisect_left(arr_timestamp, startTime)
        
        # update statistics of tweet occurrences based on specified interval
        while index < len(arr_timestamp) and arr_timestamp[index] <= endTime:
            result[ (arr_timestamp[index]-startTime) // unit_size ] += 1
            index += 1
            
        return result


# Your TweetCounts object will be instantiated and called as such:
# obj = TweetCounts()
# obj.recordTweet(tweetName,time)
# param_2 = obj.getTweetCountsPerFrequency(freq,tweetName,startTime,endTime)