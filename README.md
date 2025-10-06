# Corpus -> to -> Grammar

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
2) Splits each text of text into tokens
3) Chunks each text into 3 parts based on the transitional probabilities from all corpus texts
4) Converts each chunked text into a grammar rule
5) Combines rules based on shared chunks or other similarity criteria
6) Outputs the resulting rules to a jsgf grammar file

### Transitional Probabilities

Core to this project is the identification of commonly occurring chunks of text within a corpus. One simple approach is to use transitional probabilities between tokens to split a text up into groups of tokens that tend to be seen together. The sentence "send me three hundred dollars" might be word-tokenized as ["send", "me", "three", "hundred", "dollars"] and the transitional probabilities (as calculated from the entire corpus) might be ("send", "me", 0.5), ("me", "three", 0.1), ("three", "hundred", 0.7), ("hundred", "dollars", 0.8) where each probability is the chance that the 1st token will be followed by the second. From here, we can set a threshold below which we'll split tokens into chunks. With a threshold of 0.2, we split the tokens into ["send me", "three hundred dollars"]. We count each of these as a chunk, and count how frequently they occur accross the entire corpus for chunking purposes.

This approach works for any sequence, including part of speech and constituency tags (scroll down for more on those).

### Text Structure

One of the goals of this tool is to keep human readability and interpretability as high as possible. To keep the grammar/rule structure simple, I decided to split texts into 3 chunks: a prefix, root, and suffix. Root chunks are set first, generally prioritizing the largest/most frequently occurring chunk found in the corpus. From here, the prefix and suffix are set as all of the text ocurring before/after the root chunk. It's not uncommon for the root to cover the beginning or end of a text, in which case the prefix and/or suffix will be empty. This structure is carried over to grammar rules for merging and factoring.

### Constituency Tagging

Go doesn't have quite as much support for natural language processing tasks when compared to something like Python or Ruby, but [prose](https://github.com/jdkato/prose) provides some really useful core utilities. The part of speech tagger is exceptionally accurate, and while prose doesn't have support for constituency tagging, we can get some pretty common rules for constituency tags from the Penn Treebank corpus included in [NLTK](https://www.nltk.org/). Based on those rules, we can tag each text with constituency tags as well as part of speech tags.

### Rule Merging

Grammar compression is primarily achieved by merging rules with shared chunks. Depending on the merging strategy used, we can use different criteria to check if chunks are similar enough to be merged. C2G allows for merging based on exact matches, edit distance, syntax tag matches, and embedding distances.

### Grammar Expansion

Once the grammar has been constructed and all merging is complete, we can expand the grammar with common synonyms for key terms. Here, that takes the form of creating a new rule in the grammar (i.e. \<term\> = synonym1|synonym2...;) and factoring out all occurrences of the target term with this rule. When creating productions, there will be sentences with the term and all synonyms. You can define as many synonyms as is useful, and synonyms can cover single word or multi word phrases.

---

## Implementation Notes

- Some grammar construction methods will only result in productions found in the grammar, while some will result in productions not seen in the original corpus. These are noted in the c2g executable help text
- In the grammar (and subsequent productions), consecutive whitespaces will be replaced with a single space, except before punctuation
- Constituency rules derived from Penn Treebank are far from exhaustive, and may not reflect an optimal resolution order. External tools would provide better constituency tagging, but are outside of the scope of this project

---

## Dataset

The dataset used for the current project was pulled from the following:

- [Bitext](https://www.kaggle.com/datasets/bitext/training-dataset-for-chatbotsvirtual-assistants), mostly for testing

---

## Usage

```shell
# show general or command specific help (-h flag optional)
c2g [clone|compress|interpolate|extrapolate] [-h]

# convert example.csv to grammar and save to out.jsgf
c2g clone -outFile=out.jsgf  example.csv

# convert example.csv to a grammar, merging based on POS tags, logging to ./log, and factoring chunks occurring more than 10 times
c2g compress -chunk=posTag -logfile=log -factor=10 example.csv

# convert example.csv to a grammar, merging rules with 1 or more shared chunks, factoring based on constituency tags, and only considering texts matching the top 95% of constituency tag structures in the corpus 
c2g interpolate -preTokenized -conFactor -filterQuantile=0.95example.csv

# convert example.csv to a grammar, merging rules with 1 or more shared chunks, and expanding with synonyms from syn.json
c2g extrapolate -synFile=syn.json example.csv
```

---

## Outputs

- [c2g](./c2g) executable
- [Example](./data) corpora
- [Penn Treebank](./penn.py) scraping script
