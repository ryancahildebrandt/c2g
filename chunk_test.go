// -*- coding: utf-8 -*-

// Created on Thu Oct  2 12:57:28 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectTransitions(t *testing.T) {
	type args struct {
		c []Text
	}
	tests := []struct {
		args args
		want Transitions
	}{
		{args: args{c: []Text{}}, want: Transitions{}},
		{args: args{c: []Text{{text: ""}, {text: ""}}}, want: Transitions{}},
		{args: args{c: []Text{{text: ""}, {text: ""}, {text: ""}, {text: ""}, {text: ""}, {text: ""}}}, want: Transitions{}},
		{args: args{c: []Text{{text: "."}, {text: ","}, {text: "."}, {text: ""}, {text: "."}, {text: ""}}}, want: Transitions{".": map[string]float64{"": 1.0}, ",": map[string]float64{"": 1.0}}},
		{args: args{c: []Text{{text: "abc abc"}, {text: "d e e f"}, {text: "g ."}, {text: ". h"}, {text: "h ,"}}}, want: Transitions{"abc": map[string]float64{"abc": 1}, "d": map[string]float64{"e": 1}, "e": map[string]float64{"e": 0.5, "f": 0.5}, "g": map[string]float64{".": 1}, ".": map[string]float64{"h": 1}, "h": map[string]float64{",": 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tok := NewWordTokenizer()
			assert.Equalf(t, tt.want, CollectTransitions(tt.args.c, TokenSplit(tok)), "NewTransitions(%v)", tt.args.c)
		})
	}
}

func TestTransitionChunk(t *testing.T) {
	transitions := Transitions{
		"a": map[string]float64{},
		"b": map[string]float64{"a": 0.1, "b": 0.6, "c": 0.2, "d": 0.7, "e": 0.3, "f": 0.8},
		"c": map[string]float64{"a": 0.2, "b": 0.7, "c": 0.3, "d": 0.8, "e": 0.4, "f": 0.9},
		"d": map[string]float64{"a": 0.3, "b": 0.8, "c": 0.4, "d": 0.9, "e": 0.5, "f": 0.1},
		"e": map[string]float64{"a": 0.4, "b": 0.9, "c": 0.5, "d": 0.1, "e": 0.6, "f": 0.2},
		"f": map[string]float64{"a": 0.5, "b": 0.1, "c": 0.6, "d": 0.2, "e": 0.7, "f": 0.3},
	}

	type args struct {
		s []string
		t Transitions
		p float64
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{s: []string{}, t: transitions, p: 0.0}, want: []string{}},
		{args: args{s: []string{}, t: transitions, p: 0.5}, want: []string{}},
		{args: args{s: []string{}, t: transitions, p: 1.0}, want: []string{}},

		{args: args{s: []string{"", "", "", "", "", ""}, t: transitions, p: 0.0}, want: []string{}},
		{args: args{s: []string{"", "", "", "", "", ""}, t: transitions, p: 0.5}, want: []string{}},
		{args: args{s: []string{"", "", "", "", "", ""}, t: transitions, p: 1.0}, want: []string{}},

		{args: args{s: []string{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 0.0}, want: []string{"a b c d e f"}},
		{args: args{s: []string{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 0.5}, want: []string{"a", "b c d", "e f"}},
		{args: args{s: []string{"a", "b", "c", "d", "e", "f"}, t: transitions, p: 1.0}, want: []string{"a", "b", "c", "d", "e f"}},

		{args: args{s: []string{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 0.0}, want: []string{"a f f d d h"}},
		{args: args{s: []string{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 0.5}, want: []string{"a", "f", "f d", "d h"}},
		{args: args{s: []string{"a", "f", "f", "d", "d", "h"}, t: transitions, p: 1.0}, want: []string{"a", "f", "f", "d", "d h"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, TransitionChunk(tt.args.s, tt.args.s, tt.args.t, tt.args.p))
		})
	}
}

func TestCollectChunks(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		args args
		want []string
	}{
		{args: args{f: "./data/tests/test5.csv"}, want: []string{"I don't have an online account"}},
		{args: args{f: "./data/tests/test6.csv"}, want: []string{
			"I don't have an online account",
			"I don't understand you",
			"I got an error message when I attempted to make a payment",
			"I want an online accoynt",
			"ask an agent to notify issues with my payment",
			"can you show me information about the status of my refund ?",
			"can you show me my invoices ?",
			"can you tell me how I can get some bills ?",
			"i dont want my profile",
			"i want to know wat the email of Customer Service is",
			"where can i leave an opinion for a service ?",
		}},
		{args: args{f: "./data/tests/test7.csv"}, want: []string{}},
		{args: args{f: "./data/tests/test8.csv"}, want: []string{}},
		{args: args{f: "./data/tests/test9.csv"}, want: []string{
			"I don't have an online account",
			"I have a question",
			"I ordered an item and Id like to modify my fucking order",
			"I want to download a bill",
			"I want to know what the number of Customer Service is",
			"I want to make a review for a service",
			"how do I make changes to my shipping address ?",
			"i get an error message when i ty to make a payment for my order",
			"i want to request an invoice",
			"where do i check the delivery options ?",
			"you arent helping",
		}},
		{args: args{f: "./data/tests/test10.csv"}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = tk.normalize(t.text)
			}
			tr := CollectTransitions(tx, TokenSplit(tk))
			for i, t := range tx {
				tokens := tk.tokenize(t.text)
				tx[i].chunk = TransitionChunk(tokens, tokens, tr, 0.1)
			}
			ng := CollectChunks(tx)
			slices.Sort(ng)
			assert.Equal(t, tt.want, ng)
		})
	}
}
