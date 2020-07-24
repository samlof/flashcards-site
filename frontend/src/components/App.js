import HeaderCss from "./HeaderCss";

export default function App({ children }) {
  return (
    <main>
      <HeaderCss />
      {children}
      <style jsx global>{`
        * {
          font-family: "Raleway", Menlo, Monaco, "Lucida Console",
            "Liberation Mono", "DejaVu Sans Mono", "Bitstream Vera Sans Mono",
            "Courier New", monospace, serif;
          box-sizing: border-box;
        }
        h1,
        h2,
        h3,
        h4,
        h5 {
          font-family: "Lato", Menlo, Monaco, "Lucida Console",
            "Liberation Mono", "DejaVu Sans Mono", "Bitstream Vera Sans Mono",
            "Courier New", monospace, serif;
        }
        body {
          margin: 0;
          padding: 25px 50px;
          --color-green: #4c9f70;
          --color-blue: #0b4f6c;
          --color-orange: #fcd7ad;
          --color-brown: #6c4b5e;
          --color-red: #d55672;
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
