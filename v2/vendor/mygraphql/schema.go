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
        account(id: ID!) : [Account]!
		lead(id: ID!) : [Lead]!
        product(id: ID!) : [Product]!
        relatedproductgroup(id: ID!) : [Relatedproductgroup]!

	}

	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createDevice(device: DeviceInput!): Device
		createUser(user: UserInput!): User
        createLead(lead: LeadInput!): Lead
        createAccount(account: AccountInput!): Account
        createProduct(product: ProductInput!): Product
        createRelatedproductgroup(relatedproductgroup: RelatedproductgroupInput!): Relatedproductgroup
        deleteUser(id:ID!) : User

	}

	# The individual using xShowroom application
	type User {
		id: ID!
		name: String!
		device: Device!
        leads: [Lead!]!

	}
	input UserInput {
		id: ID
		name: String!
		device: DeviceInput
		leads: [LeadInput!]
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

    type Account {
      id: ID!
      name: String!
      leads: [Lead!]!
   }

    input AccountInput {
      id: ID
      name: String!
      leads: [LeadInput!]
   }

    type Lead {
      id: ID!
      name: String!
      leadType: String!
      typeId: ID!
      accounts: Account!
      user: User!

   }

    input LeadInput {
      id: ID
      name: String!
      leadType: String
      typeId: ID
   }

    type Product {
      id: ID!
      name: String!
      relatedproductgroups: [Relatedproductgroup!]!
   }

    input ProductInput {
      id: ID
      name: String!
      relatedproductgroupids: [ID!]
    #  relatedproductgroups: [RelatedproductgroupInput!]
   }

    type Relatedproductgroup {
      id: ID!
      grouptype: String!
      products: [Product!]!
   }

    input RelatedproductgroupInput {
      id: ID
      grouptype: String!
      productids: [ID!]
     # products: [ProductInput!]
   }
`
