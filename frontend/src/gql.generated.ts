import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: String;
};

export type CardSchedule = {
  __typename?: 'CardSchedule';
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  word: Word;
  scheduledFor: Scalars['Time'];
};

export type UpdateWord = {
  id: Scalars['ID'];
  lang1: Scalars['String'];
  lang2: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type CardLog = {
  __typename?: 'CardLog';
  createTime: Scalars['Time'];
  id: Scalars['ID'];
  word: Word;
  lastResult: CardResult;
};

export type CardStatus = {
  cardId: Scalars['ID'];
  result: CardResult;
};

export type UserSettings = {
  __typename?: 'UserSettings';
  newCardsPerDay: Scalars['Int'];
};

export type SetSettings = {
  newCardsPerDay: Scalars['Int'];
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

export type Query = {
  __typename?: 'Query';
  scheduledWords: ScheduledWordsResponse;
  userSettings: UserSettings;
  getWords: Array<Word>;
};


export type QueryScheduledWordsArgs = {
  shuffle?: Scalars['Boolean'];
};

export enum CardResult {
  Easy = 'Easy',
  Good = 'Good',
  Bad = 'Bad',
  Retry = 'Retry'
}

export type ScheduledWordsResponse = {
  __typename?: 'ScheduledWordsResponse';
  cards: Array<Word>;
};


export type NewWord = {
  lang1: Scalars['String'];
  lang2: Scalars['String'];
  word1: Scalars['String'];
  word2: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  cardStatus: CardLog;
  setSettings: UserSettings;
  createWord: Word;
  deleteWord: Scalars['ID'];
  updateWord: Word;
};


export type MutationCardStatusArgs = {
  input: CardStatus;
};


export type MutationSetSettingsArgs = {
  input: SetSettings;
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

export type AllFlashcardsQueryVariables = Exact<{ [key: string]: never; }>;


export type AllFlashcardsQuery = (
  { __typename?: 'Query' }
  & { getWords: Array<(
    { __typename?: 'Word' }
    & Pick<Word, 'lang1' | 'lang2' | 'word1' | 'word2'>
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

export type UserSettingsQueryVariables = Exact<{ [key: string]: never; }>;


export type UserSettingsQuery = (
  { __typename?: 'Query' }
  & { userSettings: (
    { __typename?: 'UserSettings' }
    & Pick<UserSettings, 'newCardsPerDay'>
  ) }
);

export type SetUserSettingsMutationVariables = Exact<{
  cardsPerDay: Scalars['Int'];
}>;


export type SetUserSettingsMutation = (
  { __typename?: 'Mutation' }
  & { setSettings: (
    { __typename?: 'UserSettings' }
    & Pick<UserSettings, 'newCardsPerDay'>
  ) }
);

export type FlashcardPageQueryVariables = Exact<{ [key: string]: never; }>;


export type FlashcardPageQuery = (
  { __typename?: 'Query' }
  & { scheduledWords: (
    { __typename?: 'ScheduledWordsResponse' }
    & { cards: Array<(
      { __typename?: 'Word' }
      & Pick<Word, 'id' | 'lang1' | 'lang2' | 'word1' | 'word2'>
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


export const AllFlashcardsDocument = gql`
    query AllFlashcards {
  getWords {
    lang1
    lang2
    word1
    word2
  }
}
    `;

/**
 * __useAllFlashcardsQuery__
 *
 * To run a query within a React component, call `useAllFlashcardsQuery` and pass it any options that fit your needs.
 * When your component renders, `useAllFlashcardsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useAllFlashcardsQuery({
 *   variables: {
 *   },
 * });
 */
export function useAllFlashcardsQuery(baseOptions?: Apollo.QueryHookOptions<AllFlashcardsQuery, AllFlashcardsQueryVariables>) {
        return Apollo.useQuery<AllFlashcardsQuery, AllFlashcardsQueryVariables>(AllFlashcardsDocument, baseOptions);
      }
export function useAllFlashcardsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AllFlashcardsQuery, AllFlashcardsQueryVariables>) {
          return Apollo.useLazyQuery<AllFlashcardsQuery, AllFlashcardsQueryVariables>(AllFlashcardsDocument, baseOptions);
        }
export type AllFlashcardsQueryHookResult = ReturnType<typeof useAllFlashcardsQuery>;
export type AllFlashcardsLazyQueryHookResult = ReturnType<typeof useAllFlashcardsLazyQuery>;
export type AllFlashcardsQueryResult = Apollo.QueryResult<AllFlashcardsQuery, AllFlashcardsQueryVariables>;
export const AddWordDocument = gql`
    mutation AddWord($word1: String!, $word2: String!) {
  createWord(input: {lang1: "fi", lang2: "en", word1: $word1, word2: $word2}) {
    id
    word1
    word2
  }
}
    `;
export type AddWordMutationFn = Apollo.MutationFunction<AddWordMutation, AddWordMutationVariables>;

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
export function useAddWordMutation(baseOptions?: Apollo.MutationHookOptions<AddWordMutation, AddWordMutationVariables>) {
        return Apollo.useMutation<AddWordMutation, AddWordMutationVariables>(AddWordDocument, baseOptions);
      }
export type AddWordMutationHookResult = ReturnType<typeof useAddWordMutation>;
export type AddWordMutationResult = Apollo.MutationResult<AddWordMutation>;
export type AddWordMutationOptions = Apollo.BaseMutationOptions<AddWordMutation, AddWordMutationVariables>;
export const DeleteWordDocument = gql`
    mutation DeleteWord($id: ID!) {
  deleteWord(id: $id)
}
    `;
export type DeleteWordMutationFn = Apollo.MutationFunction<DeleteWordMutation, DeleteWordMutationVariables>;

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
export function useDeleteWordMutation(baseOptions?: Apollo.MutationHookOptions<DeleteWordMutation, DeleteWordMutationVariables>) {
        return Apollo.useMutation<DeleteWordMutation, DeleteWordMutationVariables>(DeleteWordDocument, baseOptions);
      }
export type DeleteWordMutationHookResult = ReturnType<typeof useDeleteWordMutation>;
export type DeleteWordMutationResult = Apollo.MutationResult<DeleteWordMutation>;
export type DeleteWordMutationOptions = Apollo.BaseMutationOptions<DeleteWordMutation, DeleteWordMutationVariables>;
export const EditWordDocument = gql`
    mutation EditWord($id: ID!, $word1: String!, $word2: String!) {
  updateWord(input: {id: $id, word1: $word1, word2: $word2, lang1: "fi", lang2: "en"}) {
    id
  }
}
    `;
export type EditWordMutationFn = Apollo.MutationFunction<EditWordMutation, EditWordMutationVariables>;

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
export function useEditWordMutation(baseOptions?: Apollo.MutationHookOptions<EditWordMutation, EditWordMutationVariables>) {
        return Apollo.useMutation<EditWordMutation, EditWordMutationVariables>(EditWordDocument, baseOptions);
      }
export type EditWordMutationHookResult = ReturnType<typeof useEditWordMutation>;
export type EditWordMutationResult = Apollo.MutationResult<EditWordMutation>;
export type EditWordMutationOptions = Apollo.BaseMutationOptions<EditWordMutation, EditWordMutationVariables>;
export const UserSettingsDocument = gql`
    query UserSettings {
  userSettings {
    newCardsPerDay
  }
}
    `;

/**
 * __useUserSettingsQuery__
 *
 * To run a query within a React component, call `useUserSettingsQuery` and pass it any options that fit your needs.
 * When your component renders, `useUserSettingsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useUserSettingsQuery({
 *   variables: {
 *   },
 * });
 */
export function useUserSettingsQuery(baseOptions?: Apollo.QueryHookOptions<UserSettingsQuery, UserSettingsQueryVariables>) {
        return Apollo.useQuery<UserSettingsQuery, UserSettingsQueryVariables>(UserSettingsDocument, baseOptions);
      }
export function useUserSettingsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<UserSettingsQuery, UserSettingsQueryVariables>) {
          return Apollo.useLazyQuery<UserSettingsQuery, UserSettingsQueryVariables>(UserSettingsDocument, baseOptions);
        }
export type UserSettingsQueryHookResult = ReturnType<typeof useUserSettingsQuery>;
export type UserSettingsLazyQueryHookResult = ReturnType<typeof useUserSettingsLazyQuery>;
export type UserSettingsQueryResult = Apollo.QueryResult<UserSettingsQuery, UserSettingsQueryVariables>;
export const SetUserSettingsDocument = gql`
    mutation SetUserSettings($cardsPerDay: Int!) {
  setSettings(input: {newCardsPerDay: $cardsPerDay}) {
    newCardsPerDay
  }
}
    `;
export type SetUserSettingsMutationFn = Apollo.MutationFunction<SetUserSettingsMutation, SetUserSettingsMutationVariables>;

/**
 * __useSetUserSettingsMutation__
 *
 * To run a mutation, you first call `useSetUserSettingsMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useSetUserSettingsMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [setUserSettingsMutation, { data, loading, error }] = useSetUserSettingsMutation({
 *   variables: {
 *      cardsPerDay: // value for 'cardsPerDay'
 *   },
 * });
 */
export function useSetUserSettingsMutation(baseOptions?: Apollo.MutationHookOptions<SetUserSettingsMutation, SetUserSettingsMutationVariables>) {
        return Apollo.useMutation<SetUserSettingsMutation, SetUserSettingsMutationVariables>(SetUserSettingsDocument, baseOptions);
      }
export type SetUserSettingsMutationHookResult = ReturnType<typeof useSetUserSettingsMutation>;
export type SetUserSettingsMutationResult = Apollo.MutationResult<SetUserSettingsMutation>;
export type SetUserSettingsMutationOptions = Apollo.BaseMutationOptions<SetUserSettingsMutation, SetUserSettingsMutationVariables>;
export const FlashcardPageDocument = gql`
    query FlashcardPage {
  scheduledWords(shuffle: true) {
    cards {
      id
      lang1
      lang2
      word1
      word2
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
 *   },
 * });
 */
export function useFlashcardPageQuery(baseOptions?: Apollo.QueryHookOptions<FlashcardPageQuery, FlashcardPageQueryVariables>) {
        return Apollo.useQuery<FlashcardPageQuery, FlashcardPageQueryVariables>(FlashcardPageDocument, baseOptions);
      }
export function useFlashcardPageLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<FlashcardPageQuery, FlashcardPageQueryVariables>) {
          return Apollo.useLazyQuery<FlashcardPageQuery, FlashcardPageQueryVariables>(FlashcardPageDocument, baseOptions);
        }
export type FlashcardPageQueryHookResult = ReturnType<typeof useFlashcardPageQuery>;
export type FlashcardPageLazyQueryHookResult = ReturnType<typeof useFlashcardPageLazyQuery>;
export type FlashcardPageQueryResult = Apollo.QueryResult<FlashcardPageQuery, FlashcardPageQueryVariables>;
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
export type SetCardStatusMutationFn = Apollo.MutationFunction<SetCardStatusMutation, SetCardStatusMutationVariables>;

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
export function useSetCardStatusMutation(baseOptions?: Apollo.MutationHookOptions<SetCardStatusMutation, SetCardStatusMutationVariables>) {
        return Apollo.useMutation<SetCardStatusMutation, SetCardStatusMutationVariables>(SetCardStatusDocument, baseOptions);
      }
export type SetCardStatusMutationHookResult = ReturnType<typeof useSetCardStatusMutation>;
export type SetCardStatusMutationResult = Apollo.MutationResult<SetCardStatusMutation>;
export type SetCardStatusMutationOptions = Apollo.BaseMutationOptions<SetCardStatusMutation, SetCardStatusMutationVariables>;
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
export function useAllWordsQuery(baseOptions?: Apollo.QueryHookOptions<AllWordsQuery, AllWordsQueryVariables>) {
        return Apollo.useQuery<AllWordsQuery, AllWordsQueryVariables>(AllWordsDocument, baseOptions);
      }
export function useAllWordsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<AllWordsQuery, AllWordsQueryVariables>) {
          return Apollo.useLazyQuery<AllWordsQuery, AllWordsQueryVariables>(AllWordsDocument, baseOptions);
        }
export type AllWordsQueryHookResult = ReturnType<typeof useAllWordsQuery>;
export type AllWordsLazyQueryHookResult = ReturnType<typeof useAllWordsLazyQuery>;
export type AllWordsQueryResult = Apollo.QueryResult<AllWordsQuery, AllWordsQueryVariables>;