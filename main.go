package main

//go:generate ./words.sh

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var (
	seed = flag.Int64("seed", 1, "random seed")
	sep  = flag.String("sep", " ", "separator")
)

func main() {
	flag.Parse()
	r := rand.New(rand.NewSource(*seed))

	words := []Word{}
	for word, pos := range Words {
		words = append(words, Word{word, pos})
	}

	sort.Slice(words, func(i, j int) bool { return words[i].String() < words[j].String() })

	// shuffle
	for i := range words {
		j := r.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	pattern := []Pos{JJ, NN}
	result := make([]string, len(pattern))

	for _, word := range words {
		for i, pos := range pattern {
			if word.Pos == pos {
				result[i] = word.String()
			}
		}
	}

	fmt.Printf(strings.Join(result, *sep))
	fmt.Printf("\n")
}

type Word struct {
	Word string
	Pos  Pos
}

func (w Word) String() string {
	return w.Word
}

type Pos int

// https://www.ling.upenn.edu/courses/Fall_2003/ling001/penn_treebank_pos.html
const (
	CC Pos = iota
	CD
	DT
	EX
	FW
	IN
	JJ
	JJR
	JJS
	LS
	MD
	NN
	NNS
	NNP
	NNPS
	PDT
	POS
	PRP
	PR
	RB
	RBR
	RBS
	RP
	SYM
	TO
	UH
	VB
	VBD
	VBG
	VBN
	VBP
	VBZ
	WDT
	WP
	W
	WRB
)
