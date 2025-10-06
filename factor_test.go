// -*- coding: utf-8 -*-

// Created on Sun Nov  2 07:32:31 AM EST 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/jdkato/prose/tag"
	"github.com/stretchr/testify/assert"
)

func TestReadSynonyms(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		args      args
		want      Synonyms
		assertion assert.ErrorAssertionFunc
	}{
		{args: args{p: ""}, want: Synonyms{}, assertion: assert.Error},
		{args: args{p: "./data/tests/syn1.json"}, want: Synonyms{"I want to know what": []string{"TEST1", "TEST2"}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn2.json"}, want: Synonyms{"I want to know what": []string{"TEST5", "TEST6"}, "My bill": []string{"TEST7"}, "my bill": []string{}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn3.json"}, want: Synonyms{"a": []string{"TEST3", "TEST4"}}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn4.json"}, want: Synonyms{}, assertion: assert.NoError},
		{args: args{p: "./data/tests/syn5.json"}, want: Synonyms{}, assertion: assert.Error},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got, err := ReadSynonyms(tt.args.p)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExpressionFactor(t *testing.T) {
	type args struct {
		f  string
		ff int
	}
	tests := []struct {
		args args
		want []Rule
	}{
		{args: args{f: "./data/tests/test5.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{""}, root: []string{"<I_don't_have_an_onli_2>"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: 1}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test6.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"I don't understand you"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{}, root: []string{"I want an online accoynt"}, suf: []string{}, isPublic: false, id: 15}, {pre: []string{}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{}, isPublic: false, id: 16}, {pre: []string{}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{}, isPublic: false, id: 17}, {pre: []string{}, root: []string{"can you show me my invoices?"}, suf: []string{}, isPublic: false, id: 18}, {pre: []string{}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{}, isPublic: false, id: 19}, {pre: []string{}, root: []string{"i dont want my profile"}, suf: []string{}, isPublic: false, id: 20}, {pre: []string{}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{}, isPublic: false, id: 21}, {pre: []string{}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{}, isPublic: false, id: 22}, {pre: []string{""}, root: []string{"<I_don't_have_an_onli_12>"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"<I_don't_understand_y_13>"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"<I_got_an_error_messa_14>"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"<I_want_an_online_acc_15>"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"<ask_an_agent_to_noti_16>"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"<can_you_show_me_info_17>"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"<can_you_show_me_my_i_18>"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"<can_you_tell_me_how__19>"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"<i_dont_want_my_profi_20>"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"<i_want_to_know_wat_t_21>"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"<where_can_i_leave_an_22>"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test6.csv", ff: 1}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test6.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test7.csv", ff: 0}, want: []Rule{}},
		{args: args{f: "./data/tests/test8.csv", ff: 0}, want: []Rule{}},
		{args: args{f: "./data/tests/test9.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"I have a question"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{}, root: []string{"I want to download a bill"}, suf: []string{}, isPublic: false, id: 15}, {pre: []string{}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{}, isPublic: false, id: 16}, {pre: []string{}, root: []string{"I want to make a review for a service"}, suf: []string{}, isPublic: false, id: 17}, {pre: []string{}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{}, isPublic: false, id: 18}, {pre: []string{}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{}, isPublic: false, id: 19}, {pre: []string{}, root: []string{"i want to request an invoice"}, suf: []string{}, isPublic: false, id: 20}, {pre: []string{}, root: []string{"where do i check the delivery options?"}, suf: []string{}, isPublic: false, id: 21}, {pre: []string{}, root: []string{"you arent helping"}, suf: []string{}, isPublic: false, id: 22}, {pre: []string{""}, root: []string{"<I_don't_have_an_onli_12>"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"<I_have_a_question_13>"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"<I_ordered_an_item_an_14>"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"<I_want_to_download_a_15>"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"<I_want_to_know_what__16>"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"<I_want_to_make_a_rev_17>"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"<how_do_I_make_change_18>"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"<i_get_an_error_messa_19>"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"<i_want_to_request_an_20>"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"<where_do_i_check_the_21>"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"<you_arent_helping_22>"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: 1}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test10.csv", ff: 0}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = tk.normalize(t.text)
			}
			tr := CollectTransitions(tx, TokenSplit(tk))
			for i, t := range tx {
				tokens := tk.tokenize(t.text)
				tx[i].chunk = TransitionChunk(tokens, tokens, tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			rules = SetIDs(rules)
			res := ExpressionFactor(tt.args.ff, nilLogger)(rules)
			SortPRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSynonymFactor(t *testing.T) {
	type args struct {
		f  string
		ff string
	}
	tests := []struct {
		args args
		want []Rule
	}{
		{args: args{f: "./data/tests/test5.csv", ff: "./data/tests/syn1.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST1", "TEST2"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: "./data/tests/syn2.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST5", "TEST6"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{}, root: []string{"My bill", "TEST7"}, suf: []string{}, isPublic: false, id: 3}, {pre: []string{}, root: []string{"my bill"}, suf: []string{}, isPublic: false, id: 4}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: "./data/tests/syn3.json"}, want: []Rule{{pre: []string{}, root: []string{"TEST3", "TEST4", "a"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test6.csv", ff: "./data/tests/syn1.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST1", "TEST2"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test6.csv", ff: "./data/tests/syn2.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST5", "TEST6"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"My bill", "TEST7"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"my bill"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service ?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test6.csv", ff: "./data/tests/syn3.json"}, want: []Rule{{pre: []string{}, root: []string{"TEST3", "TEST4", "a"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make <TEST3_TEST4_a_12> payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund ?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills ?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for <TEST3_TEST4_a_12> service ?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test7.csv", ff: "./data/tests/syn1.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST1", "TEST2"}, suf: []string{}, isPublic: false, id: 1}}},
		{args: args{f: "./data/tests/test8.csv", ff: "./data/tests/syn2.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST5", "TEST6"}, suf: []string{}, isPublic: false, id: 1}, {pre: []string{}, root: []string{"My bill", "TEST7"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{}, root: []string{"my bill"}, suf: []string{}, isPublic: false, id: 3}}},
		{args: args{f: "./data/tests/test9.csv", ff: "./data/tests/syn3.json"}, want: []Rule{{pre: []string{}, root: []string{"TEST3", "TEST4", "a"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have <TEST3_TEST4_a_12> question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download <TEST3_TEST4_a_12> bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I want to make <TEST3_TEST4_a_12> review for <TEST3_TEST4_a_12> service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make <TEST3_TEST4_a_12> payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: "./data/tests/syn1.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST1", "TEST2"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{""}, root: []string{"<I_want_to_know_what__12> the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: "./data/tests/syn2.json"}, want: []Rule{{pre: []string{}, root: []string{"I want to know what", "TEST5", "TEST6"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"My bill", "TEST7"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"my bill"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{""}, root: []string{"<I_want_to_know_what__12> the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address ?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options ?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test10.csv", ff: "./data/tests/syn3.json"}, want: []Rule{{pre: []string{}, root: []string{"TEST3", "TEST4", "a"}, suf: []string{}, isPublic: false, id: 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = tk.normalize(t.text)
			}
			tr := CollectTransitions(tx, TokenSplit(tk))
			for i, t := range tx {
				tokens := tk.tokenize(t.text)
				tx[i].chunk = TransitionChunk(tokens, tokens, tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			syn, _ := ReadSynonyms(tt.args.ff)
			rules = SetIDs(rules)
			res := SynonymFactor(syn, tk, nilLogger)(rules)
			SortPRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestConstituencyFactor(t *testing.T) {
	type args struct {
		f  string
		ff int
	}
	tests := []struct {
		args args
		want []Rule
	}{
		{args: args{f: "./data/tests/test5.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 2}, {pre: []string{""}, root: []string{"<I_don't_have_an_onli_2>"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: 1}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test5.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}}},
		{args: args{f: "./data/tests/test6.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account", "I want an online accoynt"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"I don't understand you"}, suf: []string{}, isPublic: false, id: 18}, {pre: []string{}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{}, isPublic: false, id: 19}, {pre: []string{}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{}, isPublic: false, id: 20}, {pre: []string{}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"can you show me my invoices?"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{}, isPublic: false, id: 15}, {pre: []string{}, root: []string{"i dont want my profile"}, suf: []string{}, isPublic: false, id: 17}, {pre: []string{}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{}, isPublic: false, id: 16}, {pre: []string{}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{}, isPublic: false, id: 21}, {pre: []string{""}, root: []string{"<I_don't_understand_y_18>"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"<I_got_an_error_messa_19>"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"<ask_an_agent_to_noti_20>"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"<can_you_show_me_info_13>"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"<can_you_show_me_my_i_14>"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"<can_you_tell_me_how__15>"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"<i_dont_want_my_profi_17>"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"<i_want_to_know_wat_t_16>"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"<where_can_i_leave_an_21>"}, suf: []string{""}, isPublic: true, id: 10}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}}},
		{args: args{f: "./data/tests/test6.csv", ff: 1}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account", "I want an online accoynt"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test6.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I don't understand you"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I got an error message when I attempted to make a payment"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want an online accoynt"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"ask an agent to notify issues with my payment"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"can you show me information about the status of my refund?"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"can you show me my invoices?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"can you tell me how I can get some bills?"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i dont want my profile"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"i want to know wat the email of Customer Service is"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"where can i leave an opinion for a service?"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test7.csv", ff: 0}, want: []Rule{}},
		{args: args{f: "./data/tests/test8.csv", ff: 0}, want: []Rule{}},
		{args: args{f: "./data/tests/test9.csv", ff: 0}, want: []Rule{{pre: []string{}, root: []string{"I don't have an online account"}, suf: []string{}, isPublic: false, id: 20}, {pre: []string{}, root: []string{"I have a question"}, suf: []string{}, isPublic: false, id: 18}, {pre: []string{}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{}, isPublic: false, id: 19}, {pre: []string{}, root: []string{"I want to download a bill"}, suf: []string{}, isPublic: false, id: 15}, {pre: []string{}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{}, isPublic: false, id: 14}, {pre: []string{}, root: []string{"I want to make a review for a service"}, suf: []string{}, isPublic: false, id: 16}, {pre: []string{}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{}, isPublic: false, id: 21}, {pre: []string{}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{}, isPublic: false, id: 13}, {pre: []string{}, root: []string{"i want to request an invoice"}, suf: []string{}, isPublic: false, id: 12}, {pre: []string{}, root: []string{"where do i check the delivery options?"}, suf: []string{}, isPublic: false, id: 22}, {pre: []string{}, root: []string{"you arent helping"}, suf: []string{}, isPublic: false, id: 17}, {pre: []string{""}, root: []string{"<I_don't_have_an_onli_20>"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"<I_have_a_question_18>"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"<I_ordered_an_item_an_19>"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"<I_want_to_download_a_15>"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"<I_want_to_know_what__14>"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"<I_want_to_make_a_rev_16>"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"<how_do_I_make_change_21>"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"<i_get_an_error_messa_13>"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"<i_want_to_request_an_12>"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"<where_do_i_check_the_22>"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"<you_arent_helping_17>"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: 1}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test9.csv", ff: 10}, want: []Rule{{pre: []string{""}, root: []string{"I don't have an online account"}, suf: []string{""}, isPublic: true, id: 0}, {pre: []string{""}, root: []string{"I have a question"}, suf: []string{""}, isPublic: true, id: 1}, {pre: []string{""}, root: []string{"I ordered an item and Id like to modify my fucking order"}, suf: []string{""}, isPublic: true, id: 2}, {pre: []string{""}, root: []string{"I want to download a bill"}, suf: []string{""}, isPublic: true, id: 3}, {pre: []string{""}, root: []string{"I want to know what the number of Customer Service is"}, suf: []string{""}, isPublic: true, id: 4}, {pre: []string{""}, root: []string{"I want to make a review for a service"}, suf: []string{""}, isPublic: true, id: 5}, {pre: []string{""}, root: []string{"how do I make changes to my shipping address?"}, suf: []string{""}, isPublic: true, id: 6}, {pre: []string{""}, root: []string{"i get an error message when i ty to make a payment for my order"}, suf: []string{""}, isPublic: true, id: 7}, {pre: []string{""}, root: []string{"i want to request an invoice"}, suf: []string{""}, isPublic: true, id: 8}, {pre: []string{""}, root: []string{"where do i check the delivery options?"}, suf: []string{""}, isPublic: true, id: 9}, {pre: []string{""}, root: []string{"you arent helping"}, suf: []string{""}, isPublic: true, id: 10}}},
		{args: args{f: "./data/tests/test10.csv", ff: 0}, want: []Rule{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			tk := NewWordTokenizer()
			file, _ := os.Open(tt.args.f)
			defer file.Close()
			s := bufio.NewScanner(file)
			tx := ReadTexts(s)
			for i, t := range tx {
				tx[i].text = tk.normalize(t.text)
			}
			tr := CollectTransitions(tx, TokenSplit(tk))
			for i, t := range tx {
				tokens := tk.tokenize(t.text)
				tx[i].chunk = TransitionChunk(tokens, tokens, tr, 0.1)
			}
			ng := CollectChunks(tx)
			for i, t := range tx {
				tx[i] = ToTriplet(t, ng)
			}
			rules := []Rule{}
			for _, t := range tx {
				rules = append(rules, ToRule(t))
			}
			rules = SetIDs(rules)
			mod := tag.NewPerceptronTagger()
			tag := NewSyntacticTagger(mod, tk)
			res := ConstituencyFactor(tag, tt.args.ff, nilLogger)(rules)
			SortPRS(res)
			assert.Equal(t, tt.want, res)
		})
	}
}
