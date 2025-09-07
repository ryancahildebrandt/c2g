// -*- coding: utf-8 -*-

// Created on Sun Jul 27 07:52:09 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fmtGroup(t *testing.T) {
	type args struct {
		g []Expression
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{g: []Expression{}}, want: ""},
		{name: "", args: args{g: []Expression{""}}, want: ""},
		{name: "", args: args{g: []Expression{"", "", "", ""}}, want: "[||]"},
		{name: "", args: args{g: []Expression{" ", " ", " ", ""}}, want: "[ | | ]"},
		{name: "", args: args{g: []Expression{"a", "b", "c", ""}}, want: "[a|b|c]"},
		{name: "", args: args{g: []Expression{"a", "b", "c", "d"}}, want: "(a|b|c|d)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fmtGroup(tt.args.g))
		})
	}
}

func TestRule_print(t *testing.T) {
	type fields struct {
		root []Expression
		pre  []Expression
		suf  []Expression
	}
	type args struct {
		n string
		p bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "", fields: fields{pre: []Expression{}, root: []Expression{}, suf: []Expression{}}, args: args{n: "", p: false}, want: ""},
		// {name: "", fields: fields{pre: []Expression{}, root: []Expression{}, suf: []Expression{}}, args: args{n: "", p: true}, want: ""},
		// {name: "", fields: fields{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}, args: args{n: "", p: true}, want: ""},

		{name: "", fields: fields{pre: []Expression{"", "", "", ""}, root: []Expression{"", "", "", ""}, suf: []Expression{"", "", "", ""}}, args: args{n: "", p: true}, want: "public <> = [||] [||] [||]"},
		{name: "", fields: fields{pre: []Expression{"a", "b", "c", ""}, root: []Expression{"a", "b", "c", "d"}, suf: []Expression{}}, args: args{n: "1", p: false}, want: "<1> = [a|b|c] (a|b|c|d)"},
		{name: "", fields: fields{pre: []Expression{}, root: []Expression{"a", "b", "c", ""}, suf: []Expression{"a", "b", "c", "d"}}, args: args{n: "2", p: true}, want: "public <2> = [a|b|c] (a|b|c|d)"},
		{name: "", fields: fields{pre: []Expression{"a", "b", "c", "d"}, root: []Expression{}, suf: []Expression{"a", "b", "c", ""}}, args: args{n: " ", p: false}, want: "< > = (a|b|c|d) [a|b|c]"},
		{name: "", fields: fields{pre: []Expression{}, root: []Expression{"a", "b", "c", "d"}, suf: []Expression{}}, args: args{n: "  ", p: true}, want: "public <  > = (a|b|c|d)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				root: tt.fields.root,
				pre:  tt.fields.pre,
				suf:  tt.fields.suf,
			}
			assert.Equal(t, tt.want, r.print(tt.args.n, tt.args.p))
		})
	}
}

func TestPRSort(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I got an error message when I attempted to make a payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"ask an agent to notify issues with my payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me my invoices ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i dont want my profile"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where can i leave an opinion for a service ?"}, suf: []Expression{""}},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I have a question"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I ordered an item and Id like to modify my fucking order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to download a bill"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to know what the number of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to make a review for a service"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"how do I make changes to my shipping address ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i get an error message when i ty to make a payment for my order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to request an invoice"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where do i check the delivery options ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"you arent helping"}, suf: []Expression{""}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			res := ToRules(c.texts)
			PRSort(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestPSSort(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I got an error message when I attempted to make a payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"ask an agent to notify issues with my payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me my invoices ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i dont want my profile"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where can i leave an opinion for a service ?"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I have a question"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I ordered an item and Id like to modify my fucking order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to download a bill"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to know what the number of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to make a review for a service"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"how do I make changes to my shipping address ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i get an error message when i ty to make a payment for my order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to request an invoice"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where do i check the delivery options ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"you arent helping"}, suf: []Expression{""}}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			res := ToRules(c.texts)
			PSSort(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestRSSort(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I got an error message when I attempted to make a payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"ask an agent to notify issues with my payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me my invoices ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i dont want my profile"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where can i leave an opinion for a service ?"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I have a question"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I ordered an item and Id like to modify my fucking order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to download a bill"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to know what the number of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to make a review for a service"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"how do I make changes to my shipping address ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i get an error message when i ty to make a payment for my order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to request an invoice"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where do i check the delivery options ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"you arent helping"}, suf: []Expression{""}}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			res := ToRules(c.texts)
			RSSort(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestPRSSort(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I got an error message when I attempted to make a payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"ask an agent to notify issues with my payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me my invoices ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i dont want my profile"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where can i leave an opinion for a service ?"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I have a question"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I ordered an item and Id like to modify my fucking order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to download a bill"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to know what the number of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to make a review for a service"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"how do I make changes to my shipping address ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i get an error message when i ty to make a payment for my order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to request an invoice"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where do i check the delivery options ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"you arent helping"}, suf: []Expression{""}}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			res := ToRules(c.texts)
			PRSSort(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSSDMerger_check(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SSDMerger{}
			assert.Equal(t, tt.want, m.check(tt.args.r1, tt.args.r2))
		})
	}
}

