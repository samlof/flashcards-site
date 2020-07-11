import { GetStaticProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import ReactCardFlip from "react-card-flip";

let id = 1000;

const Title = styled.h1`
  text-align: center;
`;

const CenteringDiv = styled.div`
  text-align: center;
  justify-content: center;
  display: flex;
`;

const CardSide = styled.div`
  display: flex;
  align-self: center;
  align-items: center;
  box-shadow: 0 0 15px black;
  height: 30vh;
  width: 30vh;
`;

const CardText = styled.p`
  text-align: center;
  margin: auto;
`;
interface Props {}

const IndexPage = ({}: Props) => {
  const [animate, setAnim] = React.useState(false);
  return (
    <App>
      <Title>Flashcards</Title>
      <CenteringDiv style={{ marginBottom: "20px" }}>
        <ReactCardFlip isFlipped={animate} flipDirection="horizontal">
          <CardSide key="front" onClick={(e) => setAnim((a) => !a)}>
            <CardText>Jäätelö</CardText>
          </CardSide>

          <CardSide key="back" onClick={(e) => setAnim((a) => !a)}>
            <CardText>Ice cream</CardText>
          </CardSide>
        </ReactCardFlip>
      </CenteringDiv>
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
