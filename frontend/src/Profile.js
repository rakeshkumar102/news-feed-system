import React, { useEffect, useState } from "react";
import "./profile.css";
// import handmaid from "./images/handmaid.jpg";
import feedIcon from "./images/news.png";
import followingIcon from "./images/following.png";
import followersIcon from "./images/followers.png";
import sharp from "./images/sharp3.jpg";
import hermione from "./images/hermione.jpg";
import harry from "./images/harry.png";
import seb from "./images/seb.jpg";
import norton from "./images/norton.jpg";
import trooper from "./images/trooper.jpg";
import jack from "./images/jack_sparrow.jpeg";
import upload from "./images/upload.png";
import Grid from "@mui/material/Grid";
import Item from "@mui/material/Grid";
import { Avatar } from "@mui/material";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import leia from "./images/leia.jpg";
import mia from "./images/mia_cropped.jpg";
import tooper from "./images/trooper.jpg";
import handmaid from "./images/handmaid.jpg";
import dany from "./images/dany_cropped.jpg";
import indiana from "./images/indiana_jones.jpg";

// import Typography from "@mui/material/Typography";
import Modal from "@mui/material/Modal";
import axios from "axios";
import { ReactSession } from "react-client-session";
const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",

  boxShadow: 24,
  p: 2,
  border: 0,
  outline: 0,
};

const map = new Map();
map.set("dany", dany);
map.set("handmaid", handmaid);
map.set("harry", harry);
map.set("hermione", hermione);
map.set("indiana", indiana);
map.set("jack", jack);
map.set("leia", leia);
map.set("mia", mia);
map.set("norton", norton);
map.set("seb", seb);
map.set("tooper", tooper);
map.set("sharp", sharp);

