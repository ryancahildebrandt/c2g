// -*- coding: utf-8 -*-

// Created on Sun Jul 27 07:52:09 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"slices"
	"strings"
)

type Rule struct {
	pre      []Expression
	root     []Expression
	suf      []Expression
	isPublic bool
	id       int
}

func (r *Rule) isEmpty() bool {
	if len(r.pre) == 0 && len(r.root) == 0 && len(r.suf) == 0 {
		return true
	}
	if len(r.pre)+len(r.root)+len(r.suf) <= 3 {
		e := []Expression{}
		e = append(e, r.pre...)
		e = append(e, r.root...)
		e = append(e, r.suf...)
		for _, e := range e {
			if e != "" {
				return false
			}
		}
		return true
	}
	return false
}

func NewRule(e Expression) Rule {
	var r Rule
	r.root = []Expression{e}

	return r
}

func joinBoundaries(e Expression) Expression {
	for _, b := range boundaryChars {
		e = strings.Replace(e, fmt.Sprint(" ", b), b, -1)
	}
	return e
}

func fmtGroup(g []Expression) string {
	var o bool
	var i int

	for j := range g {
		g[j] = joinBoundaries(g[j])
	}

	if slices.Contains(g, "") {
		o = true
		i = slices.Index(g, "")
		g = slices.Delete(g, i, i+1)
	}
	if len(g) == 0 {
		return ""
	}
	if o {
		return fmt.Sprintf("[%s]", strings.Join(g, "|"))
	}
	return fmt.Sprintf("(%s)", strings.Join(g, "|"))
}

func (r *Rule) print(n string) string {
	var b strings.Builder

	if r.isEmpty() {
		return ""
	}

	pre := fmtGroup(r.pre)
	root := fmtGroup(r.root)
	suf := fmtGroup(r.suf)

	if r.isPublic {
		b.WriteString("public ")
	}
	b.WriteString(fmt.Sprintf("<%s> =", n))

	for _, s := range []string{pre, root, suf} {
		if s == "()" || s == "[]" || s == "" {
			continue
		}
		b.WriteString(fmt.Sprintf(" %s", s))
	}
	b.WriteString(";")

	return b.String()
}

func (r *Rule) name() string {
	base := strings.Join(r.root, "_")
	base = strings.Replace(base, " ", "_", -1)
	base = strings.Replace(base, "<", "", -1)
	base = strings.Replace(base, ">", "", -1)
	if r.id != 0 {
		return fmt.Sprintf("%.*s_%v", 20, base, r.id)
	}
	return fmt.Sprintf("%.*s", 20, base)
}

func PRSort(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.root, r2.root)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func PSSort(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func RSSort(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.root, r2.root) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.root, r2.root)
	})
}

