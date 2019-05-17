#!/usr/bin/env python3

import numpy as np
import matplotlib
import pandas as pd
import matplotlib.pyplot as plt
import matplotlib.font_manager
from sklearn.decomposition import PCA
from gensim.models import word2vec
from gensim.models import KeyedVectors

model = KeyedVectors.load_word2vec_format('/home/ubuntu/danmu/corpusWord2Vec.bin', binary=True)
vocab = model.wv.vocab
X=model[vocab]

plt.rcParams['font.sans-serif'] = ['SimHei']
plt.rcParams['savefig.dpi'] = 300
plt.rcParams['axes.unicode_minus'] = False

#reduce dimension
X_reduced = PCA(n_components=2).fit_transform(X)
df = pd.DataFrame(X_reduced, index=vocab, columns=['x','y'])

fig = plt.figure()
ax = fig.add_subplot(1, 1, 1)

ax.scatter(df['x'], df['y'])

for word, pos in df.iterrows():
    ax.annotate(word, pos)

plt.savefig('vector.png')

