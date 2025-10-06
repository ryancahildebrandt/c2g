// -*- coding: utf-8 -*-

// Created on Fri Jul 11 08:51:07 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"strings"

	"github.com/bzick/tokenizer"
)

const (
	Boundary = iota + 1
	WhiteSpace
	Sep
	POSSep
)

var (
	boundaryChars   []string = []string{".", ",", "?", "!", ":", ";"}
	whiteSpaceChars []string = []string{" ", "\t", "\n", "\r"}
	preTokenizedSep []string = []string{"<SEP>"}
)

type Tokenizer interface {
	// Splits a string based on defined tokens
	tokenize(s string) []string
	// Splits string to tokens and joins tokens with single space
	normalize(s string) string
}

type wordTokenizer struct{ *tokenizer.Tokenizer }

func NewWordTokenizer() wordTokenizer {
	tok := wordTokenizer{tokenizer.New()}

	tok.SetWhiteSpaces([]byte{})
	tok.DefineTokens(WhiteSpace, whiteSpaceChars)
	tok.DefineTokens(Boundary, boundaryChars)

	return tok
}

func (tok wordTokenizer) tokenize(s string) []string {
	var (
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = tok.ParseString(s)
	)

	if s == "" {
		return out
	}

	for stream.IsValid() {
		switch {
		case stream.CurrentToken().Is(Boundary):
			builder, out = flushBuilder(builder, out)
			res := stream.CurrentToken().ValueUnescapedString()
			out = append(out, res)
			stream.GoNext()
		case stream.CurrentToken().Is(WhiteSpace):
			builder, out = flushBuilder(builder, out)
			stream.GoNext()
		default:
			builder.WriteString(stream.CurrentToken().ValueUnescapedString())
			stream.GoNext()
		}
	}
	builder, out = flushBuilder(builder, out)

	return out
}

func (tok wordTokenizer) normalize(s string) string {
	return strings.Join(tok.tokenize(s), " ")
}

type sepTokenizer struct{ *tokenizer.Tokenizer }

func NewSepTokenizer() sepTokenizer {
	tok := sepTokenizer{tokenizer.New()}

	tok.SetWhiteSpaces([]byte{})
	tok.DefineTokens(Sep, preTokenizedSep)

	return tok
}

func (tok sepTokenizer) tokenize(s string) []string {
	var (
		builder strings.Builder
		out                       = []string{}
		stream  *tokenizer.Stream = tok.ParseString(s)
	)

	if s == "" {
		return out
	}
	if !strings.Contains(s, "<SEP>") {
		return []string{s}
	}

	for stream.IsValid() {
		switch {
		case stream.CurrentToken().Is(Sep):
			builder, out = flushBuilder(builder, out)
			stream.GoNext()
		default:
			builder.WriteString(stream.CurrentToken().ValueUnescapedString())
			stream.GoNext()
		}
	}
	builder, out = flushBuilder(builder, out)

	return out

}

func (tok sepTokenizer) normalize(s string) string {
	return strings.Join(tok.tokenize(s), " ")
}

// Helper function to get current contents of strings.Builder and reset
func flushBuilder(b strings.Builder, out []string) (strings.Builder, []string) {
	str := b.String()

	b.Reset()
	if len(str) != 0 {
		out = append(out, str)
	}

	return b, out
}
