type Counter struct {
    _freq map[rune]int
}

func NewCounter() *Counter {
    return &Counter{
        _freq: map[rune]int{},
    }
}

func (c *Counter) Len() int {
    return len(c._freq)
}

func (c *Counter) Count(word string) {
    output := map[rune]int{}
    for _ , char := range word {
        output[char] += 1
    }
    c._freq = output
}
func (c *Counter) Values() []int {
    output := []int{}
    for _, val := range c._freq {
        output = append(output, val)
    }
    return output
}

func (c *Counter) BalancedKeys(c2 *Counter) bool {
    if c.Len() != c2.Len() {
        return false
    }

    for key := range c._freq {
        if _, ok := c2._freq[key]; !ok {
            return false
        }
    }
    return true
}

func (c *Counter) BalancedValues(c2 *Counter) bool {
    v1, v2 := c.Values(), c2.Values()
    sort.Ints(v1)
    sort.Ints(v2)
    for i, val := range v1 {
        if val != v2[i] {
            return false
        }
    }
    return true
}

func closeStrings(word1 string, word2 string) bool {
        c1 := NewCounter()
        c2 := NewCounter()
        c1.Count(word1)
        c2.Count(word2)
        if c1.Len() != c2.Len() {
            return false
        }
        return c1.BalancedKeys(c2) && c1.BalancedValues(c2)
}