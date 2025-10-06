// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:51:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"testing"

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

func Test_wordTokenizer_tokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{s: ""}, want: []string{}},
		{name: "", args: args{s: " "}, want: []string{}},
		{name: "", args: args{s: " 	"}, want: []string{}},
		{name: "", args: args{s: "."}, want: []string{"."}},
		{name: "", args: args{s: "..?"}, want: []string{".", ".", "?"}},
		{name: "", args: args{s: "a.b"}, want: []string{"a", ".", "b"}},
		{name: "", args: args{s: "a . b"}, want: []string{"a", ".", "b"}},
		{name: "", args: args{s: " a. b"}, want: []string{"a", ".", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewWordTokenizer()
			assert.Equal(t, tt.want, tk.tokenize(tt.args.s))
		})
	}
}

func Test_wordTokenizer_normalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{s: ""}, want: ""},
		{name: "", args: args{s: " "}, want: ""},
		{name: "", args: args{s: " 	"}, want: ""},
		{name: "", args: args{s: "."}, want: "."},
		{name: "", args: args{s: "..?"}, want: ". . ?"},
		{name: "", args: args{s: "a.b"}, want: "a . b"},
		{name: "", args: args{s: "a . b"}, want: "a . b"},
		{name: "", args: args{s: " a. b"}, want: "a . b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewWordTokenizer()
			assert.Equal(t, tt.want, tk.normalize(tt.args.s))
		})
	}
}

func Test_sepTokenizer_tokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{s: ""}, want: []string{}},
		{name: "", args: args{s: "<SEP>"}, want: []string{}},
		{name: "", args: args{s: " <SEP>	"}, want: []string{" ", "\t"}},
		{name: "", args: args{s: "."}, want: []string{"."}},
		{name: "", args: args{s: ".<SEP>.?"}, want: []string{".", ".?"}},
		{name: "", args: args{s: "a<SEP>.b<SEP>"}, want: []string{"a", ".b"}},
		{name: "", args: args{s: "a . <SEP>b"}, want: []string{"a . ", "b"}},
		{name: "", args: args{s: " <SEP>a.<SEP> b"}, want: []string{" ", "a.", " b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewSepTokenizer()
			assert.Equal(t, tt.want, tk.tokenize(tt.args.s))
		})
	}
}

func Test_sepTokenizer_normalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{s: ""}, want: ""},
		{name: "", args: args{s: "<SEP><SEP>"}, want: ""},
		{name: "", args: args{s: " <SEP>	"}, want: " <SEP>	"},
		{name: "", args: args{s: "."}, want: "."},
		{name: "", args: args{s: ".<SEP>.?"}, want: ".<SEP>.?"},
		{name: "", args: args{s: "a<SEP><SEP>.b<SEP>"}, want: "a<SEP>.b"},
		{name: "", args: args{s: "a . <SEP>b<SEP>"}, want: "a . <SEP>b"},
		{name: "", args: args{s: " <SEP><SEP>a.<SEP> b"}, want: " <SEP>a.<SEP> b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewSepTokenizer()
			assert.Equal(t, tt.want, tk.normalize(tt.args.s))
		})
	}
}
