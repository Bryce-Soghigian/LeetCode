/**
 * @param {number[][]} rooms
 * @return {boolean}
 */
var canVisitAllRooms = function(rooms) {
    /*
    
    We build a graph. Then we walk through and see if we have the keys to access rooms 0-n
    **/
    const uniqueKeys = new Set([0]);
    const availableKeys = [0];
    
    while (availableKeys.length) {
        let currentKey = availableKeys.pop();
        rooms[currentKey].forEach(key => {
           if (!uniqueKeys.has(key)) {
               uniqueKeys.add(key);
               availableKeys.push(key);
           }
        });
    }

    return uniqueKeys.size === rooms.length;
};