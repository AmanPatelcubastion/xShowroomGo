package mygraphql

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}

	# The query type, represents all of the entry points into our object graph
	type Query {
		user(id: ID!) : [User]!
		device(id: ID!) : [Device]!
	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createDevice(device: DeviceInput!): Device
		createUser(user: UserInput!): User
	}

	# The individual using xShowroom application
	type User {
		id: ID!
		name: String!
		device: Device!
	}
	input UserInput {
		id: ID
		name: String!
		device: DeviceInput
	}

	# The Android or I-pad device used by a user
	type Device {
		id: ID!
		device_uuid: String!
		user: User!
	}
	input DeviceInput {
		id: ID
		device_uuid: String!
	}
`
