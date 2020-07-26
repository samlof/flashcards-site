import { gql } from "@apollo/client";
import { GetServerSideProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import {
  FlashcardPageDocument,
  FlashcardPageQuery,
  FlashcardPageQueryVariables,
  useFlashcardPageQuery,
} from "../gql.generated";
import { initializeApollo } from "../lib/apolloClient";
import { PageProps } from "./_app";
import Loading from "../components/Loading";
import GqlError from "../components/GqlError";

interface Props {}

gql`
  query FlashcardPage {
    getWords {
      langData
      word1
      word2
    }
  }
`;

const IndexPage = () => {
  const { loading, error, data } = useFlashcardPageQuery();
  if (loading) return <Loading />;
  if (error) return <GqlError msg="Error getting words" err={error} />;
  if (!data || !data.getWords?.length) return <span>No words</span>;

  const dbWords = data.getWords;
  const words = dbWords.map((x) => ({ fi: x.word1, en: x.word2 }));
  return (
    <App>
      <h1>Flashcards</h1>
      <Flashcard words={words} />
    </App>
  );
};

export const getServerSideProps: GetServerSideProps<
  Props & PageProps
> = async () => {
  const apolloClient = initializeApollo();

  await apolloClient.query<FlashcardPageQuery, FlashcardPageQueryVariables>({
    query: FlashcardPageDocument,
  });

  return {
    props: { initialApolloState: apolloClient.cache.extract() },
  };
};

export default IndexPage;
