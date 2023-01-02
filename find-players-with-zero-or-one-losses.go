func findWinners(matches [][]int) [][]int {

   distinctPlayers := make(map[int]bool, 1)
   losses := make(map[int]int, 1) // player -> loss count
   // potential mathematical solution if we can find the nnumber of games 
   // people played in.

    for i := range matches {
        winner, loser := matches[i][0], matches[i][1]
        losses[loser] += 1
        distinctPlayers[winner] = true
        distinctPlayers[loser] = true
    }
    undefeated, singleLoser := []int{}, []int{}
    for player, _ := range distinctPlayers {
        lossCount, ok := losses[player]
        if !ok {
            undefeated = append(undefeated, player)
        }
        if lossCount == 1 {
            singleLoser = append(singleLoser, player)
        }
    }
    sort.Ints(undefeated)
    sort.Ints(singleLoser)
    return [][]int{
        undefeated, 
        singleLoser,
    }
}