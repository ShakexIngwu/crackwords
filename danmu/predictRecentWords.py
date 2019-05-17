#!/usr/bin/env python3

import re
from gensim.models import word2vec
from gensim.models import KeyedVectors
from operator import itemgetter

filePath = '/home/ubuntu/danmu/corpusSegRecentWords.txt'
fileTrainRead = []
#read the file by line
with open(filePath) as fileTrainRaw:
    for line in fileTrainRaw:
        fileTrainRead.append(line)

#load the pre-trained word2vec vector set
model = KeyedVectors.load_word2vec_format('/home/ubuntu/danmu/corpusWord2Vec.bin', binary=True)
#predict for each word and then calculate the most frequent topic word set
wordFreq = {}
for i in range(len(fileTrainRead)):
    words = fileTrainRead[i][0].split( )
    for j, word in enumerate(words):
       # word = re.sub("[\s+\.\!\/_,$%^*(+\"\']+|[+——！，。？、~@#￥%……&*（）]+".decode("utf8"), "",word)
       # word = re.sub("[【】╮╯▽╰╭★→「」]+".decode("utf8"),"",word)
       # word = re.sub("！，❤。～《》：（）【】「」？”“；：、".decode("utf8"),"",word)
        if not re.match(r"[【】╮╯▽╰╭★→「」\s+\.\!\/_,$%^*(+\"\']+|[+——！，。？、~@#￥%……&*（）！，❤。～《》：（）【】「」？”“；：、0-9a-zA-Z]+", word):
            try:
                similarWords = model.most_similar(word, topn=10)
                for idx, similarWord in enumerate(similarWords):
                    if similarWord[0] not in wordFreq:
                        wordFreq[similarWord[0]] = 1
                    else:
                        wordFreq[similarWord[0]] += 1
            except:
                pass

top10Words = [k for k in sorted(wordFreq.items(), key=itemgetter(1), reverse=True)[:10]]

for _, word in enumerate(top10Words):
    print (word[0])
