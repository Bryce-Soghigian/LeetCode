var similarRGB = function(color) {
    let cur = 0;
    let res = "#"
    for(let i=1; i<color.length; i+=2){
        cur = parseInt(color.slice(i, i+2),16);
        let time = Math.round(cur/17)*17;
        res+=(time.toString(16).padStart(2,"0"));
    }
    return res;
};