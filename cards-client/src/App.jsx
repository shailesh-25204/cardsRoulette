
import React, {useState,useCallback,useEffect} from 'react'
import useWebSocket,{ReadyState} from 'react-use-websocket'
import Header from './components/Header/Header';
import Game from './components/Game/Game';
import {useDispatch, useSelector} from 'react-redux'
import {updateLeaderboard, updateGameState, changeUrl, updateStatus ,newJsonMsg, updateUsername} from './features/websocketSlice';


function App() {
  // const socketUrl = process.env.WS_SERVER_ENDPOINT
  // const [socketUrl, setSocketUrl] = useState('ws://localhost:8080');
  const dispatch = useDispatch()
  // const [messageHistory, setMessageHistory] = useState([]);
  const socketUrl = useSelector(state => state.socketUrl)
  // console.log("THIS IS WSURL = ", import.meta.env.VITE_WS_SERVER_ENDPOINT)
  const { sendMessage, lastJsonMessage,  readyState } = useWebSocket(socketUrl,{
    share: true,
  });
  const connectionStatus = {
    [ReadyState.CONNECTING]: 'Connecting',
    [ReadyState.OPEN]: 'Open',
    [ReadyState.CLOSING]: 'Closing',
    [ReadyState.CLOSED]: 'Closed',
    [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
  }[readyState];


  useEffect(() => {
    console.log("recieved this json message ",lastJsonMessage)
    if (lastJsonMessage != null) {
      dispatch(newJsonMsg(lastJsonMessage))
      if(lastJsonMessage.type === "regSuccess" && lastJsonMessage.username !== ""){
        dispatch(updateUsername(lastJsonMessage.data))
      }
      else if (lastJsonMessage.type === "gameState"){
        dispatch(updateGameState(lastJsonMessage))
      }
      else if(lastJsonMessage.type === "leaderboard"){
        dispatch(updateLeaderboard(lastJsonMessage))
      }

    }
  }, [lastJsonMessage ]);

  useEffect(() => {
    dispatch(updateStatus(connectionStatus))
  },[connectionStatus])

  useEffect(() => {
    dispatch(changeUrl(socketUrl))
  },[socketUrl])
  // const handleClickSendMessage = useCallback(() => sendMessage('reveal'), []);
  const [inputVal,setInputVal] = useState("")
  const userSubmitHandler = (e) => {
    e.preventDefault()
    if(inputVal === "") return;
    const msg = inputVal
    setUsername(msg)
    const req = {
      type: "newUser",
      data: msg
    }
    // console.log("sending req",req)
    sendMessage(JSON.stringify(req))
  }
  const [username,setUsername] = useState("")
  const player = useSelector(state => state.username)
  return (
    <> 
    <div className='h-dvh w-dvw bg-main-texture bg-repeat overflow-hidden'>
      {
        player === "" && (
          <div className='absolute z-50 h-full w-full backdrop-blur-md flex items-center justify-center'>
          <div>
                <form id="#username" className="m-4 flex" onSubmit={userSubmitHandler} >
                    <input onChange={(e) => setInputVal(e.target.value)} type="text" value={inputVal} className="rounded-l-lg p-4 border-t mr-0 border-b border-l text-gray-800 border-gray-200 bg-white" placeholder="enter username" />
                    <button type='submit' className="px-8 rounded-r-lg bg-yellow-400  text-gray-800 font-bold p-4 uppercase border-yellow-500 border-t border-b border-r">Play!</button>
                </form>
                </div>
          </div>
          
        )
      }
      <Header/>
      <Game/>

    </div>
    </>
  )
}

export default App
