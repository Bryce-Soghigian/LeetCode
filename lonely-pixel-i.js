/**
 * @param {character[][]} picture
 * @return {number}
 */
var findLonelyPixel = function(picture) {
    
//We want to preform a dfs for each row and column form the 
    
let numberOfLonelyPixels = 0;

for(let i = 0;i<picture.length;i++){
    for(let j = 0;j<picture[i].length;j++){
        if(picture[i][j] === "B"){
            picture[i][j]= "TEMP"
            //Check left and right
            let isLonely = true
            for(let top = 0;top < picture[i].length;top++){
                if(picture[i][top] === "B"){
                    isLonely = false
                }
            }
            //Loop through column
            for(let col = 0;col<picture.length;col++){
                if(picture[col][j]==="B"){
                    isLonely = false
                }
            }
            
            
            
            if(isLonely === true){
                numberOfLonelyPixels++
            }
            picture[i][j] = "B"
        }
    }
}

return numberOfLonelyPixels
};