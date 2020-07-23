export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Mutation = {
  __typename?: 'Mutation';
  createWord: Word;
};


export type MutationCreateWordArgs = {
  input: NewWord;
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
};
