package main

import "testing"

func TestMessage(t *testing.T) {
  result := Message()
  if result != "pong" {
    t.Errorf("Server should respond with 'pong', got '%s'", result)
  }
}
