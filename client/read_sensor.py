import socket
import json
import struct
import time

def recvall(sock, size):
    data = bytearray()
    while len(data) < size:
        packet = sock.recv(size - len(data))
        if not packet:
            raise Exception("Connection closed")
        data.extend(packet)
    return data
    
server_ip = '127.0.0.1'  
server_port = 7777        
with open("oneTimeToken.json", 'r') as fp:
  oneTimeToken = json.load(fp)

# Create a socket object
client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# Connect to the server
client_socket.connect((server_ip, server_port))


# Send Verification data
verify_message = oneTimeToken["userId"] + oneTimeToken["oneTimeToken"]
client_socket.sendall(verify_message.encode())

# Receive sensor data
while True:
  response1 = recvall(client_socket, 10)
  response2 = recvall(client_socket, 4)
  sensorVal = struct.unpack('>I', response2)[0]
  print(f"Received from sensor {response1.decode()}: {sensorVal}")


