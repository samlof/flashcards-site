import React from "react";
import styled from "styled-components";
import AddCsvWords from "../components/admin/AddCsvWords";
import AddWord from "../components/admin/AddWord";
import AllWords from "../components/admin/AllWords";
import App from "../components/App";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import { useAllWordsQuery } from "../gql.generated";
import Head from "next/head";

interface Props {}
const AdminPage = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();

  const handleStatuses = () => {
    if (loading) return <Loading />;
    if (error) return <GqlError msg="Failed to get words" err={error} />;

    if (!data) return <span>No words</span>;
    return null;
  };
  return (
    <App>
      <Head>
        <title>Admin | kieli.club</title>
      </Head>
      <h1>Admin panel</h1>
      {handleStatuses() || (
        <>
          <AddWord refetchWords={refetchWords} />
          <AddCsvWords allWords={data!} refetchWords={refetchWords} />
          <AllWords />
        </>
      )}
    </App>
  );
};

export default AdminPage;
