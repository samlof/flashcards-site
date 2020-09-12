import { useEffect, useState } from "react";
import { environment } from "./environment";

export const useAudio = (url: string): [boolean, () => void] => {
  const [audio] = useState(() => {
    const audio = new Audio(url);
    audio.autoplay = false;
    audio.controls = false;
    return audio;
  });
  const [playing, setPlaying] = useState(false);

  const start = () => {
    setPlaying(true);
    audio.play();
  };

  useEffect(() => {
    const setPlayFalse = () => setPlaying(false);
    audio.addEventListener("ended", setPlayFalse);

    document.body.appendChild(audio);
    return () => {
      audio.removeEventListener("ended", setPlayFalse);

      try {
        document.body.removeChild(audio);
      } catch (error) {
        //
      }
    };
  }, [audio]);

  return [playing, start];
};

export const makeAudioUrl = (text: string, lang: string): string => {
  const urlBase = environment.ttsUrl;

  return (
    urlBase + `/${encodeURIComponent(text)}-${encodeURIComponent(lang)}.mp3`
  );
};
