import { gql } from "@apollo/client";
import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";

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
