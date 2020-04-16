import React, {useState} from "react";

const SendMessage = ({sendMessage}) => {
  const [text, updateText] = useState('')

  let ta

  return (
    <div className="chat-message clearfix">
                <textarea ref={node => ta = node} name="message-to-send" id="message-to-send" placeholder="Type your message"
                          rows="3" onChange={e => {updateText(e.target.value)}}></textarea>

      <i className="fa fa-file-o"></i> &nbsp;&nbsp;&nbsp;
      <i className="fa fa-file-image-o"></i>

      <button onClick={() => {sendMessage(text); ta.value = ""}}>Send</button>

    </div>
  )
}

export default SendMessage