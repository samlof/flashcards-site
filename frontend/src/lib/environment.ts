function verifyHasValue(obj: any) {
  for (const key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key)) {
      const element = obj[key];
      if (typeof element === "object") verifyHasValue(element);
      else {
        if (!element) throw new Error(`environment.ts ${key} is empty`);
      }
    }
  }
}

export const environment = {
  firebase: {
    apiKey: process.env.NEXT_PUBLIC_API_KEY,
    authDomain: process.env.NEXT_PUBLIC_AUTH_DOMAIN,
    databaseURL: process.env.NEXT_PUBLIC_DATABASE_URL,
    projectId: process.env.NEXT_PUBLIC_PROJECT_ID,
    storageBucket: process.env.NEXT_PUBLIC_STORAGE_BUCKET,
    messagingSenderId: process.env.NEXT_PUBLIC_MESSAGING_SENDERID,
    appId: process.env.NEXT_PUBLIC_APP_ID,
  },
  graphqlUrl: process.env.NEXT_PUBLIC_GRAPHQL_URL,
  ssrDomain: process.env.NEXT_PUBLIC_SSR_DOMAIN,
};
verifyHasValue(environment);
