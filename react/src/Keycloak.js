import Keycloak from "keycloak-js";
import { ipAddress } from "./env";

const url0 = "http://" + ipAddress + ":8080/auth";

const keycloak = new Keycloak({
  url: url0,
  realm: "test_realm",
  clientId: "react-auth",
});

export default keycloak;
