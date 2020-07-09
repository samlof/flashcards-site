import { GetStaticProps } from "next";
import React from "react";
import styled from "styled-components";
import App from "../components/App";
import { numberToString } from "../helpers/numberToString";
import { QuestionLine, Question } from "../components/QuestionLine";
import { CSSTransition } from "react-transition-group";

let id = 1000;

const generateQuestions = (amount = 10): Question[] => {
  const ret: Question[] = [];
  const hundreds = amount / 10;
  const ones = amount / 4;
  for (let i = 0; i < hundreds; i++) {
    const num = Math.floor(Math.random() * 899 + 100);
    const answer = numberToString(num);
    ret.push({ num, answer, id: id++ });
  }
  for (let i = 0; i < ones; i++) {
    const num = Math.floor(Math.random() * 9 + 1);
    const answer = numberToString(num);
    ret.push({ num, answer, id: id++ });
  }
  for (let i = 0; i < amount - hundreds - ones; i++) {
    const num = Math.floor(Math.random() * 89 + 10);
    const answer = numberToString(num);
    ret.push({ num, answer, id: id++ });
  }

  return ret;
};

const Title = styled.h1`
  text-align: center;
`;
const NewQuestionsButton = styled.button`
  background-color: #22bad9;
  border: 0;
  color: white;
  padding: 0.3rem 0.6rem;
  transition: background-color 0.3s;
  font-size: 1.5rem;
  text-align:center;
  width:20rem
  cursor: pointer;

  &:active {
    background-color: #1b9db7;
  }

  &:disabled {
    background-color: #b5bebf;
  }
  &:focus {
    outline: none;
  }
`;
const Container = styled.div`
  display: flex;
  justify-content: center;
  margin-bottom: 1rem;
  align-items: center;
  height: 2rem;
`;

const ResultArea = styled.div`
  margin-top: 3rem;
  text-align: center;
`;
interface Props {
  initialQuestions: Question[];
}

const IndexPage = ({ initialQuestions }: Props) => {
  const [questions, setQuestions] = React.useState<Question[]>(
    initialQuestions || []
  );
  const [showResults, setShowResults] = React.useState<{
    [key: number]: boolean;
  }>({});
  const [showTotalResults, setShowTotalResults] = React.useState(false);
  const [loadingResults, setLoadingResults] = React.useState(false);
  const [rightAnswerCount, setRightAnswerCount] = React.useState(0);

  const generateNewQuestions = () => {
    setQuestions(generateQuestions());
    setRightAnswerCount(0);
    setShowResults({});
    setShowTotalResults(false);
    setLoadingResults(false);
  };

  const answersChanged = (res: boolean) => {
    if (res) setRightAnswerCount((prev) => prev + 1);
    else setRightAnswerCount((prev) => prev - 1);
  };

  const submitResults = () => {
    setLoadingResults(true);
    for (let i = 0; i < questions.length; i++) {
      const q = questions[i];
      const waitTime = i * 500;
      setTimeout(() => {
        setShowResults((prev) => ({ ...prev, [q.id]: true }));
      }, waitTime);
    }
    setTimeout(() => {
      setShowTotalResults(true);
      setLoadingResults(false);
    }, questions.length * 500);
  };

  return (
    <App>
      <Title>Write the number in Finnish and press submit</Title>
      {questions.map((question) => (
        <Container key={question.id}>
          <QuestionLine
            question={question}
            showResults={showResults[question.id]}
            resultChanged={answersChanged}
          ></QuestionLine>
        </Container>
      ))}
      <ResultArea>
        <div style={{ display: "inline-block" }}>
          <div style={{ height: "1rem" }}>
            <CSSTransition
              unmountOnExit
              in={showTotalResults}
              timeout={400}
              classNames="my-node"
            >
              <span>
                Right answers: {rightAnswerCount} out of {questions.length}
              </span>
            </CSSTransition>
          </div>
          {!showTotalResults ? (
            <NewQuestionsButton
              type="button"
              disabled={loadingResults}
              onClick={(e) => submitResults()}
            >
              Show results
            </NewQuestionsButton>
          ) : (
            <NewQuestionsButton
              type="button"
              onClick={(e) => generateNewQuestions()}
            >
              New Questions
            </NewQuestionsButton>
          )}
        </div>
      </ResultArea>
    </App>
  );
};

export const getStaticProps: GetStaticProps<Props> = async (context) => {
  id = 1;
  const questions = generateQuestions(10);
  id = 1000;
  return {
    props: { initialQuestions: questions },
    unstable_revalidate: 1,
  };
};

export default IndexPage;
