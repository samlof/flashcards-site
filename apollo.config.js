module.exports = {
  client: {
    excludes: ["**/*.generated.{ts,tsx}"],
    service: {
      name: "youlist",
      url: "http://localhost:8080",
    },
  },
};
