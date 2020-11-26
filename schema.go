package main

const rootSchema = `
type Query {
		arrivals (stop: String!): [Arrival]!
}
type Arrival {
		headsign:  String!
		stoptime:  String!
		vehicle:   String!
		estimated: String!
		latitude:  String!
		longitude: String!
		shape:     String!
		id:        String!
		trip:      String!
		canceled:  String!
		date:      String!
		route:     String!
		direction: String!
}
`
