#!/usr/bin/env python3

import jieba
import pandas as pd

filePath = '/home/ubuntu/danmu/recentWords.txt'
fileSegDone = '/home/ubuntu/danmu/corpusSegRecentWords.txt'
fileTrainRead = []
#read the file by line
with open(filePath) as fileTrainRaw:
    for line in fileTrainRaw:
        fileTrainRead.append(line)

#segment word with jieba
fileTrainSeg = []
for i in range(len(fileTrainRead)):
    fileTrainSeg.append([' '.join(list(jieba.cut(fileTrainRead[i], cut_all=False)))])

#save the result
with open(fileSegDone, 'w') as fw:
    for i in range(len(fileTrainSeg)):
        fw.write(str(fileTrainSeg[i][0]))
        fw.write('\n')
