// coding: utf-8 -*-

// Created on Tue Oct  7 06:53:46 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func levenshteinDistance(s1, s2 []string) int {
	var prev = make([]int, len(s2)+1)
	var curr = make([]int, len(s2)+1)

	for i := range prev {
		prev[i] = i
	}

	for i := 1; i <= len(s1); i++ {
		curr = make([]int, len(s2)+1)
		curr[0] = i
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1]
				continue
			}
			curr[j] = slices.Min([]int{curr[j-1], prev[j], prev[j-1]}) + 1
		}
		prev = curr
	}

	return curr[len(s2)]
}

func CharacterLevenshtein(s1, s2 string) float64 {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	switch {
	case s1 == s2:
		return 1
	case len(s1) == 0:
		return 0
	case len(s2) == 0:
		return 0
	}

	var arr1 = strings.Split(s1, "")
	var arr2 = strings.Split(s2, "")
	var dist = levenshteinDistance(arr1, arr2)

	return 1 - (float64(dist) / float64(len(arr2)))
}

func TokenLevenshtein(s1, s2 []string) float64 {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	switch {
	case slices.Equal(s1, s2):
		return 1.0
	case len(s1) == 0:
		return 0.0
	case len(s2) == 0:
		return 0.0
	}

	var arr1 = s1
	var arr2 = s2
	var dist = levenshteinDistance(arr1, arr2)

	return 1 - (float64(dist) / float64(len(arr2)))
}

func CollectVocab(t []Text, tk Tokenizer) []string {
	var v = []string{}
	var tok []string
	for i := range t {
		tok = tk.tokenize(strings.ToLower(t[i].text))
		v = append(v, tok...)
	}
	slices.Sort(v)
	v = slices.Compact(v)
	return v
}

func CollectIDF(t []Text, tk Tokenizer) map[string]float64 {
	var m = make(map[string]float64)
	var tok []string

	for i := range t {
		tok = tk.tokenize(strings.ToLower(t[i].text))
		slices.Sort(tok)
		tok = slices.Compact(tok)
		for j := range tok {
			m[tok[j]]++
		}
	}

	for k := range m {
		m[k] = math.Log(float64(len(t))/m[k] + 1)
	}

	return m
}

func CountEmbed(s string, v []string, tk Tokenizer) (mat.VecDense, error) {
	if len(v) == 0 {
		return mat.VecDense{}, fmt.Errorf("vocab is empty")
	}
	var e = mat.NewVecDense(len(v), nil)
	var tok = tk.tokenize(strings.ToLower(s))
	var counts = make(map[int]int)
	var ind int
	var err error

	for i := range tok {
		ind = slices.Index(v, tok[i])
		if ind == -1 {
			return mat.VecDense{}, fmt.Errorf("token not found in vocab")
		}
		counts[ind]++
	}
	for k, v := range counts {
		e.SetVec(k, float64(v))
	}

	return *e, err
}

func TFIDFTransform(v mat.VecDense, vv []string, idf map[string]float64) mat.VecDense {
	collectTF := func(v mat.VecDense) map[string]float64 {
		var n float64
		var m = make(map[string]float64)
		for i := range v.Len() {
			n += v.AtVec(i)
		}
		for i := range v.Len() {
			if v.AtVec(i) == 0 {
				continue
			}
			tok := vv[i]
			count := v.AtVec(i)
			m[tok] = count / n
		}
		return m
	}

	tf := collectTF(v)
	for i := range v.Len() {
		if v.AtVec(i) == 0 {
			continue
		}
		val := v.AtVec(i)
		tok := vv[i]
		weight := tf[tok] * idf[tok]
		v.SetVec(i, val*weight)
	}

	return v
}

func CosineSimilarity(v1, v2 mat.VecDense) (float64, error) {
	switch {
	case v1.Len() != v2.Len():
		return 0.0, fmt.Errorf("vectors v1 and v2 lengths differ. len(v1):%v len(v2):%v", v1.Len(), v2.Len())
	case mat.Equal(&v1, &v2):
		return 1.0, nil
	case v1.Norm(2) == 0 || v2.Norm(2) == 0:
		return 0.0, nil
	default:
	}

	return mat.Dot(&v1, &v2) / (v1.Norm(2) * v2.Norm(2)), nil
}
