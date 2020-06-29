package tree

type Trie struct {
	Val  *rune
	Next []*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	w := []rune(word)
	if len(w) <= 0 {
		return
	}
	isExistChar := false
	for _, t := range this.Next {
		if *t.Val == w[0] {
			//处理
			if len(w) > 1 {
				addChar(w[1:], t)
			}
			isExistChar = true
			break
		}
	}
	if !isExistChar {
		prefix := new(Trie)
		prefix.Val = &w[0]
		if len(w) > 1 {
			addChar(w[1:], prefix)
		}
		this.Next = append(this.Next, prefix)
	}
}

func addChar(w []rune, trie *Trie) {
	isExistChar := false
	for _, t := range trie.Next {
		if *t.Val == w[0] {
			//处理
			if len(w) > 1 {
				addChar(w[1:], t)
			}
			isExistChar = true
			break
		}
	}
	if !isExistChar {
		prefix := new(Trie)
		prefix.Val = &w[0]
		if len(w) > 1 {
			addChar(w[1:], prefix)
		}
		trie.Next = append(trie.Next, prefix)
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	w := []rune(word)
	if len(w) <= 0 {
		return false
	}
	isExistChar := false
	for _, t := range this.Next {
		if *t.Val == w[0] {
			if len(w) > 1 {
				isExistChar = isExistCharFunc(w[1:], t)
			} else {
				isExistChar = true
			}
		}
	}

	return isExistChar
}

func isExistCharFunc(w []rune, trie *Trie) bool {
	isExistChar := false
	for _, t := range trie.Next {
		if *t.Val == w[0] {
			if len(w) > 1 {
				isExistChar = isExistCharFunc(w[1:], t)
			} else {
				isExistChar = true
			}
		}
	}
	return isExistChar
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	w := []rune(prefix)
	if len(w) <= 0 {
		return false
	}
	isExistChar := false
	for _, t := range this.Next {
		if *t.Val == w[0] {
			if len(w) > 1 {
				isExistChar = isExistPrefixFunc(w[1:], t)
			} else if len(w) == 1 && t.Next != nil {
				isExistChar = true
			} else {
				isExistChar = false
			}
		}
	}

	return isExistChar
}

func isExistPrefixFunc(w []rune, trie *Trie) bool {
	isExistChar := false
	for _, t := range trie.Next {
		if *t.Val == w[0] {
			if len(w) > 1 {
				isExistChar = isExistCharFunc(w[1:], t)
			} else if len(w) == 1 && t.Next != nil {
				isExistChar = true
			} else {
				isExistChar = false
			}
		}
	}
	return isExistChar
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
