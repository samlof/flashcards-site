import React from "react";
import styled from "styled-components";
import { useAllWordsQuery, useDeleteWordMutation } from "../../gql.generated";
import Loading from "../Loading";
import GqlError from "../GqlError";
import App from "../../pages/_app";
import AddCsvWords from "./AddCsvWords";

const Title = styled.h1`
  text-align: center;
`;

const WordTable = styled.table`
  border-spacing: inherit;
`;

const WordRow = styled.tr`
  &:nth-child(odd) {
    background-color: #fff;
  }
  &:nth-child(even) {
    background-color: var(--color-white);
  }
`;

interface Props {}
const AllWords = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();
  const [
    deleteWord,
    { loading: deleteWordLoading, error: deleteWordError },
  ] = useDeleteWordMutation();

  const [showAllWords, setShowAllWords] = React.useState(false);

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

  if (!data) return <span>No words</span>;
  const words = data.getWords;

  const handleDeleteWord = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
    id: string
  ) => {
    e.preventDefault();
    deleteWord({ variables: { id: id } }).then((res) => {
      refetchWords();
    });
  };
  return (
    <>
      <h3>All current words</h3>
      {deleteWordError && (
        <GqlError msg="Failed to delete" err={deleteWordError} />
      )}
      <button type="button" onClick={(e) => setShowAllWords((s) => !s)}>
        Toggle all words
      </button>
      {showAllWords && (
        <WordTable>
          <tbody>
            {words.map((word) => (
              <WordRow key={word.id}>
                <td>{word.word1}</td>
                <td>{word.word2}</td>
                <td>
                  <button
                    type="button"
                    onClick={(e) => handleDeleteWord(e, word.id)}
                  >
                    X
                  </button>
                </td>
              </WordRow>
            ))}
          </tbody>
        </WordTable>
      )}
    </>
  );
};

export default AllWords;
