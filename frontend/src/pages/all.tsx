import Head from "next/head";
import React from "react";
import App from "../components/App";
import AllFlashcards from "../components/AllFlashcards";
import Navbar from "../components/Navbar";

interface Props {}
const IndexPage = ({}: Props) => {
  return (
    <App>
      <Head>
        <title>All Cards | kieli.club</title>
      </Head>
      <Navbar />
      <h1>All Cards</h1>
      <AllFlashcards />
    </App>
  );
};

export default IndexPage;
