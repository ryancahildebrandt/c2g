// -*- coding: utf-8 -*-

// Created on Fri Jun 27 03:42:54 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	t := NewUnigramTokenizer()
	file, err := os.Open("./data/20000-Utterances-Training-dataset-for-chatbots-virtual-assistant-Bitext-sample.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	corpus := Corpus{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			text := scanner.Text()
			corpus.texts = append(corpus.texts, Sentence{text, []Ngram{}})
		}
	}
	corpus.transitions = NewTransitions(corpus, t)
	ngram_map := make(map[Expression]Ngram)
	for _, c := range corpus.texts {
		tokens := UnigramTokenize(c.text, t)
		for _, ng := range NgramTokenize(tokens, corpus.transitions, 0.1) {
			_, ok := ngram_map[ng]
			if !ok {
				ngram_map[ng] = Ngram{ng, len(strings.Split(ng, " ")) - 1, 0}
			}
			ngram_map[ng] = Ngram{ng, ngram_map[ng].len, ngram_map[ng].count + 1}
		}
	}
	ng := []Ngram{}
	for _, v := range ngram_map {
		ng = append(ng, v)
	}
	slices.SortFunc(ng, func(a, b Ngram) int {
		switch {
		case a.len == b.len:
			return b.count - a.count
		default:
			return b.len - a.len
		}
	})

	tmp := corpus.texts
	for j, ngram := range ng {
		if j == 200 {
			break
		}
		var rule Rule = NewRule(ngram)
		for i := 0; i < len(tmp); i++ {
			c := tmp[i]
			p, s, found := strings.Cut(c.text, rule.root.text)
			if found {
				if p != "" {
					if !slices.Contains(rule.pre, p) {
						rule.pre = append(rule.pre, p)
					}
				}
				if s != "" {
					if !slices.Contains(rule.suf, s) {
						rule.suf = append(rule.suf, s)
					}
				}
				tmp = slices.Delete(corpus.texts, i, i+1)
			}
		}
		fmt.Println(j, rule.print(fmt.Sprint(j)))
	}

	fmt.Println(time.Since(start))
}
