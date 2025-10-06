// -*- coding: utf-8 -*-

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
		args args
		want string
	}{
		{args: args{s: ""}, want: ""},
		{args: args{s: "I"}, want: "PRP()"},
		{args: args{s: "can you tell me if i can regisger two accounts with a single email address?"}, want: "MD()PRP()VB()PRP()IN()VBN()MD()VB()CD()NNS()IN()DT()JJ()NN()NN().()"},
		{args: args{s: "I have no online account"}, want: "PRP()VBP()DT()NN()NN()"},
		{args: args{s: "could you ask an agent how to open an account"}, want: "MD()PRP()VB()DT()NN()WRB()TO()VB()DT()NN()"},
		{args: args{s: "i want an online account"}, want: "NN()VBP()DT()NN()NN()"},
		{args: args{s: "i want an account"}, want: "NN()VBP()DT()NN()"},
		{args: args{s: "tell me if I can register  two online accounts with the same email"}, want: "VB()PRP()IN()PRP()MD()VB()CD()JJ()NNS()IN()DT()JJ()NN()"},
		{args: args{s: "i want to know if i could create two profiles with the same email address"}, want: "NN()VBP()TO()VB()IN()NNS()MD()VB()CD()NNS()IN()DT()JJ()NN()NN()"},
		{args: args{s: "can you tell me if i can create more than one fucking user account with the same email?"}, want: "MD()PRP()VB()PRP()IN()NNS()MD()VB()JJR()IN()CD()VBG()NN()NN()IN()DT()JJ()NN().()"},
		{args: args{s: "were to create an onlind account"}, want: "VBD()TO()VB()DT()NN()NN()"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tok := NewWordTokenizer()
			mod := tag.NewPerceptronTagger()
			tag := NewSyntacticTagger(mod, tok)
			assert.Equal(t, tt.want, tag.POS(tt.args.s))
		})
	}
}

func TestSyntacticTagger_Constituency(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{s: ""}, want: ""},
		{args: args{s: "I"}, want: "PRP()"},
		{args: args{s: "can you tell me if i can regisger two accounts with a single email address?"}, want: "MD()PRP()VB()PRP()IN()VBN()MD()VB()ADJPS()IN()DT()NP()NN().()"},
		{args: args{s: "I have no online account"}, want: "PRP()VP()NN()"},
		{args: args{s: "could you ask an agent how to open an account"}, want: "MD()PRP()VP()WRB()TO()VP()"},
		{args: args{s: "i want an online account"}, want: "NN()VP()NN()"},
		{args: args{s: "i want an account"}, want: "NN()VP()"},
		{args: args{s: "tell me if I can register  two online accounts with the same email"}, want: "VB()PRP()IN()PRP()MD()VB()CD()NPS()IN()DT()NP()"},
		{args: args{s: "i want to know if i could create two profiles with the same email address"}, want: "NN()VBP()TO()VB()IN()NNS()MD()VB()ADJPS()IN()DT()NP()NN()"},
		{args: args{s: "can you tell me if i can create more than one fucking user account with the same email?"}, want: "MD()PRP()VB()PRP()IN()NNS()MD()VB()JJR()QP()VBG()NN()NN()IN()DT()NP().()"},
		{args: args{s: "were to create an onlind account"}, want: "VBD()TO()VP()NN()"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tok := NewWordTokenizer()
			mod := tag.NewPerceptronTagger()
			tag := NewSyntacticTagger(mod, tok)
			assert.Equal(t, tt.want, tag.Constituency(tt.args.s))
		})
	}
}
