// -*- coding: utf-8 -*-

// Created on Sun Nov  2 07:32:31 AM EST 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

// Function that abstracts one or more rules out from a slice of rules based on some condition (frequency/content)
type FactorFunction func(r []Rule) []Rule

// Factor expressions to rules based on their frequency (matched by literal match)
func ExpressionFactor(f int, l *log.Logger) FactorFunction {
	return func(rules []Rule) []Rule {
		getCounts := func(r []Rule) map[string]int {
			counts := make(map[string]int)
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
			var chunks []string
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
		factor := func(r Rule, f Rule) Rule {
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
				l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "ExpressionFactor", f.print(f.name()))
				for i := range rules {
					rules[i] = factor(rules[i], f)
				}
				rules = append(rules, f)
			}
		}
		return rules
	}

}

// Factor expressions to rules based on their frequency (matched by constituency tags)
func ConstituencyFactor(tag SyntacticTagger, f int, l *log.Logger) FactorFunction {
	return func(rules []Rule) []Rule {
		getCounts := func(r []Rule) map[string]int {
			counts := make(map[string]int)
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
			var chunks []string
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
			syn := make(map[string][]string)
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
		factor := func(r Rule, f Rule) Rule {
			if slices.Equal(f.root, []string{}) {
				return r
			}
			if slices.Equal(f.root, []string{""}) {
				return r
			}

			r = r.sort()
			f = f.sort()
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

		var (
			counts = getCounts(rules)
			ngs    = getChunks(counts)
			syns   = getSyns(rules, tag)
		)
		for _, n := range ngs {
			if counts[n] > f {
				f := Rule{pre: []string{}, root: syns[n], suf: []string{}, isPublic: false, id: len(rules) + 1}
				l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "ConstituencyFactor", f.print(f.name()))
				for i := range rules {
					rules[i] = factor(rules[i], f)
				}
				rules = append(rules, f)
			}
		}
		return rules
	}
}

// Factor expressions based on user provided synonyms, regardless of frequency
func SynonymFactor(syn Synonyms, tok Tokenizer, l *log.Logger) FactorFunction {
	return func(rules []Rule) []Rule {
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
		factor := func(r Rule, f Rule, tok Tokenizer) Rule {
			if slices.Equal(f.root, []string{}) {
				return r
			}
			if slices.Equal(f.root, []string{""}) {
				return r
			}

			r = r.sort()
			f = f.sort()
			for i := range f.root {
				for j := range r.pre {
					for {
						rtokens := tok.tokenize(r.pre[j])
						ftokens := tok.tokenize(f.root[i])
						ind, found := scanSubseq(rtokens, ftokens)
						if !found {
							break
						}
						r.pre[j] = strings.Join(slices.Replace(rtokens, ind[0], ind[1], fmt.Sprintf("<%s>", f.name())), " ")
					}
				}
				for j := range r.root {
					for {
						rtokens := tok.tokenize(r.root[j])
						ftokens := tok.tokenize(f.root[i])
						ind, found := scanSubseq(rtokens, ftokens)
						if !found {
							break
						}
						r.root[j] = strings.Join(slices.Replace(rtokens, ind[0], ind[1], fmt.Sprintf("<%s>", f.name())), " ")
					}
				}
				for j := range r.suf {
					for {
						rtokens := tok.tokenize(r.suf[j])
						ftokens := tok.tokenize(f.root[i])
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
			slices.Sort(vals)
			f := Rule{pre: []string{}, root: vals, suf: []string{}, isPublic: false, id: len(rules) + 1}
			l.Printf("FACTOR: factor function %s extracted %v to new rule\n", "SynonymFactor", f.print(f.name()))
			for i := range rules {
				if !strings.Contains(strings.Join([]string{strings.Join(rules[i].pre, ""), strings.Join(rules[i].root, ""), strings.Join(rules[i].suf, "")}, " "), k) {
					continue
				}
				rules[i] = factor(rules[i], f, tok)
			}
			rules = append(rules, f)
		}
		return rules
	}

}

// Set of main term and synonyms
type Synonyms map[string][]string

func ReadSynonyms(p string) (Synonyms, error) {
	var err error
	syn := Synonyms{}

	file, err := os.Open(p)
	if err != nil {
		return syn, err
	}
	dec := json.NewDecoder(file)
	err = dec.Decode(&syn)

	return syn, err
}
