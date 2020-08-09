import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import { FbApp, FbAuthUiConf } from "../lib/firebase";

interface Props {}
const Login = ({}: Props) => {
  return (
    <StyledFirebaseAuth uiConfig={FbAuthUiConf} firebaseAuth={FbApp.auth()} />
  );
};

export default Login;
