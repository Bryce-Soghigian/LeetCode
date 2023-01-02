var FrontMiddleBackQueue = function() {
    this.arr = [];    
};

/** 
 * @param {number} val
 * @return {void}
 */
FrontMiddleBackQueue.prototype.pushFront = function(val) {
    this.arr.splice(0, 0,val);
    return null;
};

/** 
 * @param {number} val
 * @return {void}
 */
FrontMiddleBackQueue.prototype.pushMiddle = function(val) {
    this.arr.splice(Math.floor(this.arr.length/2), 0, val);
    return null;
};

/** 
 * @param {number} val
 * @return {void}
 */
FrontMiddleBackQueue.prototype.pushBack = function(val) {
    this.arr.splice(this.arr.length, 0, val);
    return null;
};

/**
 * @return {number}
 */
FrontMiddleBackQueue.prototype.popFront = function() {
    if(this.arr.length == 0) return -1;
    return this.arr.splice(0,1);
};

/**
 * @return {number}
 */
FrontMiddleBackQueue.prototype.popMiddle = function() {
     if(this.arr.length == 0) return -1;
     return this.arr.splice(Math.ceil(this.arr.length/2)-1, 1);    
};

/**
 * @return {number}
 */
FrontMiddleBackQueue.prototype.popBack = function() {
     if(this.arr.length == 0) return -1;
    return this.arr.splice(this.arr.length-1,1);
};