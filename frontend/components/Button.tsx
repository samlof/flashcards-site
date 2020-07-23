import styled from "styled-components";

export const Button = styled.button`
  background-color: #22bad9;
  border: 0;
  color: white;
  padding: 0.3rem 0.6rem;
  transition: background-color 0.3s;
  font-size: 1.5rem;
  text-align: center;
  cursor: pointer;
  border-radius: 5%;

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
