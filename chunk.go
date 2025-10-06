// -*- coding: utf-8 -*-

// Created on Thu Oct  2 12:57:28 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"slices"
	"strings"

	"gonum.org/v1/gonum/floats"
)

type Transitions map[string]map[string]float64

type TransitionSplitFunction func(string) ([]string, []string)

func TokenSplit(tk Tokenizer) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		tokens := tk.tokenize(s)
		return tokens, tokens
	}
}

func POSSplit(tag SyntacticTagger) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		return tag.POS(s)
	}
}

func ConstituencySplit(tag SyntacticTagger) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		return tag.Constituency(s)
	}
}

// higher for smaller chunks, lower for larger chunks
func TransitionChunk(tok []string, tag []string, tra Transitions, p float64) []string {
	var (
		b     strings.Builder
		s     string
		out   = []string{}
		probs []float64
	)

	if len(tok) == 0 {
		return out
	}

	for i := range len(tag) - 1 {
		probs = append(probs, tra[tag[i]][tag[i+1]])
	}

	for i, pp := range probs {
		if pp < p {
			s = strings.TrimSpace(b.String())
			out = append(out, s)
			b.Reset()
		}
		b.WriteString(tok[i])
		b.WriteString(" ")
	}
	b.WriteString(tok[len(probs)])

	s = strings.TrimSpace(b.String())
	out = append(out, s)
	out = slices.DeleteFunc(out, func(i string) bool { return i == "" })

	return out
}

func CollectChunks(t []Text) []string {
	var n = []string{}
	var c = make(map[string]int)

	for _, text := range t {
		for _, ng := range text.chunk {
			n = append(n, ng)
			c[ng]++
		}
	}

	slices.SortStableFunc(n, func(a, b string) int {
		switch {
		case len(strings.Split(a, " ")) == len(strings.Split(b, " ")) && c[a] == c[b]:
			return strings.Compare(a, b)
		case len(strings.Split(a, " ")) == len(strings.Split(b, " ")):
			return c[b] - c[a]
		default:
			return len(strings.Split(b, " ")) - len(strings.Split(a, " "))
		}
	})

	return n
}

// Counts bigram co-occurrences and converts to probabilities
func CollectTransitions(t []Text, f TransitionSplitFunction) Transitions {
	toBigrams := func(e []string) [][]string {
		var b [][]string

		switch len(e) {
		case 0:
			return b
		case 1:
			return append(b, []string{e[0], ""})
		default:
			for i := 0; i < len(e)-1; i++ {
				b = append(b, e[i:i+2])
			}
			return b
		}
	}

	// Convert transition counts to transition probabilities such that all probabilities sum to 1
	normalizeCounts := func(p map[string]float64) map[string]float64 {
		var (
			out  map[string]float64 = make(map[string]float64)
			l    int                = len(p)
			ks   []string
			vs   []float64
			norm []float64 = make([]float64, l)
		)

		for k, v := range p {
			ks = append(ks, k)
			vs = append(vs, v)
		}

		floats.DivTo(norm, vs, slices.Repeat([]float64{floats.Sum(vs)}, l))

		for i := range l {
			out[ks[i]] = norm[i]
		}

		return out
	}

	var (
		b   [][]string
		bb  [][]string
		tra Transitions = make(Transitions)
	)

	for _, ss := range t {
		tags, _ := f(ss.text)
		bb = toBigrams(tags)
		if len(bb) != 0 {
			b = append(b, bb...)
		}
	}

	for _, bb := range b {
		_, ok := tra[bb[0]]
		if !ok {
			tra[bb[0]] = make(map[string]float64)
		}
		tra[bb[0]][bb[1]]++
	}

	for k, v := range tra {
		tra[k] = normalizeCounts(v)
	}

	return tra
}
