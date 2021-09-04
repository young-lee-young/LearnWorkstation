import requests

data = {
    "request_id": "ccb515a7-0925-4ddf-94df-f2ce13003663",
    "phone_id": "uphone-lee",
    "candidates": ["1", "2"]
}

cert = ("/Users/leeyoung/WorkStation/Code/signal-server/tls/server.crt", "/Users/leeyoung/WorkStation/Code/signal-server/tls/server.key")
verify = "/Users/leeyoung/WorkStation/Code/signal-server/tls/server.crt"

response = requests.post("https://127.0.0.1:443/api/v1/send_candidate", data=data, cert=cert)
print response
