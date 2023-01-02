import (
	"math/rand"
)

type RandomizedSet struct {
    state []int 
    location map[int]int // value -> index 
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        state: []int{},
        location: map[int]int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    //1. Check if its in the set. IF true return 
    if _, ok := this.location[val]; ok {
        return false
    }
    this.state = append(this.state, val)
    this.location[val] = len(this.state) - 1
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    index, okk := this.location[val] 
    if _, ok := this.location[val]; !ok || len(this.state) == 0 || !okk {
        return false
    } 

    temp := this.state[len(this.state)-1]
    this.state[index] = temp
    delete(this.location, temp)
    delete(this.location, val)

    this.state = this.state[:len(this.state)-1]
    if val != temp {
          this.location[temp] = index  
    }
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.state[rand.Intn(len(this.state))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */