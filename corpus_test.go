// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"

	"github.com/bzick/tokenizer"
	"github.com/stretchr/testify/assert"
)

func Test_corpusToTokens(t *testing.T) {
	var tok = NewUnigramTokenizer()
	var emp [][]Expression

	type args struct {
		c Corpus
		t *tokenizer.Tokenizer
	}
	tests := []struct {
		name string
		args args
		want [][]Expression
	}{	
		{name: "", args: args{c: NewCorpus([]Expression{}), t: tok}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{"", ""}), t: tok}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{"", "", "", "", "", ""}), t: tok}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{"  ", " "}), t: tok}, want: emp},
		{name: "", args: args{c: NewCorpus([]Expression{".", ",", ".", "", ".", ""}), t: tok}, want: [][]Expression{{"."}, {","}, {"."}, {"."}}},
		{name: "", args: args{c: NewCorpus([]Expression{"abc", "abc", "d e f", "g.h,i"}), t: tok}, want: [][]Expression{{"abc"}, {"abc"}, {"d", "e", "f"}, {"g", ".", "h", ",", "i"}}},
		{name: "", args: args{c: NewCorpus([]Expression{"abc abc", "d e e f", "g .", ". h", "h ,"}), t: tok}, want: [][]Expression{{"abc", "abc"}, {"d", "e", "e", "f"}, {"g", "."}, {".", "h"}, {"h", ","}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, corpusToTokens(tt.args.c, tt.args.t), "corpusToTokens(%v, %v)", tt.args.c, tt.args.t)
		})
	}
}
