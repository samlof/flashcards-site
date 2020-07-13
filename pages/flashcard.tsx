import { GetStaticProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import FlipCard, { CardWord } from "../components/FlipCard";

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
  const frontCard: CardWord = {
    text: "Jäätelö",
  };
  const backCard: CardWord = {
    text: "Ice cream",
  };
  return (
    <App>
      <Title>Flashcards</Title>
      <CenteringDiv style={{ marginBottom: "20px" }}>
        <FlipCard back={backCard} front={frontCard}></FlipCard>
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
