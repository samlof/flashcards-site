import * as Types from '../../../src/graphql-types';

import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type AllWordsQueryQueryVariables = Types.Exact<{ [key: string]: never; }>;


export type AllWordsQueryQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Types.Word, 'word1' | 'word2'>
  )> }
);


export const AllWordsQueryDocument = gql`
    query AllWordsQuery {
  getWords {
    word1
    word2
  }
}
    `;

/**
 * __useAllWordsQueryQuery__
 *
 * To run a query within a React component, call `useAllWordsQueryQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllWordsQueryQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllWordsQueryQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllWordsQueryQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<AllWordsQueryQuery, AllWordsQueryQueryVariables>) {
        return ApolloReactHooks.useQuery<AllWordsQueryQuery, AllWordsQueryQueryVariables>(AllWordsQueryDocument, baseOptions);
      }
export function useAllWordsQueryLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<AllWordsQueryQuery, AllWordsQueryQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<AllWordsQueryQuery, AllWordsQueryQueryVariables>(AllWordsQueryDocument, baseOptions);
        }
export type AllWordsQueryQueryHookResult = ReturnType<typeof useAllWordsQueryQuery>;
export type AllWordsQueryLazyQueryHookResult = ReturnType<typeof useAllWordsQueryLazyQuery>;
export type AllWordsQueryQueryResult = ApolloReactCommon.QueryResult<AllWordsQueryQuery, AllWordsQueryQueryVariables>;