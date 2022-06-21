# QuecOpen Golang Development Guide

## Introduction

This repository will introduce how to develope Golang based application with Quectel QuecOpen modules which support Linux based SDK: LTE modules (EC21-E, EC25-E, ... ), Automotive modules (AG35-E, AG35-NA, ... ), 5G modules (RG500Q, ... ).

The code base is tested with EC25-E SDK# EC25EFAR06A01M4G_OPCU_01.001.06, however it should works with other modules too as they mostly share the same API.

Note: This repo does not cover basic QuecOpen development guide, pls refer to the *getting started* manual.

## Why develop *GoLang* based application

For example, user wants to create MQTT client application, below is the normal approach:

1) Check SDK if there is any MQTT like *libmosquitto*, ... As this is missing, user need to find a way to install missing pacakges: header/libs for development and libs into rootfs
2) Create normal C application which use headers/libs from the generated libs


Above process becomes more complicated once package requires more dependencies or existing dependencies do not meet package rquirements, ...

With *Golang*  based approach, customer can create great application as:
- Can take advantage of many Golang packages to build application without worrying about addition libs, runtime environment, ... as application is a single binary file
- Can call C based libs to use Quectel API to mange network, simcard, ...
- Easy to maintain in the future
- ...

## Usage

Application development can be done in PC and cross compile to target ARM Linux, for example with `mqtt-client`

```
make
push dist/mqtt-client-linux-arm /usrdata
```

Execute application inside QuecOpen modules:
```
adb shell
cd /usrdata
chmod a+x mqtt-client-linux-arm
./mqtt-client-linux-arm
```

Result:
```
root@mdm9607-perf:/usrdata# ./mqtt-client-linux-arm
Received message on topic:  /QuecOpen/action  with payload:  Hello World
Received message on topic:  /QuecOpen/action  with payload:  Hello World
``

PS: Module `data-call` need to be executed so the `mqtt-client` can reach out to the broker.
