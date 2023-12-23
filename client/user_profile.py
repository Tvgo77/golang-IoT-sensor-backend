import requests
import json

url = "http://localhost:8080/profile"

with open("loginToken.json", 'r') as fp:
  tokenDict = json.load(fp)

payload={}

headers = {"Authorization": "Bearer " + tokenDict["accessToken"]}

response = requests.request("GET", url, headers=headers, data=payload)

print(response.text)
