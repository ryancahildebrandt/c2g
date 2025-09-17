// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTexts(t *testing.T) {
	var emp []Expression
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Expression
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Expression{"I don't have an online account"}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Expression{
			"I don't have an online account",
			"I don't understand you",
			"I got an error message when I attempted to make a payment",
			"can you tell me how I can get some bills?",
			"can you show me information about the status of my refund?",
			"i dont want my profile",
			"can you show me my invoices?",
			"i want to know wat the email of Customer Service is",
			"ask an agent to notify issues with my payment",
			"I want an online accoynt",
			"where can i leave an opinion for a service?",
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Expression{" ", "   "}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Expression{
			"I don't have an online account",
			"I want to make a review for a service",
			"where do i check the delivery options?",
			"you arent helping",
			"how do I make changes to my shipping address?",
			"I have a question",
			"i want to request an invoice",
			"I want to download a bill",
			"i get an error message when i ty to make a payment for my order",
			"I ordered an item and Id like to modify my fucking order",
			"I want to know what the number of Customer Service is",
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: emp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			assert.Equal(t, tt.want, ReadTexts(s))
		})
	}
}

func TestToNgrams(t *testing.T) {
	var emp []Ngram
	var tk = NewUnigramTokenizer()

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Ngram
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Ngram{{text: "I don't have an online account", len: 5, count: 1}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Ngram{
			{text: "I don't have an online account", len: 5, count: 1},
			{text: "I don't understand you", len: 3, count: 1},
			{text: "I got an error message when I attempted to make a payment", len: 11, count: 1},
			{text: "I want an online accoynt", len: 4, count: 1},
			{text: "ask an agent to notify issues with my payment", len: 8, count: 1},
			{text: "can you show me information about the status of my refund ?", len: 11, count: 1},
			{text: "can you show me my invoices ?", len: 6, count: 1},
			{text: "can you tell me how I can get some bills ?", len: 10, count: 1},
			{text: "i dont want my profile", len: 4, count: 1},
			{text: "i want to know wat the email of Customer Service is", len: 10, count: 1},
			{text: "where can i leave an opinion for a service ?", len: 9, count: 1},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Ngram{
			{text: "I don't have an online account", len: 5, count: 1},
			{text: "I have a question", len: 3, count: 1},
			{text: "I ordered an item and Id like to modify my fucking order", len: 11, count: 1},
			{text: "I want to download a bill", len: 5, count: 1},
			{text: "I want to know what the number of Customer Service is", len: 10, count: 1},
			{text: "I want to make a review for a service", len: 8, count: 1},
			{text: "how do I make changes to my shipping address ?", len: 9, count: 1},
			{text: "i get an error message when i ty to make a payment for my order", len: 14, count: 1},
			{text: "i want to request an invoice", len: 5, count: 1},
			{text: "where do i check the delivery options ?", len: 7, count: 1},
			{text: "you arent helping", len: 2, count: 1},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: emp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			c := NewCorpus(tx)
			tr := NewTransitions(c, tk)
			res := ToNgrams(c.texts, tk, tr, 0.1)
			slices.SortStableFunc(res, func(i, j Ngram) int {
				return strings.Compare(i.text, j.text)
			})
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSplitTriplets(t *testing.T) {
	var emps []Sentence
	var empt Triplet
	var tk = NewUnigramTokenizer()

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Sentence
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Sentence{
			{text: "I don't have an online account", pre: "", root: "I don't have an online account", suf: "", seg: Triplet{"I don't have an online account"}},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Sentence{
			{text: "I don't have an online account", pre: "", root: "I don't have an online account", suf: "", seg: Triplet{"I don't have an online account"}},
			{text: "I don't understand you", pre: "", root: "I don't understand you", suf: "", seg: Triplet{"I don't understand you"}},
			{text: "I got an error message when I attempted to make a payment", pre: "", root: "I got an error message when I attempted to make a payment", suf: "", seg: Triplet{"I got an error message when I attempted to make a payment"}},
			{text: "I want an online accoynt", pre: "", root: "I want an online accoynt", suf: "", seg: Triplet{"I want an online accoynt"}},
			{text: "ask an agent to notify issues with my payment", pre: "", root: "ask an agent to notify issues with my payment", suf: "", seg: Triplet{"ask an agent to notify issues with my payment"}},
			{text: "can you show me information about the status of my refund ?", pre: "", root: "can you show me information about the status of my refund ?", suf: "", seg: Triplet{"can you show me information about the status of my refund ?"}},
			{text: "can you show me my invoices ?", pre: "", root: "can you show me my invoices ?", suf: "", seg: Triplet{"can you show me my invoices ?"}},
			{text: "can you tell me how I can get some bills ?", pre: "", root: "can you tell me how I can get some bills ?", suf: "", seg: Triplet{"can you tell me how I can get some bills ?"}},
			{text: "i dont want my profile", pre: "", root: "i dont want my profile", suf: "", seg: Triplet{"i dont want my profile"}},
			{text: "i want to know wat the email of Customer Service is", pre: "", root: "i want to know wat the email of Customer Service is", suf: "", seg: Triplet{"i want to know wat the email of Customer Service is"}},
			{text: "where can i leave an opinion for a service ?", pre: "", root: "where can i leave an opinion for a service ?", suf: "", seg: Triplet{"where can i leave an opinion for a service ?"}},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Sentence{
			{text: "", pre: "", root: "", suf: "", seg: empt},
			{text: "", pre: "", root: "", suf: "", seg: empt},
		}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emps},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Sentence{
			{text: "I don't have an online account", pre: "", root: "I don't have an online account", suf: "", seg: Triplet{"I don't have an online account"}},
			{text: "I have a question", pre: "", root: "I have a question", suf: "", seg: Triplet{"I have a question"}},
			{text: "I ordered an item and Id like to modify my fucking order", pre: "", root: "I ordered an item and Id like to modify my fucking order", suf: "", seg: Triplet{"I ordered an item and Id like to modify my fucking order"}},
			{text: "I want to download a bill", pre: "", root: "I want to download a bill", suf: "", seg: Triplet{"I want to download a bill"}},
			{text: "I want to know what the number of Customer Service is", pre: "", root: "I want to know what the number of Customer Service is", suf: "", seg: Triplet{"I want to know what the number of Customer Service is"}},
			{text: "I want to make a review for a service", pre: "", root: "I want to make a review for a service", suf: "", seg: Triplet{"I want to make a review for a service"}},
			{text: "how do I make changes to my shipping address ?", pre: "", root: "how do I make changes to my shipping address ?", suf: "", seg: Triplet{"how do I make changes to my shipping address ?"}},
			{text: "i get an error message when i ty to make a payment for my order", pre: "", root: "i get an error message when i ty to make a payment for my order", suf: "", seg: Triplet{"i get an error message when i ty to make a payment for my order"}},
			{text: "i want to request an invoice", pre: "", root: "i want to request an invoice", suf: "", seg: Triplet{"i want to request an invoice"}},
			{text: "where do i check the delivery options ?", pre: "", root: "where do i check the delivery options ?", suf: "", seg: Triplet{"where do i check the delivery options ?"}},
			{text: "you arent helping", pre: "", root: "you arent helping", suf: "", seg: Triplet{"you arent helping"}},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: emps},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			c := NewCorpus(tx)
			c.transitions = NewTransitions(c, tk)
			c.ngrams = ToNgrams(c.texts, tk, c.transitions, 0.1)
			assert.Equal(t, tt.want, SplitTriplets(c.texts, c.ngrams))
		})
	}
}

func TestToRules(t *testing.T) {
	var emp []Rule
	var tk = NewUnigramTokenizer()

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{root: []Expression{"I don't have an online account"}, pre: []Expression{""}, suf: []Expression{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{root: []string{"I don't have an online account"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I don't understand you"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I got an error message when I attempted to make a payment"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I want an online accoynt"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"ask an agent to notify issues with my payment"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"can you show me information about the status of my refund ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"can you show me my invoices ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"can you tell me how I can get some bills ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"i dont want my profile"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"i want to know wat the email of Customer Service is"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"where can i leave an opinion for a service ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{root: []string{"I don't have an online account"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I have a question"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I ordered an item and Id like to modify my fucking order"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I want to download a bill"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I want to know what the number of Customer Service is"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"I want to make a review for a service"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"how do I make changes to my shipping address ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"i get an error message when i ty to make a payment for my order"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"i want to request an invoice"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"where do i check the delivery options ?"}, pre: []string{""}, suf: []string{""}, isPublic: true},
			{root: []string{"you arent helping"}, pre: []string{""}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: emp},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			c := NewCorpus(tx)
			c.transitions = NewTransitions(c, tk)
			c.ngrams = ToNgrams(c.texts, tk, c.transitions, 0.1)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			res := ToRules(c.texts)
			assert.Equal(t, tt.want, res)
		})
	}
}
