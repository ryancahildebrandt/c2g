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

func TestRule_print(t *testing.T) {
	type fields struct {
		root     []string
		pre      []string
		suf      []string
		isPublic bool
	}
	type args struct {
		n string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "", fields: fields{pre: []string{}, root: []string{}, suf: []string{}, isPublic: false}, args: args{n: ""}, want: ""},
		{name: "", fields: fields{pre: []string{}, root: []string{}, suf: []string{}, isPublic: true}, args: args{n: ""}, want: ""},
		{name: "", fields: fields{pre: []string{""}, root: []string{""}, suf: []string{""}, isPublic: true}, args: args{n: ""}, want: ""},
		{name: "", fields: fields{pre: []string{"", "", "", ""}, root: []string{"", "", "", ""}, suf: []string{"", "", "", ""}, isPublic: true}, args: args{n: ""}, want: "public <> = [||] [||] [||];"},
		{name: "", fields: fields{pre: []string{"a", "b", "c", ""}, root: []string{"a", "b", "c", "d"}, suf: []string{}, isPublic: false}, args: args{n: "1"}, want: "<1> = [a|b|c] (a|b|c|d);"},
		{name: "", fields: fields{pre: []string{}, root: []string{"a", "b", "c", ""}, suf: []string{"a", "b", "c", "d"}, isPublic: true}, args: args{n: "2"}, want: "public <2> = [a|b|c] (a|b|c|d);"},
		{name: "", fields: fields{pre: []string{"a", "b", "c", "d"}, root: []string{}, suf: []string{"a", "b", "c", ""}, isPublic: false}, args: args{n: " "}, want: "< > = (a|b|c|d) [a|b|c];"},
		{name: "", fields: fields{pre: []string{}, root: []string{"a", "b", "c", "d"}, suf: []string{}, isPublic: true}, args: args{n: "  "}, want: "public <  > = (a|b|c|d);"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				root:     tt.fields.root,
				pre:      tt.fields.pre,
				suf:      tt.fields.suf,
				isPublic: tt.fields.isPublic,
			}
			assert.Equal(t, tt.want, r.print(tt.args.n))
		})
	}
}

func TestRule_isEmpty(t *testing.T) {
	type fields struct {
		pre  []string
		root []string
		suf  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "", fields: fields{
			pre: []string{}, root: []string{}, suf: []string{},
		}, want: true},
		{name: "", fields: fields{
			pre: []string{""}, root: []string{}, suf: []string{},
		}, want: true},
		{name: "", fields: fields{
			pre: []string{}, root: []string{""}, suf: []string{},
		}, want: true},
		{name: "", fields: fields{
			pre: []string{}, root: []string{}, suf: []string{""},
		}, want: true},
		{name: "", fields: fields{
			pre: []string{""}, root: []string{""}, suf: []string{""},
		}, want: true},
		{name: "", fields: fields{
			pre: []string{"a"}, root: []string{"1", "2"}, suf: []string{"c"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"1", "2"}, root: []string{"a"}, suf: []string{"c"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"c"}, root: []string{"a"}, suf: []string{"1", "2"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"1", "2"}, root: []string{"1", "2"}, suf: []string{"a"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"a"}, root: []string{"1", "2"}, suf: []string{"1", "2"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"1", "2"}, root: []string{"a"}, suf: []string{"1", "2"},
		}, want: false},
		{name: "", fields: fields{
			pre: []string{"1", "2"}, root: []string{"1", "2"}, suf: []string{"1", "2"},
		}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rule{
				pre:  tt.fields.pre,
				root: tt.fields.root,
				suf:  tt.fields.suf,
			}
			assert.Equal(t, tt.want, r.isEmpty())
		})
	}
}

func TestRule_name(t *testing.T) {
	type args struct {
		r Rule
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{r: Rule{pre: []string{}, root: []string{}, suf: []string{}}}, want: ""},
		{name: "", args: args{r: Rule{pre: []string{}, root: []string{}, suf: []string{}}}, want: ""},
		{name: "", args: args{r: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}}}, want: ""},
		{name: "", args: args{r: Rule{pre: []string{"", "", "", ""}, root: []string{"", "", "", ""}, suf: []string{"", "", "", ""}}}, want: "___"},
		{name: "", args: args{r: Rule{pre: []string{"a", "b", "c", ""}, root: []string{"a", "b", "c", "d"}, suf: []string{}}}, want: "a_b_c_d"},
		{name: "", args: args{r: Rule{pre: []string{}, root: []string{"a", "b", "c", ""}, suf: []string{"a", "b", "c", "d"}}}, want: "a_b_c_"},
		{name: "", args: args{r: Rule{pre: []string{"a", "b", "c", "d"}, root: []string{}, suf: []string{"a", "b", "c", ""}}}, want: ""},
		{name: "", args: args{r: Rule{pre: []string{}, root: []string{"a", "b", "c", "d"}, suf: []string{}}}, want: "a_b_c_d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.args.r.name())
		})
	}
}

func TestSetIDs(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true, id: 5},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true, id: 6},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true, id: 7},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 10},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true, id: 6},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true, id: 9},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = tk.normalize(t.text)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(tk.tokenize(t.text), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := SetIDs(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}
