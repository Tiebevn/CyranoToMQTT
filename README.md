# CyranoToMQTT

A UDP listener that parses Cyrano (Ethernet Fencing Protocol) messages and publishes them to MQTT.

## Requirements
- Go 1.24+
- Docker (optional, for container build/run)

## Build and Run (local)
```bash
cd /Users/tiebevn/Developer/CyranoToMQTT
go build -o cyrano-to-mqtt .
MQTT_BROKER_URL=tcp://localhost:1883 \
MQTT_TOPIC_BASE=cyrano \
./cyrano-to-mqtt
```

## Environment Variables
- `MQTT_BROKER_URL` (default `tcp://localhost:1883`)
- `MQTT_CLIENT_ID` (default `cyrano-to-mqtt`)
- `MQTT_USERNAME` / `MQTT_PASSWORD` (optional)
- `MQTT_TOPIC_BASE` (default `cyrano`)
- `MQTT_QOS` (0,1,2; default 0)
- `MQTT_RETAIN` (`true`/`false`; default false)

## Docker
Build image locally:
```bash
docker build -t cyrano-to-mqtt:latest .
```

## Docker Compose (with Mosquitto)
```bash
docker compose up --build
```
- Mosquitto expects a config at `./mqtt-config` mounted to `/mosquitto/config/mosquitto.conf`.
- Cyrano listener exposes UDP `50103`.

## Testing
Subscribe to MQTT:
```bash
mosquitto_sub -h localhost -p 1883 -t 'cyrano/#' -v
```
Send a sample UDP packet:
```bash
echo "|EFP1|INFO|Piste 1|Comp|Phase|Poul|Match|Round|12:34|Running|I|Foil|A|F|Ref1|Alice|USA%%%" | nc -u -w1 127.0.0.1 50103
```
