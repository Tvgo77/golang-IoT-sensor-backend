import requests
import json

url = "http://localhost:8080/updateSensor"

with open("loginToken.json", 'r') as fp:
  tokenDict = json.load(fp)

payload = "serial=1234567890&operation=add"
# payload = "serial=1234567890&operation=remove"


headers = {"Authorization": "Bearer " + tokenDict["accessToken"], 
           "Content-Type": "application/x-www-form-urlencoded"}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)