# Notifier

Creating a service for sending notifications using **MQTT**.

## What does this project do?
In this project, we have a notifier service (written in _golang_) that allows
us to send notifications to our clients using **MQTT**.

By setting the _topic_, you can send data into different channels. On the other
hand, the clients will receive data from those channels.

## What is MQTT?
MQTT (_MQ Telemetry Transport_) is a lightweight 
open messaging protocol that provides resource-constrained 
network clients with a simple way to distribute telemetry 
information in low-bandwidth environments. 

The protocol, which employs a **publish/subscribe** 
communication pattern, is used for machine-to-machine 
(M2M) communication.

Imagine the following example:

<img src="./assets/MQTT_Schema_EN.jpeg" />

The temperature sensor (**publisher**) will send the data to _MQTT-broker_
and the broker will send the data to clients (**subscribers**).

## How to use?
Clone into the project and enter the project root:
```shell
git clone https://github.com/amirhnajafiz/Notifier.git
cd Notifier
```

Copy the config file into your main root:
```shell
cp ./configs/config.example.yml config.yml
```

Now reset the configurations to run the project in your system.

### Notifier Server
Now run the notifier server with following command:
```shell
make dev
```

Response should be like this:
```shell
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.27.0                   │ 
 │               http://127.0.0.1:5000               │ 
 │       (bound on host 0.0.0.0 and port 5000)       │ 
 │                                                   │ 
 │ Handlers ............. 5  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID .............. 7443 │ 
 └───────────────────────────────────────────────────┘ 
```

Now you can send the notifications with the following route:
```json lines
POST 'localhost:[PORT]/api/send'
```

Body:
```json
{
    "topic":   "[optional]", 
    "message": "[data]"
}
```

Response:
```json
{
	"length": 7,
	"send from": "amir",
	"status": "200|OK",
	"time": "Sat, 26 Feb 2022 19:14:10 +0330"
}
```

### Client
You can set **MQTT** options and configurations for your client,
but if you don't change anything in **config.json**, it still works.

Run your clients with the following command:
```shell
make cli
```

Now the client is listening to the MQTT-Broker to receive data.
Once you send the notifications via notifier server, clients
will get them like this:
```
Sat, 26 Feb 2022 19:13:18 +0330 #Amirs-MacBook-Pro.local: test990
```

## Dependencies
