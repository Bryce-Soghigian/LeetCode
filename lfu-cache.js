/**
 * @param {number} capacity
 */
var LFUCache = function(capacity) {
    this.capacity=capacity;/**/
    this.map=new Map();/*key:index*/
    
    this.lru=new Array();/*stack to maintain LRU*/
    this.ar=new Array();/*2d array [ [key,freq,val],[key,freq,val],...]*/
    
    /* Method to Maintain Min Heap Property*/
    this.minHeapify=function(i){
        if(i>=this.ar.length)
            return;
        let temp=i;
        
        /* if it has left child*/
        if((2*i)+1<this.ar.length){
            /* Then if left child frequency less than root*/
            if(this.ar[(2*i)+1][1]<this.ar[temp][1])
                temp=(2*i)+1;
            else if(this.ar[(2*i)+1][1]==this.ar[temp][1]){/*Else if left child frequency == root*/
                /* Then If left child is less recently used than root*/
                if(this.lru.indexOf(this.ar[(2*i)+1][0])<this.lru.indexOf(this.ar[temp][0]))
                    temp=(2*i)+1;
            }
        }
        /* if it has right child*/
        if((2*i)+2<this.ar.length){
            /* Then if right child frequency less than root*/
            if(this.ar[(2*i)+2][1]<this.ar[temp][1])
                temp=(2*i)+2;
            else if(this.ar[(2*i)+2][1]==this.ar[temp][1]){/*Else if right child frequency == root*/
                /* Then If right child is less recently used than root*/
                if(this.lru.indexOf(this.ar[(2*i)+2][0])<this.lru.indexOf(this.ar[temp][0]))
                    temp=(2*i)+2;
            }
        }
        
        if(temp!=i){            
            /*Swap and heapify*/
            let t=this.ar[i];
            this.ar[i]=this.ar[temp];
            this.ar[temp]=t;
            
            /*Store indexes of array/heap in map*/
            this.map.set(this.ar[i][0],i);
            this.map.set(this.ar[temp][0],temp);
            
            /*Heapify to maintain min-heap property*/
            this.minHeapify(temp);
        }
    }
    
    /*Method to insert new key or update previously inserted key*/
    this.insertHeap =function(key,val){
        let ind=-1;/*Index of key in array(heap) */
        
        /*get index from key in our map*/
        if(this.map.has(key))
            ind=this.map.get(key);
        
        /*If we have this key in our map*/
        if(ind!==-1){
            /*Updation*/
            this.ar[ind][1]+=1;
            this.ar[ind][2]=val;
            /*Heapify*/
            this.minHeapify(ind);
            return;
        }
        else{
            /*Insertion */
             
             if(this.ar.length>=this.capacity){                 
                 /*Capacity Full*/
                 let old=this.extractMinHeap();/*Extract LFU*/
                 
                 if(old==-1)/*No Space in Heap*/
                     return;
            }
            
            /*Push item to rear(last) of heap*/
            this.ar.push([key,1,val]);
            /*Set the index in map*/
            this.map.set(key,this.ar.length-1);
            
            /*Travel bottom up to maintain heap property*/
            let i=this.ar.length-1;            
            while((Math.floor((i-1)/2))>=0
                  &&(this.ar[Math.floor((i-1)/2)][1])>(this.ar[i][1] )){
                this.minHeapify((Math.floor((i-1)/2)));/*Heapify Parent*/
                i=(Math.floor((i-1)/2));
            }   
        }        
            
    }
    
    /*Method to extract LFU/LRU item from heap*/
    this.extractMinHeap = function(){
        if(this.ar.length<=0)/*Heap Empty*/
            return -1;
        let res=this.ar[0];/*store LFU item*/
        this.ar[0]=(this.ar[this.ar.length-1]);/*replace with last item of heap*/
        this.ar.pop();/*remove last item of heap*/
        this.minHeapify(0);/*heapify root */        
        this.lru.splice(this.lru.indexOf(res[0]),1);/*remove item from lru stack*/
        this.map.delete(res[0]);/*remove item from index map*/
        return res;/*return item*/
        
    }
};

/** 
 * @param {number} key
 * @return {number}
 */
LFUCache.prototype.get = function(key) {
    /*Update LRU Stack*/
    let ind=this.lru.indexOf(key);
    if(ind!==-1){
        this.lru.splice(ind,1);/*delete prev position*/
        this.lru.push(key);/*Move to top*/
    }
    else this.lru.push(key);/*Move to top*/
    
    /*If the key is in map*/
    if(this.map.has(key)){
        let item=this.ar[this.map.get(key)];/*Get item*/
        this.insertHeap(key,item[2]);/*Update heap (frequency update)*/
        return item[2];/*Return result*/
    }     
    return -1;
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LFUCache.prototype.put = function(key, value) {
    /*Update LRU Stack*/
    let ind=this.lru.indexOf(key);
    if(ind!==-1){
        this.lru.splice(ind,1);/*delete prev position*/
        this.lru.push(key);/*Move to top*/
    }
    else this.lru.push(key);/*Move to top*/
    
    /*Insertion*/
    this.insertHeap(key,value);
};

/** 
 * Your LFUCache object will be instantiated and called as such:
 * var obj = new LFUCache(capacity)
 * var param_1 = obj.get(key)
 * obj.put(key,value)
 */