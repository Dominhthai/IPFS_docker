import ipfshttpclient

client = ipfshttpclient.connect()  # Connects to: /dns/localhost/tcp/5001/http
res = client.add('test3.txt')
print(res)