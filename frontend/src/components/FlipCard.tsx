import React from "react";
import ReactCardFlip from "react-card-flip";
import styled from "styled-components";
import { delayMs } from "../helpers/delay";

const CardSide = styled.div<{ front: boolean }>`
  display: flex;
  align-self: center;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  box-shadow: 2px 1px 10px var(--color-brown);
  height: 18rem;
  width: 14rem;
  border-radius: 10%;
  background-color: var(--color-white);
  font-size: 2rem;
`;

const LanguageText = styled.i`
  padding-top: 1rem;
`;

export interface CardWord {
  lang: string;
  text: string;
}
interface Props {
  front: CardWord;
  back: CardWord;
}

const flipSpeed = 600;
const flipSpeedS = flipSpeed / 1000;
const flipSpeedHalf = flipSpeed / 2;

const FlipCard = ({ front, back }: Props) => {
  const [isFront, setFront] = React.useState(false);
  const [isFrontVisible, setIsFrontVisible] = React.useState(false);

  const handleFlip = async () => {
    setFront((f) => !f);
    await delayMs(flipSpeedHalf - 100);
    setIsFrontVisible((f) => !f);
  };

  return (
    <ReactCardFlip
      isFlipped={isFront}
      flipDirection="horizontal"
      flipSpeedBackToFront={flipSpeedS}
      flipSpeedFrontToBack={flipSpeedS}
    >
      <CardSide
        key="front"
        onClick={(e) => handleFlip()}
        front={isFrontVisible}
      >
        <span>{front.text}</span>
        <LanguageText>{front.lang}</LanguageText>
      </CardSide>

      <CardSide key="back" onClick={(e) => handleFlip()} front={isFrontVisible}>
        <span>{back.text}</span>
        <LanguageText>{back.lang}</LanguageText>
      </CardSide>
    </ReactCardFlip>
  );
};

export default FlipCard;
