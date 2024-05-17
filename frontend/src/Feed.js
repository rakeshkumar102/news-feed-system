import React from "react";
import "./feed.css";
import { Avatar } from "@mui/material";
import Comment from "./Comment";

function Feed({ username, image_url, caption, desc, user_image, comments }) {
  function handleComment(){
    
  }
  return (
    <div className="feed">
      {/* <img className="main__image" src={image_url} /> */}
      <div className="feed__content">
        <div className="feed__header">
          <Avatar className="feed__avatar" src={user_image} />
          <text className="user__name">{username}</text>
        </div>
        <h2 className="feed__text">{caption}</h2>
        <p className="desc__text"> {desc}</p>
        <div>
          <form>
            
          <input  className="comment_section" type="text" placeholder="Enter Your comment"/>
          <button className="button-53" onClick={() => handleComment()}>Enter</button>
          
          </form>
        </div>
        <div className="user__comments">
      {
      comments.map((c) => (
        <Comment
          user={c.user.name}
          image={c.user.image}
          comment={c.comment}/>
      ))
    }
    </div> 
      </div>
        
    </div>
  );
  
}

export default Feed;
