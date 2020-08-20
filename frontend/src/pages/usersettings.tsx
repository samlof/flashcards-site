import Head from "next/head";
import React from "react";
import App from "../components/App";
import Navbar from "../components/Navbar";
import UserSettings from "../components/UserSettings";
import { useUser } from "../lib/user";
import { useRouter } from "next/router";

interface Props {}
const UserSettingsPage = ({}: Props) => {
  const user = useUser();
  const router = useRouter();

  if (!user.loading && !user.user) {
    router.push("/all");
  }
  return (
    <App>
      <Head>
        <title>User Settings | kieli.club</title>
      </Head>
      <Navbar />

      <h1>User Settings</h1>
      <UserSettings />
    </App>
  );
};

export default UserSettingsPage;
