var reorderLogFiles = function(logs) {
    const letterLogs = [];
    const digitLogs = [];
    
    for (const log of logs) {
        const index = log.indexOf(" ");
        const id = log.substring(0, index);
        const data = log.substring(index + 1);
    
        if (areOnlyLetters(data)) letterLogs.push([id, data]);
        else digitLogs.push(log);
    }
    
    letterLogs.sort((a, b) => a[1].localeCompare(b[1]) || a[0].localeCompare(b[0]));
    
    const combinedLogs = [];
    
    for (const letterLog of letterLogs) {
        const [id, data] = letterLog;
        
        combinedLogs.push(id + " " + data);
    }
    
    combinedLogs.push(...digitLogs);
    
    return combinedLogs;
    
    function areOnlyLetters(str) {
        const firstChar = str.charAt(0);
        const regExp = /[a-z]/i;
        
        return regExp.test(firstChar);
    }
};