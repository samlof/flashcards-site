import { useApolloClient } from "@apollo/client";
import React from "react";
import styled from "styled-components";
import { FbApp } from "../lib/firebase";
import { useUser } from "../lib/user";
import Link from "next/link";
import { useRouter } from "next/router";
import NavLink from "./Navbar/NavLink";

const NavDiv = styled.nav`
  width: 100%;
  height: 3rem;
  box-shadow: 0px 1px 5px var(--color-brown);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 0 0 2rem;
  background-color: var(--color-green);
`;
const Right = styled.div`
  display: flex;
  align-items: center;
  height: 100%;
`;
const Middle = styled.div`
  display: flex;
  align-items: center;
  flex: 1;
  height: 100%;
`;
const Left = styled.div`
  display: flex;
  align-items: center;
  height: 100%;
`;
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

interface Props {}
const Navbar = ({}: Props) => {
  const router = useRouter();
  const user = useUser();
  const apolloClient = useApolloClient();

  const clickLogout = async () => {
    await FbApp.auth().signOut();
    await apolloClient.clearStore();
  };

  return (
    <NavDiv>
      <Right>
        <NavLink path="/" label="Cards" />
        <NavLink path="/all" label="All Cards" />
      </Right>
      <Middle></Middle>
      <Left>
        {!user.loading && user && (
          <>
            <NavLink path="/usersettings" label="Settings" />

            <NavItem onClick={clickLogout}>Logout</NavItem>
          </>
        )}
      </Left>
    </NavDiv>
  );
};

export default Navbar;
