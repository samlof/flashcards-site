export default function App({ children }) {
  return (
    <main>
      {children}
      <style jsx global>{`
        * {
          font-family: Menlo, Monaco, "Lucida Console", "Liberation Mono",
            "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Courier New",
            monospace, serif;
        }
        body {
          margin: 0;
          padding: 25px 50px;
        }
        .my-node-enter {
          opacity: 0;
        }
        .my-node-enter-active {
          opacity: 1;
          transition: opacity 200ms;
        }
        .my-node-exit {
          opacity: 1;
        }
        .my-node-exit-active {
          opacity: 0;
          transition: opacity 200ms;
        }
      `}</style>
    </main>
  );
}
