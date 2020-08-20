import styled from "styled-components";

const NavItem = styled.a<{ selected?: boolean }>`
  font-size: 1.3rem;
  color: var(--color-white);
  height: 100%;
  padding: 0.5rem;
  cursor: pointer;
  transition: filter 0.3s;
  text-decoration: none;

  &:hover {
    filter: brightness(85%);
  }

  ${(props) => (props.selected ? "background-color: var(--color-blue)" : "")}
`;

export default NavItem;
