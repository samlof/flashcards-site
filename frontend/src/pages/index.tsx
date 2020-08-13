import { GetServerSideProps } from "next";
import dynamic from "next/dynamic";
import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import { IdTokenCookie } from "../constants/cookieNames";
import { FlashcardPageDocument, useFlashcardPageQuery } from "../gql.generated";
import { initializeApollo } from "../lib/apolloClient";
import Loading from "../components/Loading";
import GqlError from "../components/GqlError";

const Login = dynamic(() => import("../components/Login"), { ssr: false });

interface Props {}
const IndexPage = ({}: Props) => {
  const { loading, error, data } = useFlashcardPageQuery();

  const FlashcardElement = () => {
    if (loading) return <Loading />;
    if (error) return <GqlError msg="Error getting words" err={error} />;
    if (!data || data.scheduledWords.cards.length === 0)
      return <span>No words</span>;
    return <Flashcard initialWords={data.scheduledWords.cards} />;
  };

  return (
    <App>
      <Head>
        <title>Flashcards | kieli.club</title>
      </Head>
      <Login />

      <h1>Flashcards</h1>
      <FlashcardElement />
    </App>
  );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const nextCookie = await import("next-cookies").then((x) => x.default);
  const idtoken = nextCookie(ctx)[IdTokenCookie];

  const apolloClient = initializeApollo(undefined, idtoken);
  await apolloClient.query({
    query: FlashcardPageDocument,
  });
  return {
    props: { initialApolloState: apolloClient.cache.extract() },
  };
};

export default IndexPage;
