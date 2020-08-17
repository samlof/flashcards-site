// Firebase App (the core Firebase SDK) is always required and must be listed first
import { User } from "firebase/app";
import { useEffect, useState } from "react";
import { FbApp } from "./firebase";

interface userValue {
  loading: boolean;
  user: User | null;
}
export const useUser: () => userValue = () => {
  const [user, setUser] = useState<userValue>({ loading: true, user: null });
  useEffect(() => {
    const unsub = FbApp.auth().onIdTokenChanged((user) => {
      setUser({ user, loading: false });
    });
    return unsub;
  }, []);
  return user;
};
