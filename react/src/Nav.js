import React, { useState } from "react";
import axios from "axios";
import { useKeycloak } from "@react-keycloak/web";
import { ipAddress } from "./env";

const Nav = () => {
  const { keycloak, initialized } = useKeycloak();
  const [getData, setGetData] = useState([]);

  const fetchData = async () => {
    const url0 = "http://" + ipAddress + ":8081/health_data";
    const response = await axios.get(url0);
    setGetData(response.data);
  };

  return (
    <div>
      <h1>Keycloak authentication</h1>
      <hr />
      <p>Login や Logout のボタンを押してみてください</p>
      {!keycloak.authenticated ? (
        <button type="button" onClick={() => keycloak.login()}>
          Login
        </button>
      ) : (
        <button type="button" onClick={() => keycloak.logout()}>
          Logout ({keycloak.tokenParsed.preferred_username})
        </button>
      )}

      <div>{`User is ${
        !initialized || !keycloak.authenticated ? "NOT " : ""
      }authenticated`}</div>

      <hr />

      {!keycloak.authenticated ? (
        <p>ログインするとここに「GETボタン」が現れます</p>
      ) : (
        <>
        <button onClick={fetchData}>GET</button>
        <ul>
          {getData.map((data) => (
            <li key={data.day}>
              id : {data.id}, month : {data.month}, day : {data.day}, weight :{" "}
              {data.weight}
            </li>
          ))}
        </ul>
        </>
      )}

    </div>
  );
};

export default Nav;
