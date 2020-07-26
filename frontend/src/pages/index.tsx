import { gql } from "@apollo/client";
import React from "react";
import App from "../components/App";
import Flashcard, { Word } from "../components/Flashcard";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import {
  useFlashcardPageQuery,
  FlashcardPageQuery,
  FlashcardPageQueryVariables,
  FlashcardPageDocument,
} from "../gql.generated";
import { shuffle } from "../helpers/numberToString";
import { GetServerSideProps } from "next";
import { PageProps } from "./_app";
import { initializeApollo } from "../lib/apolloClient";

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
  const [words, setWords] = React.useState<Word[]>([]);
  React.useEffect(() => {
    if (!data) return;
    const copiedWords = [...data.getWords];
    shuffle(copiedWords);
    setWords(copiedWords);
  }, [data]);
  if (loading) return <Loading />;
  if (error) return <GqlError msg="Error getting words" err={error} />;
  if (!data || !data.getWords?.length || !words.length)
    return <span>No words</span>;

  const dbWords = data.getWords;
  const langInfo = dbWords[0].langData.split("-");
  return (
    <App>
      <h1>Flashcards</h1>
      <Flashcard words={words} lang1={langInfo[0]} lang2={langInfo[1]} />
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
