var minKnightMoves = function(x, y) {
    x = Math.abs(x)
    y = Math.abs(y)
   
    let dir = [
                [-1,2],[-2,1],
                [-2,-1],[-1,-2],
                [1,-2],[2,-1],
                [2,1],[1,2]
              ]
    
    let queue = [[0,0]]//initialize origin
    let level = 0
    let visited = new Set()
    
    while(queue.length != 0){
        let next = [] //used to populate next queue
        
        while(queue.length != 0){
            const cur = queue.pop()
			//If going below -1 direction in either x or y, ignore
            if((cur[0] < -1 || cur[1] < -1)) continue
            if(cur[0] === x && cur[1] === y) return level
            
            for(let d of dir){
                let nextX = cur[0] + d[0]
                let nextY = cur[1] + d[1]
                
				//Limit search to positive direction (PS if you know further restrictions to improve let me know in comments!!!)
                if(!visited.has(nextX+','+nextY) && (nextX >= -1 && nextY >= -1 && nextX <= x+2 && nextY <= y+2)){
                    visited.add(nextX+','+nextY)
                    next.push([nextX,nextY])
                }
            }
        }
        queue = next
        level++
    }
};