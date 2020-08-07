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

export type Query = {
  __typename?: 'Query';
  scheduledWords: ScheduledWordsResponse;
  getWords: Array<Word>;
};


export type QueryScheduledWordsArgs = {
  newWordCount?: Maybe<Scalars['Int']>;
};

export type CardLog = {
  __typename?: 'CardLog';
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  word: Word;
  scheduledFor: Scalars['Time'];
  lastResult: CardResult;
};


export type Word = {
  __typename?: 'Word';
  id: Scalars['ID'];
  lang1: Scalars['String'];
  lang2: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
  createTime: Scalars['Time'];
  updateTime: Scalars['Time'];
};

export type NewWord = {
  lang1: Scalars['String'];
  lang2: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type UpdateWord = {
  id: Scalars['ID'];
  lang1: Scalars['String'];
  lang2: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  cardStatus: CardLog;
  createWord: Word;
  deleteWord: Scalars['ID'];
  updateWord: Word;
};


export type MutationCardStatusArgs = {
  input: CardStatus;
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

export enum CardResult {
  Good = 'Good',
  Average = 'Average',
  Bad = 'Bad'
}

export type ScheduledWordsResponse = {
  __typename?: 'ScheduledWordsResponse';
  reviews: Array<CardLog>;
  newWords: Array<Word>;
};

export type CardStatus = {
  cardId: Scalars['ID'];
  result: CardResult;
};

export type AddWordMutationVariables = Exact<{
  word1: Scalars['String'];
  word2: Scalars['String'];
}>;


export type AddWordMutation = (
  { __typename?: 'Mutation' }
  & { createWord: (
    { __typename?: 'Word' }
    & Pick<Word, 'id' | 'word1' | 'word2'>
  ) }
);

export type DeleteWordMutationVariables = Exact<{
  id: Scalars['ID'];
}>;


export type DeleteWordMutation = (
  { __typename?: 'Mutation' }
  & Pick<Mutation, 'deleteWord'>
);

export type EditWordMutationVariables = Exact<{
  id: Scalars['ID'];
  word1: Scalars['String'];
  word2: Scalars['String'];
}>;


export type EditWordMutation = (
  { __typename?: 'Mutation' }
  & { updateWord: (
    { __typename?: 'Word' }
    & Pick<Word, 'id'>
  ) }
);

export type FlashcardPageQueryVariables = Exact<{
  newWordCount: Scalars['Int'];
}>;


export type FlashcardPageQuery = (
  { __typename?: 'Query' }
  & { scheduledWords: (
    { __typename?: 'ScheduledWordsResponse' }
    & { newWords: Array<(
      { __typename?: 'Word' }
      & Pick<Word, 'id' | 'lang1' | 'lang2' | 'word1' | 'word2'>
    )>, reviews: Array<(
      { __typename?: 'CardLog' }
      & Pick<CardLog, 'id'>
      & { word: (
        { __typename?: 'Word' }
        & Pick<Word, 'id' | 'lang1' | 'lang2' | 'word1' | 'word2'>
      ) }
    )> }
  ) }
);

export type SetCardStatusMutationVariables = Exact<{
  cardId: Scalars['ID'];
  result: CardResult;
}>;


export type SetCardStatusMutation = (
  { __typename?: 'Mutation' }
  & { cardStatus: (
    { __typename?: 'CardLog' }
    & Pick<CardLog, 'id'>
    & { word: (
      { __typename?: 'Word' }
      & Pick<Word, 'id'>
    ) }
  ) }
);

export type AllWordsQueryVariables = Exact<{ [key: string]: never; }>;


export type AllWordsQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Word, 'id' | 'word1' | 'word2'>
  )> }
);


export const AddWordDocument = gql`
    mutation AddWord($word1: String!, $word2: String!) {
  createWord(input: {lang1: "fi", lang2: "en", word1: $word1, word2: $word2}) {
    id
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
export const EditWordDocument = gql`
    mutation EditWord($id: ID!, $word1: String!, $word2: String!) {
  updateWord(input: {id: $id, word1: $word1, word2: $word2, lang1: "fi", lang2: "en"}) {
    id
  }
}
    `;
export type EditWordMutationFn = ApolloReactCommon.MutationFunction<EditWordMutation, EditWordMutationVariables>;

/**
 * __useEditWordMutation__
 *
 * To run a mutation, you first call `useEditWordMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useEditWordMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [editWordMutation, { data, loading, error }] = useEditWordMutation({
 *   variables: {
 *      id: // value for 'id'
 *      word1: // value for 'word1'
 *      word2: // value for 'word2'
 *   },
 * });
 */
export function useEditWordMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<EditWordMutation, EditWordMutationVariables>) {
        return ApolloReactHooks.useMutation<EditWordMutation, EditWordMutationVariables>(EditWordDocument, baseOptions);
      }
export type EditWordMutationHookResult = ReturnType<typeof useEditWordMutation>;
export type EditWordMutationResult = ApolloReactCommon.MutationResult<EditWordMutation>;
export type EditWordMutationOptions = ApolloReactCommon.BaseMutationOptions<EditWordMutation, EditWordMutationVariables>;
export const FlashcardPageDocument = gql`
    query FlashcardPage($newWordCount: Int!) {
  scheduledWords(newWordCount: $newWordCount) {
    newWords {
      id
      lang1
      lang2
      word1
      word2
    }
    reviews {
      id
      word {
        id
        lang1
        lang2
        word1
        word2
      }
    }
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
 *      newWordCount: // value for 'newWordCount'
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
export const SetCardStatusDocument = gql`
    mutation SetCardStatus($cardId: ID!, $result: CardResult!) {
  cardStatus(input: {cardId: $cardId, result: $result}) {
    id
    word {
      id
    }
  }
}
    `;
export type SetCardStatusMutationFn = ApolloReactCommon.MutationFunction<SetCardStatusMutation, SetCardStatusMutationVariables>;

/**
 * __useSetCardStatusMutation__
 *
 * To run a mutation, you first call `useSetCardStatusMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSetCardStatusMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [setCardStatusMutation, { data, loading, error }] = useSetCardStatusMutation({
 *   variables: {
 *      cardId: // value for 'cardId'
 *      result: // value for 'result'
 *   },
 * });
 */
export function useSetCardStatusMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<SetCardStatusMutation, SetCardStatusMutationVariables>) {
        return ApolloReactHooks.useMutation<SetCardStatusMutation, SetCardStatusMutationVariables>(SetCardStatusDocument, baseOptions);
      }
export type SetCardStatusMutationHookResult = ReturnType<typeof useSetCardStatusMutation>;
export type SetCardStatusMutationResult = ApolloReactCommon.MutationResult<SetCardStatusMutation>;
export type SetCardStatusMutationOptions = ApolloReactCommon.BaseMutationOptions<SetCardStatusMutation, SetCardStatusMutationVariables>;
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