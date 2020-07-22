import * as Types from '../../src/graphql-types';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type AllWordsQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type AllWordsQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Types.Word, 'word1' | 'word2'>
  )> }
);


export const AllWordsDocument = gql`
    query AllWords {
  getWords {
    word1
    word2
  }
}
    `;

/**
 * __useAllWordsQuery__
 *
 * To run a query within a React component, call `useAllWordsQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllWordsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllWordsQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllWordsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<AllWordsQuery, AllWordsQueryVariables>) {
        return ApolloReactHooks.useQuery<AllWordsQuery, AllWordsQueryVariables>(AllWordsDocument, baseOptions);
      }
export function useAllWordsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AllWordsQuery, AllWordsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<AllWordsQuery, AllWordsQueryVariables>(AllWordsDocument, baseOptions);
        }
export type AllWordsQueryHookResult = ReturnType<typeof useAllWordsQuery>;
export type AllWordsLazyQueryHookResult = ReturnType<typeof useAllWordsLazyQuery>;
export type AllWordsQueryResult = ApolloReactCommon.QueryResult<AllWordsQuery, AllWordsQueryVariables>;