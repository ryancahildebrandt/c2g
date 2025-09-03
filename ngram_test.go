// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:17:15 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"

	"github.com/bzick/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestNewTransitions(t *testing.T) {
	var emp Transitions = Transitions{}
	var tok *tokenizer.Tokenizer = NewUnigramTokenizer()

	type args struct {
		c Corpus
	}
	tests := []struct {
		name string
		args args
		want Transitions
	}{
		{name: "", args: args{c: NewCorpus([]Expression{})}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{"", ""})}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{"", "", "", "", "", ""})}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{".", ",", ".", "", ".", ""})}, want: Transitions{".": TransitionProbabilities{"": 1.0}, ",": TransitionProbabilities{"": 1.0}}},
		{name: "", args: args{c: NewCorpus([]Expression{"abc abc", "d e e f", "g .", ". h", "h ,"})}, want: Transitions{"abc": TransitionProbabilities{"abc": 1}, "d": TransitionProbabilities{"e": 1}, "e": TransitionProbabilities{"e": 0.5, "f": 0.5}, "g": TransitionProbabilities{".": 1}, ".": TransitionProbabilities{"h": 1}, "h": TransitionProbabilities{",": 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewTransitions(tt.args.c, tok), "NewTransitions(%v)", tt.args.c)
		})
	}
}

func Test_probabilityNorm(t *testing.T) {
	type args struct {
		p TransitionProbabilities
	}
	tests := []struct {
		name string
		args args
		want TransitionProbabilities
	}{
		{name: "", args: args{p: TransitionProbabilities{}}, want: TransitionProbabilities{}},
		{name: "", args: args{p: TransitionProbabilities{"": 0}}, want: TransitionProbabilities{}},
		{name: "", args: args{p: TransitionProbabilities{"": 1}}, want: TransitionProbabilities{"": 1}},
		{name: "", args: args{p: TransitionProbabilities{"a": 2, "b": 3, "c": 10}}, want: TransitionProbabilities{"a": 2.0 / 15.0, "b": 3 / 15.0, "c": 10 / 15.0}},
		{name: "", args: args{p: TransitionProbabilities{"a": 1, "b": 1, "": 1}}, want: TransitionProbabilities{"a": 1.0 / 3.0, "b": 1.0 / 3.0, "": 1.0 / 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, probabilityNorm(tt.args.p), "probabilityNorm(%v)", tt.args.p)
		})
	}
}

func TestNgramTokenize(t *testing.T) {
	var emp []Expression
	transitions := Transitions{
		"a": TransitionProbabilities{},
		"b": TransitionProbabilities{"a": 0.1, "b": 0.6, "c": 0.2, "d": 0.7, "e": 0.3, "f": 0.8},
		"c": TransitionProbabilities{"a": 0.2, "b": 0.7, "c": 0.3, "d": 0.8, "e": 0.4, "f": 0.9},
		"d": TransitionProbabilities{"a": 0.3, "b": 0.8, "c": 0.4, "d": 0.9, "e": 0.5, "f": 0.1},
		"e": TransitionProbabilities{"a": 0.4, "b": 0.9, "c": 0.5, "d": 0.1, "e": 0.6, "f": 0.2},
		"f": TransitionProbabilities{"a": 0.5, "b": 0.1, "c": 0.6, "d": 0.2, "e": 0.7, "f": 0.3},
	}

	type args struct {
		s []Expression
		t Transitions
		p float64
	}
	tests := []struct {
		name string
		args args
		want []Expression
	}{
		{name: "", args: args{s: []Expression{}, t: transitions, p: 0.0}, want: emp},
		{name: "", args: args{s: []Expression{}, t: transitions, p: 0.5}, want: emp},
		{name: "", args: args{s: []Expression{}, t: transitions, p: 1.0}, want: emp},

		{name: "", args: args{s: []Expression{"", "", "", "", "", ""}, t: transitions, p: 0.0}, want: emp},
		{name: "", args: args{s: []Expression{"", "", "", "", "", ""}, t: transitions, p: 0.5}, want: emp},
		{name: "", args: args{s: []Expression{"", "", "", "", "", ""}, t: transitions, p: 1.0}, want: emp},

		{name: "", args: args{s: []Expression{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 0.0}, want: []Expression{"a b c d e f"}},
		{name: "", args: args{s: []Expression{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 0.5}, want: []Expression{"a", "b c d", "e f"}},
		{name: "", args: args{s: []Expression{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 1.0}, want: []Expression{"a", "b", "c", "d", "e f"}},

		{name: "", args: args{s: []Expression{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 0.0}, want: []Expression{"a f f d d h"}},
		{name: "", args: args{s: []Expression{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 0.5}, want: []Expression{"a", "f", "f d", "d h"}},
		{name: "", args: args{s: []Expression{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 1.0}, want: []Expression{"a", "f", "f", "d", "d h"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NgramTokenize(tt.args.s, tt.args.t, tt.args.p))
		})
	}
}
