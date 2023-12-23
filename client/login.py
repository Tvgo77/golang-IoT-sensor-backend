import requests
import json
import os

url = "http://localhost:8080/login"

payload='email=test%40gmail.com&password=test'
headers = {"Content-Type": "application/x-www-form-urlencoded"}

response = requests.request("POST", url, headers=headers, data=payload)

if not response.ok:
  os._exit(1)

with open("loginToken.json", "w") as f:
  f.write(response.text)
