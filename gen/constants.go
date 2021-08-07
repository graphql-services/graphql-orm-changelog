package gen

type key int

const (
	KeyPrincipalID         key    = iota
	KeyLoaders             key    = iota
	KeyExecutableSchema    key    = iota
	KeyJWTClaims           key    = iota
	KeyHTTPRequest         key    = iota
	KeyMutationTransaction key    = iota
	KeyMutationEvents      key    = iota
	SchemaSDL              string = `scalar Time

type Query {
  changelogChange(id: ID, q: String, filter: ChangelogChangeFilterType): ChangelogChange
  changelogChanges(offset: Int, limit: Int = 30, q: String, sort: [ChangelogChangeSortType!], filter: ChangelogChangeFilterType): ChangelogChangeResultType!
  changelog(id: ID, q: String, filter: ChangelogFilterType): Changelog
  changelogs(offset: Int, limit: Int = 30, q: String, sort: [ChangelogSortType!], filter: ChangelogFilterType): ChangelogResultType!
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

enum ObjectSortType {
  ASC
  DESC
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
  changesConnection(offset: Int, limit: Int = 30, q: String, sort: [ChangelogChangeSortType!], filter: ChangelogChangeFilterType): ChangelogChangeResultType!
}

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

input ChangelogChangeSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  column: ObjectSortType
  columnMin: ObjectSortType
  columnMax: ObjectSortType
  oldValue: ObjectSortType
  oldValueMin: ObjectSortType
  oldValueMax: ObjectSortType
  newValue: ObjectSortType
  newValueMin: ObjectSortType
  newValueMax: ObjectSortType
  logId: ObjectSortType
  logIdMin: ObjectSortType
  logIdMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  log: ChangelogSortType
}

input ChangelogChangeFilterType {
  AND: [ChangelogChangeFilterType!]
  OR: [ChangelogChangeFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_not_in: [ID!]
  idMin_not_in: [ID!]
  idMax_not_in: [ID!]
  id_null: Boolean
  column: String
  columnMin: String
  columnMax: String
  column_ne: String
  columnMin_ne: String
  columnMax_ne: String
  column_gt: String
  columnMin_gt: String
  columnMax_gt: String
  column_lt: String
  columnMin_lt: String
  columnMax_lt: String
  column_gte: String
  columnMin_gte: String
  columnMax_gte: String
  column_lte: String
  columnMin_lte: String
  columnMax_lte: String
  column_in: [String!]
  columnMin_in: [String!]
  columnMax_in: [String!]
  column_not_in: [String!]
  columnMin_not_in: [String!]
  columnMax_not_in: [String!]
  column_like: String
  columnMin_like: String
  columnMax_like: String
  column_prefix: String
  columnMin_prefix: String
  columnMax_prefix: String
  column_suffix: String
  columnMin_suffix: String
  columnMax_suffix: String
  column_null: Boolean
  oldValue: String
  oldValueMin: String
  oldValueMax: String
  oldValue_ne: String
  oldValueMin_ne: String
  oldValueMax_ne: String
  oldValue_gt: String
  oldValueMin_gt: String
  oldValueMax_gt: String
  oldValue_lt: String
  oldValueMin_lt: String
  oldValueMax_lt: String
  oldValue_gte: String
  oldValueMin_gte: String
  oldValueMax_gte: String
  oldValue_lte: String
  oldValueMin_lte: String
  oldValueMax_lte: String
  oldValue_in: [String!]
  oldValueMin_in: [String!]
  oldValueMax_in: [String!]
  oldValue_not_in: [String!]
  oldValueMin_not_in: [String!]
  oldValueMax_not_in: [String!]
  oldValue_like: String
  oldValueMin_like: String
  oldValueMax_like: String
  oldValue_prefix: String
  oldValueMin_prefix: String
  oldValueMax_prefix: String
  oldValue_suffix: String
  oldValueMin_suffix: String
  oldValueMax_suffix: String
  oldValue_null: Boolean
  newValue: String
  newValueMin: String
  newValueMax: String
  newValue_ne: String
  newValueMin_ne: String
  newValueMax_ne: String
  newValue_gt: String
  newValueMin_gt: String
  newValueMax_gt: String
  newValue_lt: String
  newValueMin_lt: String
  newValueMax_lt: String
  newValue_gte: String
  newValueMin_gte: String
  newValueMax_gte: String
  newValue_lte: String
  newValueMin_lte: String
  newValueMax_lte: String
  newValue_in: [String!]
  newValueMin_in: [String!]
  newValueMax_in: [String!]
  newValue_not_in: [String!]
  newValueMin_not_in: [String!]
  newValueMax_not_in: [String!]
  newValue_like: String
  newValueMin_like: String
  newValueMax_like: String
  newValue_prefix: String
  newValueMin_prefix: String
  newValueMax_prefix: String
  newValue_suffix: String
  newValueMin_suffix: String
  newValueMax_suffix: String
  newValue_null: Boolean
  logId: ID
  logIdMin: ID
  logIdMax: ID
  logId_ne: ID
  logIdMin_ne: ID
  logIdMax_ne: ID
  logId_gt: ID
  logIdMin_gt: ID
  logIdMax_gt: ID
  logId_lt: ID
  logIdMin_lt: ID
  logIdMax_lt: ID
  logId_gte: ID
  logIdMin_gte: ID
  logIdMax_gte: ID
  logId_lte: ID
  logIdMin_lte: ID
  logIdMax_lte: ID
  logId_in: [ID!]
  logIdMin_in: [ID!]
  logIdMax_in: [ID!]
  logId_not_in: [ID!]
  logIdMin_not_in: [ID!]
  logIdMax_not_in: [ID!]
  logId_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_not_in: [Time!]
  updatedAtMin_not_in: [Time!]
  updatedAtMax_not_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_not_in: [Time!]
  createdAtMin_not_in: [Time!]
  createdAtMax_not_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_not_in: [ID!]
  updatedByMin_not_in: [ID!]
  updatedByMax_not_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_not_in: [ID!]
  createdByMin_not_in: [ID!]
  createdByMax_not_in: [ID!]
  createdBy_null: Boolean
  log: ChangelogFilterType
}

type ChangelogChangeResultType {
  items: [ChangelogChange!]!
  count: Int!
  aggregations: ChangelogChangeResultAggregations!
}

type ChangelogChangeResultAggregations {
  columnMin: String!
  columnMax: String!
  oldValueMin: String
  oldValueMax: String
  newValueMin: String
  newValueMax: String
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

input ChangelogSortType {
  id: ObjectSortType
  idMin: ObjectSortType
  idMax: ObjectSortType
  entity: ObjectSortType
  entityMin: ObjectSortType
  entityMax: ObjectSortType
  entityID: ObjectSortType
  entityIDMin: ObjectSortType
  entityIDMax: ObjectSortType
  principalID: ObjectSortType
  principalIDMin: ObjectSortType
  principalIDMax: ObjectSortType
  type: ObjectSortType
  typeMin: ObjectSortType
  typeMax: ObjectSortType
  date: ObjectSortType
  dateMin: ObjectSortType
  dateMax: ObjectSortType
  updatedAt: ObjectSortType
  updatedAtMin: ObjectSortType
  updatedAtMax: ObjectSortType
  createdAt: ObjectSortType
  createdAtMin: ObjectSortType
  createdAtMax: ObjectSortType
  updatedBy: ObjectSortType
  updatedByMin: ObjectSortType
  updatedByMax: ObjectSortType
  createdBy: ObjectSortType
  createdByMin: ObjectSortType
  createdByMax: ObjectSortType
  changesIds: ObjectSortType
  changesIdsMin: ObjectSortType
  changesIdsMax: ObjectSortType
  changes: ChangelogChangeSortType
}

input ChangelogFilterType {
  AND: [ChangelogFilterType!]
  OR: [ChangelogFilterType!]
  id: ID
  idMin: ID
  idMax: ID
  id_ne: ID
  idMin_ne: ID
  idMax_ne: ID
  id_gt: ID
  idMin_gt: ID
  idMax_gt: ID
  id_lt: ID
  idMin_lt: ID
  idMax_lt: ID
  id_gte: ID
  idMin_gte: ID
  idMax_gte: ID
  id_lte: ID
  idMin_lte: ID
  idMax_lte: ID
  id_in: [ID!]
  idMin_in: [ID!]
  idMax_in: [ID!]
  id_not_in: [ID!]
  idMin_not_in: [ID!]
  idMax_not_in: [ID!]
  id_null: Boolean
  entity: String
  entityMin: String
  entityMax: String
  entity_ne: String
  entityMin_ne: String
  entityMax_ne: String
  entity_gt: String
  entityMin_gt: String
  entityMax_gt: String
  entity_lt: String
  entityMin_lt: String
  entityMax_lt: String
  entity_gte: String
  entityMin_gte: String
  entityMax_gte: String
  entity_lte: String
  entityMin_lte: String
  entityMax_lte: String
  entity_in: [String!]
  entityMin_in: [String!]
  entityMax_in: [String!]
  entity_not_in: [String!]
  entityMin_not_in: [String!]
  entityMax_not_in: [String!]
  entity_like: String
  entityMin_like: String
  entityMax_like: String
  entity_prefix: String
  entityMin_prefix: String
  entityMax_prefix: String
  entity_suffix: String
  entityMin_suffix: String
  entityMax_suffix: String
  entity_null: Boolean
  entityID: String
  entityIDMin: String
  entityIDMax: String
  entityID_ne: String
  entityIDMin_ne: String
  entityIDMax_ne: String
  entityID_gt: String
  entityIDMin_gt: String
  entityIDMax_gt: String
  entityID_lt: String
  entityIDMin_lt: String
  entityIDMax_lt: String
  entityID_gte: String
  entityIDMin_gte: String
  entityIDMax_gte: String
  entityID_lte: String
  entityIDMin_lte: String
  entityIDMax_lte: String
  entityID_in: [String!]
  entityIDMin_in: [String!]
  entityIDMax_in: [String!]
  entityID_not_in: [String!]
  entityIDMin_not_in: [String!]
  entityIDMax_not_in: [String!]
  entityID_like: String
  entityIDMin_like: String
  entityIDMax_like: String
  entityID_prefix: String
  entityIDMin_prefix: String
  entityIDMax_prefix: String
  entityID_suffix: String
  entityIDMin_suffix: String
  entityIDMax_suffix: String
  entityID_null: Boolean
  principalID: String
  principalIDMin: String
  principalIDMax: String
  principalID_ne: String
  principalIDMin_ne: String
  principalIDMax_ne: String
  principalID_gt: String
  principalIDMin_gt: String
  principalIDMax_gt: String
  principalID_lt: String
  principalIDMin_lt: String
  principalIDMax_lt: String
  principalID_gte: String
  principalIDMin_gte: String
  principalIDMax_gte: String
  principalID_lte: String
  principalIDMin_lte: String
  principalIDMax_lte: String
  principalID_in: [String!]
  principalIDMin_in: [String!]
  principalIDMax_in: [String!]
  principalID_not_in: [String!]
  principalIDMin_not_in: [String!]
  principalIDMax_not_in: [String!]
  principalID_like: String
  principalIDMin_like: String
  principalIDMax_like: String
  principalID_prefix: String
  principalIDMin_prefix: String
  principalIDMax_prefix: String
  principalID_suffix: String
  principalIDMin_suffix: String
  principalIDMax_suffix: String
  principalID_null: Boolean
  type: ChangelogType
  typeMin: ChangelogType
  typeMax: ChangelogType
  type_ne: ChangelogType
  typeMin_ne: ChangelogType
  typeMax_ne: ChangelogType
  type_gt: ChangelogType
  typeMin_gt: ChangelogType
  typeMax_gt: ChangelogType
  type_lt: ChangelogType
  typeMin_lt: ChangelogType
  typeMax_lt: ChangelogType
  type_gte: ChangelogType
  typeMin_gte: ChangelogType
  typeMax_gte: ChangelogType
  type_lte: ChangelogType
  typeMin_lte: ChangelogType
  typeMax_lte: ChangelogType
  type_in: [ChangelogType!]
  typeMin_in: [ChangelogType!]
  typeMax_in: [ChangelogType!]
  type_not_in: [ChangelogType!]
  typeMin_not_in: [ChangelogType!]
  typeMax_not_in: [ChangelogType!]
  type_null: Boolean
  date: Time
  dateMin: Time
  dateMax: Time
  date_ne: Time
  dateMin_ne: Time
  dateMax_ne: Time
  date_gt: Time
  dateMin_gt: Time
  dateMax_gt: Time
  date_lt: Time
  dateMin_lt: Time
  dateMax_lt: Time
  date_gte: Time
  dateMin_gte: Time
  dateMax_gte: Time
  date_lte: Time
  dateMin_lte: Time
  dateMax_lte: Time
  date_in: [Time!]
  dateMin_in: [Time!]
  dateMax_in: [Time!]
  date_not_in: [Time!]
  dateMin_not_in: [Time!]
  dateMax_not_in: [Time!]
  date_null: Boolean
  updatedAt: Time
  updatedAtMin: Time
  updatedAtMax: Time
  updatedAt_ne: Time
  updatedAtMin_ne: Time
  updatedAtMax_ne: Time
  updatedAt_gt: Time
  updatedAtMin_gt: Time
  updatedAtMax_gt: Time
  updatedAt_lt: Time
  updatedAtMin_lt: Time
  updatedAtMax_lt: Time
  updatedAt_gte: Time
  updatedAtMin_gte: Time
  updatedAtMax_gte: Time
  updatedAt_lte: Time
  updatedAtMin_lte: Time
  updatedAtMax_lte: Time
  updatedAt_in: [Time!]
  updatedAtMin_in: [Time!]
  updatedAtMax_in: [Time!]
  updatedAt_not_in: [Time!]
  updatedAtMin_not_in: [Time!]
  updatedAtMax_not_in: [Time!]
  updatedAt_null: Boolean
  createdAt: Time
  createdAtMin: Time
  createdAtMax: Time
  createdAt_ne: Time
  createdAtMin_ne: Time
  createdAtMax_ne: Time
  createdAt_gt: Time
  createdAtMin_gt: Time
  createdAtMax_gt: Time
  createdAt_lt: Time
  createdAtMin_lt: Time
  createdAtMax_lt: Time
  createdAt_gte: Time
  createdAtMin_gte: Time
  createdAtMax_gte: Time
  createdAt_lte: Time
  createdAtMin_lte: Time
  createdAtMax_lte: Time
  createdAt_in: [Time!]
  createdAtMin_in: [Time!]
  createdAtMax_in: [Time!]
  createdAt_not_in: [Time!]
  createdAtMin_not_in: [Time!]
  createdAtMax_not_in: [Time!]
  createdAt_null: Boolean
  updatedBy: ID
  updatedByMin: ID
  updatedByMax: ID
  updatedBy_ne: ID
  updatedByMin_ne: ID
  updatedByMax_ne: ID
  updatedBy_gt: ID
  updatedByMin_gt: ID
  updatedByMax_gt: ID
  updatedBy_lt: ID
  updatedByMin_lt: ID
  updatedByMax_lt: ID
  updatedBy_gte: ID
  updatedByMin_gte: ID
  updatedByMax_gte: ID
  updatedBy_lte: ID
  updatedByMin_lte: ID
  updatedByMax_lte: ID
  updatedBy_in: [ID!]
  updatedByMin_in: [ID!]
  updatedByMax_in: [ID!]
  updatedBy_not_in: [ID!]
  updatedByMin_not_in: [ID!]
  updatedByMax_not_in: [ID!]
  updatedBy_null: Boolean
  createdBy: ID
  createdByMin: ID
  createdByMax: ID
  createdBy_ne: ID
  createdByMin_ne: ID
  createdByMax_ne: ID
  createdBy_gt: ID
  createdByMin_gt: ID
  createdByMax_gt: ID
  createdBy_lt: ID
  createdByMin_lt: ID
  createdByMax_lt: ID
  createdBy_gte: ID
  createdByMin_gte: ID
  createdByMax_gte: ID
  createdBy_lte: ID
  createdByMin_lte: ID
  createdByMax_lte: ID
  createdBy_in: [ID!]
  createdByMin_in: [ID!]
  createdByMax_in: [ID!]
  createdBy_not_in: [ID!]
  createdByMin_not_in: [ID!]
  createdByMax_not_in: [ID!]
  createdBy_null: Boolean
  changes: ChangelogChangeFilterType
}

type ChangelogResultType {
  items: [Changelog!]!
  count: Int!
  aggregations: ChangelogResultAggregations!
}

type ChangelogResultAggregations {
  entityMin: String!
  entityMax: String!
  entityIDMin: String!
  entityIDMax: String!
  principalIDMin: String
  principalIDMax: String
}`
)
