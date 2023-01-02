var ambiguousCoordinates = function(S) {
    const trimmed = S.substring(1, S.length - 1);
    const nums = [];
    
    for (let i = 0; i < trimmed.length - 1; i++) {
        const frontNum = trimmed.substring(0, i + 1);
        const backNum = trimmed.substring(i + 1);
        
        nums.push(frontNum, backNum);
    }
   
    const res = [];
    
    for (let i = 0; i < nums.length; i += 2) {
        const frontNum = nums[i];
        const backNum = nums[i + 1];
        
        const xCoords = createCoords(frontNum);
        const yCoords = createCoords(backNum);
        
        for (const xCoord of xCoords) { 
            for (const yCoord of yCoords) {
                const xyCoords = `(${xCoord}, ${yCoord})`;
                res.push(xyCoords);
            }
        }
    }
    
    return res; 
    
    function createCoords(coord) {
        const coords = [];
        
        for (let i = 0; i < coord.length; i++) {
            const beforeDec = coord.substring(0, i + 1);
            const afterDec = coord.substring(i + 1);
           
            if ((!beforeDec.startsWith("0") || beforeDec == "0") && !afterDec.endsWith("0")) coords.push(parseFloat(`${beforeDec}.${afterDec}`));
        }
        
        return coords;
    }
};