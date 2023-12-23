import requests

url = "http://localhost:8080/signup"

payload='email=test%40gmail.com&password=test&name=Test%20Name'
headers = {"Content-Type": "application/x-www-form-urlencoded"}

response = requests.request("POST", url, headers=headers, data=payload)

print(response.text)
