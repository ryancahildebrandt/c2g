// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:17:15 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"slices"
	"strings"

	"gonum.org/v1/gonum/floats"
	"github.com/bzick/tokenizer"
)

type TransitionProbabilities map[Expression]float64
type Transitions map[Expression]TransitionProbabilities

type Ngram struct {
	text  Expression
	len   int
	count int
}

// Counts bigram co-occurrences
func NewTransitions(c Corpus, t *tokenizer.Tokenizer) Transitions {
	var b [][]Expression
	var tt Transitions = make(Transitions)

	for _, cc := range c.texts {
		bb := ToBigrams(UnigramTokenize(cc.text, t))
		if len(bb) != 0 {
			b = append(b, bb...)
		}
	}

	for _, bb := range b {
		_, ok := tt[bb[0]]
		if !ok {
			tt[bb[0]] = make(TransitionProbabilities)
		}
		tt[bb[0]][bb[1]]++
	}

	return tt.normalize()
}

// Convert transition counts to transition probabilities such that all probabilities sum to 1
func probabilityNorm(p TransitionProbabilities) TransitionProbabilities {
	var (
		out  TransitionProbabilities = make(TransitionProbabilities)
		l    int                     = len(p)
		ks   []Expression
		vs   []float64
		norm []float64 = make([]float64, l)
	)

	for k, v := range p {
		ks = append(ks, k)
		vs = append(vs, v)
	}
	if floats.Sum(vs) == 0 {
		return out
	}

	floats.DivTo(norm, vs, slices.Repeat([]float64{floats.Sum(vs)}, l))

	for i := range l {
		out[ks[i]] = norm[i]
	}

	return out
}

// Applies probabilityNorm function to all transitions in t
func (t Transitions) normalize() Transitions {
	var out Transitions = make(Transitions)
	for k, v := range t {
		out[k] = probabilityNorm(v)
	}
	return out
}

func NgramTokenize(s []Expression, t Transitions, p float64) []Expression {
	var b strings.Builder
	var o string
	var out []Expression
	var probs []float64

	// higher for smaller ngrams, lower for larger ngrams
	// maybe 1/thr to provide high number for low probabilities?

	if len(s) == 0 {
		return out
	}

	for i := range len(s) - 1 {
		probs = append(probs, t[s[i]][s[i+1]])
	}

	for i, pp := range probs {
		if pp < p {
			o = strings.TrimSpace(b.String())
			if o != "" {
				out = append(out, o)
				b.Reset()
			}
		}
		b.WriteString(s[i])
		b.WriteString(" ")
	}
	b.WriteString(s[len(probs)])

	o = strings.TrimSpace(b.String())
	if o != "" {
		out = append(out, o)
	}

	return out
}
