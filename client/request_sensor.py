import requests
import json
import os

url = "http://localhost:8080/requestSensor"

with open("loginToken.json", 'r') as fp:
  tokenDict = json.load(fp)

payload = {}

headers = {"Authorization": "Bearer " + tokenDict["accessToken"]}

response = requests.request("GET", url, headers=headers, data=payload)

if not response.ok:
  os._exit(1)

with open("oneTimeToken.json", "w") as f:
  f.write(response.text)