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
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{s: ""}, assertion: assert.Error},
		{args: args{s: "\x00"}, assertion: assert.Error},
		{args: args{s: "\x00\x00"}, assertion: assert.Error},
		{args: args{s: " "}, assertion: assert.NoError},
		{args: args{s: "abc"}, assertion: assert.NoError},
		{args: args{s: "()"}, assertion: assert.NoError},
		{args: args{s: "\x01"}, assertion: assert.NoError},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tt.assertion(t, ValidateTokenizerString(tt.args.s), fmt.Sprintf("ValidateTokenizerString(%v)", tt.args.s))
		})
	}
}

func Test_wordTokenizer_tokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{s: ""}, want: []string{}},
		{args: args{s: " "}, want: []string{}},
		{args: args{s: " 	"}, want: []string{}},
		{args: args{s: "."}, want: []string{"."}},
		{args: args{s: "..?"}, want: []string{".", ".", "?"}},
		{args: args{s: "a.b"}, want: []string{"a", ".", "b"}},
		{args: args{s: "a . b"}, want: []string{"a", ".", "b"}},
		{args: args{s: " a. b"}, want: []string{"a", ".", "b"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
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
		args args
		want string
	}{
		{args: args{s: ""}, want: ""},
		{args: args{s: " "}, want: ""},
		{args: args{s: " 	"}, want: ""},
		{args: args{s: "."}, want: "."},
		{args: args{s: "..?"}, want: ". . ?"},
		{args: args{s: "a.b"}, want: "a . b"},
		{args: args{s: "a . b"}, want: "a . b"},
		{args: args{s: " a. b"}, want: "a . b"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
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
		args args
		want []string
	}{
		{args: args{s: ""}, want: []string{}},
		{args: args{s: "<SEP>"}, want: []string{}},
		{args: args{s: " <SEP>	"}, want: []string{" ", "\t"}},
		{args: args{s: "."}, want: []string{"."}},
		{args: args{s: ".<SEP>.?"}, want: []string{".", ".?"}},
		{args: args{s: "a<SEP>.b<SEP>"}, want: []string{"a", ".b"}},
		{args: args{s: "a . <SEP>b"}, want: []string{"a . ", "b"}},
		{args: args{s: " <SEP>a.<SEP> b"}, want: []string{" ", "a.", " b"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
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
		args args
		want string
	}{
		{args: args{s: ""}, want: ""},
		{args: args{s: "<SEP><SEP>"}, want: ""},
		{args: args{s: " <SEP>	"}, want: "  	"},
		{args: args{s: "."}, want: "."},
		{args: args{s: ".<SEP>.?"}, want: ". .?"},
		{args: args{s: "a<SEP><SEP>.b<SEP>"}, want: "a .b"},
		{args: args{s: "a . <SEP>b<SEP>"}, want: "a .  b"},
		{args: args{s: " <SEP><SEP>a.<SEP> b"}, want: "  a.  b"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewSepTokenizer()
			assert.Equal(t, tt.want, tk.normalize(tt.args.s))
		})
	}
}

func Test_posTokenizer_tokenize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{s: ""}, want: []string{}},
		{args: args{s: "-"}, want: []string{}},
		{args: args{s: " -	"}, want: []string{" ", "\t"}},
		{args: args{s: "."}, want: []string{"."}},
		{args: args{s: ".-.?"}, want: []string{".", ".?"}},
		{args: args{s: "a-.b-"}, want: []string{"a", ".b"}},
		{args: args{s: "a . -b"}, want: []string{"a . ", "b"}},
		{args: args{s: " -a.- b"}, want: []string{" ", "a.", " b"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewPOSTokenizer()
			assert.Equal(t, tt.want, tk.tokenize(tt.args.s))
		})
	}
}

func Test_posTokenizer_normalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{s: ""}, want: ""},
		{args: args{s: "--"}, want: ""},
		{args: args{s: " -	"}, want: " -	"},
		{args: args{s: "."}, want: "."},
		{args: args{s: ".-.?"}, want: ".-.?"},
		{args: args{s: "a--.b-"}, want: "a-.b"},
		{args: args{s: "a . -b-"}, want: "a . -b"},
		{args: args{s: " --a.- b"}, want: " -a.- b"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewPOSTokenizer()
			assert.Equal(t, tt.want, tk.normalize(tt.args.s))
		})
	}
}
