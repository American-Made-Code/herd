# Herd

## Harmonized Ensemble of Responsive Dialogues

[![Go Report Card][report-card-image]][report-card-url]
[![Documentation][docs-image]][docs-url]
[![Star][star-image]][star-url]
[![Package License][package-license-image]][package-license-url]
![GitHub last commit](https://img.shields.io/github/last-commit/American-Made-Code/herd)

______________________________________________________________________

## Overview

Herd is a pioneering Go package designed to seamlessly integrate and facilitate communication between multiple Large Language Models (LLMs). This package offers developers the unique capability to orchestrate a symphony of AI-driven conversations, enabling diverse LLMs to interact, share insights, and provide cohesive responses.

## Installation

### Using CLI

~~~bash
go get github.com/American-Made-Code/herd 
~~~

### Using `go.mod`

~~~go.mod
# Go Modules
require github.com/American-Made-Code/herd v0.1.0 
~~~

## Usage

Below is an example of a gin controller implementing the roleplay service to loop through a conversation:

~~~go
package controllers

import (
 "fmt"
 "os"

 "github.com/American-Made-Code/herd/pkg/domain/roleplay"
 "github.com/American-Made-Code/herd/pkg/external/gpt"
 "github.com/gin-gonic/gin"
)

type TestController interface {
 Post(c *gin.Context)
}

type testController struct {
}

func NewTestController() TestController {
 return &testController{}
}

func (t testController) Post(c *gin.Context) {
 apiKey := os.Getenv("OPENAI_API_KEY")
 providers := []gpt.Auth{
  {Provider: gpt.OpenAi, ApiKey: apiKey},
 }

 roleplayService := roleplay.NewService(providers)

 first := roleplay.Participant{
  Index:        0,
  SystemPrompt: "You are a wizard.",
  Name:         "Gandalf",
 }

 second := roleplay.Participant{
  Index:        1,
  SystemPrompt: "You are a barbarian.",
  Name:         "Conan",
 }

 third := roleplay.Participant{
  Index:        2,
  SystemPrompt: "You are a rogue.",
  Name:         "Bilbo",
 }

 allParticipants := []roleplay.Participant{first, second, third}

 var roleplayMessages []roleplay.Message
 initialMessage := roleplay.Message{
  Content:     "Hello my friends, we have been hired by the mayor of Neverwinter to investigate the disappearance of several children. Are you ready to embark on this adventure?",
  Participant: first,
 }
 roleplayMessages = append(roleplayMessages, initialMessage)

 responseOrder := []int{1, 2, 0, 1, 2, 0}

 command := roleplay.CreateResponseMessageCommand{
  AllParticipants:          allParticipants,
  ResponseParticipantIndex: 1,
  ResponseGpt: roleplay.Gpt{
   Provider: gpt.OpenAi,
   Model:    "gpt-3.5-turbo",
  },
  Messages: roleplayMessages,
 }

 for _, index := range responseOrder {
  fmt.Printf("\n=============================\n")
  fmt.Printf("\nINDEX: %v \n", index)
  fmt.Printf("\nCURRENT MESSAGES (%v) ----------\n", len(roleplayMessages))
  for _, message := range roleplayMessages {
   fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
  }
  fmt.Printf("\n----------\n")
  command.ResponseParticipantIndex = index
  command.Messages = roleplayMessages

  res, err := roleplayService.CreateResponseMessage(command)

  if err != nil {
   c.JSON(500, gin.H{
    "error": err,
   })
   return
  }

  fmt.Printf("\nNEW MESSAGES (%v) ----------\n", len(res.GeneratedMessages))
  for _, message := range res.GeneratedMessages {
   fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
  }
  fmt.Printf("\n----------\n")

  roleplayMessages = append(roleplayMessages, res.GeneratedMessages[0])
 }

 fmt.Printf("\n FINAL: (%v) ----------\n", len(roleplayMessages))
 for _, message := range roleplayMessages {
  fmt.Printf("\n<%v> (\n%v\n)\n", message.Participant.Name, message.Content)
 }
 fmt.Printf("\n----------\n")

 // Return the user
 c.JSON(200, gin.H{
  "data": roleplayMessages,
 })
}

~~~

## Key Features

- Multi-LLM Integration: Effortlessly connect various LLMs, allowing them to exchange information and collaborate on tasks.

- Synchronized Conversations: Ensure smooth and synchronous communication among LLMs, preserving context and coherence.

- Flexible Architecture: Adaptable to various LLM platforms, Herd is built with modularity and scalability in mind.

- Real-time Interaction: Facilitate dynamic, real-time dialogues between LLMs, opening up possibilities for complex problem-solving and idea generation.

- Go Efficiency: Leverage the simplicity and efficiency of Go, making the package lightweight and performant.

[docs-image]: https://img.shields.io/badge/Documentation-grey.svg?logo=github
[docs-url]: https://github.com/American-Made-Code/herd
[star-image]: https://img.shields.io/github/stars/camel-ai/camel?label=stars&logo=github&color=brightgreen
[star-url]: https://github.com/American-Made-Code/herd/stargazers
[package-license-image]: https://img.shields.io/github/license/American-Made-Code/herd
[package-license-url]: https://github.com/American-Made-Code/herd/blob/main/LICENSE
[report-card-image]: https://goreportcard.com/badge/American-Made-Code/herd
[report-card-url]: https://goreportcard.com/report/American-Made-Code/herd
