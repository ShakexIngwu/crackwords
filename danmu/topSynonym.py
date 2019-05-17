#!/usr/bin/env python3
import sys
#sys.path.append("/home/ubuntu/danmu/")
import getopt
from gensim.models import word2vec
from gensim.models import KeyedVectors


def help():
    print ("************Usage of topSynonym*************")
    print ("-w:    the word that needs to be searched for synomyms")

def main(argv):
    #Set variables
    word = "666"

    #get passed arguments
    try:
        opts, args = getopt.getopt(argv, "w:", "word=")
    except getopt.GetoptError:
        help()
        sys.exit(2)

    for opt, arg in opts:
        if opt in ("-w", "--word"):
            word = arg

    #load the pre-trained word2vec vector set, and get most similar words
    model = KeyedVectors.load_word2vec_format('/home/ubuntu/danmu/corpusWord2Vec.bin', binary=True)
    indexes = model.most_similar(word, topn=10)
    for index in indexes:
        print(index[0])

if __name__ == "__main__":
    main(sys.argv[1:])
