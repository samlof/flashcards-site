import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
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

interface Props {
  path: string;
  label: string;
}
const NavLink = ({ path, label }: Props) => {
  const router = useRouter();

  return (
    <Link href={path} passHref>
      <NavItem selected={router.pathname === path}>{label}</NavItem>
    </Link>
  );
};

export default NavLink;
