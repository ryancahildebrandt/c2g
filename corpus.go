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

type Sentence struct {
	text Expression
	pre  Expression
	root Expression
	suf  Expression
	seg  []Expression
}

type Corpus struct {
	texts       []Sentence
	transitions Transitions
	ngrams      []Ngram
}

func ReadTexts(s *bufio.Scanner) []Expression {
	var e []Expression

	for s.Scan() {
		e = append(e, strings.TrimSpace(s.Text()))
	}

	return e
}

func ToNgrams(s []Sentence, t *tokenizer.Tokenizer, tr Transitions) []Ngram {
	var n []Ngram
	var m = make(map[Expression]Ngram)

	for i, sent := range s {
		tokens := UnigramTokenize(sent.text, t)
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
			sent.pre = ""
			sent.suf = ""
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

func ToTripletMap(s []Sentence) TripletMap {
	var t = make(TripletMap)
	for _, sent := range s {
		_, ok := t[sent.pre]
		if !ok {
			t[sent.pre] = make(map[string][]string)
		}
		t[sent.pre][sent.root] = append(t[sent.pre][sent.root], sent.suf)
	}

	return t
}

func ToRules(t TripletMap) []Rule {
	var r []Rule
	var rule Rule

	for k, v := range t {
		for kk, vv := range v {
			rule = NewRule(Ngram{kk, len(strings.Split(kk, " ")) - 1, 0})
			rule.pre = append(rule.pre, k)
			rule.suf = append(rule.suf, vv...)
			r = append(r, rule)
		}
	}

	return r
}
