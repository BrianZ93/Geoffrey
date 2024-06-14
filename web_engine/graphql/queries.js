/* eslint-disable */
// this is an auto generated file. This will be overwritten

export const getEventsGuests = /* GraphQL */ `
  query GetEventsGuests($id: String!) {
    getEventsGuests(id: $id) {
      id
      eventId
      __typename
    }
  }
`;
export const listEventsGuests = /* GraphQL */ `
  query ListEventsGuests(
    $filter: TableEventsGuestsFilterInput
    $limit: Int
    $nextToken: String
  ) {
    listEventsGuests(filter: $filter, limit: $limit, nextToken: $nextToken) {
      items {
        id
        eventId
        __typename
      }
      nextToken
      __typename
    }
  }
`;