func TestSSDMerger_merge(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want Rule
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{}, root: []string{}, suf: []string{}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c", "d"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c", "d"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a", "b"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SSDMerger{}
			assert.Equal(t, tt.want, m.merge(tt.args.r1, tt.args.r2))
		})
	}
}

func TestSSDMerger_apply(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I have a question"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I ordered an item and Id like to modify my fucking order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to download a bill"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to know what the number of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want to make a review for a service"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"how do I make changes to my shipping address ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i get an error message when i ty to make a payment for my order"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to request an invoice"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where do i check the delivery options ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"you arent helping"}, suf: []Expression{""}}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			rules := ToRules(c.texts)
			m := &SSDMerger{}
			assert.Equal(t, tt.want, m.apply(rules))
		})
	}
}

func TestSDSMerger_check(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SDSMerger{}
			assert.Equal(t, tt.want, m.check(tt.args.r1, tt.args.r2))
		})
	}
}

func TestSDSMerger_merge(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want Rule
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{}, root: []string{}, suf: []string{}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a", "b"}, suf: []Expression{"c"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"c"}, root: []Expression{"a", "b"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a", "b"}, suf: []Expression{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SDSMerger{}
			assert.Equal(t, tt.want, m.merge(tt.args.r1, tt.args.r2))
		})
	}
}

func TestSDSMerger_apply(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I don't understand you", "I got an error message when I attempted to make a payment", "I want an online accoynt", "ask an agent to notify issues with my payment", "can you show me information about the status of my refund ?", "can you show me my invoices ?", "can you tell me how I can get some bills ?", "i dont want my profile", "i want to know wat the email of Customer Service is", "where can i leave an opinion for a service ?"}, suf: []string{""}},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I have a question", "I ordered an item and Id like to modify my fucking order", "I want to download a bill", "I want to know what the number of Customer Service is", "I want to make a review for a service", "how do I make changes to my shipping address ?", "i get an error message when i ty to make a payment for my order", "i want to request an invoice", "where do i check the delivery options ?", "you arent helping"}, suf: []string{""}}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			rules := ToRules(c.texts)
			m := &SDSMerger{}
			assert.Equal(t, tt.want, m.apply(rules))
		})
	}
}

func TestDSSMerger_check(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: true},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DSSMerger{}
			assert.Equal(t, tt.want, m.check(tt.args.r1, tt.args.r2))
		})
	}
}

func TestDSSMerger_merge(t *testing.T) {
	type args struct {
		r1 Rule
		r2 Rule
	}
	tests := []struct {
		name string
		args args
		want Rule
	}{
		{name: "", args: args{
			r1: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{}, root: []string{}, suf: []string{}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			r2: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		},
			want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []string{"a", "b"}, root: []string{"1", "2"}, suf: []string{"c"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"d"}},
		},
			want: Rule{pre: []string{"1", "2"}, root: []string{"a"}, suf: []string{"c"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"d"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []string{"c", "d"}, root: []string{"a"}, suf: []string{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"b"}},
		},
			want: Rule{pre: []string{"1", "2"}, root: []string{"1", "2"}, suf: []string{"a"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"b"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []string{"a", "b"}, root: []string{"1", "2"}, suf: []string{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"b"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []string{"1", "2"}, root: []string{"a"}, suf: []string{"1", "2"}}},
		{name: "", args: args{
			r1: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
			r2: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		},
			want: Rule{pre: []string{"1", "2"}, root: []string{"1", "2"}, suf: []string{"1", "2"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DSSMerger{}
			assert.Equal(t, tt.want, m.merge(tt.args.r1, tt.args.r2))
		})
	}
}

func TestDSSMerger_apply(t *testing.T) {
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
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't have an online account"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I got an error message when I attempted to make a payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"ask an agent to notify issues with my payment"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me my invoices ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i dont want my profile"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"where can i leave an opinion for a service ?"}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			rules := ToRules(c.texts)
			m := &DSSMerger{}
			assert.Equal(t, tt.want, m.apply(rules))
		})
	}
}

