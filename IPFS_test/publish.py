
'''
This is the implementation of uploading and pinning file onto IPFS node using ipfshttpclient,
a library of Python
'''

import ipfshttpclient

# Share TCP connections using a context manager

with ipfshttpclient.connect() as client:
	res = client.add('spaceship-launch.jpg')
	hash = res["Hash"]
	print("IPFS hash is:", hash, "\nName:", res["Name"])


# Share TCP connections until the client session is closed
'''
class SomeObject:
	def __init__(self):
		self._client = ipfshttpclient.connect(session=True)

	def do_something(self):
		hash = self._client.add('spaceship-launch.jpg')['Hash']
		print(self._client.ls(hash))

	def close(self):  # Call this when your done
		self._client.close()

first_obj = SomeObject()    
first_obj.do_something()
first_obj.close()
'''