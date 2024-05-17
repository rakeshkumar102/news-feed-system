import React from "react";
import "./comment.css"
import { Avatar } from "@mui/material"

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

function Comment({ user, image, comment }) {
    return (
        <div className="all__comments">

            <div className="avatar__name">
                <Avatar src={map.get(image)} />
                
            </div>
            <div className="user__comment">
            <div className="user__name">
                    {user}
                </div>
                {comment}
            </div>

        </div>
    );
}

export default Comment