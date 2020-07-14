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

const CardDiv = styled.div`
  height: 32vh;
`;
const clampValue = (n: number, mod: number): number => {
  if (n < 0) return n + mod;
  if (n > mod - 1) return n - mod;
  return n;
};

interface Props {}

const IndexPage = ({}: Props) => {
  const [index, setIndex] = React.useState(0);
  const [cardVisible, setVisible] = React.useState(true);

  const nextCard = async () => {
    setVisible(false);
    await delayMs(500);
    setIndex((i) => clampValue(i + 1, words.length));
    setVisible(true);
  };
  const lastCard = async () => {
    setVisible(false);
    await delayMs(500);
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
        <CardDiv>
          <CSSTransition
            unmountOnExit
            in={cardVisible}
            timeout={300}
            classNames="fade-in-out"
          >
            <FlipCard key={index} back={backCard} front={frontCard}></FlipCard>
          </CSSTransition>
        </CardDiv>
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
