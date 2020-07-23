import React from "react";
import App from "../components/App";
import styled from "styled-components";
import { GetServerSideProps } from "next";
import { gql } from "@apollo/client";
import { useAllWordsQuery, useAddWordMutation } from "../generated_gql";

const Title = styled.h1`
  text-align: center;
`;

const GridContainer = styled.div`
  display: grid;
  grid-template-columns: auto auto;
  grid-column-gap: 1rem;
  grid-row-gap: 1rem;
`;
const FirstWord = styled.div`
  text-align: left;
`;
const SecondWord = styled.div`
  text-align: center;
`;
gql`
  query AllWords {
    getWords {
      id
      word1
      word2
    }
  }
`;

gql`
  mutation AddWord($word1: String!, $word2: String!) {
    createWord(input: { langData: "fi-en", word1: $word1, word2: $word2 }) {
      id
      langData
      word1
      word2
    }
  }
`;

interface Props {}
const AdminPage = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();
  const [
    addWord,
    { loading: mutationLoading, error: mutationError },
  ] = useAddWordMutation();

  const [word1, setWord1] = React.useState("");
  const [word2, setWord2] = React.useState("");

  if (loading) return <span>loading...</span>;
  if (error) {
    console.error("Error getting all words");
    console.dir(error);
    return (
      <div>
        <span>Error happened.</span>
        <pre>{JSON.stringify(error, null, 2)}</pre>
      </div>
    );
  }

  if (!data) return <span>No words</span>;

  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    addWord({ variables: { word1, word2 } }).then((res) => {
      refetchWords();
    });
    setWord1("");
    setWord2("");
  };
  const words = data.getWords;
  return (
    <App>
      <Title>See all words</Title>
      <div className="center-div">
        <form onSubmit={handleFormSubmit}>
          <label>
            Finnish:
            <input
              type="text"
              value={word1}
              onChange={(e) => setWord1(e.target.value)}
            />
          </label>
          <label>
            English:
            <input
              type="text"
              value={word2}
              onChange={(e) => setWord2(e.target.value)}
            />
          </label>
          <button type="submit">Add</button>
        </form>
        <GridContainer>
          {words.map((word) => (
            <React.Fragment key={word.id}>
              <FirstWord>{word.word1}</FirstWord>
              <SecondWord>{word.word2}</SecondWord>
            </React.Fragment>
          ))}
        </GridContainer>
      </div>
    </App>
  );
};

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  return { props: {} };
};
export default AdminPage;
