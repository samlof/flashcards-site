import React from "react";
import { CSSTransition } from "react-transition-group";
import styled from "styled-components";

const InputWrapper = styled.span`
  width: 14rem;
  text-align: right;
`;
const Input = styled.input`
  padding: 0.3rem;
  text-align: right;
`;
const QuestionSpan = styled.span`
  padding-right: 1rem;
  width: 2rem;
`;
const ResultSpan = styled.span`
  padding-left: 1rem;
  width: 2rem;
`;

export interface Question {
  num: number;
  answer: string;
  id: number;
}

interface Props {
  question: Question;
  showResults: boolean;
  resultChanged: (result: boolean) => void;
  initialAnswer?: string;
}

export const QuestionLine = ({
  question,
  showResults,
  resultChanged,
  initialAnswer,
}: Props) => {
  const [answer, setAnswer] = React.useState(initialAnswer ?? "");
  const [showAnswer, setShowAnswer] = React.useState(false);
  const [showLoading, setShowLoading] = React.useState(true);

  const lowercaseAnswer = answer.toLowerCase();
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    const lowerValue = value.toLowerCase();
    if (answer !== question.answer && lowerValue === question.answer) {
      resultChanged(true);
    } else if (answer === question.answer && lowerValue !== question.answer) {
      resultChanged(false);
    }
    setAnswer(value);
  };
  return (
    <>
      <QuestionSpan>{question.num}</QuestionSpan>
      <InputWrapper>
        {showAnswer ? (
          answer
        ) : (
          <Input type="text" value={answer ?? ""} onChange={handleChange} />
        )}
      </InputWrapper>
      <ResultSpan>
        <CSSTransition
          unmountOnExit
          in={showLoading && showResults}
          timeout={400}
          classNames="fade-in-out"
          onEntered={() => setShowAnswer(true)}
        >
          <span>{!showAnswer && "Calculating..."}</span>
        </CSSTransition>
        <CSSTransition
          unmountOnExit
          in={showAnswer}
          timeout={200}
          onEnter={() => setShowLoading(false)}
          classNames="fade-in-out"
        >
          <span>{question.answer === lowercaseAnswer ? "OK!" : "Not ok"}</span>
        </CSSTransition>
      </ResultSpan>
    </>
  );
};
