import React from "react";
import styled from "styled-components";
import App from "../components/App";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import {
  useAddWordMutation,
  useAllWordsQuery,
  useDeleteWordMutation,
} from "../gql.generated";
import AddCsvWords from "../components/admin/AddCsvWords";

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
const AdminPage = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();
  const [
    addWord,
    { loading: addWordLoading, error: addWordError },
  ] = useAddWordMutation();
  const [
    deleteWord,
    { loading: deleteWordLoading, error: deleteWordError },
  ] = useDeleteWordMutation();

  const [showAllWords, setShowAllWords] = React.useState(false);

  const [word1, setWord1] = React.useState("");
  const [word2, setWord2] = React.useState("");

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

  if (!data) return <span>No words</span>;
  const words = data.getWords;

  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    addWord({ variables: { word1, word2 } }).then((res) => {
      refetchWords();
    });
    setWord1("");
    setWord2("");
  };

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
          {addWordError && <GqlError msg="Failed to add" err={addWordError} />}
        </form>
        <AddCsvWords allWords={data} refetchWords={refetchWords} />
        {deleteWordError && (
          <GqlError msg="Failed to delete" err={deleteWordError} />
        )}
        <h3>All current words</h3>
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
      </div>
    </App>
  );
};

export default AdminPage;
