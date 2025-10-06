// -*- coding: utf-8 -*-

// Created on Mon Oct  6 04:09:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

// Function that determines if two expression groups can be considered equivalent
type EqualityFunction func(e1, e2 []string) bool

// Used for testing
func DummyEqual(l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		return true
	}
}

// Compares expressions for exact string match
func LiteralEqual(l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		if slices.Equal(e1, e2) {
			l.Printf("equality function %s matched %v and %v\n", "LiteralEqual", e1, e2)
			return true
		}
		return false
	}
}

// Compares expressions for matching sequences of POS tags
func POSTagEqual(c SyntacticTagger, l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		s1, _ := c.POS(strings.Join(e1, " "))
		s2, _ := c.POS(strings.Join(e2, " "))
		sig1 := strings.Join(s1, "-")
		sig2 := strings.Join(s2, "-")
		if sig1 == sig2 {
			l.Printf("equality function %s matched %v and %v\n", "POSTagEqual", e1, e2)
			return true
		}
		return false
	}
}

// Compares expressions for matching sequences of constituency tags
func ConstituencyTagEqual(c SyntacticTagger, l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		s1, _ := c.Constituency(strings.Join(e1, " "))
		s2, _ := c.Constituency(strings.Join(e2, " "))
		sig1 := strings.Join(s1, "-")
		sig2 := strings.Join(s2, "-")
		if sig1 == sig2 {
			l.Printf("equality function %s matched %v and %v\n", "ConstituencyTagEqual", e1, e2)
			return true
		}
		return false
	}
}

// Compares expression similarity via character level levenshtein distance
func CharacterLevenshteinThreshold(t float64, l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		sim := CharacterLevenshtein(strings.Join(e1, " "), strings.Join(e2, " "))
		if sim >= t {
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "CharacterLevenshteinThreshold", e1, e2, t, sim)
			return true
		}
		return false
	}
}

// Compares expression similarity via token level levenshtein distance
func TokenLevenshteinThreshold(t float64, l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		e1 = strings.Split(strings.Join(e1, " "), " ")
		e2 = strings.Split(strings.Join(e2, " "), " ")
		sim := TokenLevenshtein(e1, e2)
		if sim >= t {
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "TokenLevenshteinThreshold", e1, e2, t, sim)
			return true
		}
		return false
	}
}

// Compares expression similarity via tfidf embeddings and cosine similarity
func TFIDFCosineThreshold(thr float64, voc []string, tok Tokenizer, idf map[string]float64, l *log.Logger) EqualityFunction {
	return func(e1, e2 []string) bool {
		vec1, err := CountEmbed(strings.Join(e1, " "), voc, tok)
		if err != nil {
			fmt.Println(err)
			return false
		}
		vec2, err := CountEmbed(strings.Join(e2, " "), voc, tok)
		if err != nil {
			fmt.Println(err)
			return false
		}
		vec1 = TFIDFTransform(vec1, voc, idf)
		vec2 = TFIDFTransform(vec2, voc, idf)
		sim, err := CosineSimilarity(vec1, vec2)
		if err != nil {
			fmt.Println(err)
			return false
		}
		if sim >= thr {
			l.Printf("equality function %s matched %v and %v, threshold %v, similarity %v\n", "TFIDFCosineThreshold", e1, e2, thr, sim)
			return true
		}
		return false
	}
}

// Sort rules on prefixes and roots
func SortPR(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.root, r2.root)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

// Sort rules on prefixes and suffixes
func SortPS(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.pre, r2.pre) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.pre, r2.pre)
	})
}

// Sort rules on roots and suffixes
func SortRS(r []Rule) {
	slices.SortStableFunc(r, func(r1 Rule, r2 Rule) int {
		if slices.Equal(r1.root, r2.root) {
			return slices.Compare(r1.suf, r2.suf)
		}
		return slices.Compare(r1.root, r2.root)
	})
}

// Sort rules on prefixes, roots, and suffixes
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

// Merge rules based on equality function match in rule prefix
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

	SortPRS(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeP", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules based on equality function match in rule root
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

	SortPRS(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeR", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules based on equality function match in rule suffix
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

	SortPRS(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeS", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules based on equality function match in rule prefix and root
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

	SortPR(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergePR", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules based on equality function match in rule prefix and suffix
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

	SortPS(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergePS", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules based on equality function match in rule root and suffix
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

	SortRS(r)
	for i := 0; i < len(r)-1; {
		if check(r[i], r[i+1]) {
			rule := merge(r[i], r[i+1])
			l.Printf("merge function %s replaced %v and %v with new rule %v\n", "MergeRS", r[i], r[i+1], rule)
			r[i] = rule
			r = slices.Delete(r, i+1, i+2)
			continue
		}
		i++
	}

	return r
}

// Merge rules where prefix, root, and suffix are all len==1 or empty into one rule
// generally applied after all other megring strategies
func MergeMisc(r []Rule, e EqualityFunction, l *log.Logger) []Rule {
	check := func(r Rule) bool {
		return len(r.pre) <= 1 && len(r.root) <= 1 && len(r.suf) <= 1
	}
	merge := func(rules ...Rule) Rule {
		var (
			rule Rule
			b    strings.Builder
			exp  string
		)

		for _, rr := range rules {
			b.WriteString(strings.Join(rr.pre, " "))
			b.WriteString(" ")
			b.WriteString(strings.Join(rr.root, " "))
			b.WriteString(" ")
			b.WriteString(strings.Join(rr.suf, " "))
			exp = strings.TrimSpace(b.String())
			if !slices.Contains(rule.root, exp) {
				rule.root = append(rule.root, exp)
			}
			b.Reset()
		}
		if rule.root != nil {
			rule.pre = []string{}
			rule.suf = []string{}
			rule.isPublic = true
		}

		return rule
	}

	var rr []Rule
	var res Rule

	SortPRS(r)
	for i := 0; i < len(r); i++ {
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
