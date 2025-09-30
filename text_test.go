// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadTexts(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []Text
	}{
		{name: "", args: args{f: "./data/tests/test5.csv"}, want: []Text{{pre: "", root: "", suf: "", chunk: []string{}, text: "I don't have an online account"}}},
		{name: "", args: args{f: "./data/tests/test6.csv"}, want: []Text{
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I don't have an online account"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I don't understand you"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I got an error message when I attempted to make a payment"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "can you tell me how I can get some bills?"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "can you show me information about the status of my refund?"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "i dont want my profile"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "can you show me my invoices?"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "i want to know wat the email of Customer Service is"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "ask an agent to notify issues with my payment"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I want an online accoynt"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "where can i leave an opinion for a service?"},
		}},
		{name: "", args: args{f: "./data/tests/test7.csv"}, want: []Text{}},
		{name: "", args: args{f: "./data/tests/test8.csv"}, want: []Text{}},
		{name: "", args: args{f: "./data/tests/test9.csv"}, want: []Text{
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I don't have an online account"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I want to make a review for a service"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "where do i check the delivery options?"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "you arent helping"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "how do I make changes to my shipping address?"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I have a question"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "i want to request an invoice"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I want to download a bill"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "i get an error message when i ty to make a payment for my order"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I ordered an item and Id like to modify my fucking order"},
			{pre: "", root: "", suf: "", chunk: []string{}, text: "I want to know what the number of Customer Service is"},
		}},
		{name: "", args: args{f: "./data/tests/test10.csv"}, want: []Text{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			assert.ElementsMatch(t, tt.want, ReadTexts(s))
		})
	}
}

func TestToTriplet(t *testing.T) {
	type args struct {
		t Text
		n []string
	}
	tests := []struct {
		name string
		args args
		want Text
	}{
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{}, text: ""}, n: []string{}}, want: Text{pre: "", root: "", suf: "", chunk: []string{}, text: ""}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{}, text: " "}, n: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order", "i get an error message when", "i ty to make a payment", "for my order", "i get an error", "message when i ty to", "make a payment for my order", "i get", "an error message", "when i ty to make a payment for my order", "i get an", "error message when", "i ty to make a payment for my order", "i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}, want: Text{pre: "", root: " ", suf: "", chunk: []string{}, text: " "}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{}}, want: Text{pre: "", root: "i get an error message when i ty to make a payment for my order", suf: "", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order"}}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i get an error message when", "i ty to make a payment", "for my order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order", "i get an error message when", "i ty to make a payment", "for my order", "i get an error", "message when i ty to", "make a payment for my order", "i get", "an error message", "when i ty to make a payment for my order", "i get an", "error message when", "i ty to make a payment for my order", "i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}, want: Text{pre: "i get an error message when", root: "i ty to make a payment", suf: "for my order", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i get an error message when", "i ty to make a payment", "for my order"}}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i get an error", "message when i ty to", "make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order", "i get an error message when", "i ty to make a payment", "for my order", "i get an error", "message when i ty to", "make a payment for my order", "i get", "an error message", "when i ty to make a payment for my order", "i get an", "error message when", "i ty to make a payment for my order", "i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}, want: Text{pre: "", root: "i get an error", suf: "message when i ty to make a payment for my order", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i get an error", "message when i ty to", "make a payment for my order"}}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i get", "an error message", "when i ty to make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order", "i get an error message when", "i ty to make a payment", "for my order", "i get an error", "message when i ty to", "make a payment for my order", "i get", "an error message", "when i ty to make a payment for my order", "i get an", "error message when", "i ty to make a payment for my order", "i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}, want: Text{pre: "", root: "i get", suf: "an error message when i ty to make a payment for my order", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i get", "an error message", "when i ty to make a payment for my order"}}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i get an", "error message when", "i ty to make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order", "i get an error message when", "i ty to make a payment", "for my order", "i get an error", "message when i ty to", "make a payment for my order", "i get", "an error message", "when i ty to make a payment for my order", "i get an", "error message when", "i ty to make a payment for my order", "i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}, want: Text{pre: "", root: "i get an", suf: "error message when i ty to make a payment for my order", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i get an", "error message when", "i ty to make a payment for my order"}}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{"i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}, text: "i get an error message when i ty to make a payment for my order"}, n: []string{}}, want: Text{pre: "", root: "i get an error message when i ty to make a payment for my order", suf: "", text: "i get an error message when i ty to make a payment for my order", chunk: []string{"i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToTriplet(tt.args.t, tt.args.n))
		})
	}
}

func TestToRule(t *testing.T) {
	type args struct {
		t Text
	}
	tests := []struct {
		name string
		args args
		want Rule
	}{
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{}, text: ""}}, want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "", root: "", suf: "", chunk: []string{}, text: " "}}, want: Rule{pre: []string{""}, root: []string{""}, suf: []string{""}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "i", root: "get an error", suf: "message when i ty to make a payment for my order", chunk: []string{"i", "get an error", "message when", "i ty to make a payment", "for my order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{"i"}, root: []string{"get an error"}, suf: []string{"message when i ty to make a payment for my order"}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "i get an error message when", root: "i ty to make a payment", suf: "for my order", chunk: []string{"i get an error message when", "i ty to make a payment", "for my order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{"i get an error message when"}, root: []string{"i ty to make a payment"}, suf: []string{"for my order"}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "i get an error", root: "message when i ty to make a payment", suf: "for my order", chunk: []string{"i get an error", "message when i ty to", "make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{"i get an error"}, root: []string{"message when i ty to make a payment"}, suf: []string{"for my order"}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "", root: "i get an error", suf: "message when i ty to make a payment for my order", chunk: []string{"i get", "an error message", "when i ty to make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{""}, root: []string{"i get an error"}, suf: []string{"message when i ty to make a payment for my order"}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "i get an error", root: "", suf: "message when i ty to make a payment for my order", chunk: []string{"i get an", "error message when", "i ty to make a payment for my order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{"i get an error"}, root: []string{""}, suf: []string{"message when i ty to make a payment for my order"}, isPublic: true, id: 0}},
		{name: "", args: args{t: Text{pre: "i", root: "get an error message when i ty to make a payment for my order", suf: "", chunk: []string{"i", "get", "an", "error", "message", "when", "i", "ty", "to", "make", "a", "payment", "for", "my", "order"}, text: "i get an error message when i ty to make a payment for my order"}}, want: Rule{pre: []string{"i"}, root: []string{"get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ToRule(tt.args.t))
		})
	}
}
