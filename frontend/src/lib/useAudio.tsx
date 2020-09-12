import { useEffect, useState } from "react";
import { environment } from "./environment";

export const useAudio = (url: string): [boolean, () => void, () => void] => {
  const [audio] = useState(() => {
    const audio = new Audio(url);
    audio.autoplay = false;
    audio.controls = false;
    return audio;
  });
  const [playing, setPlaying] = useState(false);

  const toggle = () => setPlaying(!playing);
  const start = () => setPlaying(true);

  useEffect(() => {
    // If audio hasn't been loaded yet, then load it before playing
    if (playing && audio.networkState === 0) {
      audio.load();
    }
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
