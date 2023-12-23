#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>

#define SERVER_PORT 27777
#define SERIAL_NUM "1234567890"

int main() {
    int sock;
    struct sockaddr_in server_addr;
    char serialNum[] = SERIAL_NUM;

    // Create socket:
    sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        printf("Error in socket creation\n");
        exit(1);
    }

    // Configure settings of the server address struct
    server_addr.sin_family = AF_INET; // Address family
    server_addr.sin_port = htons(SERVER_PORT); // Port number
    server_addr.sin_addr.s_addr = inet_addr("127.0.0.1"); // Set IP address to localhost 

    // Connect to the server
    if (connect(sock, (struct sockaddr *)&server_addr, sizeof(server_addr)) < 0) {
        printf("Connection Failed\n");
        return -1;
    }
    printf("Connected to Server\n");

    // Send the serial number to server
    if (write(sock, serialNum, strlen(serialNum)) != strlen(serialNum)) {
        printf("Error sending serial number\n");
        return -1;
    }

    // Send random int to server in loop
    for (;;) {
        int sensorVal = rand();
        uint32_t buffer = htonl(sensorVal);
        if (write(sock, &buffer, sizeof(buffer)) != sizeof(buffer)) {
            printf("Error sending sensor value\n");
            return -1;
        }
        printf("Send value: %d\n", sensorVal);
        sleep(5);
    }
    return 0;
}
