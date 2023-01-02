/**
 * @param {string} homepage
 */
var BrowserHistory = function(homepage) {
    this.history = [homepage];
    this.position = 0;
};

/** 
 * @param {string} url
 * @return {void}
 */
BrowserHistory.prototype.visit = function(url) {
    this.history.length = this.position + 1; // clear up all the forward history
    this.history.push(url);
    this.position = this.history.length - 1;
    
};

/** 
 * @param {number} steps
 * @return {string}
 */
BrowserHistory.prototype.back = function(steps) {
    let possibleSteps = this.position;
    if (steps > possibleSteps) {
        this.position = 0;
    } else {
        this.position = this.position - steps;
    }
    return this.history[this.position];
};

/** 
 * @param {number} steps
 * @return {string}
 */
BrowserHistory.prototype.forward = function(steps) {
    let possibleSteps = this.history.length - 1 - this.position;
    if (steps > possibleSteps) {
        this.position = this.history.length - 1;
    } else {
        this.position = this.position + steps;
    }
    return this.history[this.position];
};