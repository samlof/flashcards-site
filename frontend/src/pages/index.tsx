import { GetServerSideProps } from "next";
import dynamic from "next/dynamic";
import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import Navbar from "../components/Navbar";
import { IdTokenCookie } from "../constants/cookieNames";
import { FlashcardPageDocument, useFlashcardPageQuery } from "../gql.generated";
import { initializeApollo } from "../lib/apolloClient";
import { useUser } from "../lib/user";
import { useRouter } from "next/router";

interface Props {}
const IndexPage = ({}: Props) => {
  const { loading, error, data } = useFlashcardPageQuery();
  const user = useUser();
  const router = useRouter();

  if (!user.loading && !user.user) {
    router.push("/all");
  }

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
      <Navbar />

      <h1>Flashcards</h1>
      <FlashcardElement />
    </App>
  );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  // console.log(ctx.req.headers.cookie);
  const nextCookie = await import("next-cookies").then((x) => x.default);
  const idtoken = nextCookie(ctx)[IdTokenCookie];

  if (!idtoken) {
    // No point trying here if no idtoken
    return { props: {} };
  }

  const apolloClient = initializeApollo(undefined, idtoken);
  try {
    await apolloClient.query({
      query: FlashcardPageDocument,
    });
  } catch (error) {}
  return {
    props: { initialApolloState: apolloClient.cache.extract() },
  };
};

export default IndexPage;
