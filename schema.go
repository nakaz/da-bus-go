package main

const rootSchema = `
type Query {
		arrivals (stop: Int!): [Arrival]!
		vehicle (num: Int!): Vehicle!
		route (route: String!): [Route]!
		headsign (headsign: String!): [Route]!
}
type Arrival {
		headsign:  String!
		estimated: String!
		latitude:  String!
		longitude: String!
		shape:     String!
		id:        String!
		trip:      String!
		canceled:  String!
		route:     String!
		direction: String!
		vehicleId: Int!
		latLng: [Float!]!
		arrivalTime: String!
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
type Route {
		headsign:  String!
		routeNum:  String!
		shapeID:   String!
		firstStop: String!
}
`
