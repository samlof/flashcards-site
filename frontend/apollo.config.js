module.exports = {
  client: {
    includes: ["./src/**/*.{ts,tsx,gql}"],
    excludes: ["./src**/*.generated.{ts,tsx}"],
    service: {
      name: "Words",
      url: "http://localhost:8080/query",
      skipSSLValidation: true,
    },
  },
};
