import nltk
import json

nltk.download('treebank')
productions = []
for f in nltk.corpus.treebank.fileids():
  for s in nltk.corpus.treebank.parsed_sents(f):
    productions += s.productions()

grammar = nltk.grammar.induce_pcfg(nltk.Nonterminal("S"), productions)
penn_nonterminals = ["ADJP","ADVP","CONJP","NP","PP","PRN","QP","VP"]

records = []
for nt in penn_nonterminals:
    res = [i for i in grammar.productions(lhs=nltk.Nonterminal(nt)) if i.prob() >= 0.01]
    for r in res:
        k = r.lhs().symbol()
        v = [i.symbol() for i in r.rhs()]
        skip = False
        if len(v) == 1:
            continue 
        for vv in v:
            if "-" in vv:
                skip = True
                break
        if skip:
            continue
        records.append(json.dumps({k:v}))

with open("./data/penn.jsonl", "w") as f:
    f.write("\n".join(records))

