package main

import (
  "fmt"
  "github.com/aws/aws-lambda-go/lambda"
)

type MyInput struct {
  Input string `json:"input"`
}

type MyOutput struct {
  Output string `json:"output"`
}

func HandleRequest(event *MyInput) (*MyOutput, error) {
  if event == nil {
    return nil, fmt.Errorf("received nil event")
  }
  return &MyOutput{ Output: event.Input }, nil
}

func main() {
  lambda.Start(HandleRequest)
}
