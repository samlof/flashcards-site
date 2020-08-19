import React from "react";
import { CSSTransition } from "react-transition-group";
import styled from "styled-components";
import { Button } from "../components/Button";
import FlipCard from "../components/FlipCard";
import { delayMs } from "../helpers/delay";
import { modulus } from "../helpers/modulus";
import Loading from "./Loading";
import GqlError from "./GqlError";
import { shuffle } from "../helpers/randomUtils";
import { useAllFlashcardsQuery } from "../gql.generated";
import { gql } from "@apollo/client";

const ButtonDiv = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

gql`
  query AllFlashcards {
    getWords {
      lang1
      lang2
      word1
      word2
    }
  }
`;

const DirButton = styled(Button)`
  &&& {
    font-size: 1.5rem;
  }
`;

interface FlashWord {
  word1: string;
  word2: string;
  lang1: string;
  lang2: string;
}
const animationSpeed = 175;
interface Props {}

const AllFlashcards = ({}: Props) => {
  const [index, setIndex] = React.useState(0);
  const [cardVisible, setVisible] = React.useState(true);
  const [animationName, setAnimationName] = React.useState("card-in-out");

  const { loading, error, data } = useAllFlashcardsQuery();
  const [words, setWords] = React.useState<FlashWord[]>([]);
  React.useEffect(() => {
    if (!data) return;
    const copiedWords = [...data.getWords];
    shuffle(copiedWords);
    setWords(copiedWords);
  }, [data]);
  if (loading) return <Loading />;
  if (error) return <GqlError msg="Error getting words" err={error} />;
  if (!data || !data.getWords?.length || !words.length)
    return <span>No words</span>;

  const nextCard = async (next: boolean) => {
    if (next) setAnimationName("card-in-out");
    else setAnimationName("card-out-in");

    setVisible(false);
    await delayMs(animationSpeed);
    if (next) setIndex((i) => modulus(i + 1, words.length));
    else setIndex((i) => modulus(i - 1, words.length));
    setVisible(true);
  };

  const word = words[index];
  return (
    <>
      <CSSTransition
        in={cardVisible}
        timeout={animationSpeed}
        classNames={animationName}
      >
        <FlipCard
          key={index}
          front={{ lang: word.lang1, text: word.word1 }}
          back={{ lang: word.lang2, text: word.word2 }}
        ></FlipCard>
      </CSSTransition>
      <div style={{ height: "2rem" }}></div>
      <ButtonDiv>
        <DirButton type="button" onClick={() => nextCard(false)}>
          Back
        </DirButton>
        <span style={{ minWidth: "5rem" }}>
          {index + 1}/{words.length}
        </span>
        <DirButton type="button" onClick={() => nextCard(true)}>
          Next
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

export default AllFlashcards;
