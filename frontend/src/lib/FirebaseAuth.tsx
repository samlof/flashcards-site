/**
 * Copyright 2017 Google Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// @flow

import React from "react";
import * as firebaseui from "firebaseui";

// Global ID for the element.
const ELEMENT_ID = "firebaseui_container";

// Promise that resolves unless the FirebaseUI instance is currently being deleted.
let firebaseUiDeletion = Promise.resolve();

interface props {
  // The Firebase UI Web UI Config object.
  // See: https://github.com/firebase/firebaseui-web#configuration
  uiConfig: any;
  // The Firebase App auth instance to use.
  firebaseAuth: firebase.default.auth.Auth;
  // Callback that will be passed the FirebaseUi instance before it is
  // started. This allows access to certain configuration options such as
  // disableAutoSignIn().
  uiCallback?: (firebaseUiWidget: any) => void;
  className?: string;
}
/**
 * React Component wrapper for the FirebaseUI Auth widget.
 */
export default class FirebaseAuth extends React.Component<props> {
  /**
   * Constructor  Firebase Auth UI component
   *
   * @constructor
   */
  constructor(props: props) {
    super(props);
  }

  unregisterAuthObserver?: () => void;
  userSignedIn = false;
  firebaseUiWidget: any;

  /**
   * @inheritDoc
   */
  componentDidMount() {
    // Import the css only on the client.
    require("firebaseui/dist/firebaseui.css");

    const { firebaseAuth, uiConfig, uiCallback } = this.props;
    // Wait in case the firebase UI instance is being deleted.
    // This can happen if you unmount/remount the element quickly.
    return firebaseUiDeletion.then(() => {
      // Get or Create a firebaseUI instance.
      this.firebaseUiWidget =
        firebaseui.auth.AuthUI.getInstance() ||
        new firebaseui.auth.AuthUI(firebaseAuth);
      if (uiConfig.signInFlow === "popup") {
        this.firebaseUiWidget.reset();
      }

      // We track the auth state to reset firebaseUi if the user signs out.
      this.userSignedIn = false;
      this.unregisterAuthObserver = firebaseAuth.onAuthStateChanged((user) => {
        if (!user && this.userSignedIn) {
          this.firebaseUiWidget.reset();
        }
        this.userSignedIn = !!user;
      });

      // Trigger the callback if any was set.
      if (uiCallback) {
        uiCallback(this.firebaseUiWidget);
      }

      // Render the firebaseUi Widget.
      this.firebaseUiWidget.start("#" + ELEMENT_ID, uiConfig);
    });
  }

  /**
   * @inheritDoc
   */
  componentWillUnmount() {
    firebaseUiDeletion = firebaseUiDeletion.then(() => {
      this.unregisterAuthObserver && this.unregisterAuthObserver();
      return this.firebaseUiWidget.delete();
    });
    return firebaseUiDeletion;
  }

  /**
   * @inheritDoc
   */
  render() {
    return <div className={this.props.className} id={ELEMENT_ID} />;
  }
}
