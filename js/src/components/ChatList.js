import React from "react";

const bg = "https://sjferret.com/wp-content/uploads/Thinking-of-getting-a-cat.png"

const ChatList = ({chats, selectChat, selectedChatId}) => (
  <ul className="list">
    {chats && chats.map((chat) =>
      <li key={chat.id} className={'clearfix selectChat ' + (chat.id === selectedChatId ? 'active' : '')} onClick={selectChat(chat.id)}>
        <div className="avatar" style={{backgroundImage: `url(${bg})`}}></div>
        <div className="about">
          <div className="name">{chat.name}</div>
          <div className="status">
            <i className='fa fa-circle '></i> online??
          </div>
        </div>
      </li>
    )}
  </ul>
)

export default ChatList