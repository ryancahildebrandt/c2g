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

func (c *SyntacticTagger) POS(s string) string {
	var tags []string
	var tokens = c.tokenize(s)

	for _, t := range c.Tag(tokens) {
		tags = append(tags, t.Tag)
	}
	tags = append(tags, "")

	return strings.Join(tags, "()")
}

func (c *SyntacticTagger) Constituency(s string) string {
	var sig = c.POS(s)

	for j := range c.rules {
		tag := c.rules[j].tag
		ss := strings.Join(c.rules[j].rule, "()")
		sig = strings.ReplaceAll(sig, ss, tag)
	}

	return sig
}

func NewSyntacticTagger(m *tag.PerceptronTagger, t Tokenizer) SyntacticTagger {
	var rules = []ConstituencyRule{
		ConstituencyRule{tag: "ADJP", rule: []string{"NP", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"JJ", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"RB", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"RB", "VBN"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"RB", "JJR"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"JJ", "PP"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"CD", "NN"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"QP", "NN"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"ADJP", "PP"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"RBR", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"RBS", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"JJ", "CC", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"NNP", ",", "JJ"}},
		ConstituencyRule{tag: "ADJP", rule: []string{"CD", "CD", "NN"}},
		ConstituencyRule{tag: "ADVP", rule: []string{"RB", "PP"}},
		ConstituencyRule{tag: "ADVP", rule: []string{"RB", "NP"}},
		ConstituencyRule{tag: "ADVP", rule: []string{"RB", "RB"}},
		ConstituencyRule{tag: "ADVP", rule: []string{"IN", "JJS"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"IN", "IN"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"RB", "RB"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"RB", "RB", "IN"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"CC", "RB"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"RB", "IN"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"RB", "TO", "VB"}},
		ConstituencyRule{tag: "CONJP", rule: []string{"RB", "JJ"}},
		ConstituencyRule{tag: "NP", rule: []string{"NNP", "NNP"}},
		ConstituencyRule{tag: "NP", rule: []string{"CD", "NNS"}},
		ConstituencyRule{tag: "NP", rule: []string{"DT", "NN"}},
		ConstituencyRule{tag: "NP", rule: []string{"DT", "JJ", "NN"}},
		ConstituencyRule{tag: "NP", rule: []string{"NP", "PP"}},
		ConstituencyRule{tag: "NP", rule: []string{"JJ", "NN"}},
		ConstituencyRule{tag: "NP", rule: []string{"NN", "NNS"}},
		ConstituencyRule{tag: "NP", rule: []string{"DT", "NN", "NN"}},
		ConstituencyRule{tag: "NP", rule: []string{"DT", "NNS"}},
		ConstituencyRule{tag: "NP", rule: []string{"NP", "SBAR"}},
		ConstituencyRule{tag: "NP", rule: []string{"NNP", "NNP", "NNP"}},
		ConstituencyRule{tag: "NP", rule: []string{"NP", "CC", "NP"}},
		ConstituencyRule{tag: "NP", rule: []string{"JJ", "NNS"}},
		ConstituencyRule{tag: "NP", rule: []string{"NP", "VP"}},
		ConstituencyRule{tag: "NP", rule: []string{"CD", "NN"}},
		ConstituencyRule{tag: "PP", rule: []string{"IN", "NP"}},
		ConstituencyRule{tag: "PP", rule: []string{"TO", "NP"}},
		ConstituencyRule{tag: "PRN", rule: []string{":", "NP"}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "SINV", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{":", "PP", ":"}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "PP", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{":", "NP", ":"}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "S", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{"SINV", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "ADVP", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "''", "SINV", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{":", "SBAR", ":"}},
		ConstituencyRule{tag: "PRN", rule: []string{",", "''", "S", ","}},
		ConstituencyRule{tag: "PRN", rule: []string{":", "S", ":"}},
		ConstituencyRule{tag: "QP", rule: []string{"RBR", "IN", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"IN", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"$", "CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"IN", "$", "CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"IN", "CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"RB", "$", "CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"RB", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"JJR", "IN", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"CD", "TO", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"JJR", "IN", "$", "CD", "CD"}},
		ConstituencyRule{tag: "QP", rule: []string{"CD", "NN", "TO", "CD", "NN"}},
		ConstituencyRule{tag: "QP", rule: []string{"#", "CD", "CD"}},
		ConstituencyRule{tag: "VP", rule: []string{"MD", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBD", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"TO", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VB", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBZ", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBN", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBD", "SBAR"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBZ", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBG", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBP", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBD", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBP", "NP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBD", "S"}},
		ConstituencyRule{tag: "VP", rule: []string{"VP", "CC", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBZ", "S"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBN", "NP", "PP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VB", "VP"}},
		ConstituencyRule{tag: "VP", rule: []string{"VBZ", "SBAR"}},
		ConstituencyRule{tag: "VP", rule: []string{"VB", "S"}},
	}
	slices.SortStableFunc(rules, func(i, j ConstituencyRule) int { return len(i.rule) - len(j.rule) })
	var c = SyntacticTagger{m, t, rules}

	return c
}
