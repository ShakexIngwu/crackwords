#!/usr/bin/env python3

import pandas as pd

rawData = pd.read_csv('danmu.csv', delimiter = ',')
content = rawData['Content']
content.to_csv(r'content.txt', sep='\t', index = None, header=False, encoding='utf-8') 
