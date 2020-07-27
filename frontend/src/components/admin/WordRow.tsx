import React from "react";
import styled from "styled-components";
import {
  useDeleteWordMutation,
  useEditWordMutation,
} from "../../gql.generated";
import GqlError from "../GqlError";

const Row = styled.tr`
  &:nth-child(odd) {
    background-color: #fff;
  }
  &:nth-child(even) {
    background-color: var(--color-white);
  }
`;
const EditInput = styled.input`
  text-align: center;
`;

interface Word {
  id: string;
  word1: string;
  word2: string;
}
interface Props {
  word: Word;
  refetchWords: () => void;
}
const WordRow = ({ word, refetchWords }: Props) => {
  const [
    deleteWord,
    { loading: deleteWordLoading, error: deleteWordError },
  ] = useDeleteWordMutation();
  const [
    editWord,
    { loading: editWordLoading, error: editWordError },
  ] = useEditWordMutation();

  const [editing, setEditing] = React.useState(false);
  const [text1, setText1] = React.useState(word.word1);
  const [text2, setText2] = React.useState(word.word2);
  const handleDeleteWord = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    e.preventDefault();
    deleteWord({ variables: { id: word.id } }).then((res) => {
      refetchWords();
    });
  };
  const handleEditWord = (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    e.preventDefault();
    editWord({ variables: { id: word.id, word1: text1, word2: text2 } }).then(
      (res) => {
        refetchWords();
        setEditing(false);
      }
    );
  };
  return (
    <Row>
      <td>
        {editing ? (
          <EditInput
            type="text"
            value={text1}
            onChange={(e) => setText1(e.target.value)}
          />
        ) : (
          <> {word.word1}</>
        )}
      </td>
      <td>
        {editing ? (
          <EditInput
            type="text"
            value={text2}
            onChange={(e) => setText2(e.target.value)}
          />
        ) : (
          <> {word.word2}</>
        )}
      </td>
      <td>
        {deleteWordError && (
          <GqlError msg="Failed to delete" err={deleteWordError} />
        )}
        {editWordError && <GqlError msg="Failed to edit" err={editWordError} />}
        <button type="button" onClick={(e) => setEditing((e) => !e)}>
          Edit
        </button>
        {editing ? (
          <button type="button" onClick={handleEditWord}>
            Save
          </button>
        ) : (
          <button type="button" onClick={handleDeleteWord}>
            X
          </button>
        )}
      </td>
    </Row>
  );
};

export default WordRow;
