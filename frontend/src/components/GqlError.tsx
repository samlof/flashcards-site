import React from "react";
import { ApolloError } from "@apollo/client";

interface Props {
  err: ApolloError;
  msg: string;
}
const GqlError = ({ err, msg }: Props) => {
  console.error("Got error", err);
  return (
    <span>
      {msg}: {err.message}
    </span>
  );
};

export default GqlError;
