// Firebase App (the core Firebase SDK) is always required and must be listed first
import firebase from "firebase/app";
import "firebase/auth";

export const firebaseConfig = {
  apiKey: "AIzaSyBk4Aia7cUwOYY9G_YB9E0Shdzbb2E24UY",
  authDomain: "kieliclub.firebaseapp.com",
  databaseURL: "https://kieliclub.firebaseio.com",
  projectId: "kieliclub",
  storageBucket: "kieliclub.appspot.com",
  messagingSenderId: "676509926663",
  appId: "1:676509926663:web:be676fcc9f702617cc31a2",
};

let app: firebase.app.App;
if (!firebase.apps.length) {
  app = firebase.initializeApp(firebaseConfig);
} else {
  app = firebase.apps[0];
}
export const FbApp = app;
export const FbAuthUiConf = {
  signInOptions: [
    firebase.auth.GoogleAuthProvider.PROVIDER_ID,
    firebase.auth.FacebookAuthProvider.PROVIDER_ID,
    {
      provider: firebase.auth.EmailAuthProvider.PROVIDER_ID,
      signInMethod: firebase.auth.EmailAuthProvider.EMAIL_LINK_SIGN_IN_METHOD,
    },
  ],
  signInFlow: "popup",
  signInSuccessUrl: "/signedIn",
};
