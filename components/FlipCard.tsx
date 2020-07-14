import React from "react";
import ReactCardFlip from "react-card-flip";
import styled from "styled-components";

const CardSide = styled.div`
  display: flex;
  align-self: center;
  align-items: center;
  flex-direction: column;
  justify-content: center;
  box-shadow: 0 0 10px rgba(38, 12, 12, 0.73);
  height: 30vh;
  width: 20vh;
  border-radius: 10%;
`;

export interface CardWord {
  lang: string;
  text: string;
}
interface Props {
  front: CardWord;
  back: CardWord;
}

const FlipCard = ({ front, back }: Props) => {
  const [animate, setAnim] = React.useState(false);

  return (
    <ReactCardFlip isFlipped={animate} flipDirection="horizontal">
      <CardSide key="front" onClick={(e) => setAnim((a) => !a)}>
        <span>{front.text}</span>
        <i>{front.lang}</i>
      </CardSide>

      <CardSide key="back" onClick={(e) => setAnim((a) => !a)}>
        <span>{back.text}</span>
        <i>{back.lang}</i>
      </CardSide>
    </ReactCardFlip>
  );
};

export default FlipCard;
