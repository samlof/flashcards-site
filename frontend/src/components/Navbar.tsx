import { useApolloClient } from "@apollo/client";
import { useRouter } from "next/router";
import React from "react";
import styled from "styled-components";
import { FbApp } from "../lib/firebase";
import { useUser } from "../lib/user";
import NavLink from "./Navbar/NavLink";
import dynamic from "next/dynamic";
import NavItem from "./Navbar/NavItem";
import { delayMs } from "../helpers/delay";

const NavLogin = dynamic(() => import("./Navbar/NavLogin"), { ssr: false });

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

interface Props {}
const Navbar = ({}: Props) => {
  const router = useRouter();
  const user = useUser();
  const apolloClient = useApolloClient();

  const loggedIn = !user.loading && !!user.user;
  const loggedOut = !user.loading && !user.user;

  const clickLogout = async () => {
    await FbApp.auth().signOut();
    await delayMs(200);
    await apolloClient.clearStore();
    await apolloClient.resetStore().catch((err) => {
      console.error("Error resetting store: ", err);
    });
  };

  return (
    <NavDiv>
      <Right>
        {loggedIn && <NavLink path="/" label="Cards" />}
        <NavLink path="/all" label="All Cards" />
      </Right>
      <Middle></Middle>
      <Left>
        {loggedOut && <NavLogin />}
        {loggedIn && (
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
