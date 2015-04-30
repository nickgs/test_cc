package main

import(
  "testing"
)

func TestVisaCard(t *testing.T) {
  card := VisaCard();

  if(card != "4111111111111111") {
    t.Fatalf("VisaCard failed: %s", card)
  }
}
