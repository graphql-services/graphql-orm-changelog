package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

type Query {
  changelogChange(id: ID, q: String, filter: ChangelogChangeFilterType): ChangelogChange
  changelogChanges(offset: Int, limit: Int = 30, q: String, sort: [ChangelogChangeSortType!], filter: ChangelogChangeFilterType): ChangelogChangeResultType
  changelog(id: ID, q: String, filter: ChangelogFilterType): Changelog
  changelogs(offset: Int, limit: Int = 30, q: String, sort: [ChangelogSortType!], filter: ChangelogFilterType): ChangelogResultType
}

type Mutation {
  createChangelogChange(input: ChangelogChangeCreateInput!): ChangelogChange!
  updateChangelogChange(id: ID!, input: ChangelogChangeUpdateInput!): ChangelogChange!
  deleteChangelogChange(id: ID!): ChangelogChange!
  deleteAllChangelogChanges: Boolean!
  createChangelog(input: ChangelogCreateInput!): Changelog!
  updateChangelog(id: ID!, input: ChangelogUpdateInput!): Changelog!
  deleteChangelog(id: ID!): Changelog!
  deleteAllChangelogs: Boolean!
}

enum ChangelogType {
  CREATED
  UPDATED
  DELETED
}

type ChangelogChange {
  id: ID!
  column: String!
  oldValue: String
  newValue: String
  log: Changelog
  logId: ID
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
}

type Changelog @key(fields: "id") {
  id: ID!
  entity: String!
  entityID: String!
  principalID: String
  type: ChangelogType!
  date: Time!
  changes: [ChangelogChange!]!
  updatedAt: Time
  createdAt: Time!
  updatedBy: ID
  createdBy: ID
  changesIds: [ID!]!
}

union _Entity = Changelog

input ChangelogChangeCreateInput {
  id: ID
  column: String!
  oldValue: String
  newValue: String
  logId: ID
}

input ChangelogChangeUpdateInput {
  column: String
  oldValue: String
  newValue: String
  logId: ID
}

enum ChangelogChangeSortType {
  ID_ASC
  ID_DESC
  COLUMN_ASC
  COLUMN_DESC
  OLD_VALUE_ASC
  OLD_VALUE_DESC
  NEW_VALUE_ASC
  NEW_VALUE_DESC
  LOG_ID_ASC
  LOG_ID_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
}

input ChangelogChangeFilterType {
  AND: [ChangelogChangeFilterType!]
  OR: [ChangelogChangeFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  column: String
  column_ne: String
  column_gt: String
  column_lt: String
  column_gte: String
  column_lte: String
  column_in: [String!]
  column_like: String
  column_prefix: String
  column_suffix: String
  oldValue: String
  oldValue_ne: String
  oldValue_gt: String
  oldValue_lt: String
  oldValue_gte: String
  oldValue_lte: String
  oldValue_in: [String!]
  oldValue_like: String
  oldValue_prefix: String
  oldValue_suffix: String
  newValue: String
  newValue_ne: String
  newValue_gt: String
  newValue_lt: String
  newValue_gte: String
  newValue_lte: String
  newValue_in: [String!]
  newValue_like: String
  newValue_prefix: String
  newValue_suffix: String
  logId: ID
  logId_ne: ID
  logId_gt: ID
  logId_lt: ID
  logId_gte: ID
  logId_lte: ID
  logId_in: [ID!]
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  log: ChangelogFilterType
}

type ChangelogChangeResultType {
  items: [ChangelogChange!]!
  count: Int!
}

input ChangelogCreateInput {
  id: ID
  entity: String!
  entityID: String!
  principalID: String
  type: ChangelogType!
  date: Time!
  changesIds: [ID!]
}

input ChangelogUpdateInput {
  entity: String
  entityID: String
  principalID: String
  type: ChangelogType
  date: Time
  changesIds: [ID!]
}

enum ChangelogSortType {
  ID_ASC
  ID_DESC
  ENTITY_ASC
  ENTITY_DESC
  ENTITY_ID_ASC
  ENTITY_ID_DESC
  PRINCIPAL_ID_ASC
  PRINCIPAL_ID_DESC
  TYPE_ASC
  TYPE_DESC
  DATE_ASC
  DATE_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  CHANGES_IDS_ASC
  CHANGES_IDS_DESC
}

input ChangelogFilterType {
  AND: [ChangelogFilterType!]
  OR: [ChangelogFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  entity: String
  entity_ne: String
  entity_gt: String
  entity_lt: String
  entity_gte: String
  entity_lte: String
  entity_in: [String!]
  entity_like: String
  entity_prefix: String
  entity_suffix: String
  entityID: String
  entityID_ne: String
  entityID_gt: String
  entityID_lt: String
  entityID_gte: String
  entityID_lte: String
  entityID_in: [String!]
  entityID_like: String
  entityID_prefix: String
  entityID_suffix: String
  principalID: String
  principalID_ne: String
  principalID_gt: String
  principalID_lt: String
  principalID_gte: String
  principalID_lte: String
  principalID_in: [String!]
  principalID_like: String
  principalID_prefix: String
  principalID_suffix: String
  type: ChangelogType
  type_ne: ChangelogType
  type_gt: ChangelogType
  type_lt: ChangelogType
  type_gte: ChangelogType
  type_lte: ChangelogType
  type_in: [ChangelogType!]
  date: Time
  date_ne: Time
  date_gt: Time
  date_lt: Time
  date_gte: Time
  date_lte: Time
  date_in: [Time!]
  updatedAt: Time
  updatedAt_ne: Time
  updatedAt_gt: Time
  updatedAt_lt: Time
  updatedAt_gte: Time
  updatedAt_lte: Time
  updatedAt_in: [Time!]
  createdAt: Time
  createdAt_ne: Time
  createdAt_gt: Time
  createdAt_lt: Time
  createdAt_gte: Time
  createdAt_lte: Time
  createdAt_in: [Time!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  changes: ChangelogChangeFilterType
}

type ChangelogResultType {
  items: [Changelog!]!
  count: Int!
}`
)