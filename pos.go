// -*- coding: utf-8 -*-

// Created on Tue Oct 21 07:10:21 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"strings"

	"github.com/jdkato/prose/tag"
)

type ConstituencyRule struct {
	rule []string
	tag  string
}

func (c *ConstituencyRule) signature() string { return "" }

var constituencyRules = []ConstituencyRule{
	ConstituencyRule{tag: "ADJP", rule: []string{"JJ", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"RB", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"RB", "VBN"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"RB", "JJR"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"CD", "NN"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"RBR", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"RBS", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"JJ", "CC", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"NNP", ",", "JJ"}},
	ConstituencyRule{tag: "ADJP", rule: []string{"CD", "CD", "NN"}},
	ConstituencyRule{tag: "ADVP", rule: []string{"RB", "RB"}},
	ConstituencyRule{tag: "ADVP", rule: []string{"IN", "JJS"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"IN", "IN"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"RB", "RB"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"RB", "RB", "IN"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"CC", "RB"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"RB", "IN"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"RB", "TO", "VB"}},
	ConstituencyRule{tag: "CONJP", rule: []string{"RB", "JJ"}},
	ConstituencyRule{tag: "INTJ", rule: []string{"VB", ",", "RB"}},
	ConstituencyRule{tag: "INTJ", rule: []string{"NN", ",", "NN", ",", "NN", ",", "NN", ",", "NN", ",", "NN"}},
	ConstituencyRule{tag: "INTJ", rule: []string{"RB", "JJ"}},
	ConstituencyRule{tag: "LST", rule: []string{"LS", ":"}},
	ConstituencyRule{tag: "LST", rule: []string{"LS", "."}},
	ConstituencyRule{tag: "NP", rule: []string{"NNP", "NNP"}},
	ConstituencyRule{tag: "NP", rule: []string{"CD", "NNS"}},
	ConstituencyRule{tag: "NP", rule: []string{"DT", "NN"}},
	ConstituencyRule{tag: "NP", rule: []string{"DT", "JJ", "NN"}},
	ConstituencyRule{tag: "NP", rule: []string{"JJ", "NN"}},
	ConstituencyRule{tag: "NP", rule: []string{"NN", "NNS"}},
	ConstituencyRule{tag: "NP", rule: []string{"DT", "NN", "NN"}},
	ConstituencyRule{tag: "NP", rule: []string{"DT", "NNS"}},
	ConstituencyRule{tag: "NP", rule: []string{"NNP", "NNP", "NNP"}},
	ConstituencyRule{tag: "NP", rule: []string{"JJ", "NNS"}},
	ConstituencyRule{tag: "NP", rule: []string{"CD", "NN"}},
	ConstituencyRule{tag: "NX", rule: []string{"NN", "NN"}},
	ConstituencyRule{tag: "NX", rule: []string{"NNP", "NNP"}},
	ConstituencyRule{tag: "NX", rule: []string{"JJ", "NNS"}},
	ConstituencyRule{tag: "NX", rule: []string{"NN", "NNS"}},
	ConstituencyRule{tag: "NX", rule: []string{"JJ", "NN", "NN"}},
	ConstituencyRule{tag: "NX", rule: []string{"JJ", "NN"}},
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
	ConstituencyRule{tag: "UCP", rule: []string{"NN", "CC", "JJ"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", "CC", "NN"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", "CC", "NNP", "NNP"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", ",", "NN", ",", "NN", "CC", "JJ"}},
	ConstituencyRule{tag: "UCP", rule: []string{"NN", ",", "NN", "CC", "JJ"}},
	ConstituencyRule{tag: "UCP", rule: []string{"NNP", "CC", "JJ"}},
	ConstituencyRule{tag: "UCP", rule: []string{"NN", ",", "JJ", "NN", ",", "NNS", "CC", "NN"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", ",", "NN", "CC", "NN"}},
	ConstituencyRule{tag: "UCP", rule: []string{"NN", "CC", "NNS"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", ",", "NN", "CC", "JJ"}},
	ConstituencyRule{tag: "UCP", rule: []string{"JJ", "CC", "NN", "NN"}},
	ConstituencyRule{tag: "UCP", rule: []string{"NN", ",", "JJ", "CC", "NN", "NNS"}},
	ConstituencyRule{tag: "WHADJP", rule: []string{"WRB", "JJ"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"WRB", "JJ"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"WDT", "NNS"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"WP$", "NN", "NN"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"DT", "NN"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"WRB", "RB"}},
	ConstituencyRule{tag: "WHNP", rule: []string{"WP$", "NNS"}},
}

func posSignature(s []string, m *tag.PerceptronTagger) string {
	var tags []string
	for _, t := range m.Tag(s) {
		tags = append(tags, t.Tag)
	}
	return strings.Join(tags, "()")
}

func conSignature() {}
