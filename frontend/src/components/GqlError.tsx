import React from "react";
import { ApolloError, ServerError } from "@apollo/client";

const parseError = (err: ApolloError): string => {
  if ((err.networkError as ServerError)?.result) {
    const errors: any[] = (err.networkError as ServerError)?.result?.errors;
    if (Array.isArray(errors)) {
      return errors
        .map((x) => x.message)
        .filter((x) => !!x)
        .join(",");
    }
  }
  return err.message;
};
interface Props {
  err: ApolloError;
  msg: string;
}
const GqlError = ({ err, msg }: Props) => {
  const parsedError = parseError(err);
  console.dir(err);
  console.error("Got error", parsedError);

  return (
    <span>
      {msg}: {parsedError}
    </span>
  );
};

export default GqlError;
