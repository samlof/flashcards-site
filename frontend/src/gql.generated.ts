import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: String;
};

export type NewWord = {
  langData: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  getWords: Array<Word>;
};

export type Word = {
  __typename?: 'Word';
  id: Scalars['ID'];
  langData: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
  createdAt: Scalars['Time'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createWord: Word;
};


export type MutationCreateWordArgs = {
  input: NewWord;
};


export type AllWordsQueryVariables = Exact<{ [key: string]: never; }>;


export type AllWordsQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Word, 'id' | 'word1' | 'word2'>
  )> }
);

export type AddWordMutationVariables = Exact<{
  word1: Scalars['String'];
  word2: Scalars['String'];
}>;


export type AddWordMutation = (
  { __typename?: 'Mutation' }
  & { createWord: (
    { __typename?: 'Word' }
    & Pick<Word, 'id' | 'langData' | 'word1' | 'word2'>
  ) }
);


export const AllWordsDocument = gql`
    query AllWords {
  getWords {
    id
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
export const AddWordDocument = gql`
    mutation AddWord($word1: String!, $word2: String!) {
  createWord(input: {langData: "fi-en", word1: $word1, word2: $word2}) {
    id
    langData
    word1
    word2
  }
}
    `;
export type AddWordMutationFn = ApolloReactCommon.MutationFunction<AddWordMutation, AddWordMutationVariables>;

/**
 * __useAddWordMutation__
 *
 * To run a mutation, you first call `useAddWordMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddWordMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addWordMutation, { data, loading, error }] = useAddWordMutation({
 *   variables: {
 *      word1: // value for 'word1'
 *      word2: // value for 'word2'
 *   },
 * });
 */
export function useAddWordMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<AddWordMutation, AddWordMutationVariables>) {
        return ApolloReactHooks.useMutation<AddWordMutation, AddWordMutationVariables>(AddWordDocument, baseOptions);
      }
export type AddWordMutationHookResult = ReturnType<typeof useAddWordMutation>;
export type AddWordMutationResult = ApolloReactCommon.MutationResult<AddWordMutation>;
export type AddWordMutationOptions = ApolloReactCommon.BaseMutationOptions<AddWordMutation, AddWordMutationVariables>;