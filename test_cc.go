package main

import(
  "fmt"
  "os"
)

func main() {
  card := VisaCard()
  fmt.Fprintln(os.Stdout, card);
}

func VisaCard() string {
  return "4111111111111111";
}
