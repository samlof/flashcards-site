import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import dynamic from "next/dynamic";
import { initializeApollo } from "../lib/apolloClient";
import nextCookie from "next-cookies";
import { GetServerSideProps } from "next";
import { AllFlashcardsDocument } from "../gql.generated";

const Login = dynamic(() => import("../components/Login"), { ssr: false });

interface Props {}
const IndexPage = ({}: Props) => {
  return (
    <App>
      <Head>
        <title>Flashcards | kieli.club</title>
      </Head>
      <Login />

      <h1>Flashcards</h1>
      <Flashcard />
    </App>
  );
};

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const idtoken = nextCookie(ctx).idToken;

  const apolloClient = initializeApollo(undefined, idtoken);

  await apolloClient.query({
    query: AllFlashcardsDocument,
  });
  return {
    props: { initialApolloState: apolloClient.cache.extract() },
  };
};

export default IndexPage;
