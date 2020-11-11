// Firebase App (the core Firebase SDK) is always required and must be listed first
import { useEffect, useState } from "react";

const FbAppPromise = import("./firebase").then((x) => x.FbApp);
interface userValue {
  loading: boolean;
  user: firebase.default.User | null;
}
export const useUser: () => userValue = () => {
  const [user, setUser] = useState<userValue>({ loading: true, user: null });
  useEffect(() => {
    let unsub: firebase.default.Unsubscribe;
    FbAppPromise.then((FbApp) => {
      unsub = FbApp.auth().onIdTokenChanged((user) => {
        setUser({ user, loading: false });
      });
    });
    return () => unsub && unsub();
  }, []);
  return user;
};
