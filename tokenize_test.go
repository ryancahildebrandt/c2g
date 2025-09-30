// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:51:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"testing"

	"github.com/bzick/tokenizer"
	"github.com/stretchr/testify/assert"
)

func TestValidateTokenizerString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{name: "", args: args{s: ""}, assertion: assert.Error},
		{name: "", args: args{s: "\x00"}, assertion: assert.Error},
		{name: "", args: args{s: "\x00\x00"}, assertion: assert.Error},
		{name: "", args: args{s: " "}, assertion: assert.NoError},
		{name: "", args: args{s: "abc"}, assertion: assert.NoError},
		{name: "", args: args{s: "()"}, assertion: assert.NoError},
		{name: "", args: args{s: "\x01"}, assertion: assert.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.assertion(t, ValidateTokenizerString(tt.args.s), fmt.Sprintf("ValidateTokenizerString(%v)", tt.args.s))
		})
	}
}

func TestUnigramTokenize(t *testing.T) {
	tok := NewUnigramTokenizer()

	type args struct {
		e string
		t *tokenizer.Tokenizer
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{e: "", t: tok}, want: []string{}},
		{name: "", args: args{e: " ", t: tok}, want: []string{}},
		{name: "", args: args{e: " 	", t: tok}, want: []string{}},
		{name: "", args: args{e: ".", t: tok}, want: []string{"."}},
		{name: "", args: args{e: "..?", t: tok}, want: []string{".", ".", "?"}},
		{name: "", args: args{e: "a.b", t: tok}, want: []string{"a", ".", "b"}},
		{name: "", args: args{e: "a . b", t: tok}, want: []string{"a", ".", "b"}},
		{name: "", args: args{e: " a. b", t: tok}, want: []string{"a", ".", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, UnigramTokenize(tt.args.e, tt.args.t))
		})
	}
}

func TestUnigramNormalize(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{e: ""}, want: ""},
		{name: "", args: args{e: " "}, want: ""},
		{name: "", args: args{e: " 	"}, want: ""},
		{name: "", args: args{e: "."}, want: "."},
		{name: "", args: args{e: "..?"}, want: ". . ?"},
		{name: "", args: args{e: "a.b"}, want: "a . b"},
		{name: "", args: args{e: "a . b"}, want: "a . b"},
		{name: "", args: args{e: " a. b"}, want: "a . b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			assert.Equal(t, tt.want, UnigramNormalize(tt.args.e, tk))
		})
	}
}
