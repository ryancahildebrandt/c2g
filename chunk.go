// -*- coding: utf-8 -*-

// Created on Thu Oct  2 12:57:28 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"slices"
	"strings"

	"gonum.org/v1/gonum/floats"
)

// Keeps track of transitional probabilities between tokens
type Transitions map[string]map[string]float64

// Function used to break a string into tokens and their corresponding pos/constituency tags
type TransitionSplitFunction func(string) ([]string, []string)

// Splits string into tokens
func TokenSplit(tok Tokenizer) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		tokens := tok.tokenize(s)
		return tokens, tokens
	}
}

// Splits string into tokens and POS tags
func POSSplit(tag SyntacticTagger) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		return tag.POS(s)
	}
}

// Splits string into tokens and constituency tags
func ConstituencySplit(tag SyntacticTagger) TransitionSplitFunction {
	return func(s string) ([]string, []string) {
		return tag.Constituency(s)
	}
}

// Splits a sequence of tokens based on transitional probabilities between tokens or tags
// higher p for smaller chunks, lower for larger chunks
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

// Collects chunks from texts, ordered by decreasing frequency
func CollectChunks(t []Text) []string {
	var chunks = []string{}
	var counts = make(map[string]int)

	for i := range t {
		for j := range t[i].chunk {
			chunks = append(chunks, t[i].chunk[j])
			counts[t[i].chunk[j]]++
		}
	}

	slices.SortStableFunc(chunks, func(i, j string) int {
		switch {
		case len(strings.Split(i, " ")) == len(strings.Split(j, " ")) && counts[i] == counts[j]:
			return strings.Compare(i, j)
		case len(strings.Split(i, " ")) == len(strings.Split(j, " ")):
			return counts[j] - counts[i]
		default:
			return len(strings.Split(j, " ")) - len(strings.Split(i, " "))
		}
	})

	return chunks
}

// Counts bigram co-occurrences and converts to probabilities
// counts are normalized such that all probabilities sum to 1
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

	normalizeCounts := func(p map[string]float64) map[string]float64 {
		var (
			out  map[string]float64 = make(map[string]float64)
			keys []string
			vals []float64
			norm []float64 = make([]float64, len(p))
		)

		for k, v := range p {
			keys = append(keys, k)
			vals = append(vals, v)
		}

		floats.DivTo(norm, vals, slices.Repeat([]float64{floats.Sum(vals)}, len(p)))

		for i := range len(p) {
			out[keys[i]] = norm[i]
		}

		return out
	}

	var bigrams [][]string
	tra := make(Transitions)

	for i := range t {
		tags, _ := f(t[i].text)
		bigrams = append(bigrams, toBigrams(tags)...)
	}

	for i := range bigrams {
		_, ok := tra[bigrams[i][0]]
		if !ok {
			tra[bigrams[i][0]] = make(map[string]float64)
		}
		tra[bigrams[i][0]][bigrams[i][1]]++
	}

	for k, v := range tra {
		tra[k] = normalizeCounts(v)
	}

	return tra
}
