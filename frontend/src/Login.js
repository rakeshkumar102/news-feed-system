import React, { useState } from "react";
import loginImage from "./images/login.png";
import "./login.css";
import axios from "axios";
import dany from "./images/dany_cropped.jpg";
import handmaid from "./images/handmaid.jpg";
import harry from "./images/harry.png";
import hermione from "./images/hermione.jpg";
import indiana from "./images/indiana_jones.jpg";
import jack from "./images/jack_sparrow.jpeg";
import leia from "./images/leia.jpg";
import mia from "./images/mia_cropped.jpg";
import norton from "./images/norton.jpg";
import seb from "./images/seb.jpg";
import tooper from "./images/trooper.jpg";
import sharp from "./images/sharp3.jpg";
import Grid from "@mui/material/Grid";
import Item from "@mui/material/Grid";
import { Avatar } from "@mui/material";
import { ReactSession } from "react-client-session";

function Login({ instance }) {
  const [imageSelection, setImageSelection] = useState(false);

  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");

  const [selected, setSelected] = useState(-1);

  const imageList = [
    {
      image: harry,
      code: "harry",
    },
    {
      image: indiana,
      code: "indiana",
    },
    {
      image: norton,
      code: "norton",
    },
    {
      image: seb,
      code: "seb",
    },
    {
      image: tooper,
      code: "tooper",
    },
    {
      image: jack,
      code: "jack",
    },
    {
      image: hermione,
      code: "hermione",
    },
    {
      image: dany,
      code: "dany",
    },
    {
      image: sharp,
      code: "sharp",
    },
    {
      image: mia,
      code: "mia",
    },
    {
      image: leia,
      code: "leia",
    },
    {
      image: handmaid,
      code: "handmaid",
    },
  ];

  const handleSignUp = async (event) => {
    event.preventDefault();

    axios
      .post(`${process.env.REACT_APP_base_url}api/auth/signup`, {
        email: email,
        password: password,
        name: name,
      })
      .then((response) => {
        if (response.data.message === "success") {
          setImageSelection(true);
          console.log(response);
          ReactSession.set("email", response.data.data.user.email);
          ReactSession.set("name", response.data.data.user.name);
          ReactSession.set("token", response.data.data.token);
          ReactSession.set("user_id", response.data.data.user.ID);
        }
      })
      .catch(function (error) {
        console.log(error);
      });
  };

  const handleSignIn = async (event) => {
    event.preventDefault();

    axios
      .post(`${process.env.REACT_APP_base_url}api/auth/login`, {
        email: email,
        password: password,
      })
      .then((response) => {
        if (response.data.message === "success") {
          ReactSession.set("email", response.data.data.user.email);
          ReactSession.set("name", response.data.data.user.name);
          ReactSession.set("token", response.data.data.token);
          ReactSession.set("user_id", response.data.data.user.ID);

          if (response.data.data.user.image === "") {
            setImageSelection(true);
          } else {
            ReactSession.set("image", response.data.data.user.image);
            window.location.reload();
          }
        }
      })
      .catch((err) => {
        console.log(err);
      });
  };

  function handleSelectImage(index) {
    setSelected(index);
  }

  function handleUnselectImage(index) {
    if (selected === index) {
      setSelected(-1);
    }
  }

  function handleConfirm() {
    if (selected !== -1) {
      const config = {
        headers: { Authorization: `Bearer ${ReactSession.get("token")}` },
      };
      axios
        .put(
          `${
            process.env.REACT_APP_base_url
          }api/user/update?user_id=${ReactSession.get("user_id")}`,
          {
            image: imageList[selected].code,
          },
          config
        )
        .then((response) => {
          if (response.data.message === "success") {
            ReactSession.set("image", response.data.data.user.image);
            window.location.reload();
          } else {
            console.error(response.data.message);
          }
        });
    }
  }

  return (
    <div className="login">
      {imageSelection ? (
        <div className="image__selection">
          <div className="choose__header">
            <h2 className="choose__text">Choose Avatar</h2>
            <button className="choose__button" onClick={() => handleConfirm()}>
              Confirm
            </button>
          </div>
          <Grid
            className="grid__image"
            container
            align="center"
            justify="center"
            alignItems="center"
            rowSpacing={4}
          >
            {imageList.map((f, index) => (
              <Grid item md={4} key={index}>
                {index === selected ? (
                  <Item
                    className="image__back"
                    onClick={() => handleUnselectImage(index)}
                  >
                    <Avatar className="avatar__select" src={f.image} />
                  </Item>
                ) : (
                  <Item
                    className="unselected__image__back"
                    onClick={() => handleSelectImage(index)}
                  >
                    <Avatar className="avatar__select" src={f.image} />
                  </Item>
                )}
              </Grid>
            ))}
          </Grid>
        </div>
      ) : (
        <div className="login__body">
          <img src={loginImage} className="login__image" />
          <div>
            {instance ? (
              <form className="login__form" onSubmit={handleSignUp}>
                <input
                  className="login__input"
                  type="text"
                  name="email"
                  required
                  value={email}
                  placeholder="Email"
                  onChange={(e) => setEmail(e.target.value)}
                />
                <br />
                <input
                  className="login__input"
                  type="text"
                  name="name"
                  required
                  value={name}
                  placeholder="Name"
                  onChange={(e) => setName(e.target.value)}
                />
                <br />
                <input
                  className="login__input"
                  type="password"
                  name="password"
                  required
                  value={password}
                  placeholder="Password"
                  onChange={(e) => setPassword(e.target.value)}
                />
                <br />
                <button
                  className="login__button"
                  type="submit"
                  variant="outlined"
                >
                  Sign Up
                </button>
              </form>
            ) : (
              <form className="login__form" onSubmit={handleSignIn}>
                <input
                  className="login__input"
                  type="text"
                  name="email"
                  required
                  value={email}
                  placeholder="Email"
                  onChange={(e) => setEmail(e.target.value)}
                />
                <br />
                <input
                  className="login__input"
                  type="password"
                  name="password"
                  required
                  value={password}
                  placeholder="Password"
                  onChange={(e) => setPassword(e.target.value)}
                />
                <br />
                <button
                  type="submit"
                  className="login__button"
                  variant="outlined"
                >
                  Sign In
                </button>
              </form>
            )}
          </div>
        </div>
      )}
    </div>
  );
}

export default Login;
