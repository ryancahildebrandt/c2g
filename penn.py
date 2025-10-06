import nltk
import json

nltk.download('treebank')
productions = []
for f in nltk.corpus.treebank.fileids():
  for s in nltk.corpus.treebank.parsed_sents(f):
    productions += s.productions()

grammar = nltk.grammar.induce_pcfg(nltk.Nonterminal("S"), productions)
penn_nonterminals = ["ADJP","ADVP","CONJP","INTJ","LST","NAC","NP","NX","PP","PRN","PRT","QP","RRC","UCP","VP","WHADJP","WHAVP","WHNP","WHPP"]
prose_terminals = ["(",")",",",":",".","''","``","#","$","CC","CD","DT","EX","FW","IN","JJ","JJR","JJS","LS","MD","NN","NNP","NNPS","NNS","PDT","POS","PRP","PRP$","RB","RBR","RBS","RP","SYM","TO","UH","VB","VBD","VBG","VBN","VBP","VBZ","WDT","WP","WP$","WRB",]

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
            if vv not in prose_terminals:
                skip = True
                break
        if skip:
            continue
        records.append(json.dumps({k:v}))

with open("./data/penn.jsonl", "w") as f:
    f.write("\n".join(records))

