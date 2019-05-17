#!/usr/bin/env python3

import jieba

filePath = 'content.txt'
fileSegWordDonePath = 'corpusSegDone.txt'
fileTrainRead = []
#read the file by line
with open(filePath) as fileTrainRaw:
    for line in fileTrainRaw:
        fileTrainRead.append(line)

#segment word with jieba
fileTrainSeg = []
for i in range(len(fileTrainRead)):
    fileTrainSeg.append([' '.join(list(jieba.cut(fileTrainRead[i], cut_all=False)))])

print(fileTrainSeg[0][0])
#save the result
with open(fileSegWordDonePath,'w') as fw:
    for i in range(len(fileTrainSeg)):
        fw.write(str(fileTrainSeg[i][0]))
        fw.write('\n')
