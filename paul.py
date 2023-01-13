def find_codes(graph, node, code, codes, memo):
    if len(code) == 9:
        codes.append(code)
        return
    if (node, code) in memo:
        return
    memo.add((node, code))
    for neighbor in graph[node]:
        if neighbor not in code:
            find_codes(graph, neighbor, code + neighbor, codes, memo)
graph = {}
graph['a'] = ['b', 'd', 'e']
graph['b'] = ['a', 'c', 'e']
graph['c'] = ['b', 'e', 'f']
graph['d'] = ['a', 'e', 'g']
graph['e'] = ['a', 'b', 'c', 'd', 'f', 'g', 'h', 'i']
graph['f'] = ['c', 'e', 'i']
graph['g'] = ['d', 'e', 'h']
graph['h'] = ['g', 'e', 'i']
graph['i'] = ['h', 'e', 'f']

def find_all_codes(graph):
    codes = []
    for node in graph:
        # Goroutines here potentially to speed up the process
        find_codes(graph, node, node, codes, set())
    return codes

expectedCodes = ['abcedghif', 'abcefihgd', 'abcfedghi', 'abcfeihgd', 'abcfihgde', 'abcfihged', 'abcfihedg', 'abcfihegd', 'abcfiedgh', 'abcfiehgd', 'abecfihgd', 'abedghifc', 'adebcfihg', 'adeghifcb', 'adgebcfih', 'adgehifcb', 'adghebcfi', 'adgheifcb', 'adghiebcf', 'adghiefcb', 'adghifcbe', 'adghifceb', 'adghifebc', 'adghifecb', 'aebcfihgd', 'aedghifcb', 'badecfihg', 'badeghifc', 'badgecfih', 'badgehifc', 'badghecfi', 'badgheifc', 'badghiecf', 'badghiefc', 'badghifce', 'badghifec', 'baecfihgd', 'baedghifc', 'bceadghif', 'bcefihgda', 'bcfeadghi', 'bcfeihgda', 'bcfihgdae', 'bcfihgdea', 'bcfihgead', 'bcfihgeda', 'bcfiheadg', 'bcfihegda', 'bcfieadgh', 'bcfiehgda', 'beadghifc', 'becfihgda', 'cbadefihg', 'cbadeghif', 'cbadgefih', 'cbadgehif', 'cbadghefi', 'cbadgheif', 'cbadghief', 'cbadghife', 'cbaedghif', 'cbaefihgd', 'cbeadghif', 'cbefihgda', 'cebadghif', 'cefihgdab', 'cfebadghi', 'cfeihgdab', 'cfihgdabe', 'cfihgdaeb', 'cfihgdeab', 'cfihgdeba', 'cfihgebad', 'cfihgedab', 'cfihebadg', 'cfihegdab', 'cfiebadgh', 'cfiehgdab', 'dabcefihg', 'dabceghif', 'dabcfeghi', 'dabcfeihg', 'dabcfihge', 'dabcfiheg', 'dabcfiegh', 'dabcfiehg', 'dabecfihg', 'dabeghifc', 'daebcfihg', 'daeghifcb', 'deabcfihg', 'deghifcba', 'dgeabcfih', 'dgehifcba', 'dgheabcfi', 'dgheifcba', 'dghieabcf', 'dghiefcba', 'dghifcbae', 'dghifcbea', 'dghifceab', 'dghifceba', 'dghifeabc', 'dghifecba', 'eabcfihgd', 'eadghifcb', 'ebadghifc', 'ebcfihgda', 'ecbadghif', 'ecfihgdab', 'edabcfihg', 'edghifcba', 'efcbadghi', 'efihgdabc', 'egdabcfih', 'eghifcbad', 'ehgdabcfi', 'ehifcbadg', 'eihgdabcf', 'eifcbadgh', 'fcbadeghi', 'fcbadeihg', 'fcbadgehi', 'fcbadgeih', 'fcbadghei', 'fcbadghie', 'fcbaedghi', 'fcbaeihgd', 'fcbeadghi', 'fcbeihgda', 'fcebadghi', 'fceihgdab', 'fecbadghi', 'feihgdabc', 'fihgdabce', 'fihgdabec', 'fihgdaebc', 'fihgdaecb', 'fihgdeabc', 'fihgdecba', 'fihgecbad', 'fihgedabc', 'fihecbadg', 'fihegdabc', 'fiecbadgh', 'fiehgdabc', 'gdabcefih', 'gdabcehif', 'gdabcfehi', 'gdabcfeih', 'gdabcfihe', 'gdabcfieh', 'gdabecfih', 'gdabehifc', 'gdaebcfih', 'gdaehifcb', 'gdeabcfih', 'gdehifcba', 'gedabcfih', 'gehifcbad', 'ghedabcfi', 'gheifcbad', 'ghiedabcf', 'ghiefcbad', 'ghifcbade', 'ghifcbaed', 'ghifcbead', 'ghifcbeda', 'ghifcebad', 'ghifcedab', 'ghifecbad', 'ghifedabc', 'hgdabcefi', 'hgdabceif', 'hgdabcfei', 'hgdabcfie', 'hgdabecfi', 'hgdabeifc', 'hgdaebcfi', 'hgdaeifcb', 'hgdeabcfi', 'hgdeifcba', 'hgedabcfi', 'hgeifcbad', 'hegdabcfi', 'heifcbadg', 'hiefcbadg', 'hiegdabcf', 'hifcbadeg', 'hifcbadge', 'hifcbaedg', 'hifcbaegd', 'hifcbeadg', 'hifcbegda', 'hifcebadg', 'hifcegdab', 'hifecbadg', 'hifegdabc', 'ihgdabcef', 'ihgdabcfe', 'ihgdabecf', 'ihgdabefc', 'ihgdaebcf', 'ihgdaefcb', 'ihgdeabcf', 'ihgdefcba', 'ihgedabcf', 'ihgefcbad', 'ihefcbadg', 'ihegdabcf', 'iefcbadgh', 'iehgdabcf', 'ifcbadegh', 'ifcbadehg', 'ifcbadgeh', 'ifcbadghe', 'ifcbaedgh', 'ifcbaehgd', 'ifcbeadgh', 'ifcbehgda', 'ifcebadgh', 'ifcehgdab', 'ifecbadgh', 'ifehgdabc']
assert find_all_codes(graph) == expectedCodes
print(find_all_codes(graph))
