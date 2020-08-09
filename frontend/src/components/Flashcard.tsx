import React from "react";
import { CSSTransition } from "react-transition-group";
import styled from "styled-components";
import { Button } from "../components/Button";
import FlipCard from "../components/FlipCard";
import { delayMs } from "../helpers/delay";
import { modulus } from "../helpers/modulus";
import Loading from "./Loading";
import GqlError from "./GqlError";
import { shuffle, randInt } from "../helpers/randomUtils";
import {
  useFlashcardPageQuery,
  Word,
  useSetCardStatusMutation,
  CardResult,
} from "../gql.generated";

const ButtonDiv = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const DirButton = styled(Button)`
  &&& {
    font-size: 1.5rem;
    margin: 0 0.3rem;
  }
`;

interface FlashWord {
  id: string;
  word1: string;
  word2: string;
  lang1: string;
  lang2: string;
}
const animationSpeed = 175;
interface Props {}

const Flashcard = ({}: Props) => {
  const [cardVisible, setVisible] = React.useState(true);
  const [animationName, setAnimationName] = React.useState("card-in-out");

  const { loading, error, data } = useFlashcardPageQuery();
  const [
    setCardState,
    { loading: loadingCardState },
  ] = useSetCardStatusMutation();

  const [words, setWords] = React.useState<FlashWord[]>([]);
  React.useEffect(() => {
    if (!data) return;
    const copiedWords = [...data.scheduledWords.cards];
    shuffle(copiedWords);
    setWords(copiedWords);
  }, [data]);
  if (loading) return <Loading />;
  if (error) return <GqlError msg="Error getting words" err={error} />;
  if (words.length === 0) return <span>No words</span>;
  const index = 0;
  const word = words[index];

  const handleClick = async (result: CardResult) => {
    setAnimationName("card-in-out");
    setVisible(false);

    setCardState({
      variables: { cardId: word.id, result: result },
    });
    await delayMs(animationSpeed);
    setWords((prev) => {
      const word = prev.splice(index, 1)[0];
      // If retry, then add the card back to deck
      if (result === CardResult.Retry) {
        const nextIndex = randInt(0, prev.length);
        prev.splice(nextIndex, 0, word);
      }
      return prev;
    });

    setVisible(true);
  };

  return (
    <>
      <CSSTransition
        in={cardVisible}
        timeout={animationSpeed}
        classNames={animationName}
      >
        <FlipCard
          front={{ lang: word.lang1, text: word.word1 }}
          back={{ lang: word.lang2, text: word.word2 }}
        ></FlipCard>
      </CSSTransition>
      <div style={{ height: "2rem" }}></div>
      <span style={{ minWidth: "5rem" }}>Left: {words.length}</span>
      <ButtonDiv>
        <DirButton type="button" onClick={(e) => handleClick(CardResult.Easy)}>
          Easy
        </DirButton>
        <DirButton type="button" onClick={(e) => handleClick(CardResult.Good)}>
          Good
        </DirButton>
        <DirButton type="button" onClick={(e) => handleClick(CardResult.Bad)}>
          Bad
        </DirButton>
        <DirButton type="button" onClick={(e) => handleClick(CardResult.Retry)}>
          Retry
        </DirButton>
      </ButtonDiv>

      <style jsx global>{`
        .card-in-out-enter {
          opacity: 0;
          transform: translateX(80px);
        }
        .card-in-out-enter-active {
          transform: translateX(0px);
          opacity: 1;
          transition: opacity ${animationSpeed}ms, transform ${animationSpeed}ms;
        }
        .card-in-out-exit {
          transform: translateX(0px);
          opacity: 1;
        }
        .card-in-out-exit-active {
          opacity: 0;
          transform: translateX(-80px);
          transition: opacity ${animationSpeed}ms, transform ${animationSpeed}ms;
        }
        .card-out-in-enter {
          opacity: 0;
          transform: translateX(-80px);
        }
        .card-out-in-enter-active {
          transform: translateX(0px);
          opacity: 1;
          transition: opacity ${animationSpeed}ms, transform ${animationSpeed}ms;
        }
        .card-out-in-exit {
          transform: translateX(0px);
          opacity: 1;
        }
        .card-out-in-exit-active {
          opacity: 0;
          transform: translateX(80px);
          transition: opacity ${animationSpeed}ms, transform ${animationSpeed}ms;
        }
      `}</style>
    </>
  );
};

export default Flashcard;
