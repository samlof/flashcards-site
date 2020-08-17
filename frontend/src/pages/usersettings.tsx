import Head from "next/head";
import React from "react";
import App from "../components/App";
import Navbar from "../components/Navbar";
import UserSettings from "../components/UserSettings";
import Login from "../components/Login";

interface Props {}
const AdminPage = ({}: Props) => {
  return (
    <App>
      <Head>
        <title>User Settings | kieli.club</title>
      </Head>
      <Navbar />

      <h3>User Settings</h3>
      <Login />
      <UserSettings />
    </App>
  );
};

export default AdminPage;
