import Head from "next/head";
import { useRouter } from "next/router";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import Navbar from "../components/Navbar";
import { useFlashcardPageQuery } from "../gql.generated";
import { useUser } from "../lib/user";

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

export default IndexPage;
