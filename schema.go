package main

const rootSchema = `
type Query {
		arrivals (stop: String!): [Arrival]
}
type Arrival {
		headsign: String
}
`
