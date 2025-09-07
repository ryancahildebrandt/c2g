// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"slices"
	"strings"

	"github.com/bzick/tokenizer"
)

type Triplet = []Expression

type Sentence struct {
	text Expression
	pre  Expression
	root Expression
	suf  Expression
	seg  Triplet
}

type Corpus struct {
	texts       []Sentence
	transitions Transitions
	ngrams      []Ngram
}

func ReadTexts(s *bufio.Scanner) []Expression {
	var e []Expression

	for s.Scan() {
		// e = append(e, strings.TrimSpace(s.Text()))
		e = append(e, s.Text())
	}

	return e
}

func ToNgrams(s []Sentence, t *tokenizer.Tokenizer, tr Transitions) []Ngram {
	var n []Ngram
	var m = make(map[Expression]Ngram)

	for i, sent := range s {
		tokens := UnigramTokenize(sent.text, t)
		sent.text = strings.Join(tokens, " ")
		sent.seg = NgramTokenize(tokens, tr, 0.1)
		for _, ng := range sent.seg {
			_, ok := m[ng]
			if !ok {
				m[ng] = Ngram{ng, len(strings.Split(ng, " ")) - 1, 0}
			}
			m[ng] = Ngram{ng, m[ng].len, m[ng].count + 1}
		}
		s[i] = sent
	}
	for _, v := range m {
		n = append(n, v)
	}
	slices.SortFunc(n, func(a, b Ngram) int {
		switch a.len {
		case b.len:
			return b.count - a.count
		default:
			return b.len - a.len
		}
	})

	return n
}

func NewCorpus(t []Expression) Corpus {
	var c Corpus

	slices.Sort(t)
	t = slices.Compact(t)
	for _, tt := range t {
		c.texts = append(c.texts, Sentence{text: tt, pre: "", root: "", suf: ""})
	}

	return c
}

func SplitTriplets(s []Sentence, n []Ngram) []Sentence {
	var ss []Sentence

	for _, sent := range s {
		i := slices.IndexFunc(n, func(ng Ngram) bool {
			return slices.Contains(sent.seg, ng.text)
		})

		if i == -1 {
			sent.root = sent.text
			ss = append(ss, sent)
			continue
		}

		ngram := n[i]
		sent.root = ngram.text
		p, s, _ := strings.Cut(sent.text, ngram.text)
		sent.pre = strings.TrimSpace(p)
		sent.suf = strings.TrimSpace(s)
		ss = append(ss, sent)
	}

	return ss
}

func ToRules(s []Sentence) []Rule {
	var r []Rule
	var rule Rule

	for _, sent := range s {
		rule = NewRule(sent.root)
		rule.pre = []Expression{sent.pre}
		rule.suf = []Expression{sent.suf}
		if !rule.isEmpty() {
			r = append(r, rule)
		}
	}

	return r
}
