export default function App({ children }) {
  return (
    <main>
      {children}
      <style jsx global>{`
        * {
          font-family: Menlo, Monaco, "Lucida Console", "Liberation Mono",
            "DejaVu Sans Mono", "Bitstream Vera Sans Mono", "Courier New",
            monospace, serif;
          box-sizing: border-box;
        }
        body {
          margin: 0;
          padding: 25px 50px;
        }
        .fade-in-out-enter {
          opacity: 0;
        }
        .fade-in-out-enter-active {
          opacity: 1;
          transition: opacity 200ms;
        }
        .fade-in-out-exit {
          opacity: 1;
        }
        .fade-in-out-exit-active {
          opacity: 0;
          transition: opacity 200ms;
        }

        .center-div {
          text-align: center;
          justify-content: center;
          display: flex;
          align-items: center;
          flex-direction: column;
        }
      `}</style>
    </main>
  );
}
