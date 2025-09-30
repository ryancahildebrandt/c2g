// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

type Text struct {
	pre   string
	root  string
	suf   string
	text  string
	chunk []string
}

func ReadTexts(s *bufio.Scanner) []Text {
	var t []Text
	var tt string

	for s.Scan() {
		tt = strings.TrimSpace(s.Text())
		if tt != "" {
			t = append(t, Text{text: tt, chunk: []string{}})
		}
	}

	slices.SortStableFunc(t, func(i, j Text) int { return strings.Compare(i.text, j.text) })
	t = slices.CompactFunc(t, func(i, j Text) bool { return i.text == j.text })

	return t
}

func ToTriplet(t Text, n []string) Text {
	i := slices.IndexFunc(n, func(s string) bool {
		return slices.Contains(t.chunk, s)
	})

	if i == -1 {
		t.root = t.text
		return t
	}

	t.root = n[i]
	p, s, found := strings.Cut(t.text, n[i])
	if !found {
		fmt.Println(t.text, "|", n[i], "|", p, "|", s)
	}
	t.pre = strings.TrimSpace(p)
	t.suf = strings.TrimSpace(s)

	return t
}

func ToRule(t Text) Rule {
	var r Rule

	r.root = []string{t.root}
	r.pre = []string{t.pre}
	r.suf = []string{t.suf}
	r.isPublic = true

	return r
}
