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

  const [word1, setWord1] = React.useState("");
  const [word2, setWord2] = React.useState("");
  const [csv, setcsv] = React.useState("");

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
  const handleCsvFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const lines = csv.split("\n");
    let pairs = lines
      .map((x) => x.trim().split("\t"))
      .filter((x) => !!x && x.length === 2)
      .map((x) => x.map((y) => y.trim()))
      .filter(
        (pair) =>
          !data.getWords.some(
            (dword) =>
              dword.word1.toLowerCase() === pair[0].toLowerCase() ||
              dword.word2.toLowerCase() === pair[1].toLowerCase()
          )
      );

    if (pairs.length === 0) return;

    // Remove headers row
    if (pairs[0][0].toLowerCase() === "finnish") {
      pairs = pairs.slice(1);
    }

    Promise.all(
      pairs.map((pair) =>
        addWord({ variables: { word1: pair[0], word2: pair[1] } })
      )
    ).then((res) => {
      refetchWords();
    });
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
          {addWordError && <GqlError msg="Failed to add" err={addWordError} />}
        </form>
        <form onSubmit={handleCsvFormSubmit}>
          <div>
            <label>
              Paste from excel
              <textarea
                style={{
                  width: "85vw",
                  height: "300px",
                }}
                cols={10}
                rows={10}
                onChange={(e) => setcsv(e.target.value)}
                value={csv}
              ></textarea>
            </label>
          </div>
          <button type="submit">Add excel words</button>
          {addWordError && <GqlError msg="Failed to add" err={addWordError} />}
        </form>
        {deleteWordError && (
          <GqlError msg="Failed to delete" err={deleteWordError} />
        )}
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
      </div>
    </App>
  );
};

export default AdminPage;
