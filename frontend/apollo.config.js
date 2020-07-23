module.exports = {
  client: {
    includes: ["./pages/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}"],
    excludes: ["**/*.generated.{ts,tsx}"],
    service: {
      name: "Words",
      url: "http://localhost:8080/query",
      skipSSLValidation: true,
    },
  },
};
