# golang-IoT-sensor-backend
Sensor side:
    Sensors send real-time monitored data to server every 5 seconds when activated.
    data format:
    {
        "serialNum": "xxxxxx" (string)
        "sensorVal": val (int[]) 
        "time": time   
    }

Server side:
    Consistently Receive data from sensors.
    Maintain sensor data and user data.
    Respond user's request for sensor data.

    server
    |-- SensorManager
    |   |-- ConnHandler
    |   |   |-- OnConnect (Create new entry in channel map)
    |   |   |-- OnData (push data to channel)
    |   |   \-- OnClose (delete the entry)
    |   
    |-- DataChannelMap (serialNum(int): channel(chan))
    |
    \-- UserManager
        |-- middleware
        |-- router
        |   |-- signup
        |   |-- login
        |   \-- task
        |       |-- addSensor
        |       |-- deleteSensor
        |       |-- getSensorVal
        |-- usecase
        \-- repository
    

Client side:
    Add or delete monitoring sensors.
    Fetch realtime sensor data.
