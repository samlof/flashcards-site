import { gql } from "@apollo/client";
import { GetServerSideProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import {
  FlascardPageDocument,
  FlascardPageQuery,
  FlascardPageQueryVariables,
  useFlascardPageQuery,
} from "../gql.generated";
import { initializeApollo } from "../lib/apolloClient";
import { ApolloProps } from "./_app";
import Loading from "../components/Loading";
import GqlError from "../components/GqlError";

const Title = styled.h1`
  color: var(--color-blue);
  text-align: center;
`;

interface Props {}

gql`
  query FlascardPage {
    getWords {
      langData
      word1
      word2
    }
  }
`;

const IndexPage = () => {
  const { loading, error, data } = useFlascardPageQuery();
  if (loading) return <Loading />;
  if (error) return <GqlError msg="Error getting words" err={error} />;
  if (!data) return <span>No words</span>;

  const dbWords = data.getWords;
  const words = dbWords.map((x) => ({ fi: x.word1, en: x.word2 }));
  return (
    <App>
      <Title>Flashcards</Title>
      <Flashcard words={words} />
    </App>
  );
};

export const getServerSideProps: GetServerSideProps<
  Props & ApolloProps
> = async () => {
  const apolloClient = initializeApollo();

  await apolloClient.query<FlascardPageQuery, FlascardPageQueryVariables>({
    query: FlascardPageDocument,
  });

  return {
    props: { initialApolloState: apolloClient.cache.extract() },
  };
};

export default IndexPage;
