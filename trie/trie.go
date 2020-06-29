package trie

//import (
//"fmt"
//"bytes"
//"os"
//"encoding/json"
//"github.com/bradleyjkemp/memviz"
//// "github.com/davecgh/go-spew/spew"
//)

//// main functions
//// TODO: accept command line arguments
//func main() {
//
//	// first argument is the name of the binary
//	args := os.Args[1:]
//
//
//	// arg1 should be the text to build a suffix tree from
//	text := args[0]
//	var tree *node = build_suffix_tree(text)
//	add_suffix_indices(tree, 0, len(text))
//
//	buf2 := &bytes.Buffer{}
//	memviz.Map(buf2, tree)
//	fmt.Println(buf2.String())
//
//	// arg2 should be the query to search (optional)
//	if len(args) > 1 {
//		query := args[1]
//		fmt.Println("query is ", query)
//		matches := tree.multi_search(query, text, 2)
//		match_json, _ := json.Marshal(matches)
//		fmt.Println("query matches - ", string(match_json))
//	}
//
//	// // var text string = "gattaca$"
//	// // var text string = "abcabxabcd$"
//	// // fmt.Println("start building tree")
//	// // var tree *node = build_suffix_tree(text)
//	// // fmt.Println("finished building tree")
//
//	// // fmt.Println("search tree")
//	// // fmt.Println(tree.search("abce", text, 3))
//	// // var tree *node = build_suffix_tree("gattaca")
//	// // var text string = "bananasbanananananananananabananas$"
//	// // var text string = "abcdefabxybcdmnabcdex"
//	// var text string = "mississippi@tississilli$"
//	// // var text string = "almasamolmaz"
//	// // var text string = "cdddcdc"
//	// // var text string = "dedododeeodoeodooedeeododooodoede"
//	// // var text string = "panamabananas"
//	// // var text string = "GAGACCTCGATCGGCCAACTCATCTGTGAAACGTCAGTCATTGTAAGACTGGACATTTAGGAAGTAAGCCTTTTTCTTATAGCCAATCCCGCTTTCAATTGAACGGCTAAACGAAGGTCGTTTGCCACTGATTAGCAATTGGTCCGGTGAAAAATTGTGTATTTTGGAAAGATGTAATCCTGCGAGACCTCGATCGGC"
//
//	// // fmt.Println("start building tree")
//	// var tree *node = build_suffix_tree(text)
//	// // fmt.Println("finished building tree")
//	// // spew.Dump(tree)
//
//	// // buf := &bytes.Buffer{}
//	// // memviz.Map(buf, tree)
//	// // fmt.Println(buf.String())
//
//	// add_suffix_indices(tree, 0, len(text))
//
//	// buf2 := &bytes.Buffer{}
//	// memviz.Map(buf2, tree)
//	// fmt.Println(buf2.String())
//
//	// // fmt.Println("search tree for get all possible indexes where the suffix matches")
//	// // var query string = "issp"
//	// // fmt.Println(query)
//	// // fmt.Println(query[0:len(query)])
//
//	// // matched_indices := tree.multi_search(query, text, 2)
//
//	// // fmt.Println("matches - ", matched_indices)
//}