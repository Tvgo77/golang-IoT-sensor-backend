# golang-IoT-sensor-backend
Sensor side:
    Sensors send real-time monitored data to server every 5 seconds when activated.
    data format: every data frame has following fields
    {
        "serialNum": "xxxxxxxxxx" (First 10 bytes should be serial string)
        "sensorVal": val (4 bytes for int value) 
        "time": time  (Optional)
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
    |-- DataChannelMap (serialNum(string): channel(chan))
    |
    \-- UserManager
        |-- ConnHandler (handler for direct tcp connection)
        |   /* User profile manage system */
        |-- middleware
        |-- router
        |   |-- signup
        |   |-- login
        |   |-- refresh_token
        |   |-- profile
        |   |-- update_sensor
        |   \-- sensor_data
        |
        |-- controller
        |-- domain
        |-- usecase
        \-- repository
    

Client side:
    Add or delete monitoring sensors.
    Fetch realtime sensor data.
