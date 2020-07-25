import React from "react";
import styled from "styled-components";
import AddCsvWords from "../components/admin/AddCsvWords";
import AddWord from "../components/admin/AddWord";
import AllWords from "../components/admin/AllWords";
import App from "../components/App";
import GqlError from "../components/GqlError";
import Loading from "../components/Loading";
import { useAllWordsQuery } from "../gql.generated";

const Title = styled.h1`
  text-align: center;
`;

const WordTable = styled.table`
  border-spacing: inherit;
`;

const WordRow = styled.tr`
  &:nth-child(odd) {
    background-color: #fff;
  }
  &:nth-child(even) {
    background-color: var(--color-white);
  }
`;

interface Props {}
const AdminPage = ({}: Props) => {
  const { data, loading, error, refetch: refetchWords } = useAllWordsQuery();

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

  if (!data) return <span>No words</span>;

  return (
    <App>
      <Title>Admin panel</Title>
      <div className="center-div">
        <AddWord refetchWords={refetchWords} />
        <AddCsvWords allWords={data} refetchWords={refetchWords} />
        <AllWords />
      </div>
    </App>
  );
};

export default AdminPage;
