// -*- coding: utf-8 -*-

// Created on Mon Oct  6 04:09:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/jdkato/prose/tag"
)

type EqualityFunction func(g1, g2 []string) bool

func DummyEqual(l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		return true
	}
}

func LiteralEqual(l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		if slices.Equal(g1, g2) {
			l.Printf("%v and %v merged with equality function %s\n", g1, g2, "LiteralEqual")
			return true
		}
		return false
	}
}

func POSSignatureEqual(t Tokenizer, m *tag.PerceptronTagger, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		s1 := strings.Join(g1, " ")
		s2 := strings.Join(g2, " ")
		sig1 := posSignature(t.tokenize(s1), m)
		sig2 := posSignature(t.tokenize(s2), m)
		if sig1 == sig2 {
			l.Printf("%v and %v merged with equality function %s\n", g1, g2, "POSSignatureEqual")
			return true
		}
		return false
	}
}

func CharacterLevenshteinThreshold(t float64, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		s1 := strings.Join(g1, " ")
		s2 := strings.Join(g2, " ")
		sim := CharacterLevenshtein(s1, s2)
		if sim >= t {
			l.Printf("%v and %v merged with equality function %s, threshold %v, similarity %v\n", g1, g2, "CharacterLevenshteinThreshold", t, sim)
			return true
		}
		return false
	}
}

func TokenLevenshteinThreshold(t float64, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		g1 = strings.Split(strings.Join(g1, " "), " ")
		g2 = strings.Split(strings.Join(g2, " "), " ")
		sim := TokenLevenshtein(g1, g2)
		if sim >= t {
			l.Printf("%v and %v merged with equality function %s, threshold %v, similarity %v\n", g1, g2, "TokenLevenshteinThreshold", t, sim)
			return true
		}
		return false
	}
}

func TFIDFCosineThreshold(t float64, v []string, tk Tokenizer, idf map[string]float64, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		v1, err := CountEmbed(strings.Join(g1, " "), v, tk)
		if err != nil {
			fmt.Println(err)
			return false
		}
		v2, err := CountEmbed(strings.Join(g2, " "), v, tk)
		if err != nil {
			fmt.Println(err)
			return false
		}
		v1 = TFIDFTransform(v1, v, idf)
		v2 = TFIDFTransform(v2, v, idf)
		sim, err := CosineSimilarity(v1, v2)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if sim >= t {
			l.Printf("%v and %v merged with equality function %s, threshold %v, similarity %v\n", g1, g2, "TFIDFCosineThreshold", t, sim)
			return true
		}
		return false
	}
}

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

func MergeP(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return e(r1.pre, r2.pre) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.isPublic = true

		r.root = append(r.root, r1.root...)
		r.root = append(r.root, r2.root...)
		slices.Sort(r.root)
		r.root = slices.Compact(r.root)

		r.suf = append(r.suf, r1.suf...)
		r.suf = append(r.suf, r2.suf...)
		slices.Sort(r.suf)
		r.suf = slices.Compact(r.suf)

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

func MergeR(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return e(r1.root, r2.root) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.root = r1.root
		r.isPublic = true

		r.pre = append(r.pre, r1.pre...)
		r.pre = append(r.pre, r2.pre...)
		slices.Sort(r.pre)
		r.pre = slices.Compact(r.pre)

		r.suf = append(r.suf, r1.suf...)
		r.suf = append(r.suf, r2.suf...)
		slices.Sort(r.suf)
		r.suf = slices.Compact(r.suf)

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

func MergeS(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool { return e(r1.suf, r2.suf) }
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.suf = r1.suf
		r.isPublic = true

		r.pre = append(r.pre, r1.pre...)
		r.pre = append(r.pre, r2.pre...)
		slices.Sort(r.pre)
		r.pre = slices.Compact(r.pre)

		r.root = append(r.root, r1.root...)
		r.root = append(r.root, r2.root...)
		slices.Sort(r.root)
		r.root = slices.Compact(r.root)

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

func MergePR(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return e(r1.pre, r2.pre) && e(r1.root, r2.root)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.root = r1.root
		r.isPublic = true

		r.suf = append(r.suf, r1.suf...)
		r.suf = append(r.suf, r2.suf...)
		slices.Sort(r.suf)
		r.suf = slices.Compact(r.suf)

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

func MergePS(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return e(r1.pre, r2.pre) && e(r1.suf, r2.suf)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.pre = r1.pre
		r.suf = r1.suf
		r.isPublic = true

		r.root = append(r.root, r1.root...)
		r.root = append(r.root, r2.root...)
		slices.Sort(r.root)
		r.root = slices.Compact(r.root)

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

func MergeRS(r []Rule, e EqualityFunction) []Rule {
	check := func(r1 Rule, r2 Rule) bool {
		return e(r1.root, r2.root) && e(r1.suf, r2.suf)
	}
	merge := func(r1 Rule, r2 Rule) Rule {
		var r Rule

		r.root = r1.root
		r.suf = r1.suf
		r.isPublic = true

		r.pre = append(r.pre, r1.pre...)
		r.pre = append(r.pre, r2.pre...)
		slices.Sort(r.pre)
		r.pre = slices.Compact(r.pre)

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

func MergeMisc(r []Rule, e EqualityFunction) []Rule {
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
