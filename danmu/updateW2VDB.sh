#!/bin/bash

cd /home/ubuntu/danmu
#seperate danmu content
./getText.py
#use jieba to seperate sentences
./sepWords.py
#use word2vec to train words set
./clusterWords.py
