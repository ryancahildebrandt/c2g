// -*- coding: utf-8 -*-

// Created on Mon Oct  6 04:09:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortPR(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			res := []Rule{}
			for _, t := range tx {
				res = append(res, ToRule(t))
			}
			SortPR(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSortPS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			res := []Rule{}
			for _, t := range tx {
				res = append(res, ToRule(t))
			}
			SortPS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSortRS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			res := []Rule{}
			for _, t := range tx {
				res = append(res, ToRule(t))
			}
			SortRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSortPRS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			res := []Rule{}
			for _, t := range tx {
				res = append(res, ToRule(t))
			}
			SortPRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergeP(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I don't understand you", "I got an error message when I attempted to make a payment", "I want an online accoynt", "ask an agent to notify issues with my payment", "can you show me information about the status of my refund ?", "can you show me my invoices ?", "can you tell me how I can get some bills ?", "i dont want my profile", "i want to know wat the email of Customer Service is", "where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 0}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I have a question", "I ordered an item and Id like to modify my fucking order", "I want to download a bill", "I want to know what the number of Customer Service is", "I want to make a review for a service", "how do I make changes to my shipping address ?", "i get an error message when i ty to make a payment for my order", "i want to request an invoice", "where do i check the delivery options ?", "you arent helping"}, suf: []string{""}, isPublic: true, id: 0}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergeP(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergeR(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergeR(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergeS(t *testing.T) {
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
			{pre: []string{""}, root: []string{"I don't have an online account", "I don't understand you", "I got an error message when I attempted to make a payment", "I want an online accoynt", "ask an agent to notify issues with my payment", "can you show me information about the status of my refund ?", "can you show me my invoices ?", "can you tell me how I can get some bills ?", "i dont want my profile", "i want to know wat the email of Customer Service is", "where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I have a question", "I ordered an item and Id like to modify my fucking order", "I want to download a bill", "I want to know what the number of Customer Service is", "I want to make a review for a service", "how do I make changes to my shipping address ?", "i get an error message when i ty to make a payment for my order", "i want to request an invoice", "where do i check the delivery options ?", "you arent helping"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergeS(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergePR(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergePR(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergePS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I don't understand you", "I got an error message when I attempted to make a payment", "I want an online accoynt", "ask an agent to notify issues with my payment", "can you show me information about the status of my refund ?", "can you show me my invoices ?", "can you tell me how I can get some bills ?", "i dont want my profile", "i want to know wat the email of Customer Service is", "where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account", "I have a question", "I ordered an item and Id like to modify my fucking order", "I want to download a bill", "I want to know what the number of Customer Service is", "I want to make a review for a service", "how do I make changes to my shipping address ?", "i get an error message when i ty to make a payment for my order", "i want to request an invoice", "where do i check the delivery options ?", "you arent helping"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergePS(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergeRS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergeRS(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestMergePRS(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Rule{
			{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true},
			{pre: []string{}, root: []string{"I don't have an online account", "I got an error message when I attempted to make a payment", "ask an agent to notify issues with my payment", "can you show me my invoices ?", "i dont want my profile", "where can i leave an opinion for a service ?"}, suf: []string{}, isPublic: true}}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Rule{
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true},
			{pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true},
			{pre: []string{}, root: []string{"I don't have an online account", "I ordered an item and Id like to modify my fucking order", "I want to know what the number of Customer Service is", "how do I make changes to my shipping address ?", "i want to request an invoice", "you arent helping"}, suf: []string{}, isPublic: true},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := MergePRS(rules)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFactor(t *testing.T) {
	type args struct {
		f  string
		ff int
	}
	tests := []struct {
		name string
		args args
		want []Rule
	}{
		{name: "", args: args{f: "./data/tests/test5.csv", ff: 0}, want: []Rule{
			{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{""}, root: []string{"<I_don't_have_an_onli>"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test5.csv", ff: 1}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test5.csv", ff: 10}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv", ff: 0}, want: []Rule{
			{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I don't understand you"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I want an online accoynt"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"can you show me my invoices?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"i dont want my profile"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{""}, root: []string{"<I_don't_have_an_onli>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_don't_understand_y>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_got_an_error_messa>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_want_an_online_acc>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<ask_an_agent_to_noti>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<can_you_show_me_info>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<can_you_show_me_my_i>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<can_you_tell_me_how_>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<i_dont_want_my_profi>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<i_want_to_know_wat_t>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<where_can_i_leave_an>"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv", ff: 1}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test6.csv", ff: 10}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv", ff: 0}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test8.csv", ff: 0}, want: []Rule{}},
		{name: "", args: args{f: "./data/tests/test9.csv", ff: 0}, want: []Rule{
			{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I have a question"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I want to download a bill"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"I want to make a review for a service"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"i want to request an invoice"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"where do i check the delivery options?"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{}, root: []string{"you arent helping"}, suf: []string{}, isPublic: false, id: 0},
			{pre: []string{""}, root: []string{"<I_don't_have_an_onli>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_have_a_question>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_ordered_an_item_an>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_want_to_download_a>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_want_to_know_what_>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<I_want_to_make_a_rev>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<how_do_I_make_change>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<i_get_an_error_messa>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<i_want_to_request_an>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<where_do_i_check_the>"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"<you_arent_helping>"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test9.csv", ff: 1}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test9.csv", ff: 10}, want: []Rule{
			{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 0},
			{pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 0},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv", ff: 0}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewUnigramTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = UnigramNormalize(t.text, tk)
			}
			tr := CollectTransitions(tx, tk)
			for i, t := range tx {
				tx[i].chunk = TransitionChunk(UnigramTokenize(t.text, tk), tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			res := Factor(rules, tt.args.ff)
			SortPRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}
