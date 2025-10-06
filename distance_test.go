// -*- coding: utf-8 -*-

// Created on Tue Oct  7 06:53:46 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestCharacterLevenshtein(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		
		args args
		want float64
	}{
		{args: args{s1: "", s2: ""}, want: 1.0},
		{args: args{s1: " ", s2: " "}, want: 1.0},
		{args: args{s1: " ", s2: ""}, want: 0.0},
		{args: args{s1: "", s2: " "}, want: 0.0},
		{args: args{s1: "a.b", s2: "a. bc"}, want: 0.6},
		{args: args{s1: "a.b", s2: "a . bc"}, want: 0.5},
		{args: args{s1: "a.b", s2: " a. bc"}, want: 0.5},
		{args: args{s1: "this is a test", s2: "this is a test"}, want: 1.0},
		{args: args{s1: "this is a test", s2: "this is a testt"}, want: 14.0 / 15.0},
		{args: args{s1: "test", s2: "this is a testting"}, want: 4.0 / 18.0},
		{args: args{s1: "test", s2: "this is a test"}, want: 4.0 / 14.0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, CharacterLevenshtein(tt.args.s1, tt.args.s2))
		})
	}
}

func TestTokenLevenshtein(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		
		args args
		want float64
	}{
		{args: args{s1: []string{""}, s2: []string{""}}, want: 1.0},
		{args: args{s1: []string{" "}, s2: []string{" "}}, want: 1.0},
		{args: args{s1: []string{" "}, s2: []string{""}}, want: 0.0},
		{args: args{s1: []string{""}, s2: []string{" "}}, want: 0.0},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{"a", ".", " ", "b", "c"}}, want: 0.6},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{"a", " ", ".", " ", "b", "c"}}, want: 0.5},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{" ", "a", ".", " ", "b", "c"}}, want: 0.5},
		{args: args{s1: []string{"this", "is", "a", "test"}, s2: []string{"this", "is", "a", "test"}}, want: 1.0},
		{args: args{s1: []string{"this", "is", "a", "test"}, s2: []string{"this", "is", "a", "testt"}}, want: 0.75},
		{args: args{s1: []string{"test"}, s2: []string{"this is a testing"}}, want: 0.0},
		{args: args{s1: []string{"test"}, s2: []string{"this", "is", "a", "test"}}, want: 0.25},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, TokenLevenshtein(tt.args.s1, tt.args.s2))
		})
	}
}

func Test_levenshteinDistance(t *testing.T) {
	type args struct {
		s1 []string
		s2 []string
	}
	tests := []struct {
		
		args args
		want int
	}{
		{args: args{s1: []string{""}, s2: []string{""}}, want: 0},
		{args: args{s1: []string{" "}, s2: []string{" "}}, want: 0},
		{args: args{s1: []string{" "}, s2: []string{""}}, want: 1},
		{args: args{s1: []string{""}, s2: []string{" "}}, want: 1},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{"a", ".", " ", "b", "c"}}, want: 2},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{"a", " ", ".", " ", "b", "c"}}, want: 3},
		{args: args{s1: []string{"a", ".", "b"}, s2: []string{" ", "a", ".", " ", "b", "c"}}, want: 3},
		{args: args{s1: []string{"this", "is", "a", "test"}, s2: []string{"this", "is", "a", "test"}}, want: 0},
		{args: args{s1: []string{"this", "is", "a", "test"}, s2: []string{"this", "is", "a", "testt"}}, want: 1},
		{args: args{s1: []string{"test"}, s2: []string{"this is a testing"}}, want: 1},
		{args: args{s1: []string{"test"}, s2: []string{"this", "is", "a", "test"}}, want: 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, tt.want, levenshteinDistance(tt.args.s1, tt.args.s2))
		})
	}
}

func TestCollectVocab(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		
		args args
		want []string
	}{
		{args: args{f: "./data/tests/test5.csv"}, want: []string{"account", "an", "don't", "have", "i", "online"}},
		{args: args{f: "./data/tests/test6.csv"}, want: []string{"?", "a", "about", "account", "accoynt", "agent", "an", "ask", "attempted", "bills", "can", "customer", "don't", "dont", "email", "error", "for", "get", "got", "have", "how", "i", "information", "invoices", "is", "issues", "know", "leave", "make", "me", "message", "my", "notify", "of", "online", "opinion", "payment", "profile", "refund", "service", "show", "some", "status", "tell", "the", "to", "understand", "want", "wat", "when", "where", "with", "you"}},
		{args: args{f: "./data/tests/test7.csv"}, want: []string{}},
		{args: args{f: "./data/tests/test8.csv"}, want: []string{}},
		{args: args{f: "./data/tests/test9.csv"}, want: []string{"?", "a", "account", "address", "an", "and", "arent", "bill", "changes", "check", "customer", "delivery", "do", "don't", "download", "error", "for", "fucking", "get", "have", "helping", "how", "i", "id", "invoice", "is", "item", "know", "like", "make", "message", "modify", "my", "number", "of", "online", "options", "order", "ordered", "payment", "question", "request", "review", "service", "shipping", "the", "to", "ty", "want", "what", "when", "where", "you"}},
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
			assert.Equal(t, tt.want, CollectVocab(tx, tk))
		})
	}
}

