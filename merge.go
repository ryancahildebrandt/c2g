// -*- coding: utf-8 -*-

// Created on Mon Oct  6 04:09:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"slices"
	"strings"
)

func SortPR(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.root, r2.root)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func SortPS(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func SortRS(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.root, r2.root) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.root, r2.root)
	})
}

func SortPRS(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		switch {
		case slices.Equal(r1.pre, r2.pre) && slices.Equal(r1.root, r2.root):
			return slices.Compare(r1.suf, r2.suf)
		case slices.Equal(r1.pre, r2.pre):
			return slices.Compare(r1.root, r2.root)
		default:
			return slices.Compare(r1.pre, r2.pre)
		}
	})
}

func MergeP(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return slices.Equal(r1.pre, r2.pre) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.isPublic = true

		for _, i := range r1.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}
		for _, i := range r2.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}
		for _, i := range r1.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}
		for _, i := range r2.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}

		return r
	}

	var i int

	SortPRS(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeR(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return slices.Equal(r1.root, r2.root) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.root = r1.root
		r.isPublic = true

		for _, i := range r1.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}
		for _, i := range r2.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}
		for _, i := range r1.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}
		for _, i := range r2.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}

		return r
	}

	var i int

	SortPRS(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeS(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return slices.Equal(r1.suf, r2.suf) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.suf = r1.suf
		r.isPublic = true

		for _, i := range r1.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}
		for _, i := range r2.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}
		for _, i := range r1.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}
		for _, i := range r2.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}

		return r
	}

	var i int

	SortPRS(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergePR(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return slices.Equal(r1.pre, r2.pre) && slices.Equal(r1.root, r2.root)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.root = r1.root
		r.suf = []string{}
		r.isPublic = true

		for _, i := range r1.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}
		for _, i := range r2.suf {
			if !slices.Contains(r.suf, i) {
				r.suf = append(r.suf, i)
			}
		}

		return r
	}

	var i int

	SortPR(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergePS(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return slices.Equal(r1.pre, r2.pre) && slices.Equal(r1.suf, r2.suf)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.root = []string{}
		r.suf = r1.suf
		r.isPublic = true

		for _, i := range r1.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}
		for _, i := range r2.root {
			if !slices.Contains(r.root, i) {
				r.root = append(r.root, i)
			}
		}

		return r
	}

	var i int

	SortPS(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeRS(r []Rule) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return slices.Equal(r1.root, r2.root) && slices.Equal(r1.suf, r2.suf)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = []string{}
		r.root = r1.root
		r.suf = r1.suf
		r.isPublic = true

		for _, i := range r1.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}
		for _, i := range r2.pre {
			if !slices.Contains(r.pre, i) {
				r.pre = append(r.pre, i)
			}
		}

		return r
	}

	var i int

	SortRS(r)
	for i < len(r)-1 {
		if check(r[i], r[i+1]) {
			r[i] = merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergePRS(r []Rule) []Rule {
	check := func(r Rule) bool {
		return len(r.pre) <= 1 && len(r.root) <= 1 && len(r.suf) <= 1
	}
	merge := func(rules ...Rule) Rule {
		var (
			r Rule
			b strings.Builder
			s string
		)

		for _, rule := range rules {
			b.WriteString(strings.Join(rule.pre, " "))
			b.WriteString(" ")
			b.WriteString(strings.Join(rule.root, " "))
			b.WriteString(" ")
			b.WriteString(strings.Join(rule.suf, " "))
			s = strings.TrimSpace(b.String())
			if !slices.Contains(r.root, s) {
				r.root = append(r.root, s)
			}
			b.Reset()
		}
		if r.root != nil {
			r.pre = []string{}
			r.suf = []string{}
			r.isPublic = true
		}

		return r
	}

	var (
		rr  []Rule
		res Rule
		i   int
	)

	SortPRS(r)
	for i = 0; i < len(r); i++ {
		if check(r[i]) {
			if !r[i].isEmpty() {
				rr = append(rr, r[i])
			}
			r = slices.Delete(r, i, i+1)
		}
	}
	res = merge(rr...)
	if !res.isEmpty() {
		r = append(r, res)
	}

	return r
}

func Factor(rules []Rule, f int) []Rule {
	getCounts := func(r []Rule) map[string]int {
		var counts = make(map[string]int)

		slices.SortStableFunc(r, func(i, j Rule) int {
			return strings.Compare(i.print(""), j.print(""))
		})

		for _, rule := range r {
			slices.Sort(rule.pre)
			slices.Sort(rule.root)
			slices.Sort(rule.suf)
			counts[fmt.Sprint(strings.Join(rule.pre, "|"))]++
			counts[fmt.Sprint(strings.Join(rule.root, "|"))]++
			counts[fmt.Sprint(strings.Join(rule.suf, "|"))]++
		}

		return counts
	}

	getChunks := func(c map[string]int) []string {
		var chunks = []string{}

		for cc := range c {
			switch {
			case cc == "":
				continue
			case strings.HasPrefix(cc, "<"):
				continue
			case strings.HasSuffix(cc, ">"):
				continue
			default:
				chunks = append(chunks, cc)
			}
		}
		slices.SortStableFunc(chunks, func(i, j string) int { return c[j] - c[i] })

		return chunks
	}

	factor := func(r Rule, f Rule) Rule {
		if slices.Equal(f.root, []string{}) {
			return r
		}
		if slices.Equal(f.root, []string{""}) {
			return r
		}

		slices.Sort(r.pre)
		slices.Sort(r.root)
		slices.Sort(r.suf)
		if slices.Equal(r.pre, f.root) {
			r.pre = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if slices.Equal(r.root, f.root) {
			r.root = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if slices.Equal(r.suf, f.root) {
			r.suf = []string{fmt.Sprintf("<%s>", f.name())}
		}

		return r
	}

	counts := getCounts(rules)
	ngs := getChunks(counts)
	for _, n := range ngs {
		if counts[n] > f {
			f := Rule{pre: []string{}, root: []string{n}, suf: []string{}, isPublic: false, id: 0}
			for j := range rules {
				rules[j] = factor(rules[j], f)
			}
			rules = append(rules, f)
		}
	}

	return rules
}
