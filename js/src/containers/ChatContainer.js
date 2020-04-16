import React, {useEffect, useState} from 'react'
import Logout from "./Logout";
import ChatList from "../components/ChatList";
import { connect } from 'react-redux'
import MessagesHeader from "../components/MessagesHeader";
import {loadChatsIfNotLoaded} from "../actions";
import ChatHistory from "../components/ChatHistory";
import SendMessage from "../components/SendMessage";
import {sendMessage} from "../helpers/wsConnection";

const ChatContainer = ({chats, loadChatsIfNotLoaded}) => {
  const [selectedChatId, setSelectedChatId] = useState(chats && chats[0] ? chats[0].id : 0)

  if (!selectedChatId && chats) {
    setSelectedChatId(chats[0].id)
  }

  useEffect(() => {
    loadChatsIfNotLoaded()
  })

  const handleSendMessage = (text) => {
    sendMessage('NEW_MESSAGE', {text, "chat_id": selectedChatId})
  }

  const selectedChat = chats && chats.filter(chat => chat.id === selectedChatId)[0]

  return (
    <div className="container clearfix">
      <div className="people-list" id="people-list">
        <div className="search">
          <input type="text" placeholder="search"/>
          <i className="fa fa-search"></i>
        </div>
        <ChatList chats={chats} selectChat={id => () => setSelectedChatId(id)} selectedChatId={selectedChatId}/>
        <Logout/>
      </div>

      <div className="chat">
        <MessagesHeader chat={selectedChat}/>

        <ChatHistory chat={selectedChat}/>

        <SendMessage sendMessage={handleSendMessage}/>
      </div>

    </div>
  )
}


const mapStateToProps = state => ({
    chats: state.chats
})

const mapDispatchToProps = dispatch => ({
  loadChatsIfNotLoaded: () => dispatch(loadChatsIfNotLoaded())
})

export default connect(mapStateToProps, mapDispatchToProps)(ChatContainer)