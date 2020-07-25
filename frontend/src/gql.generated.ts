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

export type Word = {
  __typename?: 'Word';
  id: Scalars['ID'];
  langData: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
  createdAt: Scalars['Time'];
};

export type NewWord = {
  langData: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type UpdateWord = {
  id: Scalars['ID'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};


export type Query = {
  __typename?: 'Query';
  getWords: Array<Word>;
};

export type Mutation = {
  __typename?: 'Mutation';
  createWord: Word;
  deleteWord: Scalars['ID'];
  updateWord: Word;
};


export type MutationCreateWordArgs = {
  input: NewWord;
};


export type MutationDeleteWordArgs = {
  id: Scalars['ID'];
};


export type MutationUpdateWordArgs = {
  input: UpdateWord;
};

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

export type DeleteWordMutationVariables = Exact<{
  id: Scalars['ID'];
}>;


export type DeleteWordMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'deleteWord'>
);

export type AllWordsQueryVariables = Exact<{ [key: string]: never; }>;


export type AllWordsQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Word, 'id' | 'word1' | 'word2'>
  )> }
);

export type FlashcardPageQueryVariables = Exact<{ [key: string]: never; }>;


export type FlashcardPageQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Word, 'langData' | 'word1' | 'word2'>
  )> }
);


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
export const DeleteWordDocument = gql`
    mutation DeleteWord($id: ID!) {
  deleteWord(id: $id)
}
    `;
export type DeleteWordMutationFn = ApolloReactCommon.MutationFunction<DeleteWordMutation, DeleteWordMutationVariables>;

/**
 * __useDeleteWordMutation__
 *
 * To run a mutation, you first call `useDeleteWordMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteWordMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteWordMutation, { data, loading, error }] = useDeleteWordMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteWordMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeleteWordMutation, DeleteWordMutationVariables>) {
        return ApolloReactHooks.useMutation<DeleteWordMutation, DeleteWordMutationVariables>(DeleteWordDocument, baseOptions);
      }
export type DeleteWordMutationHookResult = ReturnType<typeof useDeleteWordMutation>;
export type DeleteWordMutationResult = ApolloReactCommon.MutationResult<DeleteWordMutation>;
export type DeleteWordMutationOptions = ApolloReactCommon.BaseMutationOptions<DeleteWordMutation, DeleteWordMutationVariables>;
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
export const FlashcardPageDocument = gql`
    query FlashcardPage {
  getWords {
    langData
    word1
    word2
  }
}
    `;

/**
 * __useFlashcardPageQuery__
 *
 * To run a query within a React component, call `useFlashcardPageQuery` and pass it any options that fit your needs.
 * When your component renders, `useFlashcardPageQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useFlashcardPageQuery({
 *   variables: {
 *   },
 * });
 */
export function useFlashcardPageQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<FlashcardPageQuery, FlashcardPageQueryVariables>) {
        return ApolloReactHooks.useQuery<FlashcardPageQuery, FlashcardPageQueryVariables>(FlashcardPageDocument, baseOptions);
      }
export function useFlashcardPageLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<FlashcardPageQuery, FlashcardPageQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<FlashcardPageQuery, FlashcardPageQueryVariables>(FlashcardPageDocument, baseOptions);
        }
export type FlashcardPageQueryHookResult = ReturnType<typeof useFlashcardPageQuery>;
export type FlashcardPageLazyQueryHookResult = ReturnType<typeof useFlashcardPageLazyQuery>;
export type FlashcardPageQueryResult = ApolloReactCommon.QueryResult<FlashcardPageQuery, FlashcardPageQueryVariables>;