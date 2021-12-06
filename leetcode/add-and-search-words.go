package main

import "fmt"

type WordDictionary struct {
	root trieNode
}

type trieNode struct {
	children map[rune]*trieNode
	data     rune
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: trieNode{
			children: map[rune]*trieNode{},
			data:     0,
			isEnd:    false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	p := &this.root
	for _, c := range word {
		if child, ok := p.children[c]; !ok {
			child = &trieNode{
				children: map[rune]*trieNode{},
				data:     c,
				isEnd:    false,
			}
			p.children[c] = child
			p = child
		} else {
			p = child
		}
	}
	p.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	p := &this.root
	w := []rune(word)
	return doSearch(p, w, 0)
}

func doSearch(p *trieNode, w []rune, i int) bool {
	if i >= len(w) {
		return p.isEnd
	}
	c := w[i]
	if c != '.' {
		if nextP, ok := p.children[c]; !ok {
			return false
		} else {
			return doSearch(nextP, w, i+1)
		}
	} else {
		for _, v := range p.children {
			if doSearch(v, w, i+1) {
				return true
			}
		}
		return false
	}
}

func main() {
	w := Constructor()
	w.AddWord("abc")
	w.AddWord("abd")
	w.AddWord("abe")
	w.AddWord("acc")
	w.AddWord("abcde")
	fmt.Println(w.Search("a.c"), w.Search(".."), w.Search("ab"), w.Search("abcd."))
}