import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import { FbApp, FbAuthUiConf } from "../lib/firebase";
import { useUser } from "../lib/user";

interface Props {}
const Login = ({}: Props) => {
  const user = useUser();

  return (
    <>
      {!user && (
        <StyledFirebaseAuth
          uiConfig={FbAuthUiConf}
          firebaseAuth={FbApp.auth()}
        />
      )}
      {user && user !== "pending" && <p>{JSON.stringify(user, null, 2)}</p>}
    </>
  );
};

export default Login;
