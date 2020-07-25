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
import { levDistance } from "../helpers/levenshteinDistance";

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
const CompareWordsRow = styled.td``;

interface AddCsvWordsResults {
  duplicates: number;
  words: {
    pair: string[];
    levDist: number;
    comparedTo: string[];
    id: number;
  }[];
  newWords: number;
}

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

  const [addCsvResults, setAddCsvResults] = React.useState<
    AddCsvWordsResults | undefined
  >(undefined);

  const [word1, setWord1] = React.useState("");
  const [word2, setWord2] = React.useState("");
  const [csv, setcsv] = React.useState("");

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
  const handleCsvFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const lines = csv.split("\n");
    let duplicateCount = 0;
    let pairs = lines
      .map((x) => x.trim().split("\t"))
      .filter((x) => !!x && x.length === 2)
      .map((x) => x.map((y) => y.trim()));

    debugger;
    pairs = pairs.filter(
      (pair) =>
        !(
          data.getWords.some(
            (dword) =>
              dword.word1.toLowerCase() === pair[0].toLowerCase() &&
              dword.word2.toLowerCase() === pair[1].toLowerCase()
          ) && ++duplicateCount
        )
    );
    const ret: AddCsvWordsResults = {
      duplicates: duplicateCount,
      words: [],
      newWords: 0,
    };
    if (pairs.length === 0) {
      setAddCsvResults(ret);
      return;
    }

    // Remove headers row
    if (pairs[0][0].toLowerCase() === "finnish") {
      pairs = pairs.slice(1);
    }

    const wordsAsPair = words.map((word) => [word.word1, word.word2]);
    for (const pair of pairs) {
      const distances = wordsAsPair.map(
        (word) => levDistance(word[0], pair[0]) + levDistance(word[1], pair[1])
      );
      const levDist = Math.min(...distances);
      const closestWordIndex = distances.findIndex((x) => x === levDist);
      ret.words.push({
        pair,
        levDist,
        comparedTo: wordsAsPair[closestWordIndex],
        id: Math.round(Math.random() * 100000),
      });
    }

    ret.words.sort((a, b) => a.levDist - b.levDist);

    const newWords = ret.words.filter((x) => x.levDist > 5);
    ret.words = ret.words.filter((x) => x.levDist <= 5);
    ret.newWords = newWords.length;
    setAddCsvResults(ret);
    Promise.all(
      newWords.map((pair) =>
        addWord({ variables: { word1: pair.pair[0], word2: pair.pair[1] } })
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
  const handleAddCsvWord = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
    id: number
  ) => {
    e.preventDefault();
    const word = addCsvResults?.words.find((x) => x.id === id);
    if (!word) return;
    addWord({ variables: { word1: word.pair[0], word2: word.pair[1] } });
  };
  const handleSkipCsvWord = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>,
    id: number
  ) => {
    e.preventDefault();
    setAddCsvResults(
      (res) =>
        res && {
          ...res,
          words: res.words.filter((w) => w.id !== id),
        }
    );
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
        {addCsvResults && (
          <>
            {addCsvResults.newWords > 0 && (
              <span>Added {addCsvResults.newWords} new words</span>
            )}
            {addCsvResults.duplicates > 0 && (
              <span>
                Skipped {addCsvResults.duplicates} as they were duplicates
              </span>
            )}
            {addCsvResults.words.length > 0 && (
              <>
                <h3>Possible duplicate words</h3>
                <WordTable>
                  <thead>
                    <tr>
                      <td></td>
                      <td>Finnish</td>
                      <td>English</td>
                      <td>Diff</td>
                    </tr>
                  </thead>
                  <tbody>
                    {addCsvResults.words.map((word) => (
                      <WordRow key={word.id}>
                        <td>
                          new:
                          <br />
                          existing:
                        </td>
                        <CompareWordsRow>
                          {word.pair[0]}
                          <br />
                          {word.comparedTo[0]}
                        </CompareWordsRow>
                        <CompareWordsRow>
                          {word.pair[1]}
                          <br />
                          {word.comparedTo[1]}
                        </CompareWordsRow>
                        <td>{word.levDist}</td>
                        <td>
                          <button
                            type="button"
                            onClick={(e) => handleAddCsvWord(e, word.id)}
                          >
                            +
                          </button>
                          <button
                            type="button"
                            onClick={(e) => handleSkipCsvWord(e, word.id)}
                          >
                            -
                          </button>
                        </td>
                      </WordRow>
                    ))}
                  </tbody>
                </WordTable>
              </>
            )}
          </>
        )}
        {deleteWordError && (
          <GqlError msg="Failed to delete" err={deleteWordError} />
        )}
        <h3>All current words</h3>
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
