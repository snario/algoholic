package algoholic

// A trie is a data structure which stores words in a tree by letter, when the letters form a
// valid string, the node is marked terminal and a value can be stored.

type Trie struct {
	Char     rune
	Children map[rune]*Trie
	Parent   *Trie
	Terminal bool
	Value    interface{}
}

// A value indicating that the node is the root node, thus it doesn't denote a character
// itself.
const RootTrieChar = rune(0)

// Create a new trie with the specified character and parent.
func NewTrie(parent *Trie, chr rune) *Trie {
	return &Trie{chr, parent, make(map[rune]*Trie), false, nil}
}

func NewRootTrie() *Trie {
	return NewTrie(nil, RootTrieChar)
}

// Create a new trie with strings mapped to specified values.
func NewTrieFromMap(strMap map[string]interface{}) *Trie {
	ret := NewRootTrie()

	for str, val := range strMap {
		ret.Insert(str, val)
	}

	return ret
}

// Create a new trie with strings whose values we don't care about.
func NewTrieFromStrings(strs []string) *Trie {
	ret := NewRootTrie()

	for _, str := range strs {
		ret.Insert(str, nil)
	}

	return ret
}

// Find the specified string and return its trie node.
// O(m) worst-case where m is the length of the string searched for.
// Note this returns non-terminal nodes.
func (trie *Trie) FindTrie(str string) *Trie {
	if len(str) == 0 {
		return trie
	}

	if next := trie.Children[rune(str[0])]; next != nil {
		return next.FindTrie(str[1:])
	}

	return nil
}

// Find the specified string and return its value.
// O(m) worst-case where m is the length of the string searched for.
func (trie *Trie) Find(str string) (val interface{}, has bool) {
	ret := trie.FindTrie(str)

	if ret == nil || !ret.Terminal {
		// Not found.
		return
	}

	has = true
	val = ret.Value

	return
}

// Find all valid strings that consist of suffixes of the input prefix.
// O(m) worst-case where m is the length of the longest returned string.
func (trie *Trie) FindSuffixes(prefix string) []string {
	trie = trie.FindTrie(prefix)

	if trie == nil {
		return nil
	}

	var ret []string

	for str, _ := range trie.ToMap() {
		ret = append(ret, str)
	}

	return ret
}

// Insert string, value pair into the specified trie.
// O(m) worst-case where m is the length of the inserted string.
func (trie *Trie) Insert(str string, val interface{}) {
	var (
		i   int
		chr rune
	)

	// Search through existing nodes.
	for i, chr = range str {
		if next, has := trie.Children[chr]; has {
			trie = next
		} else {
			break
		}
	}

	// Insert nodes as necessary.
	for _, chr = range str[i:] {
		next := NewTrie(trie, chr)
		trie.Children[chr] = next
		trie = next
	}
	trie.Terminal = true
	trie.Value = val
}

// The String() function for a trie is its characters from root to node.
func (trie *Trie) String() string {
	var chrs []rune

	for ; trie != nil && trie.Char != RootTrieChar; trie = trie.Parent {
		chr := trie.Char

		chrs = append([]rune{chr}, chrs...)
	}

	return string(chrs)
}

// Recursively walk through all children of the input trie, returning a map of string, value
// pairs.
// O(n) where n is the number of nodes in the input trie.
func (trie *Trie) ToMap() map[string]interface{} {
	ret := make(map[string]interface{})

	trie.Walk(func(node *Trie) {
		if node.Terminal {
			ret[node.String()] = node.Value
		}
	})

	return ret
}

// Recursively walk through all children of the input trie in preorder executing the specified
// function on each trie node.
//O(n) where n is the number of nodes in the input trie.
func (trie *Trie) Walk(fn func(*Trie)) {
	fn(trie)
	for _, child := range trie.Children {
		child.Walk(fn)
	}
}
