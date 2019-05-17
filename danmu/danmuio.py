#!/usr/bin/env python3

import time, sys

from danmu import DanMuClient

def pp(msg):
    print(msg.encode(sys.stdin.encoding, 'ignore').
        decode(sys.stdin.encoding))

#dmc = DanMuClient('https://www.douyu.com/9999') QXJ401, 52876, 687423, 2009, 74960, 208114, 88660
dmc = DanMuClient('https://www.douyu.com/284584')
if not dmc.isValid(): print('Url not valid')

@dmc.danmu
def danmu_fn(msg):
    pp('[%s] %s' % (msg['NickName'], msg['Content']))

#@dmc.gift
#def gift_fn(msg):
#    pp('[%s] sent a gift!' % msg['NickName'])

#@dmc.other
#def other_fn(msg):
#    pp('Other message received')

dmc.start(blockThread = True)
