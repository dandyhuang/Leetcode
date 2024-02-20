package main

import (
	"fmt"
	"sort"
)

func addWordsV2(node *Trie, word string) {
	for c, i := range word {
		index := int(i) - int('a')
		if node.next[index] == nil {
			node.next[index] = &Trie{}
		}
		node = node.next[index]
		fmt.Printf("%c the %s %v \n", word[c], word, node.words)
		node.words = append(node.words, word)
		sort.Strings(node.words)
		if len(node.words) > 3 {
			node.words = node.words[0 : len(node.words)-1]
		}
	}
}

// 1268. 搜索推荐系统
type Trie struct {
	child map[byte]*Trie
	words []string
	next  [26]*Trie
	reqId string
}

func addWords(root *Trie, word string) {
	cur := root
	for i, _ := range word {
		_, ok := cur.child[word[i]]
		if !ok {
			cur.child[word[i]] = &Trie{child: make(map[byte]*Trie)}
			// 之前这样定义，每次都被覆盖了
			// cur.child = make(map[byte]*Trie)
			// cur.child[word[i]] = &Trie{}
		}
		cur = cur.child[word[i]]
		//fmt.Printf(" %c the %s %v %p %v\n", word[i], word, cur.words, cur, cur.child)
		cur.words = append(cur.words, word)
		sort.Strings(cur.words)
		if len(cur.words) > 3 {
			cur.words = cur.words[:len(cur.words)-1]
		}
	}
}
func suggestedProducts(products []string, searchWord string) [][]string {
	res := make([][]string, 0)
	root := &Trie{child: make(map[byte]*Trie)}
	for i := range products {
		addWords(root, products[i])
	}

	flag := false
	cur := root
	for i, _ := range searchWord {
		_, ok := cur.child[searchWord[i]]
		if flag || !ok {
			res = append(res, make([]string, 0))
			flag = true
		} else {
			cur = cur.child[searchWord[i]]
			fmt.Println(cur.words)
			res = append(res, cur.words)
		}
	}
	return res
}
func main() {
	var root = &Trie{child: make(map[byte]*Trie)}
	var words = []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	root.reqId = "13242332"
	for _, word := range words {
		addWords(root, word)
	}
}
