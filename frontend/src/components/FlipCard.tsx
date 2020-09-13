import React from "react";
import ReactCardFlip from "react-card-flip";
import styled from "styled-components";
import { delayMs } from "../helpers/delay";
import { useAudio, makeAudioUrl } from "../lib/useAudio";
import { Button } from "./Button";

const CardSide = styled.div`
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

const LanguageText = styled.i``;

const CardTop = styled.div``;
const CardMiddle = styled.div`
  flex: 1;
  display: flex;
  align-self: center;
  align-items: center;
`;
const CardBottom = styled.div`
  margin-bottom: 1rem;
  display: flex;
  flex-direction: column;
  align-self: center;
  align-items: center;
`;

export interface CardWord {
  lang: string;
  text: string;
}
interface Props {
  front: CardWord;
  back: CardWord;
  isFront: boolean;
  setFront: (fn: (fn: boolean) => boolean) => void;
}

const flipSpeed = 600;
const flipSpeedS = flipSpeed / 1000;

const FlipCard = ({ front, back, isFront, setFront }: Props) => {
  const [, frontAudio] = useAudio(makeAudioUrl(front.text, front.lang));
  const [, backAudio] = useAudio(makeAudioUrl(back.text, back.lang));

  const handleFlip = async (e: React.MouseEvent) => {
    if (e.isPropagationStopped()) return;
    setFront((f) => !f);
  };

  const playBackAudio = (e: React.MouseEvent | React.TouchEvent) => {
    e.stopPropagation();
    e.preventDefault();
    backAudio();
  };
  const playFrontAudio = (e: React.MouseEvent | React.TouchEvent) => {
    e.stopPropagation();
    e.preventDefault();
    frontAudio();
  };

  return (
    <ReactCardFlip
      isFlipped={isFront}
      flipDirection="horizontal"
      flipSpeedBackToFront={flipSpeedS}
      flipSpeedFrontToBack={flipSpeedS}
    >
      <CardSide key="front" onClick={(e) => handleFlip(e)}>
        <CardTop></CardTop>
        <CardMiddle>
          <span>{front.text}</span>
        </CardMiddle>
        <CardBottom>
          <LanguageText>{front.lang}</LanguageText>
          <Button onClick={playFrontAudio}>Listen</Button>
        </CardBottom>
      </CardSide>

      <CardSide key="back" onClick={(e) => handleFlip(e)}>
        <CardTop></CardTop>
        <CardMiddle>
          <span>{back.text}</span>
        </CardMiddle>
        <CardBottom>
          <LanguageText>{back.lang}</LanguageText>

          <Button onClick={playBackAudio}>Listen</Button>
        </CardBottom>
      </CardSide>
    </ReactCardFlip>
  );
};

export default FlipCard;
