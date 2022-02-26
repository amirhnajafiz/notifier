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

## Dependencies
