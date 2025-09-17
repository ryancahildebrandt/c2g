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
	pre  []Expression
	root []Expression
	suf  []Expression
}

func (r *Rule) isEmpty() bool {
	if len(r.pre) == 0 && len(r.root) == 0 && len(r.suf) == 0 {
		return true
	}
	return false
}

func NewRule(e Expression) Rule {
	var r Rule
	r.root = []Expression{e}

	return r
}

func fmtGroup(g []Expression) string {
	var o bool
	var i int

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

func (r *Rule) print(n string, p bool) string {
	var b strings.Builder

	if r.isEmpty() {
		return ""
	}

	pre := fmtGroup(r.pre)
	root := fmtGroup(r.root)
	suf := fmtGroup(r.suf)

	if p {
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

func PRSort(r []Rule) {
	slices.SortFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.root, r2.root)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func PSSort(r []Rule) {
	slices.SortFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

func RSSort(r []Rule) {
	slices.SortFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.root, r2.root) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.root, r2.root)
	})
}

func PRSSort(r []Rule) {
	slices.SortFunc(r, func(r1 Rule, r2 Rule) int {
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
	check(r1 Rule, r2 Rule) bool
	merge(r1 Rule, r2 Rule) Rule
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
