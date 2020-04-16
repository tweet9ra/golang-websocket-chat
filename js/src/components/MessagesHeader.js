import React from "react";

const bg = "https://sjferret.com/wp-content/uploads/Thinking-of-getting-a-cat.png"

const MessagesHeader = ({chat}) => {
  return (
    <div className="chat-header clearfix">
      <div className="avatar" style={{backgroundImage: `url(${bg})`}}></div>

      <div className="chat-about">
        <div className="chat-with">{chat && chat.name}</div>
        <div className="chat-num-messages">already {chat && chat.messages ? chat.messages.length : 0} messages</div>
      </div>
      <i className="fa fa-star"></i>
    </div>
  )
}

export default MessagesHeader