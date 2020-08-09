import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import dynamic from "next/dynamic";

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

export default IndexPage;
