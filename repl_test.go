package main

import "testing"

func TestCleanInput(t *testing.T) {
  cases := []struct {
    input     string
    expected  []string
  }{
    {
      input:    "      ",
      expected: []string{},
    },
    {
      input:    "hello",
      expected: []string{"hello"},
    },
    {
      input:    "hello world",
      expected: []string{"hello", "world"},
    },
    {
      input:    " HELLo wOrld  ",
      expected: []string{"hello", "world"},
    },
  }
  for _, c := range(cases) {
    result := cleanInput(c.input)
    if len(result) != len(c.expected) {
      t.Fatalf("Len of the result %v is different form the expected len: %v", len(result), len(c.expected))
    }
    for i, word := range(result) {
      if word != c.expected[i] {
        t.Fatalf("Was expecting %v, got %v", c.expected[i], word)
      }
    }
  }
}

