import React, { useCallback, useState } from 'react'
import useWebSocket from 'react-use-websocket'
import { useSelector } from 'react-redux'

const cardStyles = [
    {
        pos: "-translate-x-4  -translate-y-0.5 -rotate-12",
        col: "bg-[#a78ee4]",
    },
    {
        pos: "-translate-x-2  -translate-y-0.5 -rotate-6",
        col: "bg-[#4cbbe3]",
    },
    {
        pos: "translate-x-0  -translate-y-0.5 rotate-0",
        col: "bg-[#9ed466]",
    },
    {
        pos: "translate-x-2  translate-y-0 rotate-6",
        col: "bg-[#faca53]",
    },
    {
        pos: "translate-x-4  translate-y-0.5 rotate-12",
        col: "bg-[#e75362]",
    }
]
function Card() {
    const btnReqMap = {
        "New Game": "newGame",
        "Reveal" : "reveal"
    }
    const socketUrl = useSelector(state => state.socketUrl)
    const [btnText,setBtnText] = useState("New Game")
    const index = useSelector(state => state.gameState.len)
    if (btnText!=="New Game" && index == 0) {
        setBtnText("New Game")
    }
    if(btnText!=="Reveal" && index != 0){
        setBtnText("Reveal")
    }


    // console.log("this is new ind ", index)
    const { sendMessage } = useWebSocket(socketUrl,{
        share: true,
      });

    const handleReveal = useCallback((e)=>{
        setBtnText("Reveal")
        const req = {
            type: "move",
            data: btnReqMap[btnText],
        }
        sendMessage(JSON.stringify(req))
    })
  return (
    <div className='flex flex-col items-center justify-center'>
    <div className='size-36 translate-x-6'>
        {
            cardStyles.map((style,ind) => {
            return (
                <div key = {ind} className={`h-36 w-24 ${style.col} absolute  border-[2px] rounded-md ${style.pos} ${ind >= index ? "hidden" : ""}`}>
                </div>
            )}
            )
        }
    </div>
    <button onClick={(e) =>handleReveal(e)} className='m-5 text-slate-300  px-5 border-2 rounded-lg border-cyan-50'>
        {btnText}
    </button>
    </div>
  )
}

export default Card