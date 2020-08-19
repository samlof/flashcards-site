import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  NormalizedCache,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useMemo } from "react";
import { environment } from "./environment";
import { cookieValue } from "../helpers/cookies";
import { IdTokenCookie } from "../constants/cookieNames";
import { onError } from "@apollo/client/link/error";

let apolloClient: ApolloClient<NormalizedCache>;

function createApolloClient(staticIdToken?: string) {
  const serverUrl = environment.graphqlUrl;

  const httpLink = new HttpLink({
    uri: serverUrl, // Server URL (must be absolute)
  });
  const authLink = setContext((_, { headers }) => {
    // get the authentication token from local storage if it exists
    let idToken: string | undefined = staticIdToken;
    if (typeof window !== "undefined") {
      const cookieIdToken = cookieValue(IdTokenCookie);
      if (cookieIdToken) {
        idToken = cookieIdToken;
      }
    }
    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        authorization: idToken ? `Bearer ${idToken}` : "",
      },
    };
  });
  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (graphQLErrors)
      graphQLErrors.map(({ message, locations, path }) =>
        console.log(
          `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
        )
      );

    if (networkError) console.log(`[Network error]: ${networkError}`);
  });
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: errorLink.concat(authLink).concat(httpLink),
    cache: new InMemoryCache(),
  });
}

export function initializeApollo(
  initialState?: NormalizedCache,
  initialIdToken?: string
) {
  const _apolloClient = apolloClient ?? createApolloClient(initialIdToken);

  // If your page has Next.js data fetching methods that use Apollo Client, the initial state
  // gets hydrated here
  if (initialState) {
    _apolloClient.cache.restore(initialState);
  }
  // For SSG and SSR always create a new Apollo Client
  if (typeof window === "undefined") return _apolloClient;
  // Create the Apollo Client once in the client
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
}

export function useApollo(initialState?: NormalizedCache) {
  const store = useMemo(() => initializeApollo(initialState), [initialState]);
  return store;
}
