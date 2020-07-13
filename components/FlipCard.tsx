import React from "react";
import ReactCardFlip from "react-card-flip";
import styled from "styled-components";

const CardSide = styled.div`
  display: flex;
  align-self: center;
  align-items: center;
  box-shadow: 0 0 15px black;
  height: 30vh;
  width: 30vh;
  border-radius: 15%;
`;

const CardText = styled.p`
  text-align: center;
  margin: auto;
`;

export interface CardWord {
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
        <CardText>{front.text}</CardText>
      </CardSide>

      <CardSide key="back" onClick={(e) => setAnim((a) => !a)}>
        <CardText>{back.text}</CardText>
      </CardSide>
    </ReactCardFlip>
  );
};

export default FlipCard;
