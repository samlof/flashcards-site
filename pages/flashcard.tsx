import { GetStaticProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import FlipCard, { CardWord } from "../components/FlipCard";
import { words } from "../components/words";
import { Button } from "../components/Button";
import { CSSTransition } from "react-transition-group";
import { delayMs } from "../helpers/delay";

let id = 1000;

const Title = styled.h1`
  text-align: center;
`;

const ButtonDiv = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 13rem;
`;

const clampValue = (n: number, mod: number): number => {
  if (n < 0) return n + mod;
  if (n > mod - 1) return n - mod;
  return n;
};

const animationSpeed = 175;

interface Props {}

const IndexPage = ({}: Props) => {
  const [index, setIndex] = React.useState(0);
  const [cardVisible, setVisible] = React.useState(true);
  const [animationName, setAnimationName] = React.useState("card-in-out");

  const nextCard = async () => {
    setAnimationName("card-in-out");
    setVisible(false);
    await delayMs(animationSpeed);
    setIndex((i) => clampValue(i + 1, words.length));
    setVisible(true);
  };
  const lastCard = async () => {
    setAnimationName("card-out-in");
    setVisible(false);
    await delayMs(animationSpeed);
    setIndex((i) => clampValue(i - 1, words.length));
    setVisible(true);
  };
  const word = words[index];
  const frontCard: CardWord = {
    lang: "fi",
    text: word.fi,
  };
  const backCard: CardWord = {
    lang: "en",
    text: word.en,
  };
  return (
    <App>
      <Title>Flashcards</Title>
      <div className="center-div">
        <CSSTransition
          in={cardVisible}
          timeout={animationSpeed}
          classNames={animationName}
        >
          <FlipCard key={index} back={backCard} front={frontCard}></FlipCard>
        </CSSTransition>
        <span style={{ marginTop: "2rem" }}></span>
        <ButtonDiv>
          <Button type="button" onClick={(e) => lastCard()}>
            Back
          </Button>
          {index + 1}/{words.length}
          <Button type="button" onClick={(e) => nextCard()}>
            Next
          </Button>
        </ButtonDiv>
      </div>

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
    </App>
  );
};

export const getStaticProps: GetStaticProps<Props> = async (context) => {
  id = 1;
  return {
    props: {},
    unstable_revalidate: 1,
  };
};

export default IndexPage;
