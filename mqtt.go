package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTPublisher struct {
	client    mqtt.Client
	topicBase string
	qos       byte
	retain    bool
}

func NewMQTTPublisher() (*MQTTPublisher, error) {
	brokerURL := os.Getenv("MQTT_BROKER_URL")
	if brokerURL == "" {
		brokerURL = "tcp://localhost:1883"
	}
	clientID := os.Getenv("MQTT_CLIENT_ID")
	if clientID == "" {
		clientID = "cyrano-to-mqtt"
	}
	username := os.Getenv("MQTT_USERNAME")
	password := os.Getenv("MQTT_PASSWORD")
	topicBase := os.Getenv("MQTT_TOPIC_BASE")
	if topicBase == "" {
		topicBase = "cyrano"
	}

	// Optional publish settings
	qos := byte(0)
	switch strings.TrimSpace(os.Getenv("MQTT_QOS")) {
	case "1":
		qos = 1
	case "2":
		qos = 2
	}
	retain := strings.EqualFold(os.Getenv("MQTT_RETAIN"), "true")

	opts := mqtt.NewClientOptions().AddBroker(brokerURL).SetClientID(clientID)
	if username != "" {
		opts.SetUsername(username)
	}
	if password != "" {
		opts.SetPassword(password)
	}
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetOrderMatters(false)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if !token.WaitTimeout(10 * time.Second) {
		return nil, fmt.Errorf("mqtt connect timeout")
	}
	if err := token.Error(); err != nil {
		return nil, fmt.Errorf("mqtt connect error: %w", err)
	}

	return &MQTTPublisher{client: client, topicBase: topicBase, qos: qos, retain: retain}, nil
}

func (p *MQTTPublisher) Close() {
	if p != nil && p.client != nil && p.client.IsConnected() {
		p.client.Disconnect(250)
	}
}

func (p *MQTTPublisher) PublishMessage(msg *CyranoMessage) error {
	if p == nil || msg == nil {
		return nil
	}

	// Build base topic: cyrano/{piste}/{command}
	baseTopic := fmt.Sprintf("%s/%s/%s", p.topicBase, sanitizeTopic(msg.Piste), msg.Command.String())

	// Publish each field to its own topic
	fields := map[string]string{
		"protocol":    msg.ProtocolVersion.String(),
		"command":     msg.Command.String(),
		"piste":       msg.Piste,
		"competition": msg.Competition,
		"phase":       msg.Phase,
		"poulTab":     msg.PoulTab,
		"match":       msg.Match,
		"round":       msg.Round,
		"time":        msg.Time,
		"stopwatch":   msg.Stopwatch,
		"type":        msg.Type.String(),
		"weapon":      msg.Weapon,
		"priority":    msg.Priority,
		"state":       msg.State.String(),

		// Referee information
		"ref/id":   msg.RefId,
		"ref/name": msg.RefName,
		"ref/nat":  msg.RefNat,

		// Right fencer
		"right/id":      msg.RightId,
		"right/name":    msg.RightName,
		"right/nat":     msg.RightNat,
		"right/score":   msg.RightScore,
		"right/status":  msg.RightStatus,
		"right/ycard":   msg.RightYCard,
		"right/rcard":   msg.RightRCard,
		"right/light":   msg.RightLight,
		"right/wlight":  msg.RightWLight,
		"right/medical": msg.RightMedical,
		"right/reserve": msg.RightReserve,
		"right/pcard":   msg.RightPCard,

		// Left fencer
		"left/id":      msg.LeftId,
		"left/name":    msg.LeftName,
		"left/nat":     msg.LeftNat,
		"left/score":   msg.LeftScore,
		"left/status":  msg.LeftStatus,
		"left/ycard":   msg.LeftYCard,
		"left/rcard":   msg.LeftRCard,
		"left/light":   msg.LeftLight,
		"left/wlight":  msg.LeftWLight,
		"left/medical": msg.LeftMedical,
		"left/reserve": msg.LeftReserve,
		"left/pcard":   msg.LeftPCard,
	}

	// Publish each field to its own topic
	for field, value := range fields {
		// Skip empty values to reduce unnecessary messages
		if value == "" {
			continue
		}

		topic := fmt.Sprintf("%s/%s", baseTopic, field)
		token := p.client.Publish(topic, p.qos, p.retain, value)
		if !token.WaitTimeout(5 * time.Second) {
			return fmt.Errorf("mqtt publish timeout on topic %s", topic)
		}
		if err := token.Error(); err != nil {
			return fmt.Errorf("mqtt publish error on topic %s: %w", topic, err)
		}
	}

	return nil
}

func sanitizeTopic(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "unknown"
	}
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	return s
}
