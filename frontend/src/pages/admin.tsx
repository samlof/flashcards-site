import React from "react";
import styled from "styled-components";
import App from "../components/App";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import { useAddWordMutation, useAllWordsQuery } from "../gql.generated";

const Title = styled.h1`
  text-align: center;
`;

const WordRow = styled.tr`
  &:nth-child(even) {
    background-color: #fff;
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

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

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
          <div>
            <label>
              Finnish:
              <input
                type="text"
                value={word1}
                onChange={(e) => setWord1(e.target.value)}
              />
            </label>
          </div>
          <div>
            <label>
              English:
              <input
                type="text"
                value={word2}
                onChange={(e) => setWord2(e.target.value)}
              />
            </label>
          </div>
          <button type="submit">Add</button>
          {mutationError && (
            <GqlError msg="Failed to add" err={mutationError} />
          )}
        </form>
        <table>
          <tbody>
            {words.map((word) => (
              <WordRow key={word.id}>
                <td>{word.word1}</td>
                <td>{word.word2}</td>
              </WordRow>
            ))}
          </tbody>
        </table>
      </div>
    </App>
  );
};

export default AdminPage;