func TestSSSMerger_check(t *testing.T) {
	type args struct {
		r Rule
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{
			r: Rule{pre: []Expression{}, root: []Expression{}, suf: []Expression{}},
		}, want: true},
		{name: "", args: args{
			r: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}, want: true},
		{name: "", args: args{
			r: Rule{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
		}, want: true},
		{name: "", args: args{
			r: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"c"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"c"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"c"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"a"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"a"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"1", "2"}, root: []Expression{"a"}, suf: []Expression{"1", "2"}},
		}, want: false},
		{name: "", args: args{
			r: Rule{pre: []Expression{"1", "2"}, root: []Expression{"1", "2"}, suf: []Expression{"1", "2"}},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SSSMerger{}
			assert.Equal(t, tt.want, m.check(tt.args.r))
		})
	}
}

func TestSSSMerger_merge(t *testing.T) {
	var emp Rule
	var tk = NewUnigramTokenizer()

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: Rule{pre: []string{}, root: []string{" I don't have an online account "}, suf: []string{}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: Rule{pre: []string{}, root: []string{" I don't have an online account ", " I don't understand you ", " I got an error message when I attempted to make a payment ", " I want an online accoynt ", " ask an agent to notify issues with my payment ", " can you show me information about the status of my refund ? ", " can you show me my invoices ? ", " can you tell me how I can get some bills ? ", " i dont want my profile ", " i want to know wat the email of Customer Service is ", " where can i leave an opinion for a service ? "}, suf: []string{}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: Rule{pre: []string{}, root: []string{"  "}, suf: []string{}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: Rule{pre: []string{}, root: []string{" I don't have an online account ", " I have a question ", " I ordered an item and Id like to modify my fucking order ", " I want to download a bill ", " I want to know what the number of Customer Service is ", " I want to make a review for a service ", " how do I make changes to my shipping address ? ", " i get an error message when i ty to make a payment for my order ", " i want to request an invoice ", " where do i check the delivery options ? ", " you arent helping "}, suf: []string{}}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			rules := ToRules(c.texts)
			m := &SSSMerger{}
			assert.Equal(t, tt.want, m.merge(rules...))
		})
	}
}

func TestSSSMerger_apply(t *testing.T) {
	var emp []Rule
	var tk = NewUnigramTokenizer()

	type args struct {
		f string
	}
	tests := []struct {
		name string
		m    *SSSMerger
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []Expression{}, root: []Expression{" I don't have an online account "}, suf: []Expression{}}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{"I don't understand you"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"I want an online accoynt"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you show me information about the status of my refund ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"can you tell me how I can get some bills ?"}, suf: []Expression{""}},
			{pre: []Expression{""}, root: []Expression{"i want to know wat the email of Customer Service is"}, suf: []Expression{""}},
			{pre: []Expression{}, root: []Expression{" I don't have an online account ", " I got an error message when I attempted to make a payment ", " ask an agent to notify issues with my payment ", " can you show me my invoices ? ", " i dont want my profile ", " where can i leave an opinion for a service ? "}, suf: []Expression{}}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{
			{pre: []Expression{""}, root: []Expression{""}, suf: []Expression{""}},
			{pre: []Expression{}, root: []Expression{"  "}, suf: []Expression{}}}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: emp},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}},
			{pre: []string{}, root: []string{" I don't have an online account ", " I ordered an item and Id like to modify my fucking order ", " I want to know what the number of Customer Service is ", " how do I make changes to my shipping address ? ", " i want to request an invoice ", " you arent helping "}, suf: []string{}},
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
			c.ngrams = ToNgrams(c.texts, tk, c.transitions)
			c.texts = SplitTriplets(c.texts, c.ngrams)
			rules := ToRules(c.texts)
			m := &SSSMerger{}
			assert.Equal(t, tt.want, m.apply(rules))
		})
	}
}
