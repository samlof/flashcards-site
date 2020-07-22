import React from "react";
import App from "../components/App";
import styled from "styled-components";
import { GetServerSideProps } from "next";
import { gql } from "@apollo/client";
import { useAllWordsQuery } from "./__generated__/addword.generated";

const Title = styled.h1`
  text-align: center;
`;
interface Props {}

gql`
  query AllWords {
    getWords {
      word1
      word2
    }
  }
`;
const AdminPage = ({}: Props) => {
  const { data, loading, error } = useAllWordsQuery();
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
  return (
    <App>
      <Title>See all words</Title>
      <div className="center-div">{JSON.stringify(data)}</div>
    </App>
  );
};

export const getServerSideProps: GetServerSideProps<Props> = async () => {
  return { props: {} };
};
export default AdminPage;
