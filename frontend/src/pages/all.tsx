import Head from "next/head";
import React from "react";
import App from "../components/App";
import AllFlashcards from "../components/AllFlashcards";

interface Props {}
const IndexPage = ({}: Props) => {
  return (
    <App>
      <Head>
        <title>Flashcards | kieli.club</title>
      </Head>
      <h1>Flashcards</h1>
      <AllFlashcards />
    </App>
  );
};

export default IndexPage;
