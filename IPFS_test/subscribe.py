'''
This is the implementation of downloading file onto IPFS node using ipfshttpclient,
a library of Python
'''

import ipfshttpclient

# Connect to the IPFS daemon
client = ipfshttpclient.connect()

hash_value = input("Please input your CID:")
client.get(hash_value, "receive_test")
