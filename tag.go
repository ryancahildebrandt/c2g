// -*- coding: utf-8 -*-

// Created on Tue Oct 21 07:10:21 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"slices"
	"strings"

	"github.com/jdkato/prose/tag"
)

type SyntacticTagger struct {
	*tag.PerceptronTagger
	Tokenizer
	rules []ConstituencyRule
}

type ConstituencyRule struct {
	rule []string
	tag  string
}

func (t *SyntacticTagger) POS(s string) ([]string, []string) {
	var tags []string
	var tokens = t.tokenize(s)

	if len(tokens) == 0 {
		return []string{}, []string{}
	}

	for _, tt := range t.Tag(tokens) {
		tags = append(tags, tt.Tag)
	}

	return tags, tokens
}

func (t *SyntacticTagger) Constituency(s string) ([]string, []string) {
	scanSubseq := func(s1, s2 []string) ([]int, bool) {
		var res []int

		if len(s1) < len(s2) {
			return res, false
		}

		for i := range s2 {
			if !slices.Contains(s1, s2[i]) {
				return res, false
			}
		}

		for i := 0; i <= len(s1)-len(s2); i++ {
			win := s1[i : i+len(s2)]
			if slices.Equal(s2, win) {
				return []int{i, i + len(s2)}, true
			}
		}

		return res, false
	}

	tags, tokens := t.POS(s)

	if len(tokens) == 0 {
		return []string{}, []string{}
	}

	for j := range t.rules {
		tag := t.rules[j].tag
		rule := t.rules[j].rule
		for {
			ind, found := scanSubseq(tags, rule)
			if !found {
				break
			}
			tags = slices.Replace(tags, ind[0], ind[1], tag)
			tokens = slices.Replace(tokens, ind[0], ind[1], strings.Join(tokens[ind[0]:ind[1]], " "))
		}
	}

	return tags, tokens
}

