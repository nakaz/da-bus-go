package main

const rootSchema = `
type Query {
		arrivals (stop: String!): [Arrival]!
		vehicle (num: String!): Vehicle!
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
type Vehicle {
		driver:         String!
		longitude:      String!
		latitude:       String!
		lastMessage:    String!
		headsign:       String!
		routeShortName: String!
		number:         String!
		trip:           String!
		adherence:      String!
}
`
