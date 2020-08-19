import React from "react";
import styled from "styled-components";
import {
  useUserSettingsQuery,
  useSetUserSettingsMutation,
} from "../gql.generated";
import GqlError from "./GqlError";
import Loading from "./Loading";
import { Button } from "./Button";

const SmallInput = styled.input`
  width: 4rem;
  text-align: center;
`;
const NumberSpan = styled.span`
  display: inline-block;
  width: 4rem;
  text-align: center;
`;
const SaveButton = styled(Button)`
  margin-right: 0.5rem;
`;
const Spacer = styled.span`
  margin-left: 5.3rem;
`;

interface Props {}
const UserSettings = ({}: Props) => {
  const {
    data,
    loading,
    error,
    refetch: refetchsettings,
  } = useUserSettingsQuery();
  const [setSettings] = useSetUserSettingsMutation();

  const [editing, setEditing] = React.useState(false);
  const [cardsPerDay, setCardsPerDay] = React.useState(10);

  React.useEffect(() => {
    if (data?.userSettings?.newCardsPerDay)
      setCardsPerDay(data.userSettings.newCardsPerDay);
    else setCardsPerDay(10);
  }, [data?.userSettings?.newCardsPerDay]);

  if (loading) return <Loading />;
  if (error) return <GqlError msg="Failed to get words" err={error} />;

  const HandleSave = () => {
    setSettings({ variables: { cardsPerDay: cardsPerDay } }).then(() => {
      setEditing(false);
      return refetchsettings();
    });
  };
  return (
    <>
      <div>
        New words per day:{" "}
        {editing ? (
          <SmallInput
            min="0"
            max="200"
            type="number"
            value={cardsPerDay}
            onChange={(e) => setCardsPerDay(e.target.valueAsNumber)}
          />
        ) : (
          <NumberSpan>{cardsPerDay}</NumberSpan>
        )}{" "}
        {editing ? (
          <>
            <SaveButton onClick={() => HandleSave()}>Save</SaveButton>
            <Button onClick={() => setEditing((p) => !p)}>
              {editing ? "Cancel" : "Edit"}
            </Button>
          </>
        ) : (
          <>
            <Button onClick={() => setEditing((p) => !p)}>
              {editing ? "Cancel" : "Edit"}
            </Button>
            <Spacer />
          </>
        )}
      </div>
    </>
  );
};

export default UserSettings;
