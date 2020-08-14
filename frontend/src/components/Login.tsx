import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import { FbApp, FbAuthUiConf } from "../lib/firebase";
import { useUser } from "../lib/user";
import { Button } from "./Button";
import { useApolloClient } from "@apollo/client";

interface Props {}
const Login = ({}: Props) => {
  const user = useUser();
  const apolloClient = useApolloClient();

  const clickLogout = async () => {
    await FbApp.auth().signOut();
    await apolloClient.clearStore();
  };
  return (
    <>
      {!user && (
        <StyledFirebaseAuth
          uiConfig={FbAuthUiConf}
          firebaseAuth={FbApp.auth()}
        />
      )}
      {user && user !== "pending" && (
        <>
          <em>Logged in</em>
          <Button type="button" onClick={clickLogout}>
            Logout
          </Button>
        </>
      )}
    </>
  );
};

export default Login;
