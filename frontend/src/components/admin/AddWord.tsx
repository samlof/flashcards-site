import React from "react";
import { useAddWordMutation } from "../../gql.generated";
import GqlError from "../GqlError";

interface Props {
  refetchWords: () => void;
}
const AddWord = ({ refetchWords }: Props) => {
  const [
    addWord,
    { loading: addWordLoading, error: addWordError },
  ] = useAddWordMutation();

  const [word1, setWord1] = React.useState("");
  const [word2, setWord2] = React.useState("");

  const handleFormSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    addWord({ variables: { word1, word2 } }).then((res) => {
      refetchWords();
    });
    setWord1("");
    setWord2("");
  };
  return (
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
  );
};

export default AddWord;
