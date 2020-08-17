import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import { FbApp, FbAuthUiConf } from "../lib/firebase";
import { useUser } from "../lib/user";
import { Button } from "./Button";
import { useApollo } from "../lib/apolloClient";

interface Props {}
const Login = ({}: Props) => {
  const user = useUser();
  const apolloClient = useApollo();

  const clickLogout = async () => {
    await FbApp.auth().signOut();
    await apolloClient.clearStore();
  };

  if (user.loading) return null;

  return (
    <>
      {!user.user ? (
        <StyledFirebaseAuth
          uiConfig={FbAuthUiConf}
          firebaseAuth={FbApp.auth()}
        />
      ) : (
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
