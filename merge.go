// -*- coding: utf-8 -*-

// Created on Mon Oct  6 04:09:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"strings"
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
			l.Printf("equality function %s matched %v and %v\n", "LiteralEqual", g1, g2)
			return true
		}
		return false
	}
}

func POSTagEqual(c SyntacticTagger, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		s1 := strings.Join(g1, " ")
		s2 := strings.Join(g2, " ")
		ss1, _ := c.POS(s1)
		ss2, _ := c.POS(s2)
		sig1 := strings.Join(ss1, "-")
		sig2 := strings.Join(ss2, "-")
		if sig1 == sig2 {
			l.Printf("equality function %s matched %v and %v\n", "POSTagEqual", g1, g2)
			return true
		}
		return false
	}
}

func ConstituencyTagEqual(c SyntacticTagger, l *log.Logger) EqualityFunction {
	return func(g1, g2 []string) bool {
		s1 := strings.Join(g1, " ")
		s2 := strings.Join(g2, " ")
		ss1, _ := c.Constituency(s1)
		ss2, _ := c.Constituency(s2)
		sig1 := strings.Join(ss1, "-")
		sig2 := strings.Join(ss2, "-")
		if sig1 == sig2 {
			l.Printf("equality function %s matched %v and %v\n", "ConstituencyTagEqual", g1, g2)
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
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "CharacterLevenshteinThreshold", g1, g2, t, sim)
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
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "TokenLevenshteinThreshold", g1, g2, t, sim)
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
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "TFIDFCosineThreshold", g1, g2, t, sim)
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