function Profile() {
  const [feedItem, setFeedItem] = useState([]);
  const [followingItem, setFollowingItem] = useState([]);
  const [followerItem, setFollowerItem] = useState([]);
  const [profileHome, setProfileHome] = useState(0);
  const [profileUser, setProfileUser] = useState({});
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);

  const followingList = [];

  const followerList = [];

  const config = {
    headers: {
      Authorization: `Bearer ${ReactSession.get("token")}`,
    },
  };

  useEffect(() => {
    axios
      .get(`${process.env.REACT_APP_base_url}api/feed/user`, config)
      .then((res) => {
        setFeedItem(res.data.data.feed.map((f) => f));
      });

    axios
      .get(
        `${process.env.REACT_APP_base_url}api/user?user_id=${ReactSession.get(
          "user_id"
        )}`,
        config
      )
      .then((res) => {
        console.log(res.data.data.user);
        setProfileUser(res.data.data.user);
      });
    setFollowingItem(followingList.map((f) => f));
    setFollowerItem(followerList.map((f) => f));
  }, []);

  const setFollowingPage = (event) => {
    event.preventDefault();
    setProfileHome(1);
  };

  const setFollowerPage = (event) => {
    event.preventDefault();
    setProfileHome(2);
  };

  const setFeedPage = (event) => {
    event.preventDefault();
    setProfileHome(0);
  };

  function handleFollowerClick(index) {}

  const [imageUrl, setImageUrl] = useState(null);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const handleChange = (e) => {
    const image = e.target.files[0];

    const formdata = new FormData();
    formdata.append("image_file", image);
    const config = {
      headers: {
        Authorization: `Bearer ${ReactSession.get("token")}`,
        "Content-Type": "multipart/form-data",
      },
    };

    axios
      .post(
        `${process.env.REACT_APP_base_url}api/feed/upload/image`,
        formdata,
        config
      )
      .then((res) => {
        console.log(res.data);
        setImageUrl(res.data.data.image);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const handleUpload = () => {
    axios({
      method: "post",
      url: `${process.env.REACT_APP_base_url}api/feed/create`,
      data: {
        // image: imageUrl,
        title: title,
        description: description,
      },
      headers: {
        Authorization: `Bearer ${ReactSession.get("token")}`,
      },
    })
      .then((res) => {
        console.log(res);
      })
      .catch((e) => {
        console.log(e);
      });
  };
  return (
    <div className="profile">
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <div className="upload">
            {/* <input
              type="file"
              accept="image/jpeg, image/png, image/jpg"
              onChange={handleChange}
            /> */}
            <input
              type="text"
              placeholder="Title"
              onChange={(e) => setTitle(e.target.value)}
            />
            <textarea
              placeholder="Description"
              onChange={(e) => setDescription(e.target.value)}
            ></textarea>
            <Button variant="primary" onClick={handleUpload}>
              Upload
            </Button>
          </div>
        </Box>
      </Modal>
      <div className="profile__details">
        <img
          className="profile__avatar"
          src={map.get(profileUser.image)}
          alt="avatar"
        />
        <div>
          <div className="name__and__follow__button">
            <h2 className="profile__name">{profileUser.name}</h2>
            {/* <button className="profile__follow__button">Follow</button> */}
          </div>
          <p className="profile__description">{"desc"}</p>
        </div>
      </div>
      <div className="nav__bar">
        <img
          className="feed__image"
          src={feedIcon}
          onClick={setFeedPage}
          alt="feedimage"
        />
        <p className="feed__count" onClick={setFeedPage}>
          <strong>{feedItem.length} News</strong>
          <br />
          <text className="nav__text">Total News Count</text>
        </p>
        {/* <img
          className="following__image"
          src={followingIcon}
          onClick={setFollowingPage}
          alt="follower"
        />
        <p className="feed__count" onClick={setFollowingPage}>
          <strong>{12} Following</strong>
          <br />
          <text className="nav__text">Total Following Count</text>
        </p>
        <img
          className="following__image"
          src={followersIcon}
          onClick={setFollowerPage}
          alt="following"
        />
        <p className="feed__count" onClick={setFollowerPage}>
          <strong>{11} Followers</strong>
          <br />
          <text className="nav__text">Total Followers Count</text>
        </p> */}
        <div className="upload__image" onClick={handleOpen}>
          <img
            className="upload__img"
            src={upload}
            onClick={setFollowerPage}
            alt="upload"
          />
          <p className="upload__text" onClick={setFollowerPage}>
            <strong>Upload</strong>
            <br />
            <text className="nav__text">Click to upload news feed</text>
          </p>
        </div>
      </div>
      <div className="profile__body">
        {profileHome === 0 ? (
          <Grid
            container
            spacing={{ xs: 2, md: 3 }}
            columns={{ xs: 4, sm: 8, md: 12 }}
          >
            {feedItem.map((f, index) => (
              <Grid item xs={2} sm={4} md={4} key={index}>
                <Item className="item__item">
                  <div className="item__image">
                    {/* <img className="feed__item" src={f.image} alt="feed item" /> */}
                    <div className="item__details">
                      <text>
                        <strong>{f.title}</strong>
                      </text>
                      <br />
                      <text className="profile__feed__desc">
                        {f.description}
                      </text>
                    </div>
                  </div>
                </Item>
              </Grid>
            ))}
          </Grid>
        ) : profileHome === 1 ? (
          <Grid
            container
            spacing={{ xs: 2, md: 3 }}
            columns={{ xs: 4, sm: 8, md: 12 }}
          >
            {followingItem.map((f, index) => (
              <Grid item xs={2} sm={4} md={4} key={index}>
                <Item className="item__item">
                  <div className="item__image2">
                    <Avatar className="profile__user__avatar" src={f.image} />
                    <div className="item__details">
                      <text>
                        <strong>{f.name}</strong>
                      </text>
                      <br />
                      <text className="profile__feed__desc">
                        {f.followings} followings | {f.followers} followers
                      </text>
                    </div>
                  </div>
                </Item>
              </Grid>
            ))}
          </Grid>
        ) : (
          <Grid
            container
            spacing={{ xs: 2, md: 3 }}
            columns={{ xs: 4, sm: 8, md: 12 }}
          >
            {followerItem.map((f, index) => (
              <Grid item xs={2} sm={4} md={4} key={index}>
                <Item
                  className="item__item"
                  //   onClick={() => handleFollowerClick(index)}
                >
                  <div className="item__image2">
                    <Avatar className="profile__user__avatar" src={f.image} />
                    <div className="item__details">
                      <text>
                        <strong>{f.name}</strong>
                      </text>
                      <br />
                      <text className="profile__feed__desc">
                        {f.followings} followings | {f.followers} followers |{" "}
                        {index}
                      </text>
                    </div>
                  </div>
                </Item>
              </Grid>
            ))}
          </Grid>
        )}
      </div>
    </div>
  );
}

export default Profile;
