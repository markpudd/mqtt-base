# MQTT Base

This is a marshal/unmarshal library for MQTT.   The intended use of these are for a MQTT gateway project, however they could be used elsewhere.

Each packet type has a related go file and test.   There is a also data reader which is intended to read a stream of bytes, generate the relevant packet structure and send to a Processors process method.

To use you can create a new packet using New{packet_type}Packet, setup the data you need and then call Marshal which will return a byte array.

To unmarshal use the Unmarshal method on data_reader.go, this will return a base Packet type which can me cast ie:-

      `connect = packet.(\*mqttbase.ConnectPacket)`

Feel free to send pull request if you make changes.
