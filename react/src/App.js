import React from "react";
import Nav from "./Nav";
import { ReactKeycloakProvider } from "@react-keycloak/web";
import keycloak from "./Keycloak";

function App() {
  return (
    <div>
      <ReactKeycloakProvider authClient={keycloak}>
        <Nav />
      </ReactKeycloakProvider>
    </div>
  );
};

export default App;
