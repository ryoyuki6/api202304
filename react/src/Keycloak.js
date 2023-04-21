import Keycloak from "keycloak-js";

const keycloak = new Keycloak({
  url: "http://35.77.214.212:8080/auth",
  realm: "test_realm",
  clientId: "react-auth",
});

export default keycloak;
