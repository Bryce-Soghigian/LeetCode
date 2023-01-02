from collections import defaultdict
class UndergroundSystem:
    """
    {user_id: ["place", "time"]}
    {place,place2: [running_sum, num_of_times]}
    """
    user_state = {} 
    travel_state = defaultdict(lambda: [0, 0])
    def checkIn(self, id, station_name, t):
        """
        Add user state
        """
        self.user_state[id] = [station_name, t]

    
    def checkOut(self, id, station_name, t):
        """
        Add completed travel_state
        Remove user state
        """
        start_station, t0 = self.user_state[id]
        key = (start_station, station_name)
        time_delta = t-t0
        self.travel_state[key][0] += time_delta
        self.travel_state[key][1] += 1

        del self.user_state[id]
    def getAverageTime(self, start_station, end_station):
        """
        Retriving completed_travel state
        """
        return self.travel_state[(start_station, end_station)][0] / self.travel_state[(start_station, end_station)][1]





# Your UndergroundSystem object will be instantiated and called as such:
# obj = UndergroundSystem()
# obj.checkIn(id,stationName,t)
# obj.checkOut(id,stationName,t)
# param_3 = obj.getAverageTime(startStation,endStation)