import { useState, useEffect } from "react";
import { environment } from "./environment";

export const useAudio = (url: string): [boolean, () => void, () => void] => {
  const [audio] = useState(() => new Audio(url));
  const [playing, setPlaying] = useState(false);

  const toggle = () => setPlaying(!playing);
  const start = () => setPlaying(true);

  useEffect(() => {
    playing ? audio.play() : audio.pause();
  }, [playing, audio]);

  useEffect(() => {
    audio.addEventListener("ended", () => setPlaying(false));
    return () => {
      audio.removeEventListener("ended", () => setPlaying(false));
    };
  }, [audio]);

  return [playing, start, toggle];
};

export const makeAudioUrl = (text: string, lang: string): string => {
  const urlBase = environment.ttsUrl;

  return (
    urlBase + `/${encodeURIComponent(text)}-${encodeURIComponent(lang)}.mp3`
  );
};
