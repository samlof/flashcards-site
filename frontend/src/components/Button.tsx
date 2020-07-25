import styled from "styled-components";

export const Button = styled.button`
  background-color: var(--color-green);
  border: 0;
  color: white;
  padding: 0.3rem 0.6rem;
  transition: background-color 0.3s;
  text-align: center;
  cursor: pointer;
  border-radius: 5%;

  &:active {
    box-shadow: 0 0 5px var(--color-blue);
    transform: scale(0.97);
  }

  &:hover {
    box-shadow: 0 0 2px var(--color-blue);
  }

  &:disabled {
    background-color: #b5bebf;
  }
  &:focus {
    outline: none;
  }
`;