func TestCollectIDF(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		
		args args
		want map[string]float64
	}{
		{args: args{f: "./data/tests/test5.csv"}, want: map[string]float64{"account": 0.69, "an": 0.69, "don't": 0.69, "have": 0.69, "i": 0.69, "online": 0.69}},
		{args: args{f: "./data/tests/test6.csv"}, want: map[string]float64{"?": 1.32, "a": 1.87, "about": 2.48, "account": 2.48, "accoynt": 2.48, "agent": 2.48, "an": 1.16, "ask": 2.48, "attempted": 2.48, "bills": 2.48, "can": 1.32, "customer": 2.48, "don't": 1.87, "dont": 2.48, "email": 2.48, "error": 2.48, "for": 2.48, "get": 2.48, "got": 2.48, "have": 2.48, "how": 2.48, "i": 0.86, "information": 2.48, "invoices": 2.48, "is": 2.48, "issues": 2.48, "know": 2.48, "leave": 2.48, "make": 2.48, "me": 1.54, "message": 2.48, "my": 1.32, "notify": 2.48, "of": 1.87, "online": 1.87, "opinion": 2.48, "payment": 1.87, "profile": 2.48, "refund": 2.48, "service": 1.87, "show": 1.87, "some": 2.48, "status": 2.48, "tell": 2.48, "the": 1.87, "to": 1.54, "understand": 2.48, "want": 1.54, "wat": 2.48, "when": 2.48, "where": 2.48, "with": 2.48, "you": 1.32}},
		{args: args{f: "./data/tests/test7.csv"}, want: map[string]float64{}},
		{args: args{f: "./data/tests/test8.csv"}, want: map[string]float64{}},
		{args: args{f: "./data/tests/test9.csv"}, want: map[string]float64{"?": 1.87, "a": 1.32, "account": 2.48, "address": 2.48, "an": 1.32, "and": 2.48, "arent": 2.48, "bill": 2.48, "changes": 2.48, "check": 2.48, "customer": 2.48, "delivery": 2.48, "do": 1.87, "don't": 2.48, "download": 2.48, "error": 2.48, "for": 1.87, "fucking": 2.48, "get": 2.48, "have": 1.87, "helping": 2.48, "how": 2.48, "i": 0.74, "id": 2.48, "invoice": 2.48, "is": 2.48, "item": 2.48, "know": 2.48, "like": 2.48, "make": 1.54, "message": 2.48, "modify": 2.48, "my": 1.54, "number": 2.48, "of": 2.48, "online": 2.48, "options": 2.48, "order": 1.87, "ordered": 2.48, "payment": 2.48, "question": 2.48, "request": 2.48, "review": 2.48, "service": 1.87, "shipping": 2.48, "the": 1.87, "to": 0.94, "ty": 2.48, "want": 1.32, "what": 2.48, "when": 2.48, "where": 2.48, "you": 2.48}},
		{args: args{f: "./data/tests/test10.csv"}, want: map[string]float64{}},
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
			res := CollectIDF(tx, tk)
			for k, v := range res {
				res[k] = math.Floor(v*100) / 100
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestCountEmbed(t *testing.T) {
	type args struct {
		s string
		v []string
	}
	tests := []struct {
		
		args      args
		want      []float64
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{s: "", v: []string{}}, want: []float64{}, assertion: assert.Error},
		{args: args{s: "test", v: []string{}}, want: []float64{}, assertion: assert.Error},
		{args: args{s: "testing", v: []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", "."}}, want: []float64{0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0}, assertion: assert.NoError},
		{args: args{s: "this is a test.", v: []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", "."}}, want: []float64{1.0, 1.0, 0.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0}, assertion: assert.NoError},
		{args: args{s: "that is not a not a test...", v: []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", "."}}, want: []float64{2.0, 1.0, 0.0, 0.0, 1.0, 2.0, 1.0, 0.0, 0.0, 3.0}, assertion: assert.NoError},
		{args: args{s: "this is a testt", v: []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", "."}}, want: []float64{}, assertion: assert.Error},
		{args: args{s: "just testing", v: []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", "."}}, want: []float64{}, assertion: assert.Error},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewWordTokenizer()
			v := mat.VecDense{}
			if len(tt.want) > 0 {
				v = *mat.NewVecDense(len(tt.want), tt.want)
			}
			got, err := CountEmbed(tt.args.s, tt.args.v, tk)
			tt.assertion(t, err)
			assert.Equal(t, v, got)
		})
	}
}

func TestCosineSimilarity(t *testing.T) {
	type args struct {
		v1 []float64
		v2 []float64
	}
	tests := []struct {
		
		args      args
		want      float64
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{v1: []float64{}, v2: []float64{}}, want: 1.0, assertion: assert.NoError},
		{args: args{v1: []float64{0.0}, v2: []float64{}}, want: 0.0, assertion: assert.Error},
		{args: args{v1: []float64{}, v2: []float64{0.0}}, want: 0.0, assertion: assert.Error},
		{args: args{v1: []float64{0.0, 0.0}, v2: []float64{0.0, 0.0}}, want: 1.0, assertion: assert.NoError},
		{args: args{v1: []float64{0.0, 0.0, 1.0, 0.0, 0.0}, v2: []float64{0.0, 0.0, 1.0, 0.0, 0.0}}, want: 1.0, assertion: assert.NoError},
		{args: args{v1: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, v2: []float64{6.0, 7.0, 8.0, 9.0, 10.0}}, want: 0.96, assertion: assert.NoError},
		{args: args{v1: []float64{1.0, 0.0, 0.0, 0.0, 0.0}, v2: []float64{0.0, 0.0, 0.0, 0.0, 0.0}}, want: 0.0, assertion: assert.NoError},
		{args: args{v1: []float64{2.0, 2.0, 2.0, 2.0, 2.0}, v2: []float64{1.0, 1.0, 1.0, 1.0, 1.0}}, want: 1.0, assertion: assert.NoError},
		{args: args{v1: []float64{1.0, 1.0, 1.0, 0.0, 0.0}, v2: []float64{-1.0, -1.0, -1.0, 0.0, 0.0}}, want: -1.0, assertion: assert.NoError},
		{args: args{v1: []float64{0.1, 0.2, 0.3, 0.4, 0.5}, v2: []float64{0.6, 0.7, 0.8, 0.9, 0.0}}, want: 0.71, assertion: assert.NoError},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			v1 := mat.VecDense{}
			if len(tt.args.v1) > 0 {
				v1 = *mat.NewVecDense(len(tt.args.v1), tt.args.v1)
			}
			v2 := mat.VecDense{}
			if len(tt.args.v2) > 0 {
				v2 = *mat.NewVecDense(len(tt.args.v2), tt.args.v2)
			}
			got, err := CosineSimilarity(v1, v2)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, math.Round(got*100)/100)
		})
	}
}

func TestTFIDFTransform(t *testing.T) {
	type args struct {
		v []float64
	}
	tests := []struct {
		
		args args
		want []float64
	}{
		{args: args{v: []float64{}}, want: []float64{}},
		{args: args{v: []float64{2.0, 1.0}}, want: []float64{1.33, 0.43}},
		{args: args{v: []float64{0.0, 1.0, 1.0, 1.0, 3.0, 1.0, 4.0, 1.0, 9.0, 1.0}}, want: []float64{0, 0.05, 0.07, 0.07, 0.52, 0.07, 1.22, 0.04, 3.68, 0.05}},
		{args: args{v: []float64{0.0, 2.0, 1.0, 1.0, 4.0, 1.0, 5.0, 2.0, 6.0, 1.0, 9.0}}, want: []float64{0, 0.16, 0.05, 0.05, 0.64, 0.05, 1.32, 0.12, 1.12, 0.03, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			vv := []string{"a", "test", "testing", "this", "is", "not", "that", "these", "tested", ".", ""}
			idf := map[string]float64{"a": 1.0, "test": 1.29, "testing": 1.69, "this": 1.69, "is": 1.29, "not": 1.69, "that": 1.69, "these": 1.0, "tested": 1.0, ".": 1.2}
			v1 := mat.VecDense{}
			if len(tt.args.v) > 0 {
				v1 = *mat.NewVecDense(len(tt.args.v), tt.args.v)
			}
			v2 := mat.VecDense{}
			if len(tt.want) > 0 {
				v2 = *mat.NewVecDense(len(tt.want), tt.want)
			}
			res := TFIDFTransform(v1, vv, idf)
			for i := range res.Len() {
				res.SetVec(i, math.Floor(res.AtVec(i)*100)/100)
			}
			assert.Equal(t, v2, res)
		})
	}
}
