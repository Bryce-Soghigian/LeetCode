var spiralMatrixIII = function(R, C, r0, c0) {
    let offset = 1, cLimit = c0+offset, rLimit = r0+offset
    let i = r0, j = c0
    let res = []
    
    while(res.length < R*C){
		//if i and j are within the boundaries of the matix and to res
        if(i>=0 && j>=0 && i<R && j<C){  
            res.push([i,j])
        }
        // if i and j are both equal to their respective limits upate offset to next spiral length and direction
        if(i == rLimit && cLimit == j){ 
            offset = offset<0 ? offset-1 : offset+1
            offset *= -1
            cLimit = cLimit + offset
            rLimit = rLimit + offset
        }
        // each of the four possible movements through the matrix
        if(j<cLimit) j++
        else if(i<rLimit) i++
        else if(j>cLimit) j--
        else if(i>rLimit) i--
    }
    return res
};