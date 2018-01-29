package main

//go:generate ./words.sh

import (
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var (
	seed   = flag.Int64("seed", 1, "random seed")
	sep    = flag.String("sep", " ", "separator")
	length = flag.Int("length", 3, "name length")
)

func main() {
	flag.Parse()
	data := flag.Arg(0)

	// shuffle from a known order
	r := rand.New(rand.NewSource(*seed))
	sort.Slice(words, func(i, j int) bool { return words[i].String() < words[j].String() })
	for i := range words {
		j := r.Intn(i + 1)
		words[i], words[j] = words[j], words[i]
	}

	h := sha256.New()
	h.Write([]byte(data))
	d := binary.BigEndian.Uint32(h.Sum(nil)) //% uint64(math.Pow(2, 32))

	ns := Nouns()
	js := Adjectives()
	vs := Adverbs()
	n := uint32(len(ns))
	j := uint32(len(js))
	v := uint32(len(vs))

	result := []string{}

	noun := ns[d%n]
	adj := js[d%j]

	switch *length {
	case 1:
		result = append(result, noun)
	case 2:
		result = append(result, adj, noun)
	default:
		w := uint32(*length - 2)
		for i := uint32(1); i <= w; i++ {
			result = append(result, vs[(d+i)%v])
		}
		result = append(result, adj, noun)
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

func someWords(mask Pos) []string {
	matches := []string{}
	for _, word := range words {
		if (mask & word.Pos) != 0 {
			matches = append(matches, word.String())
		}
	}
	return matches
}

func Nouns() []string {
	return someWords(NN | NNS | NNP | NNPS)
}

func Adjectives() []string {
	return someWords(JJ | JJR | JJS | VBG | VBN | VBD)
}

func Adverbs() []string {
	return someWords(RB | RBR | RBS)
}

type Pos uint64

// https://www.ling.upenn.edu/courses/Fall_2003/ling001/penn_treebank_pos.html
const (
	CC Pos = 1 << iota
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
