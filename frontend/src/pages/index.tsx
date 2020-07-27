import { gql } from "@apollo/client";
import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";

gql`
  query FlashcardPage {
    getWords {
      lang1
      lang2
      word1
      word2
    }
  }
`;

interface Props {}
const IndexPage = ({}: Props) => {
  return (
    <App>
      <Head>
        <title>Flashcards | kieli.club</title>
      </Head>
      <h1>Flashcards</h1>
      <Flashcard />
    </App>
  );
};

export default IndexPage;
