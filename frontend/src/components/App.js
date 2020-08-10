import Head from "next/head";

export default function App({ children }) {
  return (
    <main>
      {children}
      <Head>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
      <style jsx global>{`
        * {
          box-sizing: border-box;
        }
        h1,
        h2,
        h3,
        h4,
        h5,
        h6 {
          font-family: -apple-system, BlinkMacSystemFont, "Lato", Menlo, Monaco,
            "Lucida Console", "Liberation Mono", "DejaVu Sans Mono",
            "Bitstream Vera Sans Mono", "Courier New", monospace, serif;
          color: var(--color-blue);
        }
        main {
          text-align: center;
          justify-content: center;
          display: flex;
          align-items: center;
          flex-direction: column;
        }
        body {
          font-family: -apple-system, BlinkMacSystemFont, "Raleway", Menlo,
            Monaco, "Lucida Console", "Liberation Mono", "DejaVu Sans Mono",
            "Bitstream Vera Sans Mono", "Courier New", monospace, serif;
          margin: 0;
          padding: 0;
          --color-green: #4c9f70;
          --color-blue: #0b4f6c;
          --color-orange: #fcd7ad;
          --color-brown: #6c4b5e;
          --color-red: #d55672;
          --color-white: #efefef;
          background-color: var(--color-orange);
          font-size: 16px;
          letter-spacing: -0.003em;
          line-height: 1.58;
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
