import React from "react";
import styled from "styled-components";
import { useAllWordsQuery } from "../../gql.generated";
import GqlError from "../GqlError";
import Loading from "../Loading";
import { Button } from "../Button";
import WordRow from "./WordRow";

const WordTable = styled.table`
  border-spacing: inherit;
`;

interface Props {}
const AllWords = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();

  const [showAllWords, setShowAllWords] = React.useState(false);

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

  if (!data) return <span>No words</span>;
  const words = data.getWords;

  return (
    <>
      <h3>All current words</h3>
      <Button type="button" onClick={(e) => setShowAllWords((s) => !s)}>
        Toggle all words
      </Button>
      {showAllWords && (
        <WordTable>
          <tbody>
            {words.map((word) => (
              <WordRow
                key={word.id}
                word={word}
                refetchWords={refetchWords}
              ></WordRow>
            ))}
          </tbody>
        </WordTable>
      )}
    </>
  );
};

export default AllWords;
