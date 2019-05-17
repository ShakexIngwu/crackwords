#!/bin/bash

cd /home/ubuntu/danmu

TEXT=$(curl -XGET "localhost:9200/danmu*/_search?pretty" -H "Content-Type:application/json" -d' 
	{
		"query": {
			"range": {
				"@timestamp": {
					"gte": "now-5m/m",
					"lt": "now"
				}
			}
		},
		"size": 10000,
		"_source": [ "message" ]
	}
	')

echo $TEXT |jq -M . |jq -r '.hits |.hits |.[] |._source |.message' > recentWords.txt

./getSepedRecentWords.py
