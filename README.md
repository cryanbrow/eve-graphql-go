# EVE GraphQL Go

[![build](https://github.com/cryanbrow/eve-graphql-go/actions/workflows/github-actions.yml/badge.svg)](https://github.com/cryanbrow/eve-graphql-go/actions/workflows/github-actions.yml) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=cryanbrow_eve-graphql-go&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=cryanbrow_eve-graphql-go) [![Go Report Card](https://goreportcard.com/badge/github.com/cryanbrow/eve-graphql-go)](https://goreportcard.com/report/github.com/cryanbrow/eve-graphql-go) [![Go Reference](https://pkg.go.dev/badge/github.com/cryanbrow/eve-graphql-go.svg)](https://pkg.go.dev/github.com/cryanbrow/eve-graphql-go) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## How to Run 

1. Download the Code to your machine
`go get github.com/cryanbrow/eve-graphql-go`
2. Execute the binary
`eve-graphql-go`

## Available Endpoints

Once the application is running there will be 3 available endpoints.

1. [http://localhost:8080/Graphql](http://localhost:8080/Graphql) - The Graphiql UI for querying
2. [http://localhost:8080/query](http://localhost:8080/query) - The endpoint to POST your GraphQL queries against.
2. [http://localhost:8080/voyager](http://localhost:8080/voyager) - The Voyager UI for a graphical representation of the schema
