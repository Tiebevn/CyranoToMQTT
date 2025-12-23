package main

import (
	"encoding/json"
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
	topic := fmt.Sprintf("%s/%s/%s", p.topicBase, sanitizeTopic(msg.Piste), msg.Command.String())
	payload, err := encodeMessageJSON(msg)
	if err != nil {
		return err
	}
	token := p.client.Publish(topic, p.qos, p.retain, payload)
	if !token.WaitTimeout(5 * time.Second) {
		return fmt.Errorf("mqtt publish timeout")
	}
	return token.Error()
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

func encodeMessageJSON(msg *CyranoMessage) ([]byte, error) {
	type jsonMsg struct {
		Protocol    string `json:"protocol"`
		Command     string `json:"command"`
		Piste       string `json:"piste"`
		Competition string `json:"competition"`
		Phase       string `json:"phase"`
		PoulTab     string `json:"poulTab"`
		Match       string `json:"match"`
		Round       string `json:"round"`
		Time        string `json:"time"`
		Stopwatch   string `json:"stopwatch"`
		Type        string `json:"type"`
		Weapon      string `json:"weapon"`
		Priority    string `json:"priority"`
		State       string `json:"state"`

		RefId   string `json:"refId"`
		RefName string `json:"refName"`
		RefNat  string `json:"refNat"`

		RightId      string `json:"rightId"`
		RightName    string `json:"rightName"`
		RightNat     string `json:"rightNat"`
		RightScore   string `json:"rightScore"`
		RightStatus  string `json:"rightStatus"`
		RightYCard   string `json:"rightYCard"`
		RightRCard   string `json:"rightRCard"`
		RightLight   string `json:"rightLight"`
		RightWLight  string `json:"rightWLight"`
		RightMedical string `json:"rightMedical"`
		RightReserve string `json:"rightReserve"`
		RightPCard   string `json:"rightPCard"`

		LeftId      string `json:"leftId"`
		LeftName    string `json:"leftName"`
		LeftNat     string `json:"leftNat"`
		LeftScore   string `json:"leftScore"`
		LeftStatus  string `json:"leftStatus"`
		LeftYCard   string `json:"leftYCard"`
		LeftRCard   string `json:"leftRCard"`
		LeftLight   string `json:"leftLight"`
		LeftWLight  string `json:"leftWLight"`
		LeftMedical string `json:"leftMedical"`
		LeftReserve string `json:"leftReserve"`
		LeftPCard   string `json:"leftPCard"`
	}

	out := jsonMsg{
		Protocol:     msg.ProtocolVersion.String(),
		Command:      msg.Command.String(),
		Piste:        msg.Piste,
		Competition:  msg.Competition,
		Phase:        msg.Phase,
		PoulTab:      msg.PoulTab,
		Match:        msg.Match,
		Round:        msg.Round,
		Time:         msg.Time,
		Stopwatch:    msg.Stopwatch,
		Type:         msg.Type.String(),
		Weapon:       msg.Weapon,
		Priority:     msg.Priority,
		State:        msg.State.String(),
		RefId:        msg.RefId,
		RefName:      msg.RefName,
		RefNat:       msg.RefNat,
		RightId:      msg.RightId,
		RightName:    msg.RightName,
		RightNat:     msg.RightNat,
		RightScore:   msg.RightScore,
		RightStatus:  msg.RightStatus,
		RightYCard:   msg.RightYCard,
		RightRCard:   msg.RightRCard,
		RightLight:   msg.RightLight,
		RightWLight:  msg.RightWLight,
		RightMedical: msg.RightMedical,
		RightReserve: msg.RightReserve,
		RightPCard:   msg.RightPCard,
		LeftId:       msg.LeftId,
		LeftName:     msg.LeftName,
		LeftNat:      msg.LeftNat,
		LeftScore:    msg.LeftScore,
		LeftStatus:   msg.LeftStatus,
		LeftYCard:    msg.LeftYCard,
		LeftRCard:    msg.LeftRCard,
		LeftLight:    msg.LeftLight,
		LeftWLight:   msg.LeftWLight,
		LeftMedical:  msg.LeftMedical,
		LeftReserve:  msg.LeftReserve,
		LeftPCard:    msg.LeftPCard,
	}

	return json.Marshal(out)
}
