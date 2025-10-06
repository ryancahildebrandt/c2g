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

// Helper function to calculate levenshtein distance from string slice
func levenshteinDistance(s1, s2 []string) int {
	prev := make([]int, len(s2)+1)
	curr := make([]int, len(s2)+1)

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

// Calculate levenshtein distance from 2 strings
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

	var (
		arr1 = strings.Split(s1, "")
		arr2 = strings.Split(s2, "")
		dist = levenshteinDistance(arr1, arr2)
	)

	return 1 - (float64(dist) / float64(len(arr2)))
}

// Calculate levenshtein distance from 2 slices of tokens
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

	dist := levenshteinDistance(s1, s2)

	return 1 - (float64(dist) / float64(len(s2)))
}

// Collect all unique tokens from Texts
func CollectVocab(t []Text, tok Tokenizer) []string {
	var tokens []string
	vocab := []string{}

	for i := range t {
		tokens = tok.tokenize(strings.ToLower(t[i].text))
		vocab = append(vocab, tokens...)
	}
	slices.Sort(vocab)
	vocab = slices.Compact(vocab)

	return vocab
}

// Calculate inverse document frequency scores from Texts
func CollectIDF(t []Text, tok Tokenizer) map[string]float64 {
	var tokens []string
	m := make(map[string]float64)

	for i := range t {
		tokens = tok.tokenize(strings.ToLower(t[i].text))
		slices.Sort(tokens)
		tokens = slices.Compact(tokens)
		for j := range tokens {
			m[tokens[j]]++
		}
	}

	for k := range m {
		m[k] = math.Log(float64(len(t))/m[k] + 1)
	}

	return m
}

// Calculate count embeddings for text based on provided vocabulary
func CountEmbed(s string, voc []string, tok Tokenizer) (mat.VecDense, error) {
	if len(voc) == 0 {
		return mat.VecDense{}, fmt.Errorf("vocab is empty")
	}

	var (
		emb    = mat.NewVecDense(len(voc), nil)
		tokens = tok.tokenize(strings.ToLower(s))
		counts = make(map[int]int)
		err    error
	)

	for i := range tokens {
		ind := slices.Index(voc, tokens[i])
		if ind == -1 {
			return mat.VecDense{}, fmt.Errorf("token not found in vocab")
		}
		counts[ind]++
	}

	for k, v := range counts {
		emb.SetVec(k, float64(v))
	}

	return *emb, err
}

// Transform count embeddings with IDF weights
func TFIDFTransform(vec mat.VecDense, voc []string, idf map[string]float64) mat.VecDense {
	collectTF := func(v mat.VecDense) map[string]float64 {
		var sum float64
		m := make(map[string]float64)

		for i := range v.Len() {
			sum += v.AtVec(i)
		}

		for i := range v.Len() {
			if v.AtVec(i) == 0 {
				continue
			}
			tok := voc[i]
			count := v.AtVec(i)
			m[tok] = count / sum
		}

		return m
	}

	tf := collectTF(vec)
	for i := range vec.Len() {
		if vec.AtVec(i) == 0 {
			continue
		}
		val := vec.AtVec(i)
		token := voc[i]
		weight := tf[token] * idf[token]
		vec.SetVec(i, val*weight)
	}

	return vec
}

// Calculate similarity between embedding vectors
func CosineSimilarity(vec1, vec2 mat.VecDense) (float64, error) {
	switch {
	case vec1.Len() != vec2.Len():
		return 0.0, fmt.Errorf("vectors v1 and v2 lengths differ. len(v1):%v len(v2):%v", vec1.Len(), vec2.Len())
	case mat.Equal(&vec1, &vec2):
		return 1.0, nil
	case vec1.Norm(2) == 0 || vec2.Norm(2) == 0:
		return 0.0, nil
	default:
	}

	return mat.Dot(&vec1, &vec2) / (vec1.Norm(2) * vec2.Norm(2)), nil
}
