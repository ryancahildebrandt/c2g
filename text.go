// -*- coding: utf-8 -*-

// Created on Sat Jul 12 08:18:48 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"

	"gonum.org/v1/gonum/stat"
)

// Stores information for one text within a corpus
// pre, root, make up the main structure of a text
type Text struct {
	// portion of the text preceeding the root
	pre string
	// the largest chunk present in the text
	root string
	// portion of the text following the root
	suf   string
	text  string
	chunk []string
}

// Reads each line of the input file, converting each line to a Text struct and removing duplicates
func ReadTexts(s *bufio.Scanner) []Text {
	texts := []Text{}

	for s.Scan() {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			texts = append(texts, Text{text: text, chunk: []string{}})
		}
	}

	slices.SortStableFunc(texts, func(i, j Text) int { return strings.Compare(i.text, j.text) })
	texts = slices.CompactFunc(texts, func(i, j Text) bool { return i.text == j.text })

	return texts
}

// Sets the largest chunk in c present in t as t.root, sets prefix and suffix accordingly
func ToTriplet(t Text, c []string) Text {
	ind := slices.IndexFunc(c, func(s string) bool {
		return slices.Contains(t.chunk, s)
	})

	if ind == -1 {
		t.root = t.text
		return t
	}

	t.root = c[ind]
	pre, suf, found := strings.Cut(t.text, c[ind])
	if !found {
		fmt.Println(t.text, "|", c[ind], "|", pre, "|", suf)
	}
	t.pre = strings.TrimSpace(pre)
	t.suf = strings.TrimSpace(suf)

	return t
}

// Converts Text to Rule
func ToRule(t Text) Rule {
	return Rule{pre: []string{t.pre}, root: []string{t.root}, suf: []string{t.suf}, isPublic: true}
}

// Keeps only the texts matching the most common structures found in the corpus
// structures are determined by constituency tags, texts not matching the top q quantile of structures are removed
// higher q will remove more texts
func FilterTexts(t []Text, tag SyntacticTagger, q float64) []Text {
	var (
		counts    = make(map[string]int)
		sigs      = []string{}
		vals      = []float64{}
		threshold float64
	)
	if len(t) == 0 {
		return t

	}
	for i := range t {
		sig, _ := tag.Constituency(t[i].text)
		sigs = append(sigs, strings.Join(sig, "-"))
		counts[strings.Join(sig, "-")]++
	}

	slices.SortStableFunc(sigs, func(i, j string) int { return counts[i] - counts[j] })

	for i := range sigs {
		vals = append(vals, float64(counts[sigs[i]]))
	}

	threshold = stat.Quantile(q, stat.Empirical, vals, nil)
	t = slices.DeleteFunc(t, func(i Text) bool {
		sig, _ := tag.Constituency(i.text)
		return counts[strings.Join(sig, "-")] < int(threshold)
	})

	return t
}
