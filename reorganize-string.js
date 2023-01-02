var reorganizeString = function(s) {
    var result = Array(s.length).fill(undefined);
    var map = new Map();
    s.split("").forEach(n=>map.set(n, map.get(n)+1 || 1));  
    let word = [...map.entries()].sort((a,b) => {return b[1]-a[1]}).map(word => word[0]); 
    let amount = [...map.entries()].sort((a,b) => {return b[1]-a[1]}).map(word => word[1]);
    var max = Math.max(...amount);
    if(max>Math.ceil(s.length/2)) return ""
    let value = 0
    for(i=0;i<s.length;i+=2){
           amount[value]--;
           result[i]=word[value];
           if(!amount[value]) value++;
        }
    for(i=1;i<s.length;i+=2){
           amount[value]--;
           result[i]=word[value];
           if(!amount[value]) value++;
        }
    return result.join("")
};