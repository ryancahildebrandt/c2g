// -*- coding: utf-8 -*-

// Created on Sat Oct  4 09:34:52 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/stat/combin"
)

func TestGrammarE2E(t *testing.T) {
	var mergers = []func(r []Rule, e EqualityFunction) []Rule{MergePR, MergePS, MergeRS, MergeMisc}
	var permutations [][]int
	for i := range 4 {
		permutations = append(permutations, combin.Permutations(4, i+1)...)
	}

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{f: "./data/tests/test1.csv"}},
		{name: "", args: args{f: "./data/tests/test2.csv"}},
		{name: "", args: args{f: "./data/tests/test3.csv"}},
		{name: "", args: args{f: "./data/tests/test4.csv"}},
		{name: "", args: args{f: "./data/tests/test5.csv"}},
		{name: "", args: args{f: "./data/tests/test6.csv"}},
		{name: "", args: args{f: "./data/tests/test7.csv"}},
		{name: "", args: args{f: "./data/tests/test8.csv"}},
		{name: "", args: args{f: "./data/tests/test9.csv"}},
		{name: "", args: args{f: "./data/tests/test10.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				t.text = tk.normalize(t.text)
				tx[i] = t
			}
			got := []string{}
			for _, ttx := range tx {
				got = append(got, ttx.text)
			}
			for i := range got {
				got[i] = strings.TrimSpace(got[i])
				got[i] = strings.ReplaceAll(got[i], "  ", " ")
			}
			got = slices.DeleteFunc(got, func(i string) bool { return i == "" })
			slices.Sort(got)
			got = slices.Compact(got)
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
			for _, p := range permutations {
				rules1 := rules
				for _, pp := range p {
					rules1 = mergers[pp](rules, LiteralEqual(nilLogger))
				}
				rules1 = SetIDs(rules1)
				rules1 = Factor(rules1, 5)
				g := Grammar{Rules: rules1}
				os.WriteFile("./data/tests/out/tmp.jsgf", []byte(g.body()), 0644)
				exec.Command("./data/tests/out/gsgf", "generate", "./data/tests/out/tmp.jsgf", "-o", "./data/tests/out/out.txt").Run()
				file, _ = os.Open("./data/tests/out/out.txt")
				defer file.Close()
				s = bufio.NewScanner(file)
				txx := ReadTexts(s)
				for i, t := range txx {
					t.text = tk.normalize(t.text)
					txx[i] = t
				}
				want := []string{}
				for _, ttx := range txx {
					want = append(want, ttx.text)
				}
				for i := range want {
					want[i] = strings.TrimSpace(want[i])
					want[i] = strings.ReplaceAll(want[i], "  ", " ")
				}
				want = slices.DeleteFunc(want, func(i string) bool { return i == "" })
				slices.Sort(want)
				want = slices.Compact(want)
				assert.ElementsMatch(t, got, want)
			}
		})
	}
}

func TestGrammarMainE2E(t *testing.T) {
	var mergers = []func(r []Rule, e EqualityFunction) []Rule{MergePR, MergePS, MergeRS, MergeMisc}
	var permutations [][]int
	for i := range 4 {
		permutations = append(permutations, combin.Permutations(4, i+1)...)
	}

	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "", args: args{f: "./data/tests/test1.csv"}},
		{name: "", args: args{f: "./data/tests/test2.csv"}},
		{name: "", args: args{f: "./data/tests/test3.csv"}},
		{name: "", args: args{f: "./data/tests/test4.csv"}},
		{name: "", args: args{f: "./data/tests/test5.csv"}},
		{name: "", args: args{f: "./data/tests/test6.csv"}},
		{name: "", args: args{f: "./data/tests/test7.csv"}},
		{name: "", args: args{f: "./data/tests/test8.csv"}},
		{name: "", args: args{f: "./data/tests/test9.csv"}},
		{name: "", args: args{f: "./data/tests/test10.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				t.text = tk.normalize(t.text)
				tx[i] = t
			}
			got := []string{}
			for _, ttx := range tx {
				got = append(got, ttx.text)
			}
			for i := range got {
				got[i] = strings.TrimSpace(got[i])
				got[i] = strings.ReplaceAll(got[i], "  ", " ")
			}
			got = slices.DeleteFunc(got, func(i string) bool { return i == "" })
			slices.Sort(got)
			got = slices.Compact(got)
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
			for _, p := range permutations {
				rules1 := rules
				for _, pp := range p {
					rules1 = mergers[pp](rules, LiteralEqual(nilLogger))
				}
				rules1 = SetIDs(rules1)
				rules1 = Factor(rules1, 5)
				g := Grammar{Rules: rules1}
				os.WriteFile("./data/tests/out/tmpMain.jsgf", []byte(g.bodyMain()), 0644)
				exec.Command("./data/tests/out/gsgf", "generate", "./data/tests/out/tmpMain.jsgf", "-o", "./data/tests/out/outMain.txt").Run()
				file, _ = os.Open("./data/tests/out/outMain.txt")
				defer file.Close()
				s = bufio.NewScanner(file)
				txx := ReadTexts(s)
				for i, t := range txx {
					t.text = tk.normalize(t.text)
					txx[i] = t
				}
				want := []string{}
				for _, ttx := range txx {
					want = append(want, ttx.text)
				}
				for i := range want {
					want[i] = strings.TrimSpace(want[i])
					want[i] = strings.ReplaceAll(want[i], "  ", " ")
				}
				want = slices.DeleteFunc(want, func(i string) bool { return i == "" })
				slices.Sort(want)
				want = slices.Compact(want)
				assert.ElementsMatch(t, got, want)
			}
		})
	}
}
