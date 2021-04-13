package main

import "testing"

func TestGetAddress(t *testing.T) {
	host := "127.0.0.1"
	port := "3000"
	actual := getAddress(host, port)
	expected := "127.0.0.1:3000"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestCreateRegisterMessage(t *testing.T) {
	group := "worker-group"
	instance := "instance"
	const cap uint64 = 100
	msgRegister := createRegisterMessage(
		group,
		instance,
		cap,
	)
	msg := msgRegister.GetRegister()
	expectedGroup := msg.Group
	expectedInstance := msg.Instance
	expectedFreq := msg.Frequency

	if (expectedGroup != group) {
		t.Errorf("Expected %s, got %s", expectedGroup, group)
	}
	if (expectedInstance != instance) {
		t.Errorf("Expected %s, got %s", expectedInstance, instance)
	}
	if (expectedFreq != cap) {
		t.Errorf("Expected %d, got %d", expectedFreq, cap)
	}
}