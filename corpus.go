// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"github.com/bzick/tokenizer"
)

type Sentence struct {
	text     Expression
	segments []Ngram
}

type Corpus struct {
	texts       []Sentence
	transitions Transitions
	ngrams      []Ngram
}

func NewCorpus(t []Expression) Corpus {
	var c Corpus
	for _, tt := range t {
		c.texts = append(c.texts, Sentence{text: tt, segments: []Ngram{}})
	}
	return c
}

// Tokenizes sentences from corpus c using tokenizer t
func corpusToTokens(c Corpus, t *tokenizer.Tokenizer) [][]Expression {
	var out [][]Expression

	for _, cc := range c.texts {
		tt := UnigramTokenize(cc.text, t)
		if len(tt) != 0 {
			out = append(out, tt)
		}
	}

	return out
}
