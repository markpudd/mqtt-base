# mqtt-base
MQTT Base

This is the base of some MQTT tools.   The intended use of these are for a MQTT gateway project.

Currently this code mainly deals with packaging and unpackaging MQTT packets.

Each packet type has a related go file and test.   There is a also data reader which is intended to read a stream of bytes, generate the relevant packet structure and send to a Processors process method.
