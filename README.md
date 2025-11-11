# Corpus -> Grammar

## *Condensing text corpora to context free grammars*

---

[![Open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/ryancahildebrandt/c2g)
[![This project contains 0% LLM-generated content](https://brainmade.org/88x31-dark.png)](https://brainmade.org/)

## Purpose

This project provides tools for converting a text corpus to a JSGF grammar. When used in conjunction with [gsgf](https://github.com/ryancahildebrandt/gsgf) or another grammar-to-production tool, you can edit, update, and expand text corpora in a predictable, reproducible way.

---

## Approach

At a high level, the program:

1) Reads in a newline-delimited corpus file (either txt or csv) and stores each line as its own text 
2) Splits each text into tokens
3) Chunks each text into 3 parts based on the transitional probabilities from all corpus texts
4) Converts each chunked text into a grammar rule
5) Combines rules based on shared chunks or other similarity criteria
6) Outputs the resulting rules to a jsgf grammar file

### Transitional Probabilities

Core to this project is the identification of commonly occurring chunks of text within a corpus. One simple approach is to use transitional probabilities between tokens to split a text up into groups of tokens that tend to be seen together. The sentence "send me three hundred dollars" might be word-tokenized as 
```
["send", "me", "three", "hundred", "dollars"]
```
and the transitional probabilities (as calculated from the entire corpus) might be 
```
("send", "me", 0.5), ("me", "three", 0.1), ("three", "hundred", 0.7), ("hundred", "dollars", 0.8)
```
where each probability is the chance that the 1st token will be followed by the second. From here, we can set a threshold below which we'll split tokens into chunks. With a threshold of 0.2, we split the tokens into 
```
["send me", "three hundred dollars"]
```
We count each of these as a chunk, and count how frequently they occur accross the entire corpus.This approach works for any sequence, including part of speech and constituency tags (scroll down for more on those).

### Text Structure

One of the goals of this tool is to keep human readability and interpretability as high as possible. To keep the grammar/rule structure simple, I decided to split texts into 3 chunks: a prefix, root, and suffix. Root chunks are set first, generally prioritizing the largest chunk found in the corpus. Again taking the sentence
```
["send", "me", "three", "hundred", "dollars"]
```
let's assume the corpus wide frequency of each chunk is
```
("send me", 10), ("three hundred dollars", 6)
```
So we would assign the longer chunk to the root, and the remaining bits of the sentence to the prefix and suffix
```
Text{prefix: "send me", root: "three hundred dollars", suffix: ""}
```

As in this example, it's not uncommon for the root to cover the beginning or end of a text, in which case the prefix and/or suffix will be empty. This structure is carried over to grammar rules for merging and factoring.

### Constituency Tagging

Go doesn't have quite as much support for natural language processing tasks when compared to something like Python or Ruby, but [prose](https://github.com/jdkato/prose) provides some really useful core utilities. The part of speech tagger is exceptionally accurate, and while prose doesn't have support for constituency tagging, we can get some pretty common rules for constituency tags from the Penn Treebank corpus included in [NLTK](https://www.nltk.org/). Based on those rules, we can tag each text with constituency tags as well as part of speech tags.

### Rule Merging

Grammar compression is primarily achieved by merging rules with shared chunks. For the rules
```
1) Rule{prefix: "send me", root: "three hundred dollars", suffix: ""}
2) Rule{prefix: "send me", root: "three hundred dollars", suffix: "now"}
3) Rule{prefix: "send me", root: "nothing", suffix: "please"}
```
We could merge 1 and 2 based on the shared prefix and root to get
```
4) Rule{prefix: "send me", root: "three hundred dollars", suffix: ["now", ""]}
```
Or 1/2 and 3 based on the shared prefix to get
```
5) Rule{prefix: "send me", root: ["three hundred dollars", "nothing], suffix: ["now", "please"]}
```
Important here is the behavior when merging rules based on one shared chunk versus two. If we return all permutations of the rule chunks for rules 4 and 5, we get
```
4)
send me three hundred dollars now
send me three hundred dollars
5)
send me three hundred dollars now
send me three hundred dollars please
send me nothing now
send me nothing please
```
The productions from rule 4 only cover what is provided in the corpus-derived rules. If we had done no merging, we would still get "send me three hundred dollars now" and "send me three hundred dollars" from rules 1 and 2. However, merging based on only one shared chunk produces sentences not found in the original rules. None of the original 3 rules would have produced "send me three hundred dollars please", but after merging this is a valid production. Controlling how/when rules are merged allows for different outputs from the same source corpus and derived rules

Depending on the merging strategy used, we can use different criteria to check if chunks are similar enough to be merged. C2G allows for merging based on exact matches, edit distance, syntax tag matches, and embedding distances.

### Grammar Expansion

Once the grammar has been constructed and all merging is complete, we can expand the grammar with common synonyms for key terms. Here, that takes the form of creating a new rule in the grammar 
```
<dollars> = (dollars|$$$|bucks);
```
and factoring out all occurrences of the target term with this rule. The original rule
```
<rule1> = (send me) (three hundred dollars)
```
would become
```
<rule1> = (send me) (three hundred <dollars>)
```
and would produce
```
send me three hundred dollars
send me three hundred $$$
send me three hundred bucks
```

You can define as many synonyms as is useful, and synonyms can cover single word or multi word phrases.

---

## Implementation Notes

- Some grammar construction methods will only result in productions found in the grammar, while some will result in productions not seen in the original corpus. These are noted in the c2g executable help text
- In the grammar (and subsequent productions), consecutive whitespaces will be replaced with a single space, except before punctuation
- Constituency rules derived from Penn Treebank are far from exhaustive, and may not reflect an optimal resolution order. External tools would provide better constituency tagging, but are outside of the scope of this project

---

## Usage

```shell
# show general or command specific help (-h flag optional)
c2g [clone|compress|interpolate|extrapolate] [-h]

# convert example.csv to grammar and save to out.jsgf
c2g clone -outFile=out.jsgf  example.csv

# convert example.csv to a grammar, merging based on POS tags, logging to ./log, and factoring chunks occurring more than 10 times
c2g compress -chunk=posTag -logfile=log -factorN=10 example.csv

# convert example.csv to a grammar, merging rules with 1 or more shared chunks, factoring based on constituency tags, and filtering out texts matching the bottom 5% of constituency tag structures in the corpus 
c2g interpolate -preTokenized -conFactor -filter=0.05 example.csv

# convert example.csv to a grammar, merging rules with 1 or more shared chunks, and expanding with synonyms from syn.json
c2g extrapolate -synFile=syn.json example.csv

# convert example.csv to a grammar, merging rules with 2 shared chunks and factoring expression groups with more than 200 occurrences
c2g custom -merge2 -factor -factorN=200 example.csv
```

---

## Dataset

The dataset used for the current project was pulled from the following:

- [Bitext](https://www.kaggle.com/datasets/bitext/training-dataset-for-chatbotsvirtual-assistants), mostly for testing


---

## Outputs

- [c2g](./c2g) executable
- [Example](./data) corpora
- [Penn Treebank](./penn.py) scraping script