func NewSyntacticTagger(m *tag.PerceptronTagger, t Tokenizer) SyntacticTagger {
	var rules = []ConstituencyRule{
		{tag: "ADJP", rule: []string{"NP", "JJ"}},
		{tag: "ADJP", rule: []string{"JJ", "JJ"}},
		{tag: "ADJP", rule: []string{"RB", "JJ"}},
		{tag: "ADJP", rule: []string{"RB", "VBN"}},
		{tag: "ADJP", rule: []string{"RB", "JJR"}},
		{tag: "ADJP", rule: []string{"JJ", "PP"}},
		{tag: "ADJP", rule: []string{"CD", "NN"}},
		{tag: "ADJP", rule: []string{"QP", "NN"}},
		{tag: "ADJP", rule: []string{"ADJP", "PP"}},
		{tag: "ADJP", rule: []string{"RBR", "JJ"}},
		{tag: "ADJP", rule: []string{"RBS", "JJ"}},
		{tag: "ADJP", rule: []string{"JJ", "CC", "JJ"}},
		{tag: "ADJP", rule: []string{"NNP", ",", "JJ"}},
		{tag: "ADJP", rule: []string{"CD", "CD", "NN"}},
		{tag: "ADVP", rule: []string{"RB", "PP"}},
		{tag: "ADVP", rule: []string{"RB", "NP"}},
		{tag: "ADVP", rule: []string{"RB", "RB"}},
		{tag: "ADVP", rule: []string{"IN", "JJS"}},
		{tag: "CONJP", rule: []string{"IN", "IN"}},
		{tag: "CONJP", rule: []string{"RB", "RB"}},
		{tag: "CONJP", rule: []string{"RB", "RB", "IN"}},
		{tag: "CONJP", rule: []string{"CC", "RB"}},
		{tag: "CONJP", rule: []string{"RB", "IN"}},
		{tag: "CONJP", rule: []string{"RB", "TO", "VB"}},
		{tag: "CONJP", rule: []string{"RB", "JJ"}},
		{tag: "NP", rule: []string{"NNP", "NNP"}},
		{tag: "NP", rule: []string{"CD", "NNS"}},
		{tag: "NP", rule: []string{"DT", "NN"}},
		{tag: "NP", rule: []string{"DT", "JJ", "NN"}},
		{tag: "NP", rule: []string{"NP", "PP"}},
		{tag: "NP", rule: []string{"JJ", "NN"}},
		{tag: "NP", rule: []string{"NN", "NNS"}},
		{tag: "NP", rule: []string{"DT", "NN", "NN"}},
		{tag: "NP", rule: []string{"DT", "NNS"}},
		{tag: "NP", rule: []string{"NP", "SBAR"}},
		{tag: "NP", rule: []string{"NNP", "NNP", "NNP"}},
		{tag: "NP", rule: []string{"NP", "CC", "NP"}},
		{tag: "NP", rule: []string{"JJ", "NNS"}},
		{tag: "NP", rule: []string{"NP", "VP"}},
		{tag: "NP", rule: []string{"CD", "NN"}},
		{tag: "PP", rule: []string{"IN", "NP"}},
		{tag: "PP", rule: []string{"TO", "NP"}},
		{tag: "PRN", rule: []string{":", "NP"}},
		{tag: "PRN", rule: []string{",", "SINV", ","}},
		{tag: "PRN", rule: []string{":", "PP", ":"}},
		{tag: "PRN", rule: []string{",", "PP", ","}},
		{tag: "PRN", rule: []string{":", "NP", ":"}},
		{tag: "PRN", rule: []string{",", "S", ","}},
		{tag: "PRN", rule: []string{"SINV", ","}},
		{tag: "PRN", rule: []string{",", "ADVP", ","}},
		{tag: "PRN", rule: []string{",", "''", "SINV", ","}},
		{tag: "PRN", rule: []string{":", "SBAR", ":"}},
		{tag: "PRN", rule: []string{",", "''", "S", ","}},
		{tag: "PRN", rule: []string{":", "S", ":"}},
		{tag: "QP", rule: []string{"RBR", "IN", "CD"}},
		{tag: "QP", rule: []string{"CD", "CD"}},
		{tag: "QP", rule: []string{"IN", "CD"}},
		{tag: "QP", rule: []string{"$", "CD", "CD"}},
		{tag: "QP", rule: []string{"IN", "$", "CD", "CD"}},
		{tag: "QP", rule: []string{"IN", "CD", "CD"}},
		{tag: "QP", rule: []string{"RB", "$", "CD", "CD"}},
		{tag: "QP", rule: []string{"RB", "CD"}},
		{tag: "QP", rule: []string{"JJR", "IN", "CD"}},
		{tag: "QP", rule: []string{"CD", "TO", "CD"}},
		{tag: "QP", rule: []string{"JJR", "IN", "$", "CD", "CD"}},
		{tag: "QP", rule: []string{"CD", "NN", "TO", "CD", "NN"}},
		{tag: "QP", rule: []string{"#", "CD", "CD"}},
		{tag: "VP", rule: []string{"MD", "VP"}},
		{tag: "VP", rule: []string{"VBD", "VP"}},
		{tag: "VP", rule: []string{"TO", "VP"}},
		{tag: "VP", rule: []string{"VB", "NP"}},
		{tag: "VP", rule: []string{"VBZ", "VP"}},
		{tag: "VP", rule: []string{"VBN", "NP"}},
		{tag: "VP", rule: []string{"VBD", "SBAR"}},
		{tag: "VP", rule: []string{"VBZ", "NP"}},
		{tag: "VP", rule: []string{"VBG", "NP"}},
		{tag: "VP", rule: []string{"VBP", "VP"}},
		{tag: "VP", rule: []string{"VBD", "NP"}},
		{tag: "VP", rule: []string{"VBP", "NP"}},
		{tag: "VP", rule: []string{"VBD", "S"}},
		{tag: "VP", rule: []string{"VP", "CC", "VP"}},
		{tag: "VP", rule: []string{"VBZ", "S"}},
		{tag: "VP", rule: []string{"VBN", "NP", "PP"}},
		{tag: "VP", rule: []string{"VB", "VP"}},
		{tag: "VP", rule: []string{"VBZ", "SBAR"}},
		{tag: "VP", rule: []string{"VB", "S"}},
	}
	slices.SortStableFunc(rules, func(i, j ConstituencyRule) int { return len(i.rule) - len(j.rule) })
	var c = SyntacticTagger{m, t, rules}

	return c
}
