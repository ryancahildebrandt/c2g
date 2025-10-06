// ","*"," coding: utf","8 ","*","

// Created on Tue Oct 21 07:10:21 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"

	"github.com/jdkato/prose/tag"
	"github.com/stretchr/testify/assert"
)

func TestSyntacticTagger_POS(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args  args
		want  []string
		want1 []string
	}{
		{args: args{s: ""}, want: []string{}, want1: []string{}},
		{args: args{s: "I"}, want: []string{"PRP"}, want1: []string{"I"}},
		{args: args{s: "can you tell me if i can regisger two accounts with a single email address?"}, want: []string{"MD", "PRP", "VB", "PRP", "IN", "VBN", "MD", "VB", "CD", "NNS", "IN", "DT", "JJ", "NN", "NN", "."}, want1: []string{"can", "you", "tell", "me", "if", "i", "can", "regisger", "two", "accounts", "with", "a", "single", "email", "address", "?"}},
		{args: args{s: "I have no online account"}, want: []string{"PRP", "VBP", "DT", "NN", "NN"}, want1: []string{"I", "have", "no", "online", "account"}},
		{args: args{s: "could you ask an agent how to open an account"}, want: []string{"MD", "PRP", "VB", "DT", "NN", "WRB", "TO", "VB", "DT", "NN"}, want1: []string{"could", "you", "ask", "an", "agent", "how", "to", "open", "an", "account"}},
		{args: args{s: "i want an online account"}, want: []string{"NN", "VBP", "DT", "NN", "NN"}, want1: []string{"i", "want", "an", "online", "account"}},
		{args: args{s: "i want an account"}, want: []string{"NN", "VBP", "DT", "NN"}, want1: []string{"i", "want", "an", "account"}},
		{args: args{s: "tell me if I can register  two online accounts with the same email"}, want: []string{"VB", "PRP", "IN", "PRP", "MD", "VB", "CD", "JJ", "NNS", "IN", "DT", "JJ", "NN"}, want1: []string{"tell", "me", "if", "I", "can", "register", "two", "online", "accounts", "with", "the", "same", "email"}},
		{args: args{s: "i want to know if i could create two profiles with the same email address"}, want: []string{"NN", "VBP", "TO", "VB", "IN", "NNS", "MD", "VB", "CD", "NNS", "IN", "DT", "JJ", "NN", "NN"}, want1: []string{"i", "want", "to", "know", "if", "i", "could", "create", "two", "profiles", "with", "the", "same", "email", "address"}},
		{args: args{s: "can you tell me if i can create more than one fucking user account with the same email?"}, want: []string{"MD", "PRP", "VB", "PRP", "IN", "NNS", "MD", "VB", "JJR", "IN", "CD", "VBG", "NN", "NN", "IN", "DT", "JJ", "NN", "."}, want1: []string{"can", "you", "tell", "me", "if", "i", "can", "create", "more", "than", "one", "fucking", "user", "account", "with", "the", "same", "email", "?"}},
		{args: args{s: "were to create an onlind account"}, want: []string{"VBD", "TO", "VB", "DT", "NN", "NN"}, want1: []string{"were", "to", "create", "an", "onlind", "account"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tok := NewWordTokenizer()
			mod := tag.NewPerceptronTagger()
			tag := NewSyntacticTagger(mod, tok)
			got, got1 := tag.POS(tt.args.s)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestSyntacticTagger_Constituency(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args  args
		want  []string
		want1 []string
	}{
		{args: args{s: ""}, want: []string{}, want1: []string{}},
		{args: args{s: "I"}, want: []string{"PRP"}, want1: []string{"I"}},
		{args: args{s: "can you tell me if i can regisger two accounts with a single email address?"}, want: []string{"MD", "PRP", "VB", "PRP", "IN", "VBN", "MD", "VP", "IN", "DT", "NP", "NN", "."}, want1: []string{"can", "you", "tell", "me", "if", "i", "can", "regisger two accounts", "with", "a", "single email", "address", "?"}},
		{args: args{s: "I have no online account"}, want: []string{"PRP", "VP", "NN"}, want1: []string{"I", "have no online", "account"}},
		{args: args{s: "could you ask an agent how to open an account"}, want: []string{"MD", "PRP", "VP", "WRB", "TO", "VP"}, want1: []string{"could", "you", "ask an agent", "how", "to", "open an account"}},
		{args: args{s: "i want an online account"}, want: []string{"NN", "VP", "NN"}, want1: []string{"i", "want an online", "account"}},
		{args: args{s: "i want an account"}, want: []string{"NN", "VP"}, want1: []string{"i", "want an account"}},
		{args: args{s: "tell me if I can register  two online accounts with the same email"}, want: []string{"VB", "PRP", "IN", "PRP", "MD", "VB", "CD", "NP", "IN", "DT", "NP"}, want1: []string{"tell", "me", "if", "I", "can", "register", "two", "online accounts", "with", "the", "same email"}},
		{args: args{s: "i want to know if i could create two profiles with the same email address"}, want: []string{"NN", "VBP", "TO", "VB", "IN", "NNS", "MD", "VP", "IN", "DT", "NP", "NN"}, want1: []string{"i", "want", "to", "know", "if", "i", "could", "create two profiles", "with", "the", "same email", "address"}},
		{args: args{s: "can you tell me if i can create more than one fucking user account with the same email?"}, want: []string{"MD", "PRP", "VB", "PRP", "IN", "NNS", "MD", "VB", "JJR", "QP", "VBG", "NN", "NN", "IN", "DT", "NP", "."}, want1: []string{"can", "you", "tell", "me", "if", "i", "can", "create", "more", "than one", "fucking", "user", "account", "with", "the", "same email", "?"}},
		{args: args{s: "were to create an onlind account"}, want: []string{"VBD", "TO", "VP", "NN"}, want1: []string{"were", "to", "create an onlind", "account"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tok := NewWordTokenizer()
			mod := tag.NewPerceptronTagger()
			tag := NewSyntacticTagger(mod, tok)
			got, got1 := tag.Constituency(tt.args.s)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}