func PRSSort(r []Rule) {
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

type RuleMerger interface {
	apply(r []Rule) []Rule
}

type SSDMerger struct{}

func (m *SSDMerger) check(r1 Rule, r2 Rule) bool {
	return slices.Equal(r1.pre, r2.pre) && slices.Equal(r1.root, r2.root)
}

func (m *SSDMerger) merge(r1 Rule, r2 Rule) Rule {
	var r Rule
	r.pre = r1.pre
	r.root = r1.root
	r.suf = []Expression{}
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

func (m *SSDMerger) apply(r []Rule) []Rule {
	var i int

	PRSort(r)
	for i < len(r)-1 {
		if m.check(r[i], r[i+1]) {
			r[i] = m.merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

type SDSMerger struct{}

func (m *SDSMerger) check(r1 Rule, r2 Rule) bool {
	return slices.Equal(r1.pre, r2.pre) && slices.Equal(r1.suf, r2.suf)
}

func (m *SDSMerger) merge(r1 Rule, r2 Rule) Rule {
	var r Rule
	r.pre = r1.pre
	r.root = []Expression{}
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

func (m *SDSMerger) apply(r []Rule) []Rule {
	var i int

	PSSort(r)
	for i < len(r)-1 {
		if m.check(r[i], r[i+1]) {
			r[i] = m.merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

type DSSMerger struct{}

func (m *DSSMerger) check(r1 Rule, r2 Rule) bool {
	return slices.Equal(r1.root, r2.root) && slices.Equal(r1.suf, r2.suf)
}

func (m *DSSMerger) merge(r1 Rule, r2 Rule) Rule {
	var r Rule
	r.pre = []Expression{}
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

func (m *DSSMerger) apply(r []Rule) []Rule {
	var i int

	RSSort(r)
	for i < len(r)-1 {
		if m.check(r[i], r[i+1]) {
			r[i] = m.merge(r[i], r[i+1])
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

type SSSMerger struct{}

func (m *SSSMerger) check(r Rule) bool {
	return len(r.pre) <= 1 && len(r.root) <= 1 && len(r.suf) <= 1
}

func (m *SSSMerger) merge(rules ...Rule) Rule {
	var r Rule
	var b strings.Builder

	for _, rule := range rules {
		b.WriteString(strings.Join(rule.pre, " "))
		b.WriteString(" ")
		b.WriteString(strings.Join(rule.root, " "))
		b.WriteString(" ")
		b.WriteString(strings.Join(rule.suf, " "))
		e := strings.TrimSpace(b.String())
		if !slices.Contains(r.root, e) {
			r.root = append(r.root, e)
		}
		b.Reset()
	}
	if r.root != nil {
		r.pre = []Expression{}
		r.suf = []Expression{}
		r.isPublic = true

	}

	return r
}

func (m *SSSMerger) apply(r []Rule) []Rule {
	var rr []Rule
	var res Rule

	PRSSort(r)
	for i := 0; i < len(r); i++ {
		if m.check(r[i]) {
			if !r[i].isEmpty() {
				rr = append(rr, r[i])
			}
			r = slices.Delete(r, i, i+1)
		}
	}
	res = m.merge(rr...)
	if !res.isEmpty() {
		r = append(r, res)
	}

	return r
}

func Factor(r Rule, f Rule) Rule {
	if slices.Equal(f.root, []Expression{}) {
		return r
	}
	if slices.Equal(f.root, []Expression{""}) {
		return r
	}

	slices.Sort(r.pre)
	slices.Sort(r.root)
	slices.Sort(r.suf)

	if slices.Equal(r.pre, f.root) {
		r.pre = []Expression{fmt.Sprintf("<%s>", f.name())}
	}
	if slices.Equal(r.root, f.root) {
		r.root = []Expression{fmt.Sprintf("<%s>", f.name())}
	}
	if slices.Equal(r.suf, f.root) {
		r.suf = []Expression{fmt.Sprintf("<%s>", f.name())}
	}

	return r
}

func ApplyFactor(rules []Rule) []Rule {
	Counts := func(r []Rule) map[string]int {
		counts := make(map[string]int)

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

	NGrams := func(c map[string]int) []Expression {
		ngs := []Expression{}

		for cc := range c {
			if cc == "" {
				continue
			}
			if strings.HasPrefix(cc, "<") {
				continue
			}
			if strings.HasSuffix(cc, ">") {
				continue
			}
			ngs = append(ngs, cc)
		}

		slices.SortStableFunc(ngs, func(i, j Expression) int {
			return c[j] - c[i]
		})

		return ngs
	}

	counts := Counts(rules)
	ngs := NGrams(counts)
	for _, n := range ngs {
		if counts[n] > 5 {
			f := Rule{pre: []Expression{}, root: []Expression{n}, suf: []Expression{}, isPublic: false, id: 0}
			for j := range rules {
				rules[j] = Factor(rules[j], f)
			}
			rules = append(rules, f)
		}
	}

	return rules
}

func SetIDs(rules []Rule) []Rule {
	slices.SortStableFunc(rules, func(i, j Rule) int {
		return strings.Compare(i.name(), j.name())
	})

	for i := range rules {
		rules[i].id = i
	}

	return rules
}