func MergeP(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeP", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeR(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeR", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeS(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeS", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergePR(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergePR", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergePS(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergePS", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeRS(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			rr := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeRS", r[i], r[i+1], rr)
			r[i] = rr
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

func MergeMisc(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
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
			l.Printf("merge function %s added %v to new misc rule\n", "MergeMisc", r[i])
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

func GroupFactor(rules []Rule, f int, l *log.Logger) []Rule {
	getCounts := func(r []Rule) map[string]int {
		var counts = make(map[string]int)

		slices.SortStableFunc(r, func(i, j Rule) int {
			return strings.Compare(i.print(""), j.print(""))
		})

		for _, rule := range r {
			rule = rule.sort()
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
		slices.SortStableFunc(chunks, func(i, j string) int {
			if c[j] == c[i] {
				return strings.Compare(i, j)
			}
			return c[j] - c[i]
		})

		return chunks
	}
	factor := func(r Rule, f Rule, l *log.Logger) Rule {
		if slices.Equal(f.root, []string{}) {
			return r
		}
		if slices.Equal(f.root, []string{""}) {
			return r
		}

		r = r.sort()
		f = f.sort()
		if LiteralEqual(l)(r.pre, f.root) {
			r.pre = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if LiteralEqual(l)(r.root, f.root) {
			r.root = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if LiteralEqual(l)(r.suf, f.root) {
			r.suf = []string{fmt.Sprintf("<%s>", f.name())}
		}

		return r
	}

	counts := getCounts(rules)
	ngs := getChunks(counts)
	for _, n := range ngs {
		if counts[n] > f {
			f := Rule{pre: []string{}, root: []string{n}, suf: []string{}, isPublic: false, id: len(rules) + 1}
			l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "GroupFactor", f)
			for j := range rules {
				rules[j] = factor(rules[j], f, l)
			}
			rules = append(rules, f)
		}
	}

	return rules
}

func SynonymFactor(rules []Rule, syn Synonyms, tk Tokenizer, l *log.Logger) []Rule {
	scanSubseq := func(s1, s2 []string) ([]int, bool) {
		var res []int

		if len(s1) < len(s2) {
			return res, false
		}

		for i := range s2 {
			if !slices.Contains(s1, s2[i]) {
				return res, false
			}
		}

		for i := 0; i <= len(s1)-len(s2); i++ {
			win := s1[i : i+len(s2)]
			if slices.Equal(s2, win) {
				return []int{i, i + len(s2)}, true
			}
		}

		return res, false
	}
	factor := func(r Rule, f Rule, tk Tokenizer, l *log.Logger) Rule {
		if slices.Equal(f.root, []string{}) {
			return r
		}
		if slices.Equal(f.root, []string{""}) {
			return r
		}

		r = r.sort()
		for i := range f.root {
			for j := range r.pre {
				for {
					rtokens := tk.tokenize(r.pre[j])
					ftokens := tk.tokenize(f.root[i])
					ind, found := scanSubseq(rtokens, ftokens)
					if !found {
						break
					}
					r.pre[j] = strings.Join(slices.Replace(rtokens, ind[0], ind[1], fmt.Sprintf("<%s>", f.name())), " ")
				}
			}
			for j := range r.root {
				for {
					rtokens := tk.tokenize(r.root[j])
					ftokens := tk.tokenize(f.root[i])
					ind, found := scanSubseq(rtokens, ftokens)
					if !found {
						break
					}
					r.root[j] = strings.Join(slices.Replace(rtokens, ind[0], ind[1], fmt.Sprintf("<%s>", f.name())), " ")
				}
			}
			for j := range r.suf {
				for {
					rtokens := tk.tokenize(r.suf[j])
					ftokens := tk.tokenize(f.root[i])
					ind, found := scanSubseq(rtokens, ftokens)
					if !found {
						break
					}
					r.suf[j] = strings.Join(slices.Replace(rtokens, ind[0], ind[1], fmt.Sprintf("<%s>", f.name())), " ")
				}
			}
		}

		return r
	}

	keys := slices.Sorted(maps.Keys(syn))
	for _, k := range keys {
		vals := append(syn[k], k)
		f := Rule{pre: []string{}, root: vals, suf: []string{}, isPublic: false, id: len(rules) + 1}
		l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "SynonymFactor", f)
		for j := range rules {
			if !strings.Contains(strings.Join([]string{strings.Join(rules[j].pre, ""), strings.Join(rules[j].root, ""), strings.Join(rules[j].suf, "")}, " "), k) {
				continue
			}
			rules[j] = factor(rules[j], f, tk, l)
		}
		rules = append(rules, f)
	}
	return rules
}

func ConstituencyFactor(rules []Rule, tag SyntacticTagger, f int, l *log.Logger) []Rule {
	getCounts := func(r []Rule) map[string]int {
		var counts = make(map[string]int)

		slices.SortStableFunc(r, func(i, j Rule) int {
			return strings.Compare(i.print(""), j.print(""))
		})

		for _, rule := range r {
			for i := range rule.pre {
				tags, _ := tag.Constituency(rule.pre[i])
				counts[strings.Join(tags, "-")]++
			}
			for i := range rule.root {
				tags, _ := tag.Constituency(rule.root[i])
				counts[strings.Join(tags, "-")]++
			}
			for i := range rule.suf {
				tags, _ := tag.Constituency(rule.suf[i])
				counts[strings.Join(tags, "-")]++
			}
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
		slices.SortStableFunc(chunks, func(i, j string) int {
			if c[j] == c[i] {
				return strings.Compare(i, j)
			}
			return c[j] - c[i]
		})

		return chunks
	}
	getSyns := func(r []Rule, tag SyntacticTagger) map[string][]string {
		var syn = make(map[string][]string)

		for _, rule := range r {
			for i := range rule.pre {
				tags, tokens := tag.Constituency(rule.pre[i])
				key := strings.Join(tags, "-")
				val := strings.Join(tokens, " ")
				syn[key] = append(syn[key], val)
			}
			for i := range rule.root {
				tags, tokens := tag.Constituency(rule.root[i])
				key := strings.Join(tags, "-")
				val := strings.Join(tokens, " ")
				syn[key] = append(syn[key], val)
			}
			for i := range rule.suf {
				tags, tokens := tag.Constituency(rule.suf[i])
				key := strings.Join(tags, "-")
				val := strings.Join(tokens, " ")
				syn[key] = append(syn[key], val)
			}

		}
		for k, v := range syn {
			slices.Sort(v)
			syn[k] = slices.Compact(v)
		}
		return syn
	}
	factor := func(r Rule, f Rule, l *log.Logger) Rule {
		if slices.Equal(f.root, []string{}) {
			return r
		}
		if slices.Equal(f.root, []string{""}) {
			return r
		}

		r = r.sort()

		if ConstituencyTagEqual(tag, l)(r.pre, f.root) {
			r.pre = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if ConstituencyTagEqual(tag, l)(r.root, f.root) {
			r.root = []string{fmt.Sprintf("<%s>", f.name())}
		}
		if ConstituencyTagEqual(tag, l)(r.suf, f.root) {
			r.suf = []string{fmt.Sprintf("<%s>", f.name())}
		}

		return r
	}

	counts := getCounts(rules)
	ngs := getChunks(counts)
	syns := getSyns(rules, tag)
	for _, n := range ngs {
		if counts[n] > f {
			f := Rule{pre: []string{}, root: syns[n], suf: []string{}, isPublic: false, id: len(rules) + 1}
			l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "ConstituencyFactor", f)
			for j := range rules {
				rules[j] = factor(rules[j], f, l)
			}
			rules = append(rules, f)
		}
	}
	return rules
}
