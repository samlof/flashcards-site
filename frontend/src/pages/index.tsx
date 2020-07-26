import { gql } from "@apollo/client";
import Head from "next/head";
import React from "react";
import App from "../components/App";
import Flashcard from "../components/Flashcard";

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

// export const getServerSideProps: GetServerSideProps<
//   Props & PageProps
// > = async () => {
//   const apolloClient = initializeApollo();

//   await apolloClient.query<FlashcardPageQuery, FlashcardPageQueryVariables>({
//     query: FlashcardPageDocument,
//   });

//   return {
//     props: { initialApolloState: apolloClient.cache.extract() },
//   };
// };
export default IndexPage;
