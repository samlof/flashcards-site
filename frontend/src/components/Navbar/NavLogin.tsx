import React from "react";
import StyledFirebaseAuth from "react-firebaseui/StyledFirebaseAuth";
import ReactModal from "react-modal";
import { FbApp, FbAuthUiConf } from "../../lib/firebase";
import { useUser } from "../../lib/user";
import NavItem from "./NavItem";
import { useApollo } from "../../lib/apolloClient";
import { delayMs } from "../../helpers/delay";

// Make sure to bind modal to your appElement (http://reactcommunity.org/react-modal/accessibility/)
ReactModal.setAppElement("#__next");

const customStyles = {
  content: {
    top: "10%",
    left: "30%",
    right: "auto",
    bottom: "auto",
    marginRight: "-50%",
    backgroundColor: "var(--color-white)",
  },
};

interface Props {}
const NavLogin = ({}: Props) => {
  const user = useUser();
  const [modalIsOpen, setIsOpen] = React.useState(false);
  const apollo = useApollo();

  function openModal() {
    setIsOpen(true);
  }
  function closeModal() {
    setIsOpen(false);
  }

  React.useEffect(() => {
    if (!user.loading && user.user) {
      setIsOpen(false);
    }
  }, [user]);

  if (user.loading) return null;

  FbAuthUiConf.callbacks.signInSuccessWithAuthResult = () => {
    console.log("login callback");

    closeModal();

    delayMs(100)
      .then(() => apollo.clearStore())
      .then(() => apollo.resetStore());
    return false;
  };

  return (
    <>
      <NavItem onClick={openModal}>Login</NavItem>
      <ReactModal
        isOpen={modalIsOpen}
        style={customStyles}
        onRequestClose={closeModal}
      >
        <StyledFirebaseAuth
          uiConfig={FbAuthUiConf}
          firebaseAuth={FbApp.auth()}
        />
      </ReactModal>
    </>
  );
};

export default NavLogin;
