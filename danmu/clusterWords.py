#!/usr/bin/env python3

from gensim.models import word2vec

content = word2vec.Text8Corpus('corpusSegDone.txt')
model = word2vec.Word2Vec(content, size=300, window=5)

model.wv.save_word2vec_format('corpusWord2Vec.bin', binary=True)
