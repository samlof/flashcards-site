// Firebase App (the core Firebase SDK) is always required and must be listed first
import { User } from "firebase/app";
import { useEffect, useState } from "react";
import { FbApp } from "./firebase";

export const useUser: () => User | null | "pending" = () => {
  const [user, setUser] = useState<User | null | "pending">("pending");
  useEffect(() => {
    const unsub = FbApp.auth().onAuthStateChanged((user) => {
      setUser(user);
    });
    return unsub;
  }, []);
  return user;
};
